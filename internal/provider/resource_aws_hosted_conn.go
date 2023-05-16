package provider

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAwsRequestHostConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourceAwsReqHostConnCreate,
		UpdateContext: resourceAwsReqHostConnUpdate,
		ReadContext:   resourceAwsReqHostConnRead,
		DeleteContext: resourceAwsServicesDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_AWS_ACCOUNT_ID", nil),
				ValidateFunc: validation.StringIsNotEmpty,
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
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the hosted connection should be provisioned (the cloud on-ramp).",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the PacketFabric port you want to connect to AWS. This starts with \"PF-AP-\".",
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
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired zone of the new connection.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The speed of the new connection.\n\n\tAvailable: 50Mbps 100Mbps 200Mbps 300Mbps 400Mbps 500Mbps 1Gbps 2Gbps 5Gbps 10Gbps",
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
						"bgp_settings": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"customer_asn": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The customer ASN of this connection.",
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
									"advertised_prefixes": {
										Type:        schema.TypeList,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "An array of prefixes that will be advertised. Required for public VIFs.",
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
									"allowed_prefixes": {
										Type:        schema.TypeList,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "An array of allowed prefixes. Required on the DirectConnect Gateway when the other Gateway is of type transit.",
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
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAwsReqHostConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	reqConn := extractAwsReqConn(d)
	expectedResp, err := c.CreateAwsHostedConn(reqConn)
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

	if _, ok := d.GetOk("cloud_settings"); !ok {
		time.Sleep(90 * time.Second) // wait for the connection to show on AWS
		resp, err := c.GetCloudConnInfo(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		if resp.CloudProviderConnectionID == "" || resp.Settings.VlanIDPf == 0 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Incomplete Cloud Information",
				Detail:   "The cloud_provider_connection_id and/or vlan_id_pf are currently unavailable.",
			})
			return diags
		} else {
			_ = d.Set("cloud_provider_connection_id", resp.CloudProviderConnectionID)
			_ = d.Set("vlan_id_pf", resp.Settings.VlanIDPf)
		}
	}

	if labels, ok := d.GetOk("labels"); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}
	return diags
}

func resourceAwsReqHostConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetCloudConnInfo(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("description", resp.Description)
		_ = d.Set("vlan", resp.Settings.VlanIDCust)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("pop", resp.CloudProvider.Pop)
		_ = d.Set("aws_account_id", resp.Settings.AwsAccountID)
		if _, ok := d.GetOk("po_number"); ok {
			_ = d.Set("po_number", resp.PONumber)
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
		} else {
			if resp.CloudProviderConnectionID == "" || resp.Settings.VlanIDPf == 0 {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Incomplete Cloud Information",
					Detail:   "The cloud_provider_connection_id and/or vlan_id_pf are currently unavailable.",
				})
				return diags
			} else {
				_ = d.Set("cloud_provider_connection_id", resp.CloudProviderConnectionID)
				_ = d.Set("vlan_id_pf", resp.Settings.VlanIDPf)
			}
		}
	}
	resp2, err2 := c.GetBackboneByVcCID(d.Id())
	if err2 != nil {
		return diag.FromErr(err2)
	}
	if resp2 != nil {
		_ = d.Set("port", resp2.Interfaces[0].PortCircuitID) // Port A
		if _, ok := d.GetOk("src_svlan"); ok {
			if resp2.Interfaces[0].Svlan != 0 {
				_ = d.Set("src_svlan", resp2.Interfaces[0].Svlan) // Port A if ENNI
			}
		}
		if _, ok := d.GetOk("zone"); ok {
			_ = d.Set("zone", resp2.Interfaces[1].Zone) // Port Z
		}
	}

	if _, ok := d.GetOk("labels"); ok {
		labels, err3 := getLabels(c, d.Id())
		if err3 != nil {
			return diag.FromErr(err3)
		}
		_ = d.Set("labels", labels)
	}

	etl, err4 := c.GetEarlyTerminationLiability(d.Id())
	if err4 != nil {
		return diag.FromErr(err4)
	}
	if etl > 0 {
		_ = d.Set("etl", etl)
	}

	return diags
}

func resourceAwsReqHostConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceServicesHostedUpdate(ctx, d, m)
}

