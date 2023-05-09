package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPortCloudRouterConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	cloudRouterConnectionPortResult := testutil.RHclCloudRouterConnectionPort()

	var cloudRouterCircuitId, cloudRouterConnectionCircuitId string

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterConnectionPortResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionPortResult.ResourceName, "description", cloudRouterConnectionPortResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionPortResult.ResourceName, "speed", cloudRouterConnectionPortResult.Speed),
					resource.TestCheckResourceAttr(cloudRouterConnectionPortResult.ResourceName, "vlan", strconv.Itoa(cloudRouterConnectionPortResult.Vlan)),
				),
			},
			{
				ResourceName:      cloudRouterConnectionPortResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: cloudRouterConnectionPortResult.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[cloudRouterConnectionPortResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", cloudRouterConnectionPortResult.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["id"]
					return nil
				},
			},
			{
				ResourceName:      cloudRouterConnectionPortResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := fmt.Sprintf("%s:%s", cloudRouterCircuitId, cloudRouterConnectionCircuitId)
					return id, nil
				},
			},
		},
	})

}
