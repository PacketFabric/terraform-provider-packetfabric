//go:build resource || cloud_router || all || smoke

package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudRouterBgpSessionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_AWS_ACCOUNT_ID"})

	bgpSessionResultEnabled, bgpSessionResultDisabled := testutil.RHclBgpSessionEnabledAndDisabled()
	var cloudRouterCircuitId, cloudRouterConnectionCircuitId, bgpSessionUuid string

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: bgpSessionResultEnabled.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(bgpSessionResultEnabled.ResourceName, "remote_address", bgpSessionResultEnabled.RemoteAddress),
					resource.TestCheckResourceAttr(bgpSessionResultEnabled.ResourceName, "l3_address", bgpSessionResultEnabled.L3Address),
					resource.TestCheckResourceAttr(bgpSessionResultEnabled.ResourceName, "remote_asn", strconv.Itoa(bgpSessionResultEnabled.Asn)),
					resource.TestCheckResourceAttr(bgpSessionResultEnabled.ResourceName, "prefixes.0.prefix", bgpSessionResultEnabled.Prefix1),
					resource.TestCheckResourceAttr(bgpSessionResultEnabled.ResourceName, "prefixes.0.type", bgpSessionResultEnabled.Type1),
					resource.TestCheckResourceAttr(bgpSessionResultEnabled.ResourceName, "prefixes.1.prefix", bgpSessionResultEnabled.Prefix2),
					resource.TestCheckResourceAttr(bgpSessionResultEnabled.ResourceName, "prefixes.1.type", bgpSessionResultEnabled.Type2),
					resource.TestCheckResourceAttr(bgpSessionResultEnabled.ResourceName, "disabled", "false"),
					resource.TestCheckResourceAttrSet(bgpSessionResultEnabled.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(bgpSessionResultEnabled.ResourceName, "connection_id"),
					resource.TestCheckResourceAttrSet(bgpSessionResultEnabled.ResourceName, "id"),
				),
			},
			{
				Config: bgpSessionResultEnabled.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[bgpSessionResultEnabled.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", bgpSessionResultEnabled.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["connection_id"]
					bgpSessionUuid = rs.Primary.Attributes["id"]
					disabled := rs.Primary.Attributes["disabled"]
					if "true" == disabled {
						t.Errorf("Expected 'disabled' to be false, but it is true")
					}
					return nil
				},
			},
			{
				Config: bgpSessionResultDisabled.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(bgpSessionResultDisabled.ResourceName, "remote_address", bgpSessionResultDisabled.RemoteAddress),
					resource.TestCheckResourceAttr(bgpSessionResultDisabled.ResourceName, "l3_address", bgpSessionResultDisabled.L3Address),
					resource.TestCheckResourceAttr(bgpSessionResultDisabled.ResourceName, "remote_asn", strconv.Itoa(bgpSessionResultDisabled.Asn)),
					resource.TestCheckResourceAttr(bgpSessionResultDisabled.ResourceName, "prefixes.0.prefix", bgpSessionResultDisabled.Prefix1),
					resource.TestCheckResourceAttr(bgpSessionResultDisabled.ResourceName, "prefixes.0.type", bgpSessionResultDisabled.Type1),
					resource.TestCheckResourceAttr(bgpSessionResultDisabled.ResourceName, "prefixes.1.prefix", bgpSessionResultDisabled.Prefix2),
					resource.TestCheckResourceAttr(bgpSessionResultDisabled.ResourceName, "prefixes.1.type", bgpSessionResultDisabled.Type2),
					resource.TestCheckResourceAttr(bgpSessionResultDisabled.ResourceName, "disabled", "true"),
					resource.TestCheckResourceAttrSet(bgpSessionResultDisabled.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(bgpSessionResultDisabled.ResourceName, "connection_id"),
					resource.TestCheckResourceAttrSet(bgpSessionResultDisabled.ResourceName, "id"),
				),
			},
			{
				Config: bgpSessionResultDisabled.Hcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[bgpSessionResultEnabled.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", bgpSessionResultEnabled.ResourceName)
					}
					cloudRouterCircuitId = rs.Primary.Attributes["circuit_id"]
					cloudRouterConnectionCircuitId = rs.Primary.Attributes["connection_id"]
					bgpSessionUuid = rs.Primary.Attributes["id"]
					disabled := rs.Primary.Attributes["disabled"]
					if "false" == disabled {
						t.Errorf("Expected 'disabled' to be true, but it is false")
					}
					return nil
				},
			},
			{
				ResourceName:      bgpSessionResultDisabled.ResourceName,
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