func extractAwsReqConn(d *schema.ResourceData) packetfabric.HostedAwsConnection {
	hostedAwsConn := packetfabric.HostedAwsConnection{}
	if awsAccountID, ok := d.GetOk("aws_account_id"); ok {
		hostedAwsConn.AwsAccountID = awsAccountID.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		hostedAwsConn.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		hostedAwsConn.Description = description.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		hostedAwsConn.Pop = pop.(string)
	}
	if port, ok := d.GetOk("port"); ok {
		hostedAwsConn.Port = port.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		hostedAwsConn.Vlan = vlan.(int)
	}
	if srcSvlan, ok := d.GetOk("src_svlan"); ok {
		hostedAwsConn.SrcSvlan = srcSvlan.(int)
	}
	if zone, ok := d.GetOk("zone"); ok {
		hostedAwsConn.Zone = zone.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		hostedAwsConn.Speed = speed.(string)
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		hostedAwsConn.PONumber = poNumber.(string)
	}
	if cloudSettingsList, ok := d.GetOk("cloud_settings"); ok {
		cs := cloudSettingsList.([]interface{})[0].(map[string]interface{})
		hostedAwsConn.CloudSettings = &packetfabric.CloudSettings{}
		hostedAwsConn.CloudSettings.CredentialsUUID = cs["credentials_uuid"].(string)
		if awsRegion, ok := cs["aws_region"]; ok {
			hostedAwsConn.CloudSettings.AwsRegion = awsRegion.(string)
		}
		if mtu, ok := cs["mtu"]; ok {
			hostedAwsConn.CloudSettings.Mtu = mtu.(int)
		}
		hostedAwsConn.CloudSettings.AwsVifType = cs["aws_vif_type"].(string)
		if bgpSettings, ok := cs["bgp_settings"]; ok {
			bgpSettingsMap := bgpSettings.([]interface{})[0].(map[string]interface{})
			hostedAwsConn.CloudSettings.BgpSettings = &packetfabric.BgpSettings{}
			hostedAwsConn.CloudSettings.BgpSettings.CustomerAsn = bgpSettingsMap["customer_asn"].(int)
			hostedAwsConn.CloudSettings.BgpSettings.AddressFamily = bgpSettingsMap["address_family"].(string)
		}
		if awsGateways, ok := cs["aws_gateways"]; ok {
			hostedAwsConn.CloudSettings.AwsGateways = extractAwsGateways(awsGateways.([]interface{}))
		}
	}
	return hostedAwsConn
}

func extractAwsGateways(gateways []interface{}) []packetfabric.AwsGateway {
	var awsGateways []packetfabric.AwsGateway
	for _, gw := range gateways {
		gateway := gw.(map[string]interface{})

		subnetIDsInterface, subnetIDsExist := gateway["subnet_ids"].([]interface{})
		var subnetIDs []string
		if subnetIDsExist {
			subnetIDs = make([]string, len(subnetIDsInterface))
			for i, elem := range subnetIDsInterface {
				subnetIDs[i] = elem.(string)
			}
		}

		allowedPrefixesInterface, allowedPrefixesExist := gateway["allowed_prefixes"].([]interface{})
		var allowedPrefixes []string
		if allowedPrefixesExist {
			allowedPrefixes = make([]string, len(allowedPrefixesInterface))
			for i, elem := range allowedPrefixesInterface {
				allowedPrefixes[i] = elem.(string)
			}
		}

		awsGateway := packetfabric.AwsGateway{}

		if t, ok := gateway["type"].(string); ok {
			awsGateway.Type = t
		}
		if name, ok := gateway["name"].(string); ok {
			awsGateway.Name = name
		}
		if id, ok := gateway["id"].(string); ok {
			awsGateway.ID = id
		}
		if asn, ok := gateway["asn"].(int); ok {
			awsGateway.Asn = asn
		}
		if vpcID, ok := gateway["vpc_id"].(string); ok {
			awsGateway.VpcID = vpcID
		}

		awsGateway.SubnetIDs = subnetIDs
		awsGateway.AllowedPrefixes = allowedPrefixes

		awsGateways = append(awsGateways, awsGateway)
	}
	return awsGateways
}
