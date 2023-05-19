//go:build resource || core || all

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPortRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	portDetails := testutil.CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: portTestResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(portTestResult.ResourceName, "description", portTestResult.Description),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, "speed", portTestResult.Speed),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, "media", portTestResult.Media),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, "subscription_term", strconv.Itoa(portTestResult.SubscriptionTerm)),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, "pop", portTestResult.Pop),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, "zone", portTestResult.Zone),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, "enabled", strconv.FormatBool(portTestResult.Enabled)),
				),
			},
			{
				ResourceName:      portTestResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
