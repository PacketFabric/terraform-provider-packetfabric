//go:build resource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAddSpeedBurstRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	backboneVirtualCircuitSpeedBurstResult := testutil.RHclBackboneVirtualCircuitSpeedBurst()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             backboneVirtualCircuitSpeedBurstResult.Hcl,
				ExpectNonEmptyPlan: true, // i.e. 150Mbps -> 50Mbps
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(backboneVirtualCircuitSpeedBurstResult.ResourceName, "speed", backboneVirtualCircuitSpeedBurstResult.Speed),
				),
			},
		},
	})

}
