//go:build resource || hosted_cloud || all

package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmHostedConnectionRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_IBM_ACCOUNT_ID", "IC_API_KEY", "IAAS_CLASSIC_USERNAME", "IAAS_CLASSIC_API_KEY", "TF_VAR_ibm_resource_group"})
	csIbmHostedConnectionResult := testutil.RHclCsIbmHostedConnection()
	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: csIbmHostedConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csIbmHostedConnectionResult.ResourceName, "account_uuid", csIbmHostedConnectionResult.AccountUuid),
					resource.TestCheckResourceAttr(csIbmHostedConnectionResult.ResourceName, "pop", csIbmHostedConnectionResult.Pop),
					resource.TestCheckResourceAttr(csIbmHostedConnectionResult.ResourceName, "speed", csIbmHostedConnectionResult.Speed),
					resource.TestCheckResourceAttr(csIbmHostedConnectionResult.ResourceName, "vlan", strconv.Itoa(csIbmHostedConnectionResult.Vlan)),
					resource.TestCheckResourceAttr(csIbmHostedConnectionResult.ResourceName, "ibm_bgp_asn", strconv.Itoa(csIbmHostedConnectionResult.IbmBgpAsn)),
				),
			},
			{
				ResourceName:      csIbmHostedConnectionResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
