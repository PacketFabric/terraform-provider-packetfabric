package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclDataSourceLocationsMarkets() (hcl string, resourceName string) {
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "data.packetfabric_locations_markets." + hclName
	hcl = fmt.Sprintf(`
	data "packetfabric_locations_markets" "%s" {
	}`, hclName)
	return
}

func TestAccDataSourceLocationsMarkets(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)
	hcl, resourceName := hclDataSourceLocationsMarkets()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutil.PreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "locations_markets.0.name"),
					resource.TestCheckResourceAttrSet(resourceName, "locations_markets.0.code"),
					resource.TestCheckResourceAttrSet(resourceName, "locations_markets.0.country"),
				),
			},
		},
	})
}
