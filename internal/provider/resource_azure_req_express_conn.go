package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAzureReqExpressConn() *schema.Resource {
	return &schema.Resource{
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
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid S-VLAN range is from 4-4094, inclusive.",
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
	resp, err := c.CreateAzureExpressRoute(azureExpress)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.CustomerUUID)
	return diags
}

func extractAzureExpressConn(d *schema.ResourceData) packetfabric.AzureExpressRoute {
	azureExpress := packetfabric.AzureExpressRoute{
		AzureServiceKey: d.Get("azure_service_key").(string),
		AccountUUID:     d.Get("account_uuid").(string),
		Description:     d.Get("description").(string),
		Port:            d.Get("port").(string),
		VlanPrivate:     d.Get("vlan_private").(int),
		VlanMicrosoft:   d.Get("vlan_microsoft").(int),
		SrcSvlan:        d.Get("src_svlan").(int),
		Speed:           d.Get("speed").(string),
	}
	return azureExpress
}
