package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceLocationsRegionsComputedRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	locationsRegionsResult := testutil.DHclDataSourceLocationsRegions()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
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
