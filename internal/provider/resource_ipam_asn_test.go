//go:build resource || ipam_asn || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIpamAsnRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	ipamAsnResult := testutil.RHclIpamAsn()

	resource.ParallelTest(t, resource.TestCase{
		ExternalProviders: testAccExternalProviders,
		Providers:         testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: ipamAsnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(ipamAsnResult.ResourceName, "asn_byte_type"),
					resource.TestCheckResourceAttrSet(ipamAsnResult.ResourceName, "asn"),
					resource.TestCheckResourceAttrSet(ipamAsnResult.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(ipamAsnResult.ResourceName, "time_updated"),
				),
			},
		},
	})
}
