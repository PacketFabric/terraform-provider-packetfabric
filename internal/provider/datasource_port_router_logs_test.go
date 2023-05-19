//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePortRouterLogsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{})

	datasourcePortRouterLogsResult := testutil.DHclRouterLogs()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourcePortRouterLogsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "time_from"),
					resource.TestCheckResourceAttrSet(datasourcePortRouterLogsResult.ResourceName, "time_to"),
				),
			},
		},
	})

}
