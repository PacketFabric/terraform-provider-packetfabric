package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAzureExpressRouteConn() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAzureExpressRouteConnCreate,
		ReadContext:   resourceAzureExpressRouteConnRead,
		UpdateContext: resourceAzureExpressRouteConnUpdate,
		DeleteContext: resourceAzureExpressRouteConnDelete,
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
				Description:  "Circuit ID of the target cloud router. This starts with \"PF-L3-CUST-\".",
			},
			"maybe_nat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set this to true if you intend to use NAT on this connection. ",
			},
			"maybe_dnat": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Set this to true if you intend to use DNAT on this connection. ",
			},
			"azure_service_key": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The Service Key provided by Microsoft Azure when you created your ExpressRoute circuit.",
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
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
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(speedOptions(), true),
				Description:  "The speed of the new connection.\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\"]",
			},
			"is_public": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether PacketFabric should allocate a public IP address for this connection. Set this to true if you intend to set up peering with Microsoft public services (such as Microsoft 365). ",
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

func resourceAzureExpressRouteConnCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	expressRoute := extractAzureExpressRouteConn(d)
	if cid, ok := d.GetOk("circuit_id"); ok {
		resp, err := c.CreateAzureExpressRouteConn(expressRoute, cid.(string))
		if err != nil {
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

func resourceAzureExpressRouteConnRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceAzureExpressRouteConnUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		desc := d.Get("description")
		descUpdate := packetfabric.DescriptionUpdate{
			Description: desc.(string),
		}
		if _, err := c.UpdateCloudRouterConnection(cid.(string), cloudConnCID.(string), descUpdate); err != nil {
			diags = diag.FromErr(err)
		}
	}
	return diags
}

func resourceAzureExpressRouteConnDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cid, ok := d.GetOk("circuit_id"); ok {
		cloudConnCID := d.Get("id")
		if _, err := c.DeleteCloudRouterConnection(cid.(string), cloudConnCID.(string)); err != nil {
			diags = diag.FromErr(err)
		} else {
			d.SetId("")
		}
	}
	return diags
}

func extractAzureExpressRouteConn(d *schema.ResourceData) packetfabric.AzureExpressRouteConn {
	expressRoute := packetfabric.AzureExpressRouteConn{}
	if maybeNat, ok := d.GetOk("maybe_nat"); ok {
		expressRoute.MaybeNat = maybeNat.(bool)
	}
	if azureServiceKey, ok := d.GetOk("azure_service_key"); ok {
		expressRoute.AzureServiceKey = azureServiceKey.(string)
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		expressRoute.AccountUUID = accountUUID.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		expressRoute.Description = description.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		expressRoute.Speed = speed.(string)
	}
	if isPublic, ok := d.GetOk("is_public"); ok {
		expressRoute.IsPublic = isPublic.(bool)
	}
	if publishedQuoteLine, ok := d.GetOk("published_quote_line_uuid"); ok {
		expressRoute.PublishedQuoteLineUUID = publishedQuoteLine.(string)
	}
	return expressRoute
}
