//go:build resource || hosted_cloud || all || smoke

package provider

import (
	"os"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAwsHostedConnectionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_AWS_ACCOUNT_ID"})
	awsHostedConnectionResult := testutil.RHclCsAwsHostedConnection()
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: awsHostedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "aws_account_id", os.Getenv("PF_AWS_ACCOUNT_ID")),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "account_uuid", awsHostedConnectionResult.AccountUuid),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "speed", awsHostedConnectionResult.Speed),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "pop", awsHostedConnectionResult.Pop),
					resource.TestCheckResourceAttr(awsHostedConnectionResult.ResourceName, "zone", awsHostedConnectionResult.Zone),
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
