package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceHostedGoogleConnComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceHostedGoogleConnResult := testutil.DHclDatasourceHostedGoogleConn()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceHostedGoogleConnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "user_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "state"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "service_provider"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "service_class"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "port_type"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "speed"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "description"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "cloud_provider_pop"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "cloud_provider_region"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "pop"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "site"),
					resource.TestCheckResourceAttrSet(datasourceHostedGoogleConnResult.ResourceName, "is_awaiting_onramp"),
				),
			},
		},
	})

}
