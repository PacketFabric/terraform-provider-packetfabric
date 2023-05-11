//go:build resource || core || all || smoke

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclVcBackboneVlanRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)
	backboneVirtualCircuitResult := testutil.RHclBackboneVirtualCircuitVlan()
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: backboneVirtualCircuitResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "description", backboneVirtualCircuitResult.Desc),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "epl", strconv.FormatBool(backboneVirtualCircuitResult.Epl)),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "interface_a.0.vlan", strconv.Itoa(backboneVirtualCircuitResult.InterfaceBackboneA.Vlan)),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "interface_z.0.vlan", strconv.Itoa(backboneVirtualCircuitResult.InterfaceBackboneZ.Vlan)),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "bandwidth.0.longhaul_type", backboneVirtualCircuitResult.LonghaulType),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "bandwidth.0.speed", backboneVirtualCircuitResult.Speed),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "bandwidth.0.subscription_term", strconv.Itoa(backboneVirtualCircuitResult.SubscriptionTerm)),
				),
			},
			{
				ResourceName:      backboneVirtualCircuitResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
