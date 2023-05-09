package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclVcBackboneRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	backboneVirtualCircuitResult := testutil.RHclBackboneVirtualCircuit()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: backboneVirtualCircuitResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "description", backboneVirtualCircuitResult.Desc),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "epl", strconv.FormatBool(backboneVirtualCircuitResult.Epl)),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "interface_a.0.untagged", strconv.FormatBool(backboneVirtualCircuitResult.InterfaceBackboneA.Untagged)),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "interface_a.0.vlan", strconv.Itoa(backboneVirtualCircuitResult.InterfaceBackboneA.Vlan)),
					resource.TestCheckResourceAttr(backboneVirtualCircuitResult.ResourceName, "interface_z.0.untagged", strconv.FormatBool(backboneVirtualCircuitResult.InterfaceBackboneZ.Untagged)),
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
