package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclPortLOA(pop, zone, speed, destinationEmail, loaCustomerName string) (hcl string, resourceName string) {
	portHcl, portResourceName := hclPort(
		testutil.GenerateUniqueName(testPrefix),
		testutil.GetAccountUUID(),
		speed,
		"LX",
		"1",
		pop,
		zone,
		true,
		false,
	)

	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_port_loa." + hclName
	loaHcl := fmt.Sprintf(`
	resource "packetfabric_port_loa" "%s" {
		port_circuit_id   = %s.id
		destination_email = "%s"
		loa_customer_name = "%s"
	}`, hclName, portResourceName, destinationEmail, loaCustomerName)

	hcl = fmt.Sprintf("%s\n%s", portHcl, loaHcl)
	return
}

func TestAccPortLOA(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	pop, zone, err := testutil.GetPopAndZoneWithAvailablePort("1Gbps", "LX")
	if err != nil {
		t.Fatalf("Unable to find pop and zone with available port: %s", err)
	}
	hcl, resourceName := hclPortLOA(pop, zone, "1Gbps", "test@packetfabric.com", "Test Customer")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "destination_email", "test@packetfabric.com"),
					resource.TestCheckResourceAttr(resourceName, "loa_customer_name", "Test Customer"),
				),
			},
		},
	})
}
