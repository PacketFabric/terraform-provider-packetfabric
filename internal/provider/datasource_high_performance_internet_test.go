//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceHighPerformanceInternetComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	hclHighPerformanceInternetResult := testutil.DHclHighPerformanceInternet()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclHighPerformanceInternetResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "port_circuit_id"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "speed"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "vlan"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "description"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "market"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "routing_type"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "state"),
					// TODO: need to verify the routing configuration settings
				),
			},
		},
	})
}
