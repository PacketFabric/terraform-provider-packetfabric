//go:build datasource || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceLocationsPortAvailabilityComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)
	locationsPortAvailabilityResult := testutil.DHclDataSourceLocationsPortAvailability()
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: locationsPortAvailabilityResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(locationsPortAvailabilityResult.ResourceName, "pop"),
					resource.TestCheckResourceAttrSet(locationsPortAvailabilityResult.ResourceName, "ports_available.0.zone"),
					resource.TestCheckResourceAttrSet(locationsPortAvailabilityResult.ResourceName, "ports_available.0.speed"),
					resource.TestCheckResourceAttrSet(locationsPortAvailabilityResult.ResourceName, "ports_available.0.media"),
					resource.TestCheckResourceAttrSet(locationsPortAvailabilityResult.ResourceName, "ports_available.0.count"),
					resource.TestCheckResourceAttrSet(locationsPortAvailabilityResult.ResourceName, "ports_available.0.partial"),
					resource.TestCheckResourceAttrSet(locationsPortAvailabilityResult.ResourceName, "ports_available.0.enni"),
				),
			},
		},
	})
}
