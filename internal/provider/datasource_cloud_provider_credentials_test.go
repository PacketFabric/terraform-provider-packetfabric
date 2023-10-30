//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCloudProviderCredentialsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{PfeAwsAccessKeyId, PfeAwsSecretAccessKey})

	cloudProviderCredentialResult := testutil.DHclCloudProviderCredentials()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudProviderCredentialResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(cloudProviderCredentialResult.ResourceName, "cloud_credentials.0.cloud_provider"),
					resource.TestCheckResourceAttrSet(cloudProviderCredentialResult.ResourceName, "cloud_credentials.0.cloud_provider_credential_uuid"),
					resource.TestCheckResourceAttrSet(cloudProviderCredentialResult.ResourceName, "cloud_credentials.0.description"),
					resource.TestCheckResourceAttrSet(cloudProviderCredentialResult.ResourceName, "cloud_credentials.0.is_unused"),
					resource.TestCheckResourceAttrSet(cloudProviderCredentialResult.ResourceName, "cloud_credentials.0.time_created"),
					resource.TestCheckResourceAttrSet(cloudProviderCredentialResult.ResourceName, "cloud_credentials.0.time_updated"),
				),
			},
		},
	})
}
