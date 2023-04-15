package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePortDeviceInfoComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourcePortDeviceInfoResult := testutil.DHclDataSourcePortDeviceInfo()

	resource.ParallelTest(t, resource.TestCase{

		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourcePortDeviceInfoResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, "port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, "adjacent_router"),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, "device_name"),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, "device_make"),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, "admin_status"),
				),
			},
		},
	})

}
