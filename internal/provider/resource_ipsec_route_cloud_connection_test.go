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

func TestAccCloudRouterConnectionIpsecRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{})
	cloudRouterConnectionIpsecResult := testutil.RHclCloudRouterConnectionIpsec()
	var cloudRouterCircuitId, cloudRouterConnectionCircuitId string
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterConnectionIpsecResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "description", cloudRouterConnectionIpsecResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "pop", cloudRouterConnectionIpsecResult.Pop),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "gateway_address", cloudRouterConnectionIpsecResult.GatewayAddress),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "speed", cloudRouterConnectionIpsecResult.Speed),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "ike_version", strconv.Itoa(cloudRouterConnectionIpsecResult.IkeVersion)),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase1_authentication_method", cloudRouterConnectionIpsecResult.Phase1AuthenticationMethod),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase1_group", cloudRouterConnectionIpsecResult.Phase1Group),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase1_encryption_algo", cloudRouterConnectionIpsecResult.Phase1EncryptionAlgo),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase1_authentication_algo", cloudRouterConnectionIpsecResult.Phase1AuthenticationAlgo),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase1_lifetime", strconv.Itoa(cloudRouterConnectionIpsecResult.Phase1Lifetime)),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase2_pfs_group", cloudRouterConnectionIpsecResult.Phase2PfsGroup),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase2_encryption_algo", cloudRouterConnectionIpsecResult.Phase2EncryptionAlgo),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase2_authentication_algo", cloudRouterConnectionIpsecResult.Phase2AuthenticationAlgo),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "phase2_lifetime", strconv.Itoa(cloudRouterConnectionIpsecResult.Phase2Lifetime)),
					resource.TestCheckResourceAttr(cloudRouterConnectionIpsecResult.ResourceName, "shared_key", cloudRouterConnectionIpsecResult.SharedKey),
				),
			},
			{
				Config: cloudRouterConnectionIpsecResult.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[cloudRouterConnectionIpsecResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", cloudRouterConnectionIpsecResult.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["id"]
					return nil
				},
			},
			{
				ResourceName:      cloudRouterConnectionIpsecResult.ResourceName,
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
