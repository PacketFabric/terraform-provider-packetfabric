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
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.port_circuit_id"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.pop"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.site"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.site_name"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.customer_site_code"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.customer_site_name"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.speed"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.media"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.zone"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.description"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.vlan"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.untagged"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.provisioning_status"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.admin_status"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.operational_status"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.customer_uuid"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.customer_name"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.region"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.is_cloud"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.is_ptp"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.time_created"),
					resource.TestCheckResourceAttrSet(pointToPointResult.ResourceName, "point_to_points.0.interfaces.0.time_updated"),
				),
			},
		},
	})

}
