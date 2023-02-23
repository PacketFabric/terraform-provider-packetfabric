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
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
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
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The POP in which the connection should be provisioned (the cloud on-ramp).",
			},
			"port": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The circuit ID of the PacketFabric port you wish to connect to Oracle. This starts with \"PF-AP-\".",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired availability zone of the connection.",
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
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with this connection should be associated.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
		_ = d.Set("vlan", resp.Settings.VlanIDPf)
		_ = d.Set("speed", resp.Speed)
		_ = d.Set("pop", resp.Pop)
		_ = d.Set("vc_ocid", resp.Settings.VcOcid)
		_ = d.Set("region", resp.Settings.OracleRegion)
	}
	// _unsetFields := []string{"port", "zone", "src_svlan", "published_quote_line_uuid"}
	// showWarningForUnsetFields(_unsetFields, &diags)
	return diags
}

func resourceOracleHostedConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	return resourceServicesHostedUpdate(ctx, d, m, c.UpdateServiceHostedConn)
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
