//go:build resource || dedicated_cloud || all

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAzureDedicatedConnectionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	csAzureDedicatedConnectionResult := testutil.RHclCsAzureDedicatedConnection()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: csAzureDedicatedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "description", csAzureDedicatedConnectionResult.Desc),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "pop", csAzureDedicatedConnectionResult.Pop),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "zone", csAzureDedicatedConnectionResult.Zone),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "subscription_term", strconv.Itoa(csAzureDedicatedConnectionResult.SubscriptionTerm)),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "service_class", csAzureDedicatedConnectionResult.ServiceClass),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "speed", csAzureDedicatedConnectionResult.Speed),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "encapsulation", csAzureDedicatedConnectionResult.Encapsulation),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "port_category", csAzureDedicatedConnectionResult.PortCategory),
				),
			},
			{
				ResourceName:      csAzureDedicatedConnectionResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
