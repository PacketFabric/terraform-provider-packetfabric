package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclPortLOA(destinationEmail, loaCustomerName string) (hcl string, resourceName string) {

	portHcl, portResourceName := hclPort(
		os.Getenv(testutil.PF_PORT_DESCR),
		os.Getenv(testutil.PF_PORT_MEDIA_KEY),
		os.Getenv(testutil.PF_PORT_SPEED_KEY),
		os.Getenv(testutil.PF_PORT_POP1_KEY),
		os.Getenv(testutil.PF_PORT_SUBTERM_KEY),
	)
	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_port_loa." + hclName
	portLOAhcl := fmt.Sprintf(testutil.RResourcePortLoa, hclName, portResourceName, destinationEmail, loaCustomerName)
	hcl = fmt.Sprintf("%s\n%s", portHcl, portLOAhcl)
	return
}

func TestAccPortLOARequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	destinationEmail := os.Getenv(testutil.PF_DEST_EMAIL_KEY)
	loaCustomerName := os.Getenv(testutil.PF_LOA_CUSTOMER_KEY)

	hcl, resourceName := hclPortLOA(
		destinationEmail,
		loaCustomerName,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_DEST_EMAIL_KEY,
				testutil.PF_LOA_CUSTOMER_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "destination_email", destinationEmail),
					resource.TestCheckResourceAttr(resourceName, "loa_customer_name", loaCustomerName),
				),
			},
		},
	})
}
