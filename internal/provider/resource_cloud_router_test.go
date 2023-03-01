package provider

import (
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudRouterRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	hcl, resourceName := testutil.HclCloudRouter()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_CR_CAPACITY_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_uuid", os.Getenv(testutil.PF_ACCOUNT_ID_KEY)),
					resource.TestCheckResourceAttr(resourceName, "capacity", os.Getenv(testutil.PF_CR_CAPACITY_KEY)),
				),
			},
		},
	})
}
