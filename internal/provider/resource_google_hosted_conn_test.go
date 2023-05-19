//go:build resource || hosted_cloud || all || smoke

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGoogleHostedConnectionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"GOOGLE_CREDENTIALS", "TF_VAR_gcp_project_id"})

	googleHostedConnectionResult := testutil.RHclCsGoogleHostedConnection()

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: googleHostedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(googleHostedConnectionResult.ResourceName, "account_uuid", googleHostedConnectionResult.AccountUuid),
					resource.TestCheckResourceAttr(googleHostedConnectionResult.ResourceName, "pop", googleHostedConnectionResult.Pop),
					resource.TestCheckResourceAttr(googleHostedConnectionResult.ResourceName, "speed", googleHostedConnectionResult.Speed),
					resource.TestCheckResourceAttr(googleHostedConnectionResult.ResourceName, "vlan", strconv.Itoa(googleHostedConnectionResult.Vlan)),
				),
			},
			{
				ResourceName:            googleHostedConnectionResult.ResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"google_pairing_key", "google_vlan_attachment_name"},
			},
		},
	})
}
