package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCsAzureHostedMktConnRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	oracleMktConn := testutil.RHclCsAzureHostedMktConn()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CRC_AZURE_SERVICE_KEY,
				testutil.PF_ROUTING_ID_KEY,
				testutil.PF_MARKET_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: oracleMktConn.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(oracleMktConn.ResourceName, "description", oracleMktConn.Description),
					resource.TestCheckResourceAttr(oracleMktConn.ResourceName, "routing_id", oracleMktConn.RoutingID),
					resource.TestCheckResourceAttr(oracleMktConn.ResourceName, "market", oracleMktConn.Market),
					resource.TestCheckResourceAttr(oracleMktConn.ResourceName, "speed", oracleMktConn.Speed),
				),
			},
		},
	})
}
