//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamContactsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	hclIpamContactResult := testutil.DHclIpamContacts()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclIpamContactResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.address"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.country_code"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.apnic_org_id"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.apnic_ref"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.name"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.email"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.phone"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.ripe_org_id"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.ripe_ref"),
					resource.TestCheckResourceAttrSet(hclIpamContactResult.ResourceName, "ipam_contacts.0.uuid"),
				),
			},
		},
	})
}
