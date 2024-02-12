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
	fmt.Printf(">>>>>%s\n<<<<<", hclHighPerformanceInternetResult.Hcl)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclHighPerformanceInternetResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.address"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.country_code"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.apnic_org_id"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.apnic_ref"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.name"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.email"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.phone"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.ripe_org_id"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.ripe_ref"),
					resource.TestCheckResourceAttrSet(hclHighPerformanceInternetResult.ResourceName, "high_performance_internets.0.uuid"),
				),
			},
		},
	})
}
