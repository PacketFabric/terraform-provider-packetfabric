package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMHostedConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	csIBMHostedConnectionResult := testutil.RHclCsIBMHostedConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_IBM_ACCOUNT_ID_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: csIBMHostedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csIBMHostedConnectionResult.ResourceName, "description", csIBMHostedConnectionResult.Desc),
					resource.TestCheckResourceAttr(csIBMHostedConnectionResult.ResourceName, "pop", csIBMHostedConnectionResult.Pop),
					resource.TestCheckResourceAttr(csIBMHostedConnectionResult.ResourceName, "ibm_bgp_asn", strconv.Itoa(csIBMHostedConnectionResult.IbmBgpAsn)),
					resource.TestCheckResourceAttr(csIBMHostedConnectionResult.ResourceName, "vlan", strconv.Itoa(csIBMHostedConnectionResult.Vlan)),
					resource.TestCheckResourceAttr(csIBMHostedConnectionResult.ResourceName, "speed", csIBMHostedConnectionResult.Speed),
				),
			},
			{
				ResourceName:      csIBMHostedConnectionResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
