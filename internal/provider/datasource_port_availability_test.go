package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclDataSourceLocationsPortAvailability(pop string) (hcl string, resourceName string) {
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "data.packetfabric_locations_port_availability." + hclName
	hcl = fmt.Sprintf(`
	data "packetfabric_locations_port_availability" "%s" {
		pop = "%s"
	}`, hclName, pop)
	return
}

func TestAccDataSourceLocationsPortAvailability(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	pop, _, err := testutil.GetPopAndZoneWithAvailablePort("1Gbps")
	if err != nil {
		t.Fatalf("Unable to find pop and zone with available port: %s", err)
	}
	hcl, resourceName := hclDataSourceLocationsPortAvailability(pop)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "ports_available.0.zone"),
					resource.TestCheckResourceAttrSet(resourceName, "ports_available.0.speed"),
					resource.TestCheckResourceAttrSet(resourceName, "ports_available.0.media"),
					resource.TestCheckResourceAttrSet(resourceName, "ports_available.0.count"),
					resource.TestCheckResourceAttrSet(resourceName, "ports_available.0.partial"),
					resource.TestCheckResourceAttrSet(resourceName, "ports_available.0.enni"),
				),
			},
		},
	})
}
