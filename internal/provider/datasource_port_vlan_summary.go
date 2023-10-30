package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePortVlanSummary() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortVlanSummaryRead,
		Schema: map[string]*schema.Schema{
			PfPortCircuitId:       schemaStringRequiredNotEmpty(PfPortCircuitIdDescription2),
			PfLowestAvailableVlan: schemaIntComputed(PfLowestAvailableVlanDescription),
			PfMaxVlan:             schemaIntComputed(PfMaxVlanDescription),
		},
	}
}

func dataSourcePortVlanSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	portCID, ok := d.GetOk(PfPortCircuitId)
	if !ok {
		return diag.Errorf(MessageMissingPortCiruitId)
	}
	summary, err := c.GetPortVlanSummary(portCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	_ = d.Set(PfLowestAvailableVlan, summary.LowestAvailableVlan)
	_ = d.Set(PfMaxVlan, summary.MaxVlan)
	d.SetId(portCID.(string) + "-data")
	return diags
}
