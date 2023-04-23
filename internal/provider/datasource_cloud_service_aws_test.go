package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceHostedAwsConnComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceHostedAwsConnResult := testutil.DHclDatasourceHostedAwsConn()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
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
					resource.TestCheckResourceAttrSet(datasourceHostedAwsConnResult.ResourceName, "cloud_provider_region"),
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
