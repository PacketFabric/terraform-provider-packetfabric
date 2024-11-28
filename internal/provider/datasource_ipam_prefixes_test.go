//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamPrefixesComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	hclIpamPrefixes := testutil.DHclIpamPrefixes()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclIpamPrefixes.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ip_address"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.circuit_id"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.linked_object_circuit_id"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.type"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.org_id"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.address"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.city"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.postal_code"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.admin_ipam_contact_uuid"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.tech_ipam_contact_uuid"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.state"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.#"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.%"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.current_prefixes.#"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.current_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.current_prefixes.0.description"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.current_prefixes.0.ips_in_use"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.current_prefixes.0.isp_name"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.current_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.current_prefixes.0.will_renumber"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefix.description"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefix.location"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefix.usage_1y"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefix.usage_30d"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefix.usage_3m"),
				),
			},
		},
	})
}
