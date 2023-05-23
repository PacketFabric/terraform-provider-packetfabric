//go:build resource || cloud_router || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudProviderCredentialgoogleRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"GOOGLE_CREDENTIALS"})

	googleProviderCredentialsResult := testutil.RHclCloudProviderCredentialGoogle()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: googleProviderCredentialsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(googleProviderCredentialsResult.ResourceName, "description", googleProviderCredentialsResult.Desc),
				),
			},
		},
	})
}
