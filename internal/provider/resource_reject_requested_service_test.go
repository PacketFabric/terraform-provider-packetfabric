package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclRejectReqService(vcRequestUuid string) (hcl string, resourceName string) {

	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_marketplace_service_port_reject_request." + hclName
	hcl = fmt.Sprintf(testutil.RResourceMarketplaceServicePortRejectRequest, hclName, vcRequestUuid)
	return
}

func TestAccRejectReqServiceRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	vcRequestUuid := os.Getenv(testutil.PF_VC_REQUEST_UUID_KEY)

	hcl, resourceName := hclRejectReqService(vcRequestUuid)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_VC_REQUEST_UUID_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "vc_request_uuid", vcRequestUuid),
				),
			},
		},
	})

}
