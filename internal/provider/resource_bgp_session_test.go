package provider

import (
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPBgpSessionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	hcl, resourceName := testutil.HclBgpSession()
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CRBS_ADDRESS_FMLY_KEY,
				testutil.PF_CRBS_REMOTE_ASN_KEY,
				testutil.PF_CRBS_PRFX1_KEY,
				testutil.PF_CRBS_PRFX2_KEY,
				testutil.PF_CRBS_TYPE1_KEY,
				testutil.PF_CRBS_TYPE2_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "address_family", os.Getenv(testutil.PF_CRBS_ADDRESS_FMLY_KEY)),
					resource.TestCheckResourceAttr(resourceName, "remote_asn", os.Getenv(testutil.PF_CRBS_REMOTE_ASN_KEY)),
					resource.TestCheckResourceAttr(resourceName, "prefixes.0.prefix", os.Getenv(testutil.PF_CRBS_PRFX1_KEY)),
					resource.TestCheckResourceAttr(resourceName, "prefixes.0.type", os.Getenv(testutil.PF_CRBS_TYPE1_KEY)),
					resource.TestCheckResourceAttr(resourceName, "prefixes.1.prefix", os.Getenv(testutil.PF_CRBS_PRFX2_KEY)),
					resource.TestCheckResourceAttr(resourceName, "prefixes.1.type", os.Getenv(testutil.PF_CRBS_TYPE2_KEY)),
				),
			},
		},
	})
}
