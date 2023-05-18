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
	crConnGoogleResult := testutil.RHclCloudRouterConnectionAzure()
	var cloudRouterCircuitId, cloudRouterConnectionCircuitId string
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: crConnGoogleResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(crConnGoogleResult.ResourceName, "description", crConnGoogleResult.Desc),
					resource.TestCheckResourceAttr(crConnGoogleResult.ResourceName, "account_uuid", crConnGoogleResult.AccountUuid),
					resource.TestCheckResourceAttr(crConnGoogleResult.ResourceName, "speed", crConnGoogleResult.Speed),
				),
			},
			{
				Config: crConnGoogleResult.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[crConnGoogleResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", crConnGoogleResult.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["id"]
					return nil
				},
			},
			{
				ResourceName:            crConnGoogleResult.ResourceName,
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
