package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceHostedAwsConnComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceHostedAzureConnResult := testutil.DHclDatasourceHostedAzureConn()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceHostedAzureConnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "user_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "state"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "service_provider"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "service_class"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "port_type"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "speed"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "description"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "cloud_provider_pop"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "cloud_provider_region"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "pop"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "site"),
					resource.TestCheckResourceAttrSet(datasourceHostedAzureConnResult.ResourceName, "is_awaiting_onramp"),
				),
			},
		},
	})

}
