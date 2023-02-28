package provider

import (
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAwsHostedConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	hcl, resourceName := testutil.HclAwsHostedConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CS_VLAN2_KEY,
				testutil.PF_CS_SPEED2_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "vlan", os.Getenv(testutil.PF_CS_VLAN2_KEY)),
					resource.TestCheckResourceAttr(resourceName, "speed", os.Getenv(testutil.PF_CS_SPEED2_KEY)),
				),
			},
		},
	})
}
