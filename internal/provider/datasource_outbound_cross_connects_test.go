//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceOutboundCrossConnectsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourceOutboundCrossConnectsResult := testutil.DHclDataSourceOutboundCrossConnects()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceOutboundCrossConnectsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.port"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.site"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.document_uuid"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.outbound_cross_connect_id"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.obcc_status"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.description"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.user_description"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.time_updated"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectsResult.ResourceName, "outbound_cross_connects.0.time_created"),
				)},
		},
	})

}
