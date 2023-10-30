//go:build resource || cloud_router || all

package provider

import (
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudProviderCredentialAwsRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{PfeAwsAccessKeyId, PfeAwsSecretAccessKey})

	awsProviderCredentialsResult := testutil.RHclCloudProviderCredentialAws()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: awsProviderCredentialsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awsProviderCredentialsResult.ResourceName, PfDescription, awsProviderCredentialsResult.Desc),
					resource.TestCheckResourceAttr(awsProviderCredentialsResult.ResourceName, PfAwsAccessKey, os.Getenv(PfeAwsAccessKeyId)),
					resource.TestCheckResourceAttr(awsProviderCredentialsResult.ResourceName, PfAwsSecretKey, os.Getenv(PfeAwsSecretAccessKey)),
				),
			},
		},
	})
}
