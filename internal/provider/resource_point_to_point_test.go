package provider

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclPointToPoint(description, speed, media, subscriptionTerm, pop1, zone1, pop2, zone2 string, autoneg1, autonegg2 bool) (hcl string, resourceName string) {
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_point_to_point." + hclName
	hcl = fmt.Sprintf(testutil.RResourcePointToPoint, hclName, description, speed, media, subscriptionTerm, pop1, zone1, autoneg1, pop2, zone2, autonegg2)
	return
}

func TestAccPointToPointRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	description := os.Getenv(testutil.PF_PTP_DESCR)
	speed := os.Getenv(testutil.PF_PTP_SPEED_KEY)
	media := os.Getenv(testutil.PF_PTP_MEDIA_KEY)
	subscriptionTerm := os.Getenv(testutil.PF_PTP_SUBTERM_KEY)
	pop1 := os.Getenv(testutil.PF_PTP_POP1_KEY)
	zone1 := os.Getenv(testutil.PF_PTP_ZONE1_KEY)
	pop2 := os.Getenv(testutil.PF_PTP_POP2_KEY)
	zone2 := os.Getenv(testutil.PF_PTP_ZONE2_KEY)
	autoneg1 := testutil.GetEnvBool(testutil.PF_PTP_AUTONEG_KEY)
	autoneg2 := testutil.GetEnvBool(testutil.PF_PTP_AUTONEG_KEY)

	hcl, resourceName := hclPointToPoint(
		description,
		speed,
		media,
		subscriptionTerm,
		pop1,
		zone1,
		pop2,
		zone2,
		autoneg1,
		autoneg2,
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_PTP_DESCR,
				testutil.PF_PTP_SPEED_KEY,
				testutil.PF_PTP_SUBTERM_KEY,
				testutil.PF_PTP_MEDIA_KEY,
				testutil.PF_PTP_POP1_KEY,
				testutil.PF_PTP_ZONE1_KEY,
				testutil.PF_PTP_POP2_KEY,
				testutil.PF_PTP_ZONE2_KEY,
				testutil.PF_PTP_AUTONEG_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "speed", speed),
					resource.TestCheckResourceAttr(resourceName, "media", media),
					resource.TestCheckResourceAttr(resourceName, "subscription_term", subscriptionTerm),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.pop", pop1),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.zone", zone1),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.autoneg", strconv.FormatBool(autoneg1)),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.pop", pop2),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.zone", zone2),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.autoneg", strconv.FormatBool(autoneg2)),
				),
			},
		},
	})
}
