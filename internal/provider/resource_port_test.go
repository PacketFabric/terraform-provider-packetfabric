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
					resource.TestCheckResourceAttr(portTestResult.ResourceName, PfDescription, portTestResult.Description),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, PfSpeed, portTestResult.Speed),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, PfMedia, portTestResult.Media),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, PfSubscriptionTerm, strconv.Itoa(portTestResult.SubscriptionTerm)),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, PfPop, portTestResult.Pop),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, PfZone, portTestResult.Zone),
					resource.TestCheckResourceAttr(portTestResult.ResourceName, PfEnabled, strconv.FormatBool(portTestResult.Enabled)),
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
