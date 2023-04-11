package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: awsHostedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "pop", awsHostedConnectionResult.Pop),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "description", awsHostedConnectionResult.Desc),
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
