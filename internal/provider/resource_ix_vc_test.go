package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclIXVirtualCircuitMktRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	ixVirtualCircuitMktResult := testutil.RHclIXVirtualCircuitMarketplace()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ROUTING_ID_IX_KEY,
				testutil.PF_MARKET_IX_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: ixVirtualCircuitMktResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(ixVirtualCircuitMktResult.ResourceName, "routing_id", ixVirtualCircuitMktResult.RoutingID),
					resource.TestCheckResourceAttr(ixVirtualCircuitMktResult.ResourceName, "market", ixVirtualCircuitMktResult.Market),
					resource.TestCheckResourceAttr(ixVirtualCircuitMktResult.ResourceName, "asn", strconv.Itoa(ixVirtualCircuitMktResult.Asn)),
				),
			},
			{
				ResourceName:      ixVirtualCircuitMktResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}
