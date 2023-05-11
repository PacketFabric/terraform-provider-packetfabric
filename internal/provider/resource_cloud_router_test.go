//go:build resource || cloud_router || all

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudRouterRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)
	cloudRouterResult := testutil.RHclCloudRouter()
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterResult.ResourceName, "asn", strconv.Itoa(cloudRouterResult.Asn)),
					resource.TestCheckResourceAttr(cloudRouterResult.ResourceName, "capacity", cloudRouterResult.Capacity),
					resource.TestCheckResourceAttr(cloudRouterResult.ResourceName, "regions.0", cloudRouterResult.Regions[0]),
					resource.TestCheckResourceAttr(cloudRouterResult.ResourceName, "regions.1", cloudRouterResult.Regions[1]),
				),
			},
			{
				ResourceName:      cloudRouterResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
