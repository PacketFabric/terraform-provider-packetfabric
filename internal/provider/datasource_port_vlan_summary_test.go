//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPortVlansComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourcePortVlansResult := testutil.DHclDataSourcePortVlans()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourcePortVlansResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourcePortVlansResult.ResourceName, "port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourcePortVlansResult.ResourceName, "lowest_available_vlan"),
					resource.TestCheckResourceAttrSet(datasourcePortVlansResult.ResourceName, "max_vlan"),
				),
			},
		},
	})

}
