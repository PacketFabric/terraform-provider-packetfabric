//go:build datasource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIpamAsnsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	hclIpamAsnResult := testutil.DHclIpamAsns()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hclIpamAsnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(hclIpamAsnResult.ResourceName, "ipam_asns.0.asn_byte_type"),
					resource.TestCheckResourceAttrSet(hclIpamAsnResult.ResourceName, "ipam_asns.0.asn"),
					resource.TestCheckResourceAttrSet(hclIpamAsnResult.ResourceName, "ipam_asns.0.time_created"),
					resource.TestCheckResourceAttrSet(hclIpamAsnResult.ResourceName, "ipam_asns.0.time_updated"),
				),
			},
		},
	})
}
