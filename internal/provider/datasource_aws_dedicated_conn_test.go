package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceAwsDedicatedConnComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	awsDedicatedConnectionResult := testutil.DHclAwsDedicatedConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: awsDedicatedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.uuid"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.customer_uuid"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.user_uuid"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.service_provider"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.port_type"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.deleted"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.time_updated"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.time_created"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.cloud_circuit_id"),
					resource.TestCheckResourceAttrSet(awsDedicatedConnectionResult.ResourceName, "dedicated_connections.0.account_uuid"),
				),
			},
		},
	})

}
