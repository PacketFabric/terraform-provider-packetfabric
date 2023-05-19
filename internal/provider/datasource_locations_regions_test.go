//go:build datasource || location || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceLocationsRegionsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	locationsRegionsResult := testutil.DHclLocationsRegions()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: locationsRegionsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(locationsRegionsResult.ResourceName, "locations_regions.0.name"),
					resource.TestCheckResourceAttrSet(locationsRegionsResult.ResourceName, "locations_regions.0.code"),
				),
			},
		},
	})

}
