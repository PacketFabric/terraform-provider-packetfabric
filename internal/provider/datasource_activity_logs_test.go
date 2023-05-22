//go:build datasource || other || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceActivityLogsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourceActivityLogsResult := testutil.DHclActivityLogs()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceActivityLogsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceActivityLogsResult.ResourceName, "activity_logs.0.log_uuid"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogsResult.ResourceName, "activity_logs.0.user"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogsResult.ResourceName, "activity_logs.0.level"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogsResult.ResourceName, "activity_logs.0.category"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogsResult.ResourceName, "activity_logs.0.event"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogsResult.ResourceName, "activity_logs.0.message"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogsResult.ResourceName, "activity_logs.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogsResult.ResourceName, "activity_logs.0.log_level_name"),
				),
			},
		},
	})

}
