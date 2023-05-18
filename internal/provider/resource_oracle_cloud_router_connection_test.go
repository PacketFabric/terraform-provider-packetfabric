package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHclOracleRouteConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	cloudRouterConnectionOracleResult := testutil.RHclCloudRouterConnectionOracle()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_CS_ORACLE_REGION_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterConnectionOracleResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionOracleResult.ResourceName, "vc_ocid", cloudRouterConnectionOracleResult.VcOcid),
					resource.TestCheckResourceAttr(cloudRouterConnectionOracleResult.ResourceName, "region", cloudRouterConnectionOracleResult.Region),
					resource.TestCheckResourceAttr(cloudRouterConnectionOracleResult.ResourceName, "description", cloudRouterConnectionOracleResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionOracleResult.ResourceName, "pop", cloudRouterConnectionOracleResult.Pop),
				),
			},
			{
				ResourceName:      cloudRouterConnectionOracleResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}
