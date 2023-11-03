//go:build datasource || cloud_router || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCloudRouterConnectionIpsecComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourceIpsecResult := testutil.DHclCloudRouterConnIpsec()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceIpsecResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "customer_gateway_address"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "local_gateway_address"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "ike_version"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "phase1_authentication_method"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "phase1_group"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "phase1_encryption_algo"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "phase1_authentication_algo"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "phase1_lifetime"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "phase2_pfs_group"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "phase2_encryption_algo"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "phase2_lifetime"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "pre_shared_key"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(datasourceIpsecResult.ResourceName, "time_updated"),
				),
			},
		},
	})

}
