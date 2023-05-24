//go:build resource || cloud_router || marketplace || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudRouterQuickConnectRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{
		"PF_AWS_ACCOUNT_ID",
		"PF_QUICK_CONNECT_SERVICE_UUID",
	})

	cloudRouterQuickConnect := testutil.RHclCloudRouterQuickConnect()
	var cloudRouterCircuitId, cloudRouterConnectionCircuitId, importCircuitID string

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterQuickConnect.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(cloudRouterQuickConnect.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(cloudRouterQuickConnect.ResourceName, "route_set_circuit_id"),
					resource.TestCheckResourceAttrSet(cloudRouterQuickConnect.ResourceName, "state"),
				),
			},
			{
				Config: cloudRouterQuickConnect.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[cloudRouterQuickConnect.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", cloudRouterQuickConnect.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["connection_id"]
					importCircuitID = rs.Primary.Attributes["id"]
					return nil
				},
			},
			{
				ResourceName:      cloudRouterQuickConnect.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := fmt.Sprintf("%s:%s:%s", cloudRouterCircuitId, cloudRouterConnectionCircuitId, importCircuitID)
					return id, nil
				},
			},
		},
	})
}
