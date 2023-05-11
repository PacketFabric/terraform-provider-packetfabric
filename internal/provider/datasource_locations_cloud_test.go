//go:build datasource || all || smoke

package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceLocationsCloudComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)
	t.Parallel()
	testCases := []struct {
		cloudProvider       string
		cloudConnectionType string
	}{
		{"aws", "hosted"},
		{"aws", "dedicated"},
		{"google", "hosted"},
		{"google", "dedicated"},
	}
	for _, testCase := range testCases {

		dataSourceLocationsCloudResult := testutil.DHclDataSourceLocationsCloud(testCase.cloudProvider, testCase.cloudConnectionType)

		resource.Test(t, resource.TestCase{
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: dataSourceLocationsCloudResult.Hcl,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(dataSourceLocationsCloudResult.ResourceName, "cloud_locations.0.pop"),
						resource.TestCheckResourceAttrSet(dataSourceLocationsCloudResult.ResourceName, "cloud_locations.0.region"),
						resource.TestCheckResourceAttrSet(dataSourceLocationsCloudResult.ResourceName, "cloud_locations.0.market"),
						resource.TestCheckResourceAttr(dataSourceLocationsCloudResult.ResourceName, "cloud_locations.0.cloud_provider", testCase.cloudProvider),
						resource.TestCheckResourceAttr(dataSourceLocationsCloudResult.ResourceName, "cloud_locations.0.cloud_connection_hosted_type", fmt.Sprintf("%s-connection", testCase.cloudConnectionType)),
					),
				},
			},
		})
	}
}
