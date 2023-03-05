package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPointToPointRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	pointToPointResult := testutil.RHclPointToPoint()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             pointToPointResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "description", pointToPointResult.Desc),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "speed", pointToPointResult.Speed),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "media", pointToPointResult.Media),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "subscription_term", strconv.Itoa(pointToPointResult.SubscriptionTerm)),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "endpoints.0.pop", pointToPointResult.Pop1),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "endpoints.0.zone", pointToPointResult.Zone1),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "endpoints.0.autoneg", strconv.FormatBool(pointToPointResult.Autoneg1)),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "endpoints.1.pop", pointToPointResult.Pop2),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "endpoints.1.zone", pointToPointResult.Zone2),
					resource.TestCheckResourceAttr(pointToPointResult.ResourceName, "endpoints.1.autoneg", strconv.FormatBool(pointToPointResult.Autoneg2)),
				),
			},
		},
	})
}
