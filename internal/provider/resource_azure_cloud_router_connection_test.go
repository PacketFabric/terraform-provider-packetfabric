//go:build resource || cloud_router || all || smoke

package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudRouterConnectionAzureRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"ARM_SUBSCRIPTION_ID", "ARM_CLIENT_ID", "ARM_CLIENT_SECRET", "ARM_TENANT_ID"})

	crConnAzureResult := testutil.RHclCloudRouterConnectionAzure()
	var cloudRouterCircuitId, cloudRouterConnectionCircuitId string

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: crConnAzureResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(crConnAzureResult.ResourceName, "description", crConnAzureResult.Desc),
					resource.TestCheckResourceAttr(crConnAzureResult.ResourceName, "account_uuid", crConnAzureResult.AccountUuid),
					resource.TestCheckResourceAttr(crConnAzureResult.ResourceName, "speed", crConnAzureResult.Speed),
					resource.TestCheckResourceAttrSet(crConnAzureResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(crConnAzureResult.ResourceName, "id"),
				),
			},
			{
				Config: crConnAzureResult.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[crConnAzureResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", crConnAzureResult.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["id"]
					return nil
				},
			},
			{
				Config: crConnAzureResult.Hcl,
				Check: resource.TestCheckResourceAttr(crConnAzureResult.ResourceName,"is_public","true"),
			},		
			{
				ResourceName:            crConnAzureResult.ResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"azure_service_key"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := fmt.Sprintf("%s:%s", cloudRouterCircuitId, cloudRouterConnectionCircuitId)
					return id, nil
				},
			},
		},
	})
}

func TestAccCloudRouterConnectionAzureNoPublicIPs(t *testing.T) {
    testutil.PreCheck(t, []string{"ARM_SUBSCRIPTION_ID", "ARM_CLIENT_ID", "ARM_CLIENT_SECRET", "ARM_TENANT_ID"})

    crConnAzureResult := testutil.RHclCloudRouterConnectionAzureNoPublicIPs()
    var cloudRouterCircuitId, cloudRouterConnectionCircuitId string

    resource.ParallelTest(t, resource.TestCase{
        Providers:         testAccProviders,
        ExternalProviders: testAccExternalProviders,
        Steps: []resource.TestStep{
            {
                Config: crConnAzureResult.Hcl,
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr(crConnAzureResult.ResourceName, "description", crConnAzureResult.Desc),
                    resource.TestCheckResourceAttr(crConnAzureResult.ResourceName, "account_uuid", crConnAzureResult.AccountUuid),
                    resource.TestCheckResourceAttr(crConnAzureResult.ResourceName, "speed", crConnAzureResult.Speed),
                    resource.TestCheckResourceAttrSet(crConnAzureResult.ResourceName, "circuit_id"),
                    resource.TestCheckResourceAttrSet(crConnAzureResult.ResourceName, "id"),
                ),
            },
            {
				Config: crConnAzureResult.Hcl,
				Check: resource.TestCheckResourceAttr(crConnAzureResult.ResourceName,"is_public","false"),
			},
        },
    })
}