package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclDataSourceLocationsCloud(cloudProvider, cloudConnectionType string) (hcl string, resourceName string) {
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "data.packetfabric_locations_cloud." + hclName
	hcl = fmt.Sprintf(`
	data "packetfabric_locations_cloud" "%s" {
		cloud_provider        = "%s"
		cloud_connection_type = "%s"
	}`, hclName, cloudProvider, cloudConnectionType)
	return
}

func TestAccDataSourceLocationsCloud(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)
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
		hcl, resourceName := hclDataSourceLocationsCloud(testCase.cloudProvider, testCase.cloudConnectionType)
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { testutil.PreCheck(t, nil) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: hcl,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(resourceName, "cloud_locations.0.pop"),
						resource.TestCheckResourceAttrSet(resourceName, "cloud_locations.0.region"),
						resource.TestCheckResourceAttrSet(resourceName, "cloud_locations.0.market"),
						resource.TestCheckResourceAttr(resourceName, "cloud_locations.0.cloud_provider", testCase.cloudProvider),
						resource.TestCheckResourceAttr(resourceName, "cloud_locations.0.cloud_connection_hosted_type", fmt.Sprintf("%s-connection", testCase.cloudConnectionType)),
					),
				},
			},
		})
	}
}
