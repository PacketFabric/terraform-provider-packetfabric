//go:build resource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPortLOARequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_USER_EMAIL"})

	portLoaResult := testutil.RHclPortLoa()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: portLoaResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(portLoaResult.ResourceName, "destination_email", portLoaResult.DestinationEmail),
					resource.TestCheckResourceAttr(portLoaResult.ResourceName, "loa_customer_name", portLoaResult.LoaCustomerName),
				),
			},
		},
	})
}
