package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMRouteCloudConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	cloudRouterConnectionIBMResult := testutil.RHclCloudRouterConnectionIBM()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_IBM_ACCOUNT_ID_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             cloudRouterConnectionIBMResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionIBMResult.ResourceName, "description", cloudRouterConnectionIBMResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionIBMResult.ResourceName, "pop", cloudRouterConnectionIBMResult.Pop),
					resource.TestCheckResourceAttr(cloudRouterConnectionIBMResult.ResourceName, "ibm_bgp_asn", strconv.Itoa(cloudRouterConnectionIBMResult.IbmBgpAsn)),
					resource.TestCheckResourceAttr(cloudRouterConnectionIBMResult.ResourceName, "speed", cloudRouterConnectionIBMResult.Speed),
				),
			},
		},
	})
}
