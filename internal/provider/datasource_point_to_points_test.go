//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePointToPointsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	pointToPointResult := testutil.DHclPointToPoints()

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: pointToPointResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.ptp_uuid"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.ptp_circuit_id"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.description"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.speed"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.media"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.state"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.billing"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.time_created"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.time_updated"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.deleted"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.service_class"),
				),
			},
		},
	})

}
