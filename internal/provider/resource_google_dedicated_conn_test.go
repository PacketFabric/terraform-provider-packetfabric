package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGoogleDedicatedConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	csGoogleDedicatedConnectionResult := testutil.RHclCSGoogleDedicatedConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
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
