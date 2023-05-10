package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPort(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	portDetails := testutil.CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: portTestResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(portTestResult.ResourceReference, "description", portTestResult.Description),
					resource.TestCheckResourceAttr(portTestResult.ResourceReference, "speed", portTestResult.Speed),
					resource.TestCheckResourceAttr(portTestResult.ResourceReference, "media", portTestResult.Media),
					resource.TestCheckResourceAttr(portTestResult.ResourceReference, "subscription_term", strconv.Itoa(portTestResult.SubscriptionTerm)),
					resource.TestCheckResourceAttr(portTestResult.ResourceReference, "pop", portTestResult.Pop),
					resource.TestCheckResourceAttr(portTestResult.ResourceReference, "enabled", strconv.FormatBool(portTestResult.Enabled)),
				),
			},
			{
				ResourceName:      portTestResult.ResourceReference,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
