package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceHostedIbmConnComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceHostedIbmConnResult := testutil.DHclDatasourceHostedIbmConn()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceHostedIbmConnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceHostedIbmConnResult.ResourceName, "cloud_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceHostedIbmConnResult.ResourceName, "uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedIbmConnResult.ResourceName, "account_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedIbmConnResult.ResourceName, "customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedIbmConnResult.ResourceName, "user_uuid"),
					resource.TestCheckResourceAttrSet(datasourceHostedIbmConnResult.ResourceName, "state"),
					resource.TestCheckResourceAttrSet(datasourceHostedIbmConnResult.ResourceName, "deleted"),
					resource.TestCheckResourceAttrSet(datasourceHostedIbmConnResult.ResourceName, "cloud_provider.0.pop"),
				),
			},
		},
	})

}
