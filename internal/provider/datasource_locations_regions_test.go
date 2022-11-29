package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclDataSourceLocationsRegions() (hcl string, resourceName string) {
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "data.packetfabric_locations_regions." + hclName
	hcl = fmt.Sprintf(`
	data "packetfabric_locations_regions" "%s" {
	}`, hclName)
	return
}

func TestAccDataSourceLocationsRegions(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)
	hcl, resourceName := hclDataSourceLocationsRegions()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "locations_regions.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "locations_regions.0.code"),
				),
			},
		},
	})

}
