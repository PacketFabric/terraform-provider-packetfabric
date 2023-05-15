//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePortsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	hclPortResult := testutil.DHclDataSourcePorts()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclPortResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.autoneg"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.port_circuit_id"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.state"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.status"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.speed"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.media"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.zone"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.region"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.market"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.market_description"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.pop"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.site"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.site_code"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.mtu"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.description"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.vc_mode"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.is_lag"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.is_lag_member"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.is_cloud"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.is_ptp"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.is_nni"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.member_count"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.account_uuid"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.subscription_term"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.customer_name"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.customer_uuid"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.time_created"),
					resource.TestCheckResourceAttrSet(hclPortResult.ResourceName, "interfaces.0.time_updated"),
				),
			},
		},
	})
}
