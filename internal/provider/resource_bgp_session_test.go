package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPBgpSessionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	bgpSessionResult := testutil.RHclBgpSession()

	var cloudRouterCircuitId, cloudRouterConnectionCircuitId, bgpSessionUuid string

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CRC_AWS_ACCOUNT_ID_KEY,
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: bgpSessionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "remote_address", bgpSessionResult.RemoteAddress),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "l3_address", bgpSessionResult.L3Address),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "remote_asn", strconv.Itoa(bgpSessionResult.Asn)),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "prefixes.0.prefix", bgpSessionResult.Prefix1),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "prefixes.0.type", bgpSessionResult.Type1),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "prefixes.1.prefix", bgpSessionResult.Prefix2),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "prefixes.1.type", bgpSessionResult.Type2),
				),
			},
			{
				Config: bgpSessionResult.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[bgpSessionResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", bgpSessionResult.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["connection_id"]
					bgpSessionUuid = rs.Primary.Attributes["id"]
					return nil
				},
			},
			{
				ResourceName:      bgpSessionResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := fmt.Sprintf("%s:%s:%s", cloudRouterCircuitId, cloudRouterConnectionCircuitId, bgpSessionUuid)
					return id, nil
				},
			},
		},
	})
}
