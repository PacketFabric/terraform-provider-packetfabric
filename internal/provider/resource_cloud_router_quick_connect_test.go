package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudRouterQuickConnectRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{
		"PF_AWS_ACCOUNT_ID",
		"PF_QUICK_CONNECT_SERVICE_UUID",
	})

	cloudRouterQuickConnect := testutil.RHclCloudRouterQuickConnect()

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterQuickConnect.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(cloudRouterQuickConnect.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(cloudRouterQuickConnect.ResourceName, "route_set_circuit_id"),
					resource.TestCheckResourceAttrSet(cloudRouterQuickConnect.ResourceName, "time_created"),
				),
			},
			{
				ResourceName:      cloudRouterQuickConnect.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
