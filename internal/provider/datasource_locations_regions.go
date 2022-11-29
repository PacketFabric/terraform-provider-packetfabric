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
			"locations_regions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The location region name.",
						},
						"code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The location region code.",
						},
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
	err = d.Set("locations_regions", flattenLocationsRegions(&regions))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenLocationsRegions(regions *[]packetfabric.LocationRegion) []interface{} {
	flattens := make([]interface{}, len(*regions))
	for i, reg := range *regions {
		flatten := make(map[string]interface{})
		flatten["name"] = reg.Name
		flatten["code"] = reg.Code
		flattens[i] = flatten
	}
	return flattens
}
