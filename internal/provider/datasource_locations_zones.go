package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLocationsZones() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLocationsZonesRead,
		Schema: map[string]*schema.Schema{
			PfPop:            schemaStringRequiredNotEmpty(PfPopDescription9),
			PfLocationsZones: schemaStringListComputedNotEmpty(PfLocationsZonesDescription),
		},
	}
}

func dataSourceLocationsZonesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	pop, ok := d.GetOk(PfPop)
	if !ok {
		return diag.Errorf(MessageMissingPop)
	}
	var diags diag.Diagnostics
	if zones, err := c.GetLocationsZones(pop.(string)); err != nil {
		return diag.FromErr(err)
	} else {
		_ = d.Set(PfLocationsZones, zones)
	}
	d.SetId(uuid.New().String())
	return diags
}
