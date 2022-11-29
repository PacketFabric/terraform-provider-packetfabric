package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclCloudRouterConnectionAws(description, accountUUID, awsAccountID, pop, zone, speed string) (hcl string, resourceName string) {
	hclCloudRouter, crResourceName := hclCloudRouter(
		testutil.GenerateUniqueName(testPrefix),
		accountUUID,
		"US",
		"100Mbps",
		"4556",
	)
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_cloud_router_connection_aws." + hclName
	crcHcl := fmt.Sprintf(`
	resource "packetfabric_cloud_router_connection_aws" "%s" {
		circuit_id     = %s.id
		account_uuid   = "%s"
		aws_account_id = "%s"
		maybe_nat      = false
		description    = "%s"
		pop            = "%s"
		zone           = "%s"
		is_public      = false
		speed          = "%s"
	}`, hclName, crResourceName, accountUUID, awsAccountID, description, pop, zone, speed)

	hcl = fmt.Sprintf("%s\n%s", hclCloudRouter, crcHcl)
	return
}

func TestAccCloudRouterConnectionAws(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	description := testutil.GenerateUniqueName(testPrefix)
	// TODO(rjouhann): add function to get pop / zone automatically when packetfabric_locations_cloud available (#200)
	hcl, resourceName := hclCloudRouterConnectionAws(
		description,
		testutil.GetAccountUUID(),
		os.Getenv("PF_AWS_ACCOUNT_ID"),
		"SFO6",
		"B",
		"50Mbps",
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				"PF_AWS_ACCOUNT_ID",
			})
		},
		ProviderFactories: testAccProviderFactories,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "account_uuid", testutil.GetAccountUUID()),
					resource.TestCheckResourceAttr(resourceName, "pop", "SFO6"),
					resource.TestCheckResourceAttr(resourceName, "zone", "B"),
					resource.TestCheckResourceAttr(resourceName, "speed", "50Mbps"),
				),
			},
		},
	})
}
