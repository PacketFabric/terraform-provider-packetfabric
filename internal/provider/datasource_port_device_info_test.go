//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePortDeviceInfoComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourcePortDeviceInfoResult := testutil.DHclPortDeviceInfo()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourcePortDeviceInfoResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfPortCircuitId),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfDeviceName),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfDeviceMake),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfAdminStatus),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfOperStatus),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfAutoNegotiation),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfIfaceName),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfSpeed),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfPolltime),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTimeFlapped),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficRxBps),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficRxBytes),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficRxIpv6Bytes),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficRxIpv6Packets),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficRxPackets),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficRxPps),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficTxBps),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficTxBytes),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficTxIpv6Bytes),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficTxIpv6Packets),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficTxPackets),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfTrafficTxPps),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfWiringMedia),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfWiringModule),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfWiringPanel),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfWiringPosition),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfWiringReach),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfWiringType),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfLagSpeed),
					resource.TestCheckResourceAttrSet(datasourcePortDeviceInfoResult.ResourceName, PfDeviceCanLag),

				),
			},
		},
	})
}
