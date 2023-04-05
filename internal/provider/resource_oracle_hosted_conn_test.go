package provider

import (
	"log"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestOracleHostedConnectRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	csOracleHostedConnectionResult := testutil.RHclCsOracleHostedConnection()
	log.Fatal(csOracleHostedConnectionResult.Hcl)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: csOracleHostedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "vc_ocid", csOracleHostedConnectionResult.VcOcid),
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "region", csOracleHostedConnectionResult.Region),
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "description", csOracleHostedConnectionResult.Desc),
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "zone", csOracleHostedConnectionResult.Zone),
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "pop", csOracleHostedConnectionResult.Pop),
					resource.TestCheckResourceAttr(csOracleHostedConnectionResult.ResourceName, "vlan", strconv.Itoa(csOracleHostedConnectionResult.Vlan)),
				),
			},
		},
	})

}
