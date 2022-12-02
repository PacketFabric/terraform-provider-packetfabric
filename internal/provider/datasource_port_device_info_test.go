package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclDataSourcePortDeviceInfo(pop, zone, speed string) (hcl string, resourceName string) {
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
	resourceName = "data.packetfabric_port_device_info." + hclName
	portDeviceInfoHcl := fmt.Sprintf(`
	data "packetfabric_port_device_info" "%s" {
		port_circuit_id = %s.id
	}`, hclName, portResourceName)
	hcl = fmt.Sprintf("%s\n%s", portHcl, portDeviceInfoHcl)
	return
}

func TestAccDataSourcePortDeviceInfo(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)
	pop, zone, err := testutil.GetPopAndZoneWithAvailablePort("1Gbps", "LX")
	if err != nil {
		t.Fatalf("Unable to find pop and zone with available port: %s", err)
	}
	hcl, resourceName := hclDataSourcePortDeviceInfo(pop, zone, "1Gbps")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "device_id"),
					resource.TestCheckResourceAttrSet(resourceName, "device_name"),
					resource.TestCheckResourceAttrSet(resourceName, "device_make"),
					resource.TestCheckResourceAttrSet(resourceName, "admin_status"),
					resource.TestCheckResourceAttrSet(resourceName, "oper_status"),
					resource.TestCheckResourceAttrSet(resourceName, "auto_negotiation"),
					resource.TestCheckResourceAttrSet(resourceName, "iface_name"),
					resource.TestCheckResourceAttrSet(resourceName, "speed"),
					resource.TestCheckResourceAttrSet(resourceName, "site_id"),
				),
			},
		},
	})

}
