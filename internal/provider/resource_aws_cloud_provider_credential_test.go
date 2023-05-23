//go:build resource || cloud_router || all

package provider

import (
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudProviderCredentialAwsRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"})

	awsProviderCredentialsResult := testutil.RHclCloudProviderCredentialAws()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: awsProviderCredentialsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awsProviderCredentialsResult.ResourceName, "description", awsProviderCredentialsResult.Desc),
					resource.TestCheckResourceAttr(awsProviderCredentialsResult.ResourceName, "aws_access_key", os.Getenv("AWS_ACCESS_KEY_ID")),
					resource.TestCheckResourceAttr(awsProviderCredentialsResult.ResourceName, "aws_secret_key", os.Getenv("AWS_SECRET_ACCESS_KEY")),
				),
			},
		},
	})
}
