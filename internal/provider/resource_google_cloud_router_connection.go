package provider

import (
	"context"
	"errors"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGoogleCloudRouterConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGoogleCloudRouterConnCreate,
		ReadContext:   resourceGoogleCloudRouterConnRead,
		UpdateContext: resourceGoogleCloudRouterConnUpdate,
		DeleteContext: resourceGoogleCloudRouterConnDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"bgp_settings_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "BGP session ID generated when the cloud-side connection is provisioned by PacketFabric.",
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
			"google_pairing_key": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Google pairing key to use for this connection. This is provided when you create the VLAN attachment from the Google Cloud console. Required if not using cloud_settings.",
			},
			"google_vlan_attachment_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The name you used for your VLAN attachment in Google.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"pop": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The POP in which you want to provision the connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The desired speed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"labels": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Label value linked to an object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cloud_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Provision the cloud side of the connection with PacketFabric.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credentials_uuid": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.IsUUID,
							Description:  "The UUID of the credentials to be used with this connection.",
						},
						"mtu": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1500,
							ValidateFunc: validation.IntInSlice([]int{1500, 1440}),
							Description:  "Maximum Transmission Unit this port supports (size of the largest supported PDU).\n\n\tEnum: [\"1500\", \"1440\"] ",
						},
						"google_region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Google region that should be used.\n\n\tEnum: Enum: [\"asia-east1\", \"asia-east2\", \"asia-northeast1\", \"asia-northeast2\", \"asia-northeast3\", \"asia-south1\", \"asia-southeast1\", \"asia-southeast2\", \"australia-southeast1\", \"europe-north1\", \"europe-west1\", \"europe-west2\", \"europe-west3\", \"europe-west4\", \"europe-west6\", \"northamerica-northeast1\", \"southamerica-east1\", \"us-central1\", \"us-east1\", \"us-east4\", \"us-west1\", \"us-west2\", \"us-west3\", \"us-west4\"]",
							ValidateFunc: validation.StringInSlice([]string{
								"asia-east1",
								"asia-east2",
								"asia-northeast1",
								"asia-northeast2",
								"asia-northeast3",
								"asia-south1",
								"asia-southeast1",
								"asia-southeast2",
								"australia-southeast1",
								"europe-north1",
								"europe-west1",
								"europe-west2",
								"europe-west3",
								"europe-west4",
								"europe-west6",
								"northamerica-northeast1",
								"southamerica-east1",
								"us-central1",
								"us-east1",
								"us-east4",
								"us-west1",
								"us-west2",
								"us-west3",
								"us-west4",
							}, false),
						},
						"google_project_id": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The Google Project Id to be used. If not present the project id of the credentials will be used.",
						},
						"google_vlan_attachment_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The Google Interconnect Attachment name. No whitespace allowed.",
						},
						"google_pairing_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The Google pairing key to use for this connection. This is provided when you create the VLAN attachment from the Google Cloud console. Required if not using cloud_settings.",
						},
						"google_cloud_router_name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The Google Cloud Router Attachment name. No whitespace allowed.",
						},
						"google_vpc_name": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The Google VPC name. Required if a new router needs to be created.",
						},
						"google_edge_availability_domain": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validation.IntInSlice([]int{1, 2}),
							Description:  "The Google Edge Availability Domain. Must be 1 or 2.\n\n\tEnum: [\"1\", \"2\"] ",
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
									"remote_asn": {
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     16550,
										Description: "The Google ASN of this connection. Must be 16550, between 64512 and 65534, or between 4200000000 and 4294967294.",
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
													Description:  "If using NAT overload, the direction of the NAT connection. \n\t\tEnum: output, input. ",
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
																Description:  "The private prefix of this DNAT mapping.",
															},
															"public_prefix": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validateIPAddressWithPrefix,
																Description:  "The public prefix of this DNAT mapping.",
															},
															"conditional_prefix": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validateIPAddressWithPrefix,
																Description:  "The conditional prefix prefix of this DNAT mapping.",
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
									"google_keepalive_interval": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      20,
										ValidateFunc: validation.IntBetween(20, 60),
										Description:  "The Keepalive Interval. Must be between 20 and 60. ",
									},
								},
							},
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: CloudRouterImportStatePassthroughContext,
		},
	}
}

func resourceGoogleCloudRouterConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	googleRoute := extractGoogleRouterConn(d)
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

		resp, err := c.CreateGoogleCloudRouterConn(googleRoute, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		if err := checkCloudRouterConnectionStatus(c, cid.(string), resp.CloudCircuitID); err != nil {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.CloudCircuitID)
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

func resourceGoogleCloudRouterConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		resp, err := c.ReadCloudRouterConnection(cid.(string), cloudConnCID.(string))
		if err != nil {
			diags = diag.FromErr(err)
			return diags
		}

		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("circuit_id", resp.CloudRouterCircuitID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("pop", resp.CloudProvider.Pop)
		if _, ok := d.GetOk("google_pairing_key"); ok {
			_ = d.Set("google_pairing_key", resp.CloudSettings.GooglePairingKey)
		}
		if _, ok := d.GetOk("google_vlan_attachment_name"); ok {
			_ = d.Set("google_vlan_attachment_name", resp.CloudSettings.GoogleVlanAttachmentName)
		}
		_ = d.Set("po_number", resp.PONumber)

		if _, ok := d.GetOk("cloud_settings"); ok {
			// Extract the BGP settings UUID
			var bgpSettingsUUID string
			if len(resp.BgpStateList) > 0 {
				bgpSettingsUUID = resp.BgpStateList[0].BgpSettingsUUID
				_ = d.Set("bgp_settings_uuid", bgpSettingsUUID)
			}
			bgp, err := c.GetBgpSessionBy(cid.(string), cloudConnCID.(string), bgpSettingsUUID)
			if err != nil {
				return diag.FromErr(errors.New("could not retrieve bgp session"))
			}
			cloudSettings := make(map[string]interface{})
			cloudSettings["credentials_uuid"] = resp.CloudSettings.CredentialsUUID
			cloudSettings["google_region"] = resp.CloudSettings.GoogleRegion
			if googleProjectID, ok := d.GetOk("cloud_settings.0.google_project_id"); ok {
				cloudSettings["google_project_id"] = googleProjectID
			}
			cloudSettings["google_vlan_attachment_name"] = resp.CloudSettings.GoogleVlanAttachmentName
			cloudSettings["google_pairing_key"] = resp.CloudSettings.GooglePairingKey
			cloudSettings["google_cloud_router_name"] = resp.CloudSettings.GoogleCloudRouterName
			cloudSettings["google_edge_availability_domain"] = resp.CloudSettings.GoogleEdgeAvailabilityDomain
			if googleVpcName, ok := d.GetOk("cloud_settings.0.google_vpc_name"); ok {
				cloudSettings["google_vpc_name"] = googleVpcName
			}
			cloudSettings["mtu"] = resp.CloudSettings.Mtu
			bgpSettings := make(map[string]interface{})
			bgpSettings["google_keepalive_interval"] = resp.CloudSettings.BgpSettings.GoogleKeepaliveInterval
			bgpSettings["remote_asn"] = bgp.RemoteAsn
			bgpSettings["disabled"] = bgp.Disabled
			bgpSettings["orlonger"] = bgp.Orlonger
			bgpSettings["md5"] = bgp.Md5
			bgpSettings["med"] = bgp.Med
			bgpSettings["as_prepend"] = bgp.AsPrepend
			bgpSettings["local_preference"] = bgp.LocalPreference
			bgpSettings["bfd_interval"] = bgp.BfdInterval
			bgpSettings["bfd_multiplier"] = bgp.BfdMultiplier
			if bgp.Nat != nil {
				nat := flattenNatConfiguration(bgp.Nat)
				bgpSettings["nat"] = nat
			}
			prefixes := flattenPrefixConfiguration(bgp.Prefixes)
			bgpSettings["prefixes"] = prefixes
			cloudSettings["bgp_settings"] = []interface{}{bgpSettings}
			_ = d.Set("cloud_settings", []interface{}{cloudSettings})
		}
		// unsetFields: published_quote_line_uuid
	}

	labels, err2 := getLabels(c, d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	_ = d.Set("labels", labels)
	return diags
}

func resourceGoogleCloudRouterConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceGoogleCloudRouterConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractGoogleRouterConn(d *schema.ResourceData) packetfabric.GoogleCloudRouterConn {
	googleRoute := packetfabric.GoogleCloudRouterConn{}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		googleRoute.AccountUUID = accountUUID.(string)
	}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		googleRoute.MaybeNat = maybeNat.(bool)
	}
	if maybeDNat, ok := d.GetOk("maybe_dnat"); ok {
		googleRoute.MaybeDNat = maybeDNat.(bool)
	}
	if pairingKey, ok := d.GetOk("google_pairing_key"); ok {
		googleRoute.GooglePairingKey = pairingKey.(string)
	}
	if vlanAttName, ok := d.GetOk("google_vlan_attachment_name"); ok {
		googleRoute.GoogleVlanAttachmentName = vlanAttName.(string)
	}
	if desc, ok := d.GetOk("description"); ok {
		googleRoute.Description = desc.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		googleRoute.Pop = pop.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		googleRoute.Speed = speed.(string)
	}
	if publishedQuoteLine, ok := d.GetOk("published_quote_line_uuid"); ok {
		googleRoute.PublishedQuoteLineUUID = publishedQuoteLine.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		googleRoute.PONumber = poNumber.(string)
	}
	if cloudSettingsList, ok := d.GetOk("cloud_settings"); ok {
		cs := cloudSettingsList.([]interface{})[0].(map[string]interface{})
		googleRoute.CloudSettings = extractGoogleRouterCloudConnSettings(cs)
	}
	return googleRoute
}

func extractGoogleRouterCloudConnSettings(cs map[string]interface{}) *packetfabric.CloudSettings {
	cloudSettings := &packetfabric.CloudSettings{}
	cloudSettings.CredentialsUUID = cs["credentials_uuid"].(string)
	cloudSettings.GoogleRegion = cs["google_region"].(string)

	if googleProjectID, ok := cs["google_project_id"]; ok {
		cloudSettings.GoogleProjectID = googleProjectID.(string)
	}
	cloudSettings.GoogleVlanAttachmentName = cs["google_vlan_attachment_name"].(string)
	cloudSettings.GoogleCloudRouterName = cs["google_cloud_router_name"].(string)

	if googleVPCName, ok := cs["google_vpc_name"]; ok {
		cloudSettings.GoogleVPCName = googleVPCName.(string)
	}
	if googleEdgeAvailabilityDomain, ok := cs["google_edge_availability_domain"]; ok {
		cloudSettings.GoogleEdgeAvailabilityDomain = googleEdgeAvailabilityDomain.(int)
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
