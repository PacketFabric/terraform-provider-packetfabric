//go:build resource || ipam_contact || all

package provider

import (
	"strings"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIpamContactRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	ipamContactResult := testutil.RHclIpamContact()
	updatedHcl := strings.Replace(ipamContactResult.Hcl, "1234", "5678", -1)

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: ipamContactResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "name"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "address"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "phone"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "email"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "apnic_org_id"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "ripe_org_id"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "apnic_ref"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "ripe_ref"),
				),
			},
			{
				Config: updatedHcl,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[ipamContactResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", ipamContactResult.ResourceName)
					}
					expected := "5678 Peachtree St, Atlanta, GA"
					actual := rs.Primary.Attributes["address"]
					if expected != actual {
						return fmt.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
					}
					return nil
				},
			},
		},
	})
}
