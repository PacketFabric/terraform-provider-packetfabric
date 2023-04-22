//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceOutboundCrossConnectComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourceOutboundCrossConnectResult := testutil.DHclDataSourceOutboundCrossConnect()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceOutboundCrossConnectResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.port"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.site"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.document_uuid"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.outbound_cross_connect_id"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.obcc_status"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.description"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.user_description"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.destination_name"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.destination_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.panel"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.module"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.position"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.data_center_cross_connect_id"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.z_loc_cfa"),
					resource.TestCheckResourceAttrSet(datasourceOutboundCrossConnectResult.ResourceName, "outbound_cross_connects.0.time_created"),
				)},
		},
	})

}
