package provider

import (
	"log"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestOutboundCrossConnect(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	outboundCrossConnectResult := testutil.RHclOutboundCrossConnect()
	log.Fatal(outboundCrossConnectResult.Hcl)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_DOCUMENT_UUID1_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: outboundCrossConnectResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(outboundCrossConnectResult.ResourceName, "description", outboundCrossConnectResult.Desc),
					resource.TestCheckResourceAttr(outboundCrossConnectResult.ResourceName, "document_uuid", outboundCrossConnectResult.DocumentUuid),
					resource.TestCheckResourceAttr(outboundCrossConnectResult.ResourceName, "site", outboundCrossConnectResult.Site),
				),
			},
		},
	})
}
