package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGoogleHostedMktConnRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	googleMktConn := testutil.RHclCsGoogleHostedMktConn()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CRC_GOOGLE_PAIRING_KEY,
				testutil.PF_CRC_GOOGLE_VLAN_ATTACHMENT_NAME_KEY,
				testutil.PF_ROUTING_ID_KEY,
				testutil.PF_MARKET_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: googleMktConn.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(googleMktConn.ResourceName, "description", googleMktConn.Description),
					resource.TestCheckResourceAttr(googleMktConn.ResourceName, "routing_id", googleMktConn.RoutingID),
					resource.TestCheckResourceAttr(googleMktConn.ResourceName, "market", googleMktConn.Market),
					resource.TestCheckResourceAttr(googleMktConn.ResourceName, "speed", googleMktConn.Speed),
					resource.TestCheckResourceAttr(googleMktConn.ResourceName, "pop", googleMktConn.Pop),
				),
			},
		},
	})
}
