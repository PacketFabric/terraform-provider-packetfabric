package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceLocationsComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceLocationsResult := testutil.DHclDataSourceLocations()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceLocationsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.pop"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.region"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.market"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.market_description"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.vendor"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.site"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.site_code"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.type"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.status"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.latitude"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.longitude"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.timezone"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.notes"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.pcode"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.lead_time"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.single_armed"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.address1"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.address2"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.city"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.state"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.postal"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.country"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.network_provider"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceLocationsResult.ResourceName, "locations.0.enni_supported"),
				),
			},
		},
	})

}
