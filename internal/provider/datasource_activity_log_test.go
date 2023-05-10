package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceActivityLogComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceActivityLogResult := testutil.DHclDataSourceActivityLog()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceActivityLogResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceActivityLogResult.ResourceName, "activity_logs.0.log_uuid"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogResult.ResourceName, "activity_logs.0.user"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogResult.ResourceName, "activity_logs.0.level"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogResult.ResourceName, "activity_logs.0.category"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogResult.ResourceName, "activity_logs.0.event"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogResult.ResourceName, "activity_logs.0.message"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogResult.ResourceName, "activity_logs.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceActivityLogResult.ResourceName, "activity_logs.0.log_level_name"),
				),
			},
		},
	})

}
