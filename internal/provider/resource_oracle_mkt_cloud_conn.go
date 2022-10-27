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
		ReadContext:   resourceOracleMktCloudConnRead,
		UpdateContext: resourceOracleMktCloudConnUpdate,
		DeleteContext: resourceOracleMktCloudConnDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
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
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The routing ID of the customer to whom this VC will be connected.",
			},
			"market": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The market that the VC will be requested in.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Oracle Marketplace Cloud connection.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID of the contact that will be billed.",
			},
			"vc_ocid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "OCID for the Oracle VC to use in this hosted connection.",
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Oracle region for this connection.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired location for the new Oracle Hosted Connection.",
			},
			"service_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
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

func resourceOracleMktCloudConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceOracleMktCloudConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceOracleMktCloudConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	vcRequestUUID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid VC Request UUID to delete")
	}
	err := c.DeleteHostedMktConnection(vcRequestUUID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
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
