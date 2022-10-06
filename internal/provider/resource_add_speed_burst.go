package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAddSpeedBurst() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAddSpeedBurstCreate,
		ReadContext:   resourceAddSpeedBurstRead,
		UpdateContext: resourceAddSpeedBurstUpdate,
		DeleteContext: resourceAddSpeedBurstDelete,
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
			"vc_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Target's Circuit ID.",
			},
			"speed": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Speed in Mbps of the burst.",
			},
		},
	}
}

func resourceAddSpeedBurstCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCID, ok := d.GetOk("vc_circuit_id"); ok {
		if speed, ok := d.GetOk("speed"); ok {
			if _, err := c.AddSpeedBurstToCircuit(vcCID.(string), speed.(string)); err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return diags
}

func resourceAddSpeedBurstRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCID, ok := d.GetOk("vc_circuit_id"); ok {
		if _, err := c.GetBackboneByVcCID(vcCID.(string)); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAddSpeedBurstUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	settings := extractServiceSettings(d)
	if vcCID, ok := d.GetOk("vc_circuit_id"); ok {
		if _, err := c.UpdateServiceSettings(vcCID.(string), settings); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAddSpeedBurstDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if _, err := c.DeleteSpeedBurst(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
