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
			"locations_markets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The market name.",
						},
						"code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The market code.",
						},
						"country": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The market country.",
						},
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
	err = d.Set("locations_markets", flattenLocationsMarkets(&markets))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenLocationsMarkets(markets *[]packetfabric.LocationMarket) []interface{} {
	flattens := make([]interface{}, len(*markets))
	for i, market := range *markets {
		flatten := make(map[string]interface{})
		flatten["name"] = market.Name
		flatten["code"] = market.Code
		flatten["country"] = market.Country
		flattens[i] = flatten
	}
	return flattens
}
