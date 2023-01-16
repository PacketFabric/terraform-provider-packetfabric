package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceQuickConnectRejectRequest() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		CreateContext: resourceRejectRequestCreate,
		ReadContext:   resourceRejectRequestRead,
		UpdateContext: resourceRejectRequestUpdate,
		DeleteContext: resourceRejectRequestDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Circuit ID of the Quick Connect import.",
			},
			"rejection_reason": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The reason that you are rejecting the request.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceRejectRequestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	rejectionReason := ""
	if reason, ok := d.GetOk("rejection_reason"); ok {
		rejectionReason = reason.(string)
	}
	if circuitID, ok := d.GetOk("circuit_id"); ok {
		if _, err := c.RejectCloudRouterService(circuitID.(string), rejectionReason); err != nil {
			return diag.FromErr(err)
		} else {
			d.SetId(circuitID.(string))
		}
	}
	return diags
}

func resourceRejectRequestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceRejectRequestUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceRejectRequestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return diag.Diagnostics{}
}
