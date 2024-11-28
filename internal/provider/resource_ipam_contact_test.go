//go:build resource || ipam_contact || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIpamContactRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	ipamContactResult := testutil.RHclIpamContact()

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: ipamContactResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "name"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "address"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "country_code"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "phone"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "email"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "apnic_org_id"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "ripe_org_id"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "apnic_ref"),
					resource.TestCheckResourceAttrSet(ipamContactResult.ResourceName, "ripe_ref"),
				),
			},
		},
	})
}
