package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclPointToPoint(description, accountUUID, speed, media, subscriptionTerm, pop1, zone1, pop2, zone2 string) (hcl string, resourceName string) {
	hcl = fmt.Sprintf(`
	resource "packetfabric_point_to_point" "p2p" {
		description       = "%s"
		account_uuid      = "%s"
		speed             = "%s"
		media             = "%s"
		subscription_term = %s
		endpoints {
			pop     = "%s"
			zone    = "%s"
			autoneg = false
		}
		endpoints {
			pop     = "%s"
			zone    = "%s"
			autoneg = false
		}
	}`, description, accountUUID, speed, media, subscriptionTerm, pop1, zone1, pop2, zone2)
	resourceName = "packetfabric_point_to_point.p2p"
	return
}

func TestAccPointToPoint(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	description := testutil.GenerateUniqueName(testPrefix)
	hcl, resourceName := hclPointToPoint(
		description,
		testutil.GetAccountUUID(),
		"1Gbps",
		"LX",
		"1",
		"AMS1",
		"A",
		"AMS2",
		"A",
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "account_uuid", testutil.GetAccountUUID()),
					resource.TestCheckResourceAttr(resourceName, "speed", "1Gbps"),
					resource.TestCheckResourceAttr(resourceName, "media", "LX"),
					resource.TestCheckResourceAttr(resourceName, "subscription_term", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.pop", "AMS1"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.0.zone", "A"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.pop", "AMS2"),
					resource.TestCheckResourceAttr(resourceName, "endpoints.1.zone", "A"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
