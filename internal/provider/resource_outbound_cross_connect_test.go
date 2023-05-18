//go:build resource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestOutboundCrossConnect(t *testing.T) {
	testutil.PreCheck(t, []string{})

	outboundCrossConnectResult := testutil.RHclOutboundCrossConnect()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: outboundCrossConnectResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(outboundCrossConnectResult.ResourceName, "description", outboundCrossConnectResult.Desc),
					resource.TestCheckResourceAttrSet(outboundCrossConnectResult.ResourceName, "document_uuid"),
					resource.TestCheckResourceAttr(outboundCrossConnectResult.ResourceName, "site", outboundCrossConnectResult.Site),
				),
			},
		},
	})
}
