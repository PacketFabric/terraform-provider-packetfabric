//go:build datasource || location || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceLocationsZonesComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourceZonesResult := testutil.DHclZones()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceZonesResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceZonesResult.ResourceName, "pop"),
					resource.TestCheckResourceAttrSet(datasourceZonesResult.ResourceName, "locations_zones.0"),
				),
			},
		},
	})

}
