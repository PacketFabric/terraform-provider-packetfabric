package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePortDisable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePortDisableCreate,
		ReadContext:   resourcePortDisableRead,
		UpdateContext: resourcePortDisableUpdate,
		DeleteContext: resourcePortDisableDelete,
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
			"port_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port circuit ID.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePortDisableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	portCID := d.Get("port_circuit_id")
	d.SetId(portCID.(string))
	return _callPortStatusFunc(d, c.DisablePort)
}

func resourcePortDisableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourcePortDisableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourcePortDisableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	d.SetId("")
	return _callPortStatusFunc(d, c.EnablePort)
}

func _callPortStatusFunc(d *schema.ResourceData, fn func(portCID string) (*packetfabric.PortMessageResp, error)) diag.Diagnostics {
	var diags diag.Diagnostics
	if portCID, ok := d.GetOk("port_circuit_id"); ok {
		if _, err := fn(portCID.(string)); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}
