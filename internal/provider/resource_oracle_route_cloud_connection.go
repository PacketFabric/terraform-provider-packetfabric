package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceOracleCloudRouteConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOracleCloudRouteConnCreate,
		ReadContext:   resourceOracleCloudRouteConnRead,
		UpdateContext: resourceOracleCloudRouteConnUpdate,
		DeleteContext: resourceOracleCloudRouteConnDelete,
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
			"circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Set this to true if you intend to use NAT on this connection. ",
			},
			"maybe_dnat": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Set this to true if you intend to use DNAT on this connection. ",
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
				Description:  "The POP in which you want to provision the connection.",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired availability zone of the new connection.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "UUID of the published quote line with which this connection should be associated.",
			},
		},
	}
}

func resourceOracleCloudRouteConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		oracleRouter := extractOracleCloudRouterConn(d)
		resp, err := c.CreateOracleCloudRouerConnection(oracleRouter, cid.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		createOkCh := make(chan bool)
		defer close(createOkCh)
		fn := func() (*packetfabric.ServiceState, error) {
			return c.GetCloudConnectionStatus(cid.(string), resp.CloudCircuitID)
		}
		go c.CheckServiceStatus(createOkCh, fn)
		if !<-createOkCh {
			return diag.FromErr(err)
		}
		if resp != nil {
			d.SetId(resp.CloudCircuitID)
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

func resourceOracleCloudRouteConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		resp, err := c.ReadCloudRouterConnection(cid.(string), cloudConnCID.(string))
		if err != nil {
			diags = diag.FromErr(err)
		}

		_ = d.Set("account_uuid", resp.AccountUUID)
		_ = d.Set("circuit_id", resp.CloudRouterCircuitID)
		_ = d.Set("maybe_nat", resp.NatCapable)
		_ = d.Set("maybe_dnat", resp.DNatCapable)
		_ = d.Set("vc_ocid", resp.CloudSettings.VcOcid)
		_ = d.Set("region", resp.CloudSettings.OracleRegion)
		_ = d.Set("description", resp.Description)
		_ = d.Set("pop", resp.Pop)
		_ = d.Set("zone", resp.Zone)

		_unsetFields := []string{"published_quote_line_uuid"}
		showWarningForUnsetFields(_unsetFields, &diags)
	}
	return diags
}

func resourceOracleCloudRouteConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnUpdate(ctx, d, m)
}

func resourceOracleCloudRouteConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractOracleCloudRouterConn(d *schema.ResourceData) packetfabric.OracleCloudRouterConn {
	oracleRouter := packetfabric.OracleCloudRouterConn{}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		oracleRouter.MaybeNat = maybeNat.(bool)
	}
	if maybeDNat, ok := d.GetOk("maybe_dnat"); ok {
		oracleRouter.MaybeDNat = maybeDNat.(bool)
	}
	if vcOcid, ok := d.GetOk("vc_ocid"); ok {
		oracleRouter.VcOcid = vcOcid.(string)
	}
	if region, ok := d.GetOk("region"); ok {
		oracleRouter.Region = region.(string)
	}
	if desc, ok := d.GetOk("description"); ok {
		oracleRouter.Description = desc.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		oracleRouter.AccountUUID = accountUUID.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		oracleRouter.Pop = pop.(string)
	}
	if zone, ok := d.GetOk("zone"); ok {
		oracleRouter.Zone = zone.(string)
	}
	if publishedQuote, ok := d.GetOk("published_quote_line_uuid"); ok {
		oracleRouter.PublishedQuoteLineUUID = publishedQuote.(string)
	}
	return oracleRouter
}
