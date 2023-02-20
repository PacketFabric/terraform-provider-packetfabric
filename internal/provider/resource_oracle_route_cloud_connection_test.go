package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclOracleRouteConnection(description, region, vcOcid, pop string) (hcl string, resourceName string) {

	hclCloudRouter, crResourceName := hclCloudRouter(
		os.Getenv(testutil.PF_CR_DESCR),
		os.Getenv(testutil.PF_ACCOUNT_ID),
		os.Getenv(testutil.PF_CR_CAPACITY_KEY),
	)
	hclName := testutil.GenerateUniqueResourceName()

	resourceName = "packetfabric_cloud_router_connection_oracle." + hclName
	oracleRouteConnectionHcl := fmt.Sprintf(testutil.RResourceCloudRouterconnectionOracle, hclName, description, crResourceName, region, vcOcid, pop)
	hcl = fmt.Sprintf("%s\n%s", hclCloudRouter, oracleRouteConnectionHcl)
	return
}

func TestAccHclOracleRouteConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	description := os.Getenv(testutil.PF_CR_DESCR)
	vcOcid := os.Getenv(testutil.PF_CRC_ORACLE_VC_OCID_KEY)
	pop := os.Getenv(testutil.PF_CRC_POP5_KEY)
	region := os.Getenv(testutil.PF_CS_ORACLE_REGION_KEY)

	hcl, resourceName := hclOracleRouteConnection(
		description,
		region,
		vcOcid,
		pop,
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CR_DESCR,
				testutil.PF_CRC_ORACLE_VC_OCID_KEY,
				testutil.PF_CRC_POP5_KEY,
				testutil.PF_CS_ORACLE_REGION_KEY,
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
					resource.TestCheckResourceAttr(resourceName, "pop", pop),
				),
			},
		},
	})

}
