package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterConnectionAws() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterConnectionAwsCreate,
		ReadContext:   resourceRouterConnectionAwsRead,
		UpdateContext: resourceRouterConnectionAwsUpdate,
		DeleteContext: resourceRouterConnectionAwsDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"circuit_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"bgp_settings_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "BGP session ID generated when the cloud-side connection is provisioned by PacketFabric.",
			},
			"aws_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				DefaultFunc: schema.EnvDefaultFunc("PF_AWS_ACCOUNT_ID", nil),
				Description: "The AWS account ID to connect with. Must be 12 characters long. " +
					"Can also be set with the PF_AWS_ACCOUNT_ID environment variable.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set this to true if you intend to use NAT on this connection. Defaults: false",
			},
			"maybe_dnat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Set this to true if you intend to use DNAT on this connection. Defaults: false",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A brief description of this connection.",
			},
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The POP in which you want to provision the connection.",
			},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The desired availability zone of the connection.\n\n\tExample: \"A\"",
			},
			"is_public": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether PacketFabric should allocate a public IP address for this connection. Set this to true if you intend to use a public VIF on the AWS side. ",
			},
			"speed": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The desired speed of the new connection.\n\n\t Available: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line which this connection should be associated.",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Label value linked to an object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cloud_provider_connection_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cloud provider specific connection ID, eg. the Amazon connection ID of the cloud router connection.\n\t\tExample: dxcon-fgadaaa1",
			},
			"vlan_id_pf": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "PacketFabric VLAN ID.",
			},
			"cloud_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Provision the Cloud side of the connection with PacketFabric.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credentials_uuid": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.IsUUID,
							Description:  "The UUID of the credentials to be used with this connection.",
						},
						"aws_region": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The AWS region that should be used.",
						},
						"mtu": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1500,
							ValidateFunc: validation.IntInSlice([]int{1500, 9001}),
							Description:  "Maximum Transmission Unit this port supports (size of the largest supported PDU).\n\n\tEnum: [\"1500\", \"9001\"] ",
						},
						"aws_vif_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"private", "transit", "public"}, false),
							Description:  "The type of VIF to use for this connection.",
						},
						"aws_gateways": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    2,
							Description: "Only for Private or Transit VIF.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"directconnect", "private", "transit"}, false),
										Description:  "The type of this AWS Gateway.",
									},
									"name": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The name of the AWS Gateway, required if creating a new DirectConnect Gateway.",
									},
									"id": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The ID of the AWS Gateway to be used.",
									},
									"asn": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The ASN of the AWS Gateway to be used.",
									},
									"vpc_id": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The AWS VPC ID this Gateway should be associated with. Required for private and transit Gateways.",
									},
									"subnet_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "An array of subnet IDs to associate with this Gateway. Required for transit Gateways.",
									},
								},
							},
						},
						"bgp_settings": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"md5": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The MD5 value of the authenticated BGP sessions. Required for AWS.",
									},
									"l3_address": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The prefix of the customer router. Required for public VIFs.",
									},
									"remote_address": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The prefix of the remote router. Required for public VIFs.",
									},
									"local_preference": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "The local preference for this instance. When the same route is received in multiple locations, those with a higher local preference value are preferred by the cloud router. It is used when type = in.\n\n\tAvailable range is 1 through 4294967295. ",
									},
									"med": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     0,
										Description: "The Multi-Exit Discriminator of this instance. When the same route is advertised in multiple locations, those with a lower MED are preferred by the peer AS. It is used when type = out.\n\n\tAvailable range is 1 through 4294967295. ",
									},
									"as_prepend": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      0,
										ValidateFunc: validation.IntBetween(1, 5),
										Description:  "The BGP prepend value for this instance. It is used when type = out.\n\n\tAvailable range is 1 through 5. ",
									},
									"orlonger": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether to use exact match or longer for all prefixes. ",
									},
									"bfd_interval": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      0,
										ValidateFunc: validation.IntBetween(3, 30000),
										Description:  "If you are using BFD, this is the interval (in milliseconds) at which to send test packets to peers.\n\n\tAvailable range is 3 through 30000. ",
									},
									"bfd_multiplier": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      0,
										ValidateFunc: validation.IntBetween(2, 16),
										Description:  "If you are using BFD, this is the number of consecutive packets that can be lost before BFD considers a peer down and shuts down BGP.\n\n\tAvailable range is 2 through 16. ",
									},
									"disabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
										Description: "Whether this BGP session is disabled. ",
									},
									"nat": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Translate the source or destination IP address.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pre_nat_sources": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "If using NAT overload, this is the prefixes from the cloud that you want to associate with the NAT pool.\n\n\tExample: 10.0.0.0/24",
													Elem: &schema.Schema{
														Type:        schema.TypeString,
														Description: "IP prefix using CIDR format.",
													},
												},
												"pool_prefixes": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "If using NAT overload, all prefixes that are NATed on this connection will be translated to the pool prefix address.\n\n\tExample: 10.0.0.0/32",
													Elem: &schema.Schema{
														Type:        schema.TypeString,
														Description: "IP prefix using CIDR format.",
													},
												},
												"direction": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "output",
													ValidateFunc: validation.StringInSlice([]string{"output", "input"}, true),
													Description:  "If using NAT overload, the direction of the NAT connection (input=ingress, output=egress). \n\t\tEnum: output, input. ",
												},
												"nat_type": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "overload",
													ValidateFunc: validation.StringInSlice([]string{"overload", "inline_dnat"}, true),
													Description:  "The NAT type of the NAT connection, source NAT (overload) or destination NAT (inline_dnat). \n\t\tEnum: overload, inline_dnat. ",
												},
												"dnat_mappings": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "Translate the destination IP address.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"private_prefix": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validateIPAddressWithPrefix,
																Description:  "Post-translation IP prefix.",
															},
															"public_prefix": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validateIPAddressWithPrefix,
																Description:  "Pre-translation IP prefix.",
															},
															"conditional_prefix": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validateIPAddressWithPrefix,
																Description:  "Post-translation prefix must be equal to or included within the conditional IP prefix.",
															},
														},
													},
												},
											},
										},
									},
									"prefixes": {
										Type:        schema.TypeSet,
										Required:    true,
										Description: "The list of BGP prefixes",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"prefix": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validateIPAddressWithPrefix,
													Description:  "The actual IP Prefix of this instance.",
												},
												"match_type": {
													Type:         schema.TypeString,
													Optional:     true,
													Default:      "exact",
													ValidateFunc: validation.StringInSlice([]string{"exact", "orlonger"}, true),
													Description:  "The match type of this prefix.\n\n\tEnum: `\"exact\"` `\"orlonger\"` ",
												},
												"as_prepend": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      0,
													ValidateFunc: validation.IntBetween(1, 5),
													Description:  "The BGP prepend value of this prefix. It is used when type = out.\n\n\tAvailable range is 1 through 5. ",
												},
												"med": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "The MED of this prefix. It is used when type = out.\n\n\tAvailable range is 1 through 4294967295. ",
												},
												"local_preference": {
													Type:        schema.TypeInt,
													Optional:    true,
													Default:     0,
													Description: "The local_preference of this prefix. It is used when type = in.\n\n\tAvailable range is 1 through 4294967295. ",
												},
												"type": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.StringInSlice([]string{"in", "out"}, true),
													Description:  "Whether this prefix is in (Allowed Prefixes from Cloud) or out (Allowed Prefixes to Cloud).\n\t\tEnum: in, out.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"etl": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Early Termination Liability (ETL) fees apply when terminating a service before its term ends. ETL is prorated to the remaining contract days.",
			},
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, d *schema.ResourceDiff, m interface{}) error {
				if d.Id() == "" {
					return nil
				}
				if _, ok := d.GetOk("cloud_settings"); !ok {
					return nil
				}

				attributes := []string{
					"cloud_settings.0.aws_region",
					"cloud_settings.0.aws_vif_type",
					"cloud_settings.0.aws_gateways",
				}

				for _, attribute := range attributes {
					oldRaw, newRaw := d.GetChange(attribute)
					if oldRaw != nil && !reflect.DeepEqual(oldRaw, newRaw) {
						return fmt.Errorf("updating %s in-place is not supported, delete and recreate the resource with the updated values", attribute)
					}
				}
				return nil
			},
		),
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceRouterConnectionAwsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	awsConn := extractAwsConnection(d)

	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudSettingsList := d.Get("cloud_settings").([]interface{})
		if len(cloudSettingsList) != 0 {
			cloudSettings := cloudSettingsList[0].(map[string]interface{})
			bgpSettingsList := cloudSettings["bgp_settings"].([]interface{})
			if len(bgpSettingsList) != 0 {
				bgpSettings := bgpSettingsList[0].(map[string]interface{})
				prefixesValue := bgpSettings["prefixes"]
				prefixesSet := prefixesValue.(*schema.Set)
				prefixesList := prefixesSet.List()
				if err := validatePrefixes(prefixesList); err != nil {
					return diag.FromErr(err)
				}
			}
		}
		resp, err := c.CreateAwsConnection(awsConn, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		if err := checkCloudRouterConnectionStatus(c, cid.(string), resp.CloudCircuitID); err != nil {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.CloudCircuitID)

			if _, ok := d.GetOk("cloud_settings"); !ok {
				time.Sleep(90 * time.Second) // wait for the connection to show on AWS
				resp, err := c.ReadCloudRouterConnection(cid.(string), resp.CloudCircuitID)
				if err != nil {
					diags = diag.FromErr(err)
					return diags
				}

				if resp.CloudProviderConnectionID == "" || resp.CloudSettings.VlanIDPf == 0 {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  "Incomplete Cloud Information",
						Detail:   "The cloud_provider_connection_id and/or vlan_id_pf are currently unavailable.",
					})
					return diags
				} else {
					_ = d.Set("cloud_provider_connection_id", resp.CloudProviderConnectionID)
					_ = d.Set("vlan_id_pf", resp.CloudSettings.VlanIDPf)
				}
			}

			if _, ok := d.GetOk("cloud_settings"); ok {
				// Extract the BGP settings UUID
				resp, err := c.ReadCloudRouterConnection(cid.(string), resp.CloudCircuitID)
				if err != nil {
					diags = diag.FromErr(err)
					return diags
				}
				if len(resp.BgpStateList) > 0 {
					_ = d.Set("bgp_settings_uuid", resp.BgpStateList[0].BgpSettingsUUID)
				}
			}
			if labels, ok := d.GetOk("labels"); ok {
				diagnostics, created := createLabels(c, d.Id(), labels)
				if !created {
					return diagnostics
				}
			}
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Circuit ID not present",
			Detail:   "Please provide a valid Circuit ID.",
		})
	}
	return diags
}

func resourceRouterConnectionAwsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	circuitID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.FromErr(errors.New("please provide a valid Circuit ID"))
	}

	cloudConnCID := d.Get("id")
	resp, err := c.ReadCloudRouterConnection(circuitID.(string), cloudConnCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("account_uuid", resp.AccountUUID)
	_ = d.Set("circuit_id", resp.CloudRouterCircuitID)
	_ = d.Set("description", resp.Description)
	_ = d.Set("speed", resp.Speed)
	_ = d.Set("pop", resp.CloudProvider.Pop)
	_ = d.Set("zone", resp.Zone)
	_ = d.Set("aws_account_id", resp.CloudSettings.AwsAccountID)
	_ = d.Set("po_number", resp.PONumber)

	if resp.CloudSettings.PublicIP != "" {
		_ = d.Set("is_public", true)
	} else {
		_ = d.Set("is_public", false)
	}

	if _, ok := d.GetOk("cloud_settings"); ok {
		// Extract the BGP settings UUID
		var bgpSettingsUUID string
		if len(resp.BgpStateList) > 0 {
			bgpSettingsUUID = resp.BgpStateList[0].BgpSettingsUUID
			_ = d.Set("bgp_settings_uuid", bgpSettingsUUID)
		}
		bgp, err := c.GetBgpSessionBy(circuitID.(string), cloudConnCID.(string), bgpSettingsUUID)
		if err != nil {
			return diag.FromErr(errors.New("could not retrieve bgp session"))
		}
		cloudSettings := make(map[string]interface{})
		cloudSettings["credentials_uuid"] = resp.CloudSettings.CredentialsUUID
		cloudSettings["aws_region"] = resp.CloudSettings.AwsRegion
		if _, ok := d.GetOk("cloud_settings.0.mtu"); ok {
			cloudSettings["mtu"] = resp.CloudSettings.Mtu
		}
		cloudSettings["aws_vif_type"] = resp.CloudSettings.AwsVifType
		bgpSettings := make(map[string]interface{})
		if bgp != nil {
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.remote_asn"); ok {
				bgpSettings["remote_asn"] = bgp.RemoteAsn
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.l3_address"); ok {
				bgpSettings["l3_address"] = bgp.L3Address
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.remote_address"); ok {
				bgpSettings["remote_address"] = bgp.RemoteAddress
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.disabled"); ok {
				bgpSettings["disabled"] = bgp.Disabled
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.orlonger"); ok {
				bgpSettings["orlonger"] = bgp.Orlonger
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.md5"); ok {
				bgpSettings["md5"] = bgp.Md5
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.med"); ok {
				bgpSettings["med"] = bgp.Med
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.as_prepend"); ok {
				bgpSettings["as_prepend"] = bgp.AsPrepend
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.local_preference"); ok {
				bgpSettings["local_preference"] = bgp.LocalPreference
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.bfd_interval"); ok {
				bgpSettings["bfd_interval"] = bgp.BfdInterval
			}
			if _, ok := d.GetOk("cloud_settings.0.bgp_settings.0.bfd_multiplier"); ok {
				bgpSettings["bfd_multiplier"] = bgp.BfdMultiplier
			}
			if bgp.Nat != nil {
				nat := flattenNatConfiguration(bgp.Nat)
				bgpSettings["nat"] = nat
			}
			prefixes := flattenPrefixConfiguration(bgp.Prefixes)
			bgpSettings["prefixes"] = prefixes
		}
		cloudSettings["bgp_settings"] = bgpSettings

		awsGateways := make([]map[string]interface{}, len(resp.CloudSettings.AwsGateways))
		for i, gateway := range resp.CloudSettings.AwsGateways {
			awsGateway := make(map[string]interface{})
			awsGateway["type"] = gateway.Type
			awsGateway["name"] = gateway.Name
			awsGateway["id"] = gateway.ID
			awsGateway["asn"] = gateway.Asn
			awsGateway["vpc_id"] = gateway.VpcID
			awsGateway["subnet_ids"] = gateway.SubnetIDs
			awsGateway["allowed_prefixes"] = gateway.AllowedPrefixes
			awsGateways[i] = awsGateway
		}
		cloudSettings["aws_gateways"] = awsGateways
		_ = d.Set("cloud_settings", cloudSettings)
	} else {
		if resp.CloudProviderConnectionID == "" || resp.CloudSettings.VlanIDPf == 0 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Incomplete Cloud Information",
				Detail:   "The cloud_provider_connection_id and/or vlan_id_pf are currently unavailable.",
			})
			return diags
		} else {
			_ = d.Set("cloud_provider_connection_id", resp.CloudProviderConnectionID)
			_ = d.Set("vlan_id_pf", resp.CloudSettings.VlanIDPf)
		}
	}
	// unsetFields: published_quote_line_uuid

	if _, ok := d.GetOk("labels"); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set("labels", labels)
	}

	etl, err := c.GetEarlyTerminationLiability(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if etl > 0 {
		_ = d.Set("etl", etl)
	}
	return diags
}

func resourceRouterConnectionAwsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceRouterConnectionAwsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractAwsConnection(d *schema.ResourceData) packetfabric.AwsConnection {
	awsConn := packetfabric.AwsConnection{}
	if awsAccountID, ok := d.GetOk("aws_account_id"); ok {
		awsConn.AwsAccountID = awsAccountID.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		awsConn.AccountUUID = accountUUID.(string)
	}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		awsConn.MaybeNat = maybeNat.(bool)
	}
	if maybeDNat, ok := d.GetOk("maybe_dnat"); ok {
		awsConn.MaybeDNat = maybeDNat.(bool)
	}
	if description, ok := d.GetOk("description"); ok {
		awsConn.Description = description.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		awsConn.Pop = pop.(string)
	}
	if zone, ok := d.GetOk("zone"); ok {
		awsConn.Zone = zone.(string)
	}
	if isPublic, ok := d.GetOk("is_public"); ok {
		awsConn.IsPublic = isPublic.(bool)
	}
	if speed, ok := d.GetOk("speed"); ok {
		awsConn.Speed = speed.(string)
	}
	if publishedQuoteLineUUID, ok := d.GetOk("published_quote_line_uuid"); ok {
		awsConn.PublishedQuoteLineUUID = publishedQuoteLineUUID.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		awsConn.PONumber = poNumber.(string)
	}
	if cloudSettingsList, ok := d.GetOk("cloud_settings"); ok {
		cs := cloudSettingsList.([]interface{})[0].(map[string]interface{})
		awsConn.CloudSettings = extractAwsRouterCloudConnSettings(cs)
	}
	return awsConn
}

func extractAwsRouterCloudConnSettings(cs map[string]interface{}) *packetfabric.CloudSettings {
	cloudSettings := &packetfabric.CloudSettings{}
	cloudSettings.CredentialsUUID = cs["credentials_uuid"].(string)

	if awsRegion, ok := cs["aws_region"]; ok {
		cloudSettings.AwsRegion = awsRegion.(string)
	}
	cloudSettings.AwsVifType = cs["aws_vif_type"].(string)
	if awsGateways, ok := cs["aws_gateways"]; ok {
		cloudSettings.AwsGateways = extractAwsGateways(awsGateways.([]interface{}))
	}
	if mtu, ok := cs["mtu"]; ok {
		cloudSettings.Mtu = mtu.(int)
	}
	if bgpSettings, ok := cs["bgp_settings"]; ok {
		bgpSettingsMap := bgpSettings.([]interface{})[0].(map[string]interface{})
		cloudSettings.BgpSettings = extractRouterConnBgpSettings(bgpSettingsMap)
	}
	return cloudSettings
}
