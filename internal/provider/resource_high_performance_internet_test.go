//go:build resource || high_performance_internet || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccHighPerformanceInternetRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	highPerformanceInternetResult := testutil.RHclHighPerformanceInternet()

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: highPerformanceInternetResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "circuit_id"), 
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "account_uuid"),
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "port_circuit_id"),
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "speed"),
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "vlan"),
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "description"),
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "market"),
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_type"),
					resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "state"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.#"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.%"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.#"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.%"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.asn"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.l3_address"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.remote_address"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.md5"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.bgp_state"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.address_family"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.prefixes.#"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.prefixes.0.%"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.prefixes.0.prefix"),
                    resource.TestCheckResourceAttrSet(highPerformanceInternetResult.ResourceName, "routing_configuration.0.bgp_v4.0.prefixes.0.local_preference"),
				),
			},
		},
	})
}
