package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLocationsMarkets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLocationsMarketsRead,
		Schema: map[string]*schema.Schema{
			PfLocationsMarkets: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfName:    schemaStringComputed(PfNameDescription6),
						PfCode:    schemaStringComputed(PfCodeDescription2),
						PfCountry: schemaStringComputed(PfCountryDescription3),
					},
				},
			},
		},
	}
}

func dataSourceLocationsMarketsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	markets, err := c.GetLocationsMarkets()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfLocationsMarkets, flattenLocationsMarkets(&markets))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenLocationsMarkets(markets *[]packetfabric.LocationMarket) []interface{} {
	fields := stringsToMap(PfName, PfCode, PfCountry)
	flattens := make([]interface{}, len(*markets))
	for i, market := range *markets {
		flattens[i] = structToMap(market, fields)
	}
	return flattens
}
