package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclPort(description, accountUUID, speed, media, subscriptionTerm, pop, zone string, autoneg, nni bool) (hcl string, resourceName string) {
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_port." + hclName
	hcl = fmt.Sprintf(`
	resource "packetfabric_port" "%s" {
		description       = "%s"
		account_uuid      = "%s"
		speed             = "%s"
		media             = "%s"
		subscription_term = "%s"
		pop               = "%s"
		zone              = "%s"
		autoneg           = %t
		nni               = %t
	}`, hclName, description, accountUUID, speed, media, subscriptionTerm, pop, zone, autoneg, nni)
	return
}

func TestAccPort(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	description := testutil.GenerateUniqueName(testPrefix)
	pop, zone, err := testutil.GetPopAndZoneWithAvailablePort("1Gbps", "LX")
	if err != nil {
		t.Fatalf("Unable to find pop and zone with available port: %s", err)
	}
	hcl, resourceName := hclPort(
		description,
		testutil.GetAccountUUID(),
		"1Gbps",
		"LX",
		"1",
		pop,
		zone,
		true,
		false,
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
					resource.TestCheckResourceAttr(resourceName, "pop", pop),
					resource.TestCheckResourceAttr(resourceName, "zone", zone),
					resource.TestCheckResourceAttr(resourceName, "autoneg", "true"),
					resource.TestCheckResourceAttr(resourceName, "nni", "false"),
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
