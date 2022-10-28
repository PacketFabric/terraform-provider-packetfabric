package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePointToPoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePointToPointCreate,
		ReadContext:   resourcePointToPointRead,
		UpdateContext: resourcePointToPointUpdate,
		DeleteContext: resourcePointToPointDelete,
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
			"ptp_circuit_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The PTP Circuit ID.",
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Point to Point description.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"1Gbps", "10Gbps", "40Gbps", "100Gbps"}, true),
				Description:  "The point to point service speed.",
			},
			"media": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(pointToPointMediaOptions(), true),
			},
			"endpoints": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pop": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The POP name.",
						},
						"zone": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The availability zone.",
						},
						"customer_site_code": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Unique site code of the customer's equipment.",
						},
						"autoneg": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "The autoneg state.",
						},
						"loa": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringIsBase64,
							Description:  "A base64 encoded string of a PDF of a LOA.",
						},
					},
				},
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsUUID,
				Description:  "The account UUID.",
			},
			"subscription_term": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
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

func resourcePointToPointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ptpService := extractPtpService(d)
	resp, err := c.CreatePointToPointService(ptpService)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.PtpUUID)

	return diags
}

func resourcePointToPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cID, ok := d.GetOk("ptp_circuit_id"); ok {
		if _, err := c.GetPointToPointInfo(cID.(string)); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePointToPointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if desc, ok := d.GetOk("description"); ok {
		if _, err := c.UpdatePointToPoint(d.Id(), desc.(string)); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePointToPointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if cID := d.Id(); cID != "" {
		if err := c.DeletePointToPointService(cID); err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId("")
	return diags
}

func extractPtpService(d *schema.ResourceData) packetfabric.PointToPoint {
	ptpService := packetfabric.PointToPoint{}
	if description, ok := d.GetOk("description"); ok {
		ptpService.Description = description.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		ptpService.Speed = speed.(string)
	}
	if media, ok := d.GetOk("media"); ok {
		ptpService.Media = media.(string)
	}
	if endpoints, ok := d.GetOk("endpoints"); ok {
		edps := make([]packetfabric.Endpoints, 0)
		for _, endpoint := range endpoints.(*schema.Set).List() {
			edps = append(edps, packetfabric.Endpoints{
				Pop:              endpoint.(map[string]interface{})["pop"].(string),
				Zone:             endpoint.(map[string]interface{})["zone"].(string),
				CustomerSiteCode: endpoint.(map[string]interface{})["customer_site_code"].(string),
				Autoneg:          endpoint.(map[string]interface{})["autoneg"].(bool),
				Loa:              endpoint.(map[string]interface{})["loa"].(string),
			})
		}
		ptpService.Endpoints = edps
	}
	if accountUUID, ok := d.GetOk("account_uuid"); ok {
		ptpService.AccountUUID = accountUUID.(string)
	}
	if subTerm, ok := d.GetOk("subscription_term"); ok {
		ptpService.SubscriptionTerm = subTerm.(int)
	}
	if quote, ok := d.GetOk("published_quote_line_uuid"); ok {
		ptpService.PublishedQuoteLineUUID = quote.(string)
	}
	return ptpService
}

func pointToPointMediaOptions() []string {
	return []string{"LX", "EX", "ZX", "LR",
		"ER", "ER DWDM", "ZR",
		"ZE DWDM", "LR4", "ER4",
		"CWDM4", "LR4", "ER4 Lite"}
}
