//go:build resource || cloud_router || all

package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudRouterConnectionOracleRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"TF_VAR_tenancy_ocid", "TF_VAR_user_ocid", "TF_VAR_fingerprint", "TF_VAR_private_key", "TF_VAR_parent_compartment_id", "TF_VAR_pf_cs_oracle_drg_ocid"})
	cloudRouterConnectionOracleResult := testutil.RHclCloudRouterConnectionOracle()
	var cloudRouterCircuitId, cloudRouterConnectionCircuitId string
	resource.ParallelTest(t, resource.TestCase{

		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterConnectionOracleResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionOracleResult.ResourceName, "description", cloudRouterConnectionOracleResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionOracleResult.ResourceName, "pop", cloudRouterConnectionOracleResult.Pop),
					resource.TestCheckResourceAttr(cloudRouterConnectionOracleResult.ResourceName, "zone", cloudRouterConnectionOracleResult.Zone),
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
				ImportStateVerifyIgnore: []string{"vc_ocid", "region"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := fmt.Sprintf("%s:%s", cloudRouterCircuitId, cloudRouterConnectionCircuitId)
					return id, nil
				},
			},
		},
	})

}
