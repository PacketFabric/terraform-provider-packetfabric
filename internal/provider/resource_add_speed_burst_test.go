package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAddSpeedBurstRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	backboneVirtualCircuitSpeedBurstResult := testutil.RHclBackboneVirtualCircuitSpeedBurst()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: backboneVirtualCircuitSpeedBurstResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(backboneVirtualCircuitSpeedBurstResult.ResourceName, "vc_circuit_id", backboneVirtualCircuitSpeedBurstResult.VcCircuitId),
					resource.TestCheckResourceAttr(backboneVirtualCircuitSpeedBurstResult.ResourceName, "speed", backboneVirtualCircuitSpeedBurstResult.Speed),
				),
			},
			{
				ResourceName:      backboneVirtualCircuitSpeedBurstResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}
