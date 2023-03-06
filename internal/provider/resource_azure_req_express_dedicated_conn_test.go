package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAzureDedicatedConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	csAzureDedicatedConnectionResult := testutil.RHclCsAzureDedicatedConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:                    csAzureDedicatedConnectionResult.Hcl,
				ExpectNonEmptyPlan:        true,
				Destroy:                   !testutil.IsDevEnv(),
				PreventPostDestroyRefresh: testutil.IsDevEnv(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "description", csAzureDedicatedConnectionResult.Desc),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "pop", csAzureDedicatedConnectionResult.Pop),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "subscription_term", strconv.Itoa(csAzureDedicatedConnectionResult.SubscriptionTerm)),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "service_class", csAzureDedicatedConnectionResult.ServiceClass),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "speed", csAzureDedicatedConnectionResult.Speed),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "encapsulation", csAzureDedicatedConnectionResult.Encapsulation),
					resource.TestCheckResourceAttr(csAzureDedicatedConnectionResult.ResourceName, "port_category", csAzureDedicatedConnectionResult.PortCategory),
				),
			},
		},
	})
}
