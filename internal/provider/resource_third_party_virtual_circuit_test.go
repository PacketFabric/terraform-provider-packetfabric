package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccThirdPartyVCRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	vcMkt := testutil.RHclThirdPartyVirtualCircuitMkt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ROUTING_ID_KEY,
				testutil.PF_MARKET_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: vcMkt.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(vcMkt.ResourceName, "description", vcMkt.Description),
					resource.TestCheckResourceAttr(vcMkt.ResourceName, "routing_id", vcMkt.RoutingID),
					resource.TestCheckResourceAttr(vcMkt.ResourceName, "market", vcMkt.Market),
					resource.TestCheckResourceAttr(vcMkt.ResourceName, "asn", strconv.Itoa(vcMkt.Asn)),
				),
			},
		},
	})
}
