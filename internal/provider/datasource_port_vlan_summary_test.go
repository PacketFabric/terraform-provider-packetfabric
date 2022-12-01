package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclDataSourcePortVlans(pop, zone, speed string) (hcl string, resourceName string) {
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
	resourceName = "data.packetfabric_port_vlans." + hclName
	portDeviceInfoHcl := fmt.Sprintf(`
	data "packetfabric_port_vlans" "%s" {
		port_circuit_id = %s.id
	}`, hclName, portResourceName)
	hcl = fmt.Sprintf("%s\n%s", portHcl, portDeviceInfoHcl)
	return
}

func TestAccPortVlans(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)
	pop, zone, err := testutil.GetPopAndZoneWithAvailablePort("1Gbps", "LX")
	if err != nil {
		t.Fatalf("Unable to find pop and zone with available port: %s", err)
	}
	hcl, resourceName := hclDataSourcePortVlans(pop, zone, "1Gbps")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "lowest_available_vlan"),
					resource.TestCheckResourceAttrSet(resourceName, "max_vlan"),
				),
			},
		},
	})

}
