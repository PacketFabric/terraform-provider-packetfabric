//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePointToPointsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	dPointToPointResult := testutil.DHclPointToPoints()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dPointToPointResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.port_circuit_id"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.pop"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.site"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.site_name"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.customer_site_code"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.customer_site_name"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.speed"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.media"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.zone"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.description"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.vlan"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.untagged"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.provisioning_status"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.admin_status"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.operational_status"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.customer_uuid"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.customer_name"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.region"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.is_cloud"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.is_ptp"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.time_created"),
					resource.TestCheckResourceAttrSet(dPointToPointResult.ResourceName, "point_to_points.0.interfaces.0.time_updated"),
				),
			},
		},
	})

}
