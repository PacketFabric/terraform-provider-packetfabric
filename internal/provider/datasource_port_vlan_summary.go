package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourcePortVlanSummary() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortVlanSummaryRead,
		Schema: map[string]*schema.Schema{
			"port_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port circuit ID.",
			},
			"lowest_available_vlan": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The lowest available vlan.",
			},
			"max_vlan": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The max vlan.",
			},
		},
	}
}

func dataSourcePortVlanSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	portCID, ok := d.GetOk("port_circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid port circuit ID")
	}
	summary, err := c.GetPortVlanSummary(portCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	_ = d.Set("lowest_available_vlan", summary.LowestAvailableVlan)
	_ = d.Set("max_vlan", summary.MaxVlan)
	d.SetId(portCID.(string) + "-data")
	return diags
}
