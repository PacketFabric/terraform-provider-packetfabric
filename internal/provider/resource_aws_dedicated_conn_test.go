package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAwsDedicatedConnectionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)
	csAwsDedicatedConnectionResult := testutil.RHclCsAwsDedicatedConnection()
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: csAwsDedicatedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csAwsDedicatedConnectionResult.ResourceName, "description", csAwsDedicatedConnectionResult.Description),
					resource.TestCheckResourceAttr(csAwsDedicatedConnectionResult.ResourceName, "pop", csAwsDedicatedConnectionResult.Pop),
					resource.TestCheckResourceAttr(csAwsDedicatedConnectionResult.ResourceName, "subscription_term", strconv.Itoa(csAwsDedicatedConnectionResult.SubscriptionTerm)),
					resource.TestCheckResourceAttr(csAwsDedicatedConnectionResult.ResourceName, "service_class", csAwsDedicatedConnectionResult.ServiceClass),
					resource.TestCheckResourceAttr(csAwsDedicatedConnectionResult.ResourceName, "autoneg", strconv.FormatBool(csAwsDedicatedConnectionResult.Autoneg)),
				),
			},
			{
				ResourceName:      csAwsDedicatedConnectionResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
