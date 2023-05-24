//go:build datasource || location || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceLocationsMarketsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	locationsMarketsResult := testutil.DHclLocationsMarkets()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: locationsMarketsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(locationsMarketsResult.ResourceName, "locations_markets.0.name"),
					resource.TestCheckResourceAttrSet(locationsMarketsResult.ResourceName, "locations_markets.0.code"),
					resource.TestCheckResourceAttrSet(locationsMarketsResult.ResourceName, "locations_markets.0.country"),
				),
			},
		},
	})
}
