package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclOracleMktConn(description, vcOcid, region, market, pop string) (hcl string, resourceName string) {

	hclCloudRouter, crResourceName := hclCloudRouter(
		os.Getenv(testutil.PF_CR_DESCR),
		os.Getenv(testutil.PF_ACCOUNT_ID),
		os.Getenv(testutil.PF_CR_CAPACITY_KEY),
	)
	hclName := testutil.GenerateUniqueResourceName()

	resourceName = "packetfabric_cs_oracle_hosted_marketplace_connection." + hclName
	oracleMktConnHcl := fmt.Sprintf(testutil.RResourceCSOracleHostedMarketplaceConnection, hclName, description, vcOcid, region, crResourceName, market, pop)
	hcl = fmt.Sprintf("%s\n%s", hclCloudRouter, oracleMktConnHcl)
	return
}

func TestAccOracleMktConnRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	description := os.Getenv(testutil.DESCR_HC_C)
	vcOcid := os.Getenv(testutil.PF_CS_ORACLE_VC_OCID_KEY)
	pop := os.Getenv(testutil.PF_CS_POP6_KEY)
	market := os.Getenv(testutil.PF_MARKET_KEY)
	region := os.Getenv(testutil.PF_CS_ORACLE_REGION_KEY)

	hcl, resourceName := hclOracleMktConn(
		description,
		vcOcid,
		region,
		market,
		pop,
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CS_ORACLE_VC_OCID_KEY,
				testutil.PF_CS_ORACLE_REGION_KEY,
				testutil.DESCR_HC_C,
				testutil.PF_CS_POP6_KEY,
				testutil.PF_MARKET_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "vc_ocid", vcOcid),
					resource.TestCheckResourceAttr(resourceName, "region", region),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "market", market),
					resource.TestCheckResourceAttr(resourceName, "pop", pop),
				),
			},
		},
	})

}
