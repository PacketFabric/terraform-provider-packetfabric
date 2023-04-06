package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPortLOARequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	portLoaResult := testutil.RHclPortLoa()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             portLoaResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(portLoaResult.ResourceName, "destination_email", portLoaResult.DestinationEmail),
					resource.TestCheckResourceAttr(portLoaResult.ResourceName, "loa_customer_name", portLoaResult.LoaCustomerName),
				),
			},
			{
				ResourceName:      portLoaResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
