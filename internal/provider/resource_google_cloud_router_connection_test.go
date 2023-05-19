//go:build resource || cloud_router || all

package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudRouterConnectionGoogleRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"GOOGLE_CREDENTIALS", "TF_VAR_gcp_project_id"})

	crConnGoogleResult := testutil.RHclCloudRouterConnectionGoogle()
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
					resource.TestCheckResourceAttr(crConnGoogleResult.ResourceName, "pop", crConnGoogleResult.Pop),
					resource.TestCheckResourceAttr(crConnGoogleResult.ResourceName, "speed", crConnGoogleResult.Speed),
					resource.TestCheckResourceAttrSet(crConnGoogleResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(crConnGoogleResult.ResourceName, "id"),
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
				ImportStateVerifyIgnore: []string{"google_pairing_key", "google_vlan_attachment_name"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := fmt.Sprintf("%s:%s", cloudRouterCircuitId, cloudRouterConnectionCircuitId)
					return id, nil
				},
			},
		},
	})
}
