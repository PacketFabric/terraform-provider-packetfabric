//go:build resource || cloud_router || all

package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudRouterConnectionIbmRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_IBM_ACCOUNT_ID", "IC_API_KEY", "IAAS_CLASSIC_USERNAME", "IAAS_CLASSIC_API_KEY", "TF_VAR_ibm_resource_group"})
	cloudRouterConnectionIbmResult := testutil.RHclCloudRouterConnectionIbm()
	var cloudRouterCircuitId, cloudRouterConnectionCircuitId string
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterConnectionIbmResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionIbmResult.ResourceName, "description", cloudRouterConnectionIbmResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionIbmResult.ResourceName, "account_uuid", cloudRouterConnectionIbmResult.AccountUuid),
					resource.TestCheckResourceAttr(cloudRouterConnectionIbmResult.ResourceName, "pop", cloudRouterConnectionIbmResult.Pop),
					resource.TestCheckResourceAttr(cloudRouterConnectionIbmResult.ResourceName, "speed", cloudRouterConnectionIbmResult.Speed),
					resource.TestCheckResourceAttr(cloudRouterConnectionIbmResult.ResourceName, "ibm_bgp_asn", strconv.Itoa(cloudRouterConnectionIbmResult.IbmBgpAsn)),
				),
			},
			{
				Config: cloudRouterConnectionIbmResult.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[cloudRouterConnectionIbmResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", cloudRouterConnectionIbmResult.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["id"]
					return nil
				},
			},
			{
				ResourceName:      cloudRouterConnectionIbmResult.ResourceName,
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
