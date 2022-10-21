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
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the target cloud router.",
			},
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not this connection is intended for NAT later.",
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
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description for this connection.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "Circuit ID of the target cloud router.",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired location for the new Oracle Hosted Connection.",
			},
			"zone": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The desired zone of the new connection.",
			},
			"published_quote_line_uuid": {
				Type:         schema.TypeString,
				Optional:     true,
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
		_, err := c.ReadAwsConnection(cid.(string), cloudConnCID.(string))
		if err != nil {
			diags = diag.FromErr(err)
		}
	}
	return diags
}

func resourceOracleCloudRouteConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		desc := d.Get("description")
		descUpdate := packetfabric.DescriptionUpdate{
			Description: desc.(string),
		}
		if _, err := c.UpdateAwsConnection(cid.(string), cloudConnCID.(string), descUpdate); err != nil {
			diags = diag.FromErr(err)
		}
	}
	return diags
}

func resourceOracleCloudRouteConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCloudRouterConnDelete(ctx, d, m)
}

func extractOracleCloudRouterConn(d *schema.ResourceData) packetfabric.OracleCloudRouterConn {
	oracleRouter := packetfabric.OracleCloudRouterConn{}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		oracleRouter.MaybeNat = maybeNat.(bool)
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
