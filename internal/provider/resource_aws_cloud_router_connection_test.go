package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudRouterConnectionAwsRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	crConnAwsResult := testutil.RHclCloudRouterConnectioAws()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CRC_AWS_ACCOUNT_ID_KEY,
				testutil.PF_CRC_SPEED_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             crConnAwsResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(crConnAwsResult.ResourceName, "aws_account_id", crConnAwsResult.AwsAccountID),
					resource.TestCheckResourceAttr(crConnAwsResult.ResourceName, "account_uuid", crConnAwsResult.AccountUuid),
					resource.TestCheckResourceAttr(crConnAwsResult.ResourceName, "speed", crConnAwsResult.Speed),
				),
			},
			{
				ResourceName:      crConnAwsResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
