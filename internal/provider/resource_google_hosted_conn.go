package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGoogleRequestHostConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourceGoogleReqHostConnCreate,
		UpdateContext: resourceGoogleReqHostConnUpdate,
		ReadContext:   resourceGoogleReqHostConnRead,
		DeleteContext: resourceGoogeReqHostConnDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
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
			"google_pairing_key": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Google pairing key to use for this connection. This is provided when you create the VLAN attachment from the Google Cloud console. Optional if set in cloud_settings, otherwise Required.",
			},
			"google_vlan_attachment_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The name you used for your VLAN attachment in Google. Optional if set in cloud_settings, otherwise Required.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The circuit ID of the PacketFabric port you wish to connect to Google. This starts with \"PF-AP-\".",
			},
			"vlan": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid VLAN range is from 4-4094, inclusive.",
			},
			"src_svlan": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid S-VLAN range is from 4-4094, inclusive.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the hosted connection should be provisioned (the cloud on-ramp).",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The speed of the new connection.\n\n\t Available: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
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
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
									"customer_asn": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validatePublicOrPrivateASN,
										Description:  "The customer ASN of this connection.",
									},
									"remote_asn": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      16550,
										ValidateFunc: validatePublicOrPrivateASN,
										Description:  "The Google ASN of this connection. Must be 16550, between 64512 and 65534, or between 4200000000 and 4294967294.",
									},
									"md5": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The MD5 value of the authenticated BGP sessions.",
									},
									"google_keepalive_interval": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      20,
										ValidateFunc: validation.IntBetween(20, 60),
										Description:  "The Keepalive Interval. Must be between 20 and 60. ",
									},
									"google_advertised_ip_ranges": {
										Type:        schema.TypeList,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "An array of prefixes that will be advertised. Advertise Mode set to DEFAULT if not set.",
									},
								},
							},
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceGoogleReqHostConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	reqConn := extractGoogleReqConn(d)
	expectedResp, err := c.CreateRequestHostedGoogleConn(reqConn)
	if err != nil {
		return diag.FromErr(err)
	}
	// Cloud Everywhere: if cloud_circuit_id is null display error
	if expectedResp.CloudCircuitID == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Hosted location Requested",
			Detail: "On-ramp location does not have a Hosted port currently available. " +
				"Check in the Portal when your hosted cloud is provisioned and import the resource into your Terraform state file.",
		})
		return diags
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			dedicatedConns, err := c.GetCurrentCustomersHosted()
			if dedicatedConns != nil && err == nil && len(dedicatedConns) > 0 {
				for _, conn := range dedicatedConns {
					if expectedResp.UUID == conn.UUID && conn.State == "active" {
						expectedResp.CloudCircuitID = conn.CloudCircuitID
						ticker.Stop()
						createOk <- true
					}
				}
			}
		}
	}()
	<-createOk
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(expectedResp.CloudCircuitID)

	if labels, ok := d.GetOk("labels"); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}
	return diags
}

func resourceGoogleReqHostConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("cloud_circuit_id", resp.CloudCircuitID)
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("vlan", resp.Settings.VlanIDCust)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("pop", resp.CloudProvider.Pop)
		if _, ok := d.GetOk("google_pairing_key"); ok {
			_ = d.Set("google_pairing_key", resp.Settings.GooglePairingKey)
		}
		if _, ok := d.GetOk("google_vlan_attachment_name"); ok {
			_ = d.Set("google_vlan_attachment_name", resp.Settings.GoogleVlanAttachmentName)
		}
		_ = d.Set("po_number", resp.PONumber)

		if _, ok := d.GetOk("cloud_settings"); ok {
			cloudSettings := make(map[string]interface{})
			cloudSettings["credentials_uuid"] = resp.CloudSettings.CredentialsUUID
			cloudSettings["google_region"] = resp.CloudSettings.GoogleRegion
			cloudSettings["google_project_id"] = resp.CloudSettings.GoogleProjectID
			cloudSettings["google_vlan_attachment_name"] = resp.CloudSettings.GoogleVlanAttachmentName
			cloudSettings["google_cloud_router_name"] = resp.CloudSettings.GoogleCloudRouterName
			cloudSettings["google_vpc_name"] = resp.CloudSettings.GoogleVPCName
			cloudSettings["google_edge_availability_domain"] = resp.CloudSettings.GoogleEdgeAvailabilityDomain
			cloudSettings["mtu"] = resp.CloudSettings.Mtu

			bgpSettings := make(map[string]interface{})
			bgpSettings["customer_asn"] = resp.CloudSettings.BgpSettings.CustomerAsn
			bgpSettings["remote_asn"] = resp.CloudSettings.BgpSettings.RemoteAsn
			bgpSettings["md5"] = resp.CloudSettings.BgpSettings.Md5
			bgpSettings["google_keepalive_interval"] = resp.CloudSettings.BgpSettings.GoogleKeepaliveInterval
			bgpSettings["google_advertised_ip_ranges"] = resp.CloudSettings.BgpSettings.GoogleAdvertisedIPRanges

			cloudSettings["bgp_settings"] = bgpSettings
		}
	}
	resp2, err2 := c.GetBackboneByVcCID(d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	if resp2 != nil {
		_ = d.Set("port", resp2.Interfaces[0].PortCircuitID) // Port A
		if resp2.Interfaces[0].Svlan != 0 {
			_ = d.Set("src_svlan", resp2.Interfaces[0].Svlan) // Port A if ENNI
		}
		_ = d.Set("zone", resp2.Interfaces[1].Zone) // Port Z
	}

	labels, err3 := getLabels(c, d.Id())
	if err3 != nil {
		return diag.FromErr(err3)
	}
	_ = d.Set("labels", labels)
	return diags
}

func resourceGoogleReqHostConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesHostedUpdate(ctx, d, m)
}

func resourceGoogeReqHostConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudSourceDelete(ctx, d, m, "Google Service Delete")
}

func extractGoogleReqConn(d *schema.ResourceData) packetfabric.GoogleReqHostedConn {
	hostedGoogleConn := packetfabric.GoogleReqHostedConn{}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedGoogleConn.AccountUUID = accountUUID.(string)
	}
	if pairingKey, ok := d.GetOk("google_pairing_key"); ok {
		hostedGoogleConn.GooglePairingKey = pairingKey.(string)
	}
	if vlanAttach, ok := d.GetOk("google_vlan_attachment_name"); ok {
		hostedGoogleConn.GoogleVlanAttachmentName = vlanAttach.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		hostedGoogleConn.Description = description.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		hostedGoogleConn.Pop = pop.(string)
	}
	if port, ok := d.GetOk("port"); ok {
		hostedGoogleConn.Port = port.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		hostedGoogleConn.Vlan = vlan.(int)
	}
	if srcSvlan, ok := d.GetOk("src_svlan"); ok {
		hostedGoogleConn.SrcSvlan = srcSvlan.(int)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedGoogleConn.Speed = speed.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		hostedGoogleConn.PONumber = poNumber.(string)
	}
	if cloudSettingsList, ok := d.GetOk("cloud_settings"); ok {
		cs := cloudSettingsList.([]interface{})[0].(map[string]interface{})
		hostedGoogleConn.CloudSettings = &packetfabric.CloudSettingsHosted{}
		hostedGoogleConn.CloudSettings.CredentialsUUID = cs["credentials_uuid"].(string)
		hostedGoogleConn.CloudSettings.GoogleRegion = cs["google_region"].(string)
		if googleProjectID, ok := cs["google_project_id"]; ok {
			hostedGoogleConn.CloudSettings.GoogleProjectID = googleProjectID.(string)
		}
		hostedGoogleConn.CloudSettings.GoogleVlanAttachmentName = cs["google_vlan_attachment_name"].(string)
		hostedGoogleConn.CloudSettings.GoogleCloudRouterName = cs["google_cloud_router_name"].(string)
		if googleVPCName, ok := cs["google_vpc_name"]; ok {
			hostedGoogleConn.CloudSettings.GoogleVPCName = googleVPCName.(string)
		}
		if googleEdgeAvailabilityDomain, ok := cs["google_edge_availability_domain"]; ok {
			hostedGoogleConn.CloudSettings.GoogleEdgeAvailabilityDomain = googleEdgeAvailabilityDomain.(int)
		}
		if mtu, ok := cs["mtu"]; ok {
			hostedGoogleConn.CloudSettings.Mtu = mtu.(int)
		}
		if bgpSettings, ok := cs["bgp_settings"]; ok {
			bgpSettingsMap := bgpSettings.([]interface{})[0].(map[string]interface{})
			hostedGoogleConn.CloudSettings.BgpSettings = &packetfabric.BgpSettings{}
			hostedGoogleConn.CloudSettings.BgpSettings.CustomerAsn = bgpSettingsMap["customer_asn"].(int)
			hostedGoogleConn.CloudSettings.BgpSettings.RemoteAsn = bgpSettingsMap["remote_asn"].(int)
			if md5, ok := bgpSettingsMap["md5"]; ok {
				hostedGoogleConn.CloudSettings.BgpSettings.Md5 = md5.(string)
			}
			if googleKeepaliveInterval, ok := bgpSettingsMap["google_keepalive_interval"]; ok {
				hostedGoogleConn.CloudSettings.BgpSettings.GoogleKeepaliveInterval = googleKeepaliveInterval.(int)
			}
			if googleAdvertisedIPRangesInterface, ok := bgpSettingsMap["google_advertised_ip_ranges"].([]interface{}); ok {
				googleAdvertisedIPRanges := make([]string, len(googleAdvertisedIPRangesInterface))
				for i, elem := range googleAdvertisedIPRangesInterface {
					googleAdvertisedIPRanges[i] = elem.(string)
				}
				hostedGoogleConn.CloudSettings.BgpSettings.GoogleAdvertisedIPRanges = googleAdvertisedIPRanges
			}
		}
	}
	return hostedGoogleConn
}
