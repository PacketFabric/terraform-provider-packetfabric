//go:build datasource || core || all

package provider

import (
	"fmt"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceHighPerformanceInternetsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	hclHighPerformanceInternetResult := testutil.DHclHighPerformanceInternets()
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n%s\n<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n", hclHighPerformanceInternetResult.Hcl)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclHighPerformanceInternetResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.circuit_id"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.port_circuit_id"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.speed"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.vlan"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.description"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.market"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.routing_type"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.state"),
				),
			},
		},
	})
}
