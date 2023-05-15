//go:build resource || dedicated_cloud || all

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGoogleDedicatedConnectionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)
	csGoogleDedicatedConnectionResult := testutil.RHclCsGoogleDedicatedConnection()
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: csGoogleDedicatedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csGoogleDedicatedConnectionResult.ResourceName, "description", csGoogleDedicatedConnectionResult.Desc),
					resource.TestCheckResourceAttr(csGoogleDedicatedConnectionResult.ResourceName, "autoneg", strconv.FormatBool(csGoogleDedicatedConnectionResult.Autoneg)),
					resource.TestCheckResourceAttr(csGoogleDedicatedConnectionResult.ResourceName, "subscription_term", strconv.Itoa(csGoogleDedicatedConnectionResult.SubscriptionTerm)),
					resource.TestCheckResourceAttr(csGoogleDedicatedConnectionResult.ResourceName, "pop", csGoogleDedicatedConnectionResult.Pop),
					resource.TestCheckResourceAttr(csGoogleDedicatedConnectionResult.ResourceName, "service_class", csGoogleDedicatedConnectionResult.ServiceClass),
					resource.TestCheckResourceAttr(csGoogleDedicatedConnectionResult.ResourceName, "speed", csGoogleDedicatedConnectionResult.Speed),
					resource.TestCheckResourceAttr(csGoogleDedicatedConnectionResult.ResourceName, "zone", csGoogleDedicatedConnectionResult.Zone),
				),
			},
			{
				ResourceName:      csGoogleDedicatedConnectionResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
