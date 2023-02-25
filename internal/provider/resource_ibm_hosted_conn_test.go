package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclIBMHostedConnection(description, accountUUID, ibmAccountID, ibmBgpAsn, pop, zone, vlan, speed string) (hcl string, resourceName string) {
	portHcl, portResourceName := hclPort(
		testutil.GenerateUniqueName(testPrefix),
		accountUUID,
		speed,
		"LX",
		"1",
		pop,
		zone,
		true,
		false,
	)
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_cs_ibm_hosted_connection." + hclName
	ibmHostedConnectionHcl := fmt.Sprintf(`
	resource "packetfabric_cs_ibm_hosted_connection" "%s" {
		description    = "%s"
		account_uuid   = "%s"
		ibm_account_id = "%s"
		ibm_bgp_asn    = "%s"
		pop            = "%s"
		port           = %s.id
		vlan           = "%s"
		speed          = "%s"
	}`, hclName, description, accountUUID, ibmAccountID, ibmBgpAsn, pop, portResourceName, vlan, speed)
	hcl = fmt.Sprintf("%s\n%s", portHcl, ibmHostedConnectionHcl)
	return
}

func TestAccIBMHostedConnection(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)
	t.Skip("Skipped as it causes problems with PacketFabric API.")

	description := testutil.GenerateUniqueName(testPrefix)
	// TODO(medzi): add function to get pop / zone automatically when packetfabric_locations_cloud available (#200)
	hcl, resourceName := hclIBMHostedConnection(
		description,
		testutil.GetAccountUUID(),
		os.Getenv("PF_IBM_ACCOUNT_ID"),
		"12345",
		"SFO1",
		"C",
		"4",
		"1Gbps",
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{"PF_IBM_ACCOUNT_ID"})
		},
		ProviderFactories: testAccProviderFactories,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "account_uuid", testutil.GetAccountUUID()),
					resource.TestCheckResourceAttr(resourceName, "pop", "SFO1"),
					resource.TestCheckResourceAttr(resourceName, "zone", "C"),
					resource.TestCheckResourceAttr(resourceName, "ibm_bgp_asn", "12345"),
					resource.TestCheckResourceAttr(resourceName, "vlan", "4"),
					resource.TestCheckResourceAttr(resourceName, "speed", "1Gbps"),
				),
			},
		},
		{
			ResourceName:      resourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}
