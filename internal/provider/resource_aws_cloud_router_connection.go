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

func resourceRouterConnectionAws() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterConnectionAwsCreate,
		ReadContext:   resourceRouterConnectionAwsRead,
		UpdateContext: resourceRouterConnectionAwsUpdate,
		DeleteContext: resourceRouterConnectionAwsDelete,
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
				Optional:    true,
				ForceNew:    true,
				Description: "The desired AWS availability zone of the new connection.",
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
						"bgp_settings": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
									"address_family": {
										Type:         schema.TypeString,
										Optional:     true,
										Default:      "ipv4",
										Description:  "The address family that should be used. ",
										ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),
									},
									"md5": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The MD5 value of the authenticated BGP sessions.",
									},
								},
							},
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
					},
				},
			},
		},
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
		cloudSettings := make(map[string]interface{})
		cloudSettings["credentials_uuid"] = resp.CloudSettings.CredentialsUUID
		cloudSettings["aws_region"] = resp.CloudSettings.AwsRegion
		cloudSettings["mtu"] = resp.CloudSettings.Mtu
		cloudSettings["aws_vif_type"] = resp.CloudSettings.AwsVifType

		bgpSettings := make(map[string]interface{})
		bgpSettings["customer_asn"] = resp.CloudSettings.BgpSettings.CustomerAsn
		bgpSettings["address_family"] = resp.CloudSettings.BgpSettings.AddressFamily
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
	}
	// unsetFields: published_quote_line_uuid

	labels, err2 := getLabels(c, d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	_ = d.Set("labels", labels)
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
