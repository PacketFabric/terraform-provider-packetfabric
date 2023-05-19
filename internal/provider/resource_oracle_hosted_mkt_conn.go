package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOracleMktCloudConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOracleMktCloudConnCreate,
		ReadContext:   resourceThirdPartyVirtualCircuitRead,
		DeleteContext: resourceOracleMktCloudConnDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"routing_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The routing ID of the marketplace provider that will be receiving this request.\n\n\tExample: TR-1RI-OQ85",
			},
			"market": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The market code (e.g. \"ATL\" or \"DAL\") in which you would like the marketplace provider to provision their side of the connection.\n\n\tIf the marketplace provider has services published in the marketplace, you can use the PacketFabric portal to see which POPs they are in. Simply remove the number from the POP to get the market code (e.g. if they offer services in \"DAL5\", enter \"DAL\" for the market).",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "A brief description of this connection.",
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
			"vc_ocid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "OCID of the FastConnect virtual circuit that you created from the Oracle side.",
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The region in which you created the FastConnect virtual circuit.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the connection should be provisioned (the cloud on-ramp).",
			},
			"service_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the marketplace service being requested.",
			},
		},
	}
}

func resourceOracleMktCloudConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	oracleMktConn := extractOracleMktConn(d)
	res, err := c.RequestHostedOracleMktConn(oracleMktConn)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(res.VcRequestUUID)
	return diags
}

func resourceOracleMktCloudConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	vcRequestUUID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid VC Request UUID to delete")
	}
	_, err := c.DeleteHostedMktConnection(vcRequestUUID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func extractOracleMktConn(d *schema.ResourceData) packetfabric.CloudServiceOracle {
	oracle := packetfabric.CloudServiceOracle{}
	if routingID, ok := d.GetOk("routing_id"); ok {
		oracle.RoutingID = routingID.(string)
	}
	if market, ok := d.GetOk("market"); ok {
		oracle.Market = market.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		oracle.Description = description.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		oracle.AccountUUID = accountUUID.(string)
	}
	if vcOcid, ok := d.GetOk("vc_ocid"); ok {
		oracle.VcOcid = vcOcid.(string)
	}
	if region, ok := d.GetOk("region"); ok {
		oracle.Region = region.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		oracle.Pop = pop.(string)
	}
	if serviceUUID, ok := d.GetOk("service_uuid"); ok {
		oracle.ServiceUUID = serviceUUID.(string)
	}
	return oracle
}
