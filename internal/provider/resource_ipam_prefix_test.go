//go:build resource || ipam_prefix || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIpamPrefixRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	ipamPrefixResult := testutil.RHclIpamPrefix()

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: ipamPrefixResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "length"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "market"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "family"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ip_address"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "linked_object_circuit_id"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "type"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "org_id"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "address"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "city"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "postal_code"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "admin_ipam_contact_uuid"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "tech_ipam_contact_uuid"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "state"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.#"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.current_prefixes.#"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.current_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.current_prefixes.0.description"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.current_prefixes.0.ips_in_use"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.current_prefixes.0.isp_name"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.current_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.current_prefixes.0.will_renumber"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefix.description"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefix.location"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefix.prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefix.usage_1y"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefix.usage_30d"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefix.usage_3m"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "ipj_details.0.planned_prefix.usage_6m"),
				),
			},
		},
	})
}
