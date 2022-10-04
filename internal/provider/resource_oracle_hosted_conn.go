package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOracleHostedConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOracleHostedConnCreate,
		ReadContext:   resourceOracleHostedConnRead,
		UpdateContext: resourceOracleHostedConnUpdate,
		DeleteContext: resourceOracleHostedConnDelete,
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
			"vc_ocid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "OCID for the Oracle VC to use in this hosted connection.",
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"us-ashburn-1", "us-phoenix-1"}, true),
				Description:  "The description of this connection.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of this connection.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The UUID of the contact that will be billed.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired loction for the new Oracle Hosted Connection.",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Tge port to connect to Oracle.",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired zone of the new connection.",
			},
			"vlan": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid VLAN range is from 4-4094, inclusive.",
			},
			"src_svlan": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(4, 4094),
				Description:  "Valid S-VLAN range is from 4-4094, inclusive.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with this connection should be associated.",
			},
		},
	}
}

func resourceOracleHostedConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	hostedConn := extractOracleHostedConn(d)
	resp, err := c.RequestNewHostedOracleConn(hostedConn)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.CloudCircuitID)
	return diags
}

func resourceOracleHostedConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesRead(ctx, d, m, c.GetCurrentCustomersDedicated)
}

func resourceOracleHostedConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesUpdate(ctx, d, m, c.UpdateServiceConn)
}

func resourceOracleHostedConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	vcRequestUUID, ok := d.GetOk("id")
	if !ok {
		return diag.Errorf("please provide a valid VC Request UUID to delete")
	}
	err := c.DeleteRequestedHostedMktService(vcRequestUUID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func extractOracleHostedConn(d *schema.ResourceData) packetfabric.CloudServiceOracleConn {
	oracleConn := packetfabric.CloudServiceOracleConn{}
	if vcOcid, ok := d.GetOk("vc_ocid"); ok {
		oracleConn.VcOcid = vcOcid.(string)
	}
	if region, ok := d.GetOk("region"); ok {
		oracleConn.Region = region.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		oracleConn.Description = description.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		oracleConn.AccountUUID = accountUUID.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		oracleConn.Pop = pop.(string)
	}
	if port, ok := d.GetOk("port"); ok {
		oracleConn.Port = port.(string)
	}
	if zone, ok := d.GetOk("zone"); ok {
		oracleConn.Zone = zone.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		oracleConn.Vlan = vlan.(int)
	}
	if srcSvlan, ok := d.GetOk("src_svlan"); ok {
		oracleConn.SrcSvlan = srcSvlan.(int)
	}
	if pubQuoteLineUUID, ok := d.GetOk("published_quote_line_uuid"); ok {
		oracleConn.PublishedQuoteLineUUID = pubQuoteLineUUID.(string)
	}
	return oracleConn
}
