package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclCloudRouterConnectionIpsecRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	cloudRouterConnectionIpsecResult := testutil.RHclCloudRouterConnectionIpsec()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
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
		},
	})
}
