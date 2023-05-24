//go:build resource || cloud_router || all

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudRouterConnectionAwsRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_AWS_ACCOUNT_ID"})

	crConnAwsResult := testutil.RHclCloudRouterConnectionAws()
	var cloudRouterCircuitId, cloudRouterConnectionCircuitId string

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: crConnAwsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(crConnAwsResult.ResourceName, "aws_account_id", os.Getenv("PF_AWS_ACCOUNT_ID")),
					resource.TestCheckResourceAttr(crConnAwsResult.ResourceName, "account_uuid", crConnAwsResult.AccountUuid),
					resource.TestCheckResourceAttr(crConnAwsResult.ResourceName, "speed", crConnAwsResult.Speed),
					resource.TestCheckResourceAttr(crConnAwsResult.ResourceName, "pop", crConnAwsResult.Pop),
					resource.TestCheckResourceAttr(crConnAwsResult.ResourceName, "zone", crConnAwsResult.Zone),
					resource.TestCheckResourceAttrSet(crConnAwsResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(crConnAwsResult.ResourceName, "id"),
				),
			},
			{
				Config: crConnAwsResult.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[crConnAwsResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", crConnAwsResult.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["id"]
					return nil
				},
			},
			{
				ResourceName:      crConnAwsResult.ResourceName,
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
