package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclOracleCSHostedMarketplaceRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	oracleCSHostedMarketplaceResult := testutil.RHclCsOracleHostedMarketplaceConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CS_ORACLE_MKT_CONN_REGION_KEY,
				testutil.PF_CS_ORACLE_MKT_CONN_OCID_KEY,
				testutil.PF_CS_ORACLE_MKT_CONN_ROUTING_ID_KEY,
				testutil.PF_CS_ORACLE_MKT_CONN_MARKET_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: oracleCSHostedMarketplaceResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(oracleCSHostedMarketplaceResult.ResourceName, "vc_ocid", oracleCSHostedMarketplaceResult.VcOCID),
					resource.TestCheckResourceAttr(oracleCSHostedMarketplaceResult.ResourceName, "region", oracleCSHostedMarketplaceResult.Region),
					resource.TestCheckResourceAttr(oracleCSHostedMarketplaceResult.ResourceName, "description", oracleCSHostedMarketplaceResult.Description),
					resource.TestCheckResourceAttr(oracleCSHostedMarketplaceResult.ResourceName, "pop", oracleCSHostedMarketplaceResult.Pop),
				),
			},
		},
	})

}
