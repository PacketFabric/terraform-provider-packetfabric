//go:build datasource || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceHostedConnComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_AWS_ACCOUNT_ID"})
	datasourceHostedAwsConnResult := testutil.DHclDatasourceHostedAwsConn()
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceHostedAwsConnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "user_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "state"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "service_provider"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "service_class"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "port_type"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "speed"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "description"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "cloud_provider_pop"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "pop"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "site"),
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "is_awaiting_onramp"),
				),
			},
		},
	})

}
