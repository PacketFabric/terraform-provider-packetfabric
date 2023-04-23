package provider

import (
	"log"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceHostedAwsConnComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceHostedOracleConnResult := testutil.DHclDatasourceHostedOracleConn()
	log.Fatal(datasourceHostedOracleConnResult.Hcl)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceHostedOracleConnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "user_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "state"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "service_provider"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "service_class"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "port_type"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "speed"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "description"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "cloud_provider_pop"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "cloud_provider_region"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "pop"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "site"),
					resource.TestCheckResourceAttrSet(datasourceHostedOracleConnResult.ResourceName, "is_awaiting_onramp"),
				),
			},
		},
	})

}
