//go:build resource || cloud_router || marketplace || all

package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
					resource.TestCheckResourceAttrSet(cloudRouterQuickConnect.ResourceName, "subscription_term"),
					resource.TestCheckResourceAttrSet(cloudRouterQuickConnect.ResourceName, "state"),
					resource.TestCheckResourceAttr(cloudRouterQuickConnect.ResourceName, "return_filters.0.prefix", cloudRouterQuickConnect.ReturnFilterPrefix1),
					resource.TestCheckResourceAttr(cloudRouterQuickConnect.ResourceName, "return_filters.0.match_type", cloudRouterQuickConnect.ReturnFilterType1),
					resource.TestCheckResourceAttr(cloudRouterQuickConnect.ResourceName, "return_filters.1.prefix", cloudRouterQuickConnect.ReturnFilterPrefix2),
					resource.TestCheckResourceAttr(cloudRouterQuickConnect.ResourceName, "return_filters.1.match_type", cloudRouterQuickConnect.ReturnFilterType2),
				),
			},
			{
				Config: cloudRouterQuickConnect.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[cloudRouterQuickConnect.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", cloudRouterQuickConnect.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["cr_circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["connection_circuit_id"]
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
