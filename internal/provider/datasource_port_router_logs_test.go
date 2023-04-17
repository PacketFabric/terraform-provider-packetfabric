package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePortRouterLogsComputedRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	datasourcePortRouterLogsResult := testutil.DHclDataSourcePortRouterLogs()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_DTS_TIME_FROM_KEY,
				testutil.PF_DTS_TIME_TO_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourcePortRouterLogsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "time_from"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "time_to"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "port_router_logs.0.device_name"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "port_router_logs.0.iface_name"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "port_router_logs.0.message"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "port_router_logs.0.severity"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "port_router_logs.0.severity_name"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "port_router_logs.0.timestamp"),
				),
			},
		},
	})

}
