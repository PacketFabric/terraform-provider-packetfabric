//go:build resource || ipam_prefix || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIpamPrefixequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	ipamPrefixResult := testutil.RHclIpamPrefix()

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: ipamPrefixResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "admin_contact_uuid"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "bgp_region"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "length"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "tech_contact_uuid"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "version"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.#"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.currently_used_prefixes.#"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.description"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.ips_in_use"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.isp_name"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.will_renumber"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.#"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.0.usage_1y"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.0.usage_30d"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.0.usage_3m"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.0.usage_6m"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.1.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.1.description"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.1.location"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.1.prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.1.usage_1y"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.1.usage_30d"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.1.usage_3m"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefixes.1.usage_6m"),

				),
			},
		},
	})
}
