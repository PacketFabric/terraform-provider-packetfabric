package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPBgpSessionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	bgpSessionResult := testutil.RHclBgpSession()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             bgpSessionResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "address_family", bgpSessionResult.AddressFamily),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "remote_asn", strconv.Itoa(bgpSessionResult.Asn)),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "prefixes.0.prefix", bgpSessionResult.Prefix1),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "prefixes.0.type", bgpSessionResult.Type1),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "prefixes.1.prefix", bgpSessionResult.Prefix2),
					resource.TestCheckResourceAttr(bgpSessionResult.ResourceName, "prefixes.1.type", bgpSessionResult.Type2),
				),
			},
		},
	})
}
