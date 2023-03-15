package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclVcBackboneRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	backboneVirtulalCircuitResult := testutil.RHclBackboneVirtulalCircuit()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: backboneVirtulalCircuitResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "description", backboneVirtulalCircuitResult.Desc),
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "epl", strconv.FormatBool(backboneVirtulalCircuitResult.Epl)),
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "interface_a.0.untagged", strconv.FormatBool(backboneVirtulalCircuitResult.InterfaceBackboneA.Untagged)),
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "interface_a.0.vlan", strconv.Itoa(backboneVirtulalCircuitResult.InterfaceBackboneA.Vlan)),
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "interface_z.0.untagged", strconv.FormatBool(backboneVirtulalCircuitResult.InterfaceBackboneZ.Untagged)),
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "interface_z.0.vlan", strconv.Itoa(backboneVirtulalCircuitResult.InterfaceBackboneZ.Vlan)),
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "bandwidth.0.longhaul_type", backboneVirtulalCircuitResult.LonghaulType),
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "bandwidth.0.speed", backboneVirtulalCircuitResult.Speed),
					resource.TestCheckResourceAttr(backboneVirtulalCircuitResult.ResourceName, "bandwidth.0.subscription_term", strconv.Itoa(backboneVirtulalCircuitResult.SubscriptionTerm)),
				),
			},
		},
	})
}
