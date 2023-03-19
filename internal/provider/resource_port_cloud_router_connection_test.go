package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPortCloudRouterConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	cloudRouterConnectionPortResult := testutil.RHclCloudRouterConnectionPort()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             cloudRouterConnectionPortResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionPortResult.ResourceName, "description", cloudRouterConnectionPortResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionPortResult.ResourceName, "speed", cloudRouterConnectionPortResult.Speed),
				),
			},
		},
	})

}
