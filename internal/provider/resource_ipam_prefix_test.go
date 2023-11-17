//go:build resource || ipam_prefix || all

package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIpamPrefixRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	ipamPrefixResult := testutil.RHclIpamPrefix()
    updatedHcl := strings.Replace(ipamPrefixResult.Hcl, "Optional", "Totally optional", -1)

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: ipamPrefixResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "admin_contact_uuid"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "bgp_region"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "length"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixResult.ResourceName, "state"),
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
			{
				Config: updatedHcl,
				ExpectNonEmptyPlan: true,
				Check: func(s *terraform.State) error {
					rs, ok := s.RootModule().Resources[ipamPrefixResult.ResourceName]
					if !ok {
						return fmt.Errorf("Not found: %s", ipamPrefixResult.ResourceName)
					}
					expectations := map[string]string{
						"ipj_details.0.currently_used_prefixes.0.description": "Totally optional description",
						"ipj_details.0.currently_used_prefixes.0.isp_name": "Totally optional ISP Name",
						"ipj_details.0.planned_prefixes.1.location": "Totally optional Location",
					}
					for key, expected := range expectations {
						actual := rs.Primary.Attributes[key]
						if actual != expected {
							return fmt.Errorf("For \"%s\", expected \"%s\", but got \"%s\"", key, expected, actual)
						}
					}
					return nil
				},
			},
		},
	})
}
