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
	defaultInput := testutil.DefaultRHclCloudRouterInput()
	cloudRouterResult1 := testutil.RHclCloudRouter(defaultInput)
	cloudRouterResult2 := testutil.RHclCloudRouter(testutil.RHclCloudRouterInput{
		ResourceName: defaultInput.ResourceName,
		HclName:      defaultInput.HclName,
		Capacity:     testutil.CloudRouterCapacityChange,
	})
	cloudRouterResult2.ResourceName = cloudRouterResult1.ResourceName
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterResult1.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterResult1.ResourceName, "asn", strconv.Itoa(cloudRouterResult1.Asn)),
					resource.TestCheckResourceAttr(cloudRouterResult1.ResourceName, "capacity", cloudRouterResult1.Capacity),
					resource.TestCheckResourceAttr(cloudRouterResult1.ResourceName, "regions.0", cloudRouterResult1.Regions[0]),
					resource.TestCheckResourceAttr(cloudRouterResult1.ResourceName, "regions.1", cloudRouterResult1.Regions[1]),
				),
			},
			{
				Config: cloudRouterResult2.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterResult2.ResourceName, "capacity", testutil.CloudRouterCapacityChange),
				),
			},
			{
				ResourceName:      cloudRouterResult2.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
