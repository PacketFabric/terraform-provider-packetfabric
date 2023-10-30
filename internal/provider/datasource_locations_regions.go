package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceLocationsRegions() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceLocationsRegionsRead,
		Schema: map[string]*schema.Schema{
			PfLocationsRegions: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfName: schemaStringComputed(PfNameDescription7),
						PfCode: schemaStringComputed(PfCodeDescription3),
					},
				},
			},
		},
	}
}

func datasourceLocationsRegionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	regions, err := c.GetLocationRegions()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfLocationsRegions, flattenLocationsRegions(&regions))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenLocationsRegions(regions *[]packetfabric.LocationRegion) []interface{} {
	fields := stringsToMap(PfName, PfCode)
	flattens := make([]interface{}, len(*regions))
	for i, reg := range *regions {
		flattens[i] = structToMap(&reg, fields)
	}
	return flattens
}
