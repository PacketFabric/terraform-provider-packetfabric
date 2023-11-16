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
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.#"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.admin_contact_uuid"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.bgp_region"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.id"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.#"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.%"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.currently_used_prefixes.#"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.currently_used_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.currently_used_prefixes.0.description"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.currently_used_prefixes.0.ips_in_use"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.currently_used_prefixes.0.isp_name"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.currently_used_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.currently_used_prefixes.0.will_renumber"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.#"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.0.usage_1y"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.0.usage_30d"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.0.usage_3m"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.0.usage_6m"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.1.%"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.1.description"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.1.location"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.1.prefix"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.1.usage_1y"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.1.usage_30d"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.1.usage_3m"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.ipj_details.0.planned_prefixes.1.usage_6m"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.length"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.state"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.tech_contact_uuid"),
					resource.TestCheckResourceAttrSet(hclIpamPrefixes.ResourceName, "ipam_prefixes.0.version"),
				),
			},
		},
	})
}
