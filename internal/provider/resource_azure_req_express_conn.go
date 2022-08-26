package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAzureReqExpressConn() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: resourceAzureReqExpressConnCreate,
		ReadContext:   resourceAzureProvisionRead,
		UpdateContext: resourceAzureProvisionUpdate,
		DeleteContext: resourceAzureProvisionDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"azure_service_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The Service Key provided by Micosoft Azure.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID of the contact that will be billed.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of the connection.",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port to connect to Azure.",
			},
			"vlan_private": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Private peering vlan. Valid VLAN range is from 4-4094, inclusive.",
			},
			"vlan_microsoft": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Microsoft Peering VLAN.",
			},
			"src_svlan": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Valid S-VLAN range is from 4-4094, inclusive.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The desired speed of the new connection.\n\t\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAzureReqExpressConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	azureExpress := extractAzureExpressConn(d)
	expectedResp, err := c.CreateAzureExpressRoute(azureExpress)
	if err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(10 * time.Second)
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
	return diags
}

func extractAzureExpressConn(d *schema.ResourceData) packetfabric.AzureExpressRoute {
	azureExpress := packetfabric.AzureExpressRoute{}
	if azureServiceKey, ok := d.GetOk("azure_service_key"); ok {
		azureExpress.AzureServiceKey = azureServiceKey.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		azureExpress.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		azureExpress.Description = description.(string)
	}
	if port, ok := d.GetOk("port"); ok {
		azureExpress.Port = port.(string)
	}
	if vlanPrivate, ok := d.GetOk("vlan_private"); ok {
		azureExpress.VlanPrivate = vlanPrivate.(int)
	}
	if vlanMicrosoft, ok := d.GetOk("vlan_microsoft"); ok {
		azureExpress.VlanMicrosoft = vlanMicrosoft.(int)
	}
	if srcSvlan, ok := d.GetOk("src_svlan"); ok {
		azureExpress.SrcSvlan = srcSvlan.(int)
	}
	if speed, ok := d.GetOk("speed"); ok {
		azureExpress.Speed = speed.(string)
	}
	return azureExpress
}
