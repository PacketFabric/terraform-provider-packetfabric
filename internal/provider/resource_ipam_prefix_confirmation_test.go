//go:build resource || ipam_prefix || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIpamPrefixConfirmationRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	ipamPrefixConfirmationResult := testutil.RHclIpamPrefixConfirmation()

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: ipamPrefixConfirmationResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "admin_contact_uuid"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "tech_contact_uuid"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "prefix_uuid"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.#"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.currently_used_prefixes.#"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.description"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.ips_in_use"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.isp_name"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.currently_used_prefixes.0.will_renumber"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.#"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.0.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.0.usage_1y"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.0.usage_30d"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.0.usage_3m"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.0.usage_6m"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.1.%"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.1.description"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.1.location"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.1.prefix"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.1.usage_1y"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.1.usage_30d"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.1.usage_3m"),
					resource.TestCheckResourceAttrSet(ipamPrefixConfirmationResult.ResourceName, "ipj_details.0.planned_prefixes.1.usage_6m"),
				),
			},
		},
	})
}
