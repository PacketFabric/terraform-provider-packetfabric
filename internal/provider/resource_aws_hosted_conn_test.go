package provider

import (
	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"strconv"
	"testing"
)

func TestAccAwsHostedConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	awsHostedConnectionResult := testutil.RHclAwsHostedConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CRC_AWS_ACCOUNT_ID_KEY,
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: awsHostedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "aws_account_id", os.Getenv(testutil.PF_CRC_AWS_ACCOUNT_ID_KEY)),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "account_uuid", awsHostedConnectionResult.AccountUuid),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "speed", awsHostedConnectionResult.Speed),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "pop", awsHostedConnectionResult.Pop),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "vlan", strconv.Itoa(awsHostedConnectionResult.Vlan)),
				),
			},
			{
				ResourceName:      awsHostedConnectionResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
