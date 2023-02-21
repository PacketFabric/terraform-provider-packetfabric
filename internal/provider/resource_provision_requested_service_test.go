package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclProvisionReqService(types, cloudProvider, vlan, vcRequestUuid string) (hcl string, resourceName string) {

	portHcl, portResourceName := hclPort(
		os.Getenv(testutil.PF_PORT_DESCR),
		os.Getenv(testutil.PF_PORT_MEDIA_KEY),
		os.Getenv(testutil.PF_PORT_SPEED_KEY),
		os.Getenv(testutil.PF_PORT_POP1_KEY),
		os.Getenv(testutil.PF_PORT_SUBTERM_KEY),
	)

	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_marketplace_service_port_accept_request." + hclName
	provisionReqServiceHcl := fmt.Sprintf(testutil.RResourceMarketplaceServicePortAcceptRequest, hclName, types, cloudProvider, portResourceName, vlan, vcRequestUuid)
	hcl = fmt.Sprintf("%s\n%s", portHcl, provisionReqServiceHcl)
	return
}

func TestAcchclProvisionReqServiceRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	vcRequestUuid := os.Getenv(testutil.PF_VC_REQUEST_UUID_KEY)
	types := os.Getenv(testutil.PF_TYPE_KEY)
	cloudProvider := os.Getenv(testutil.PF_CLOUD_PROVIDER_KEY)
	vlan := os.Getenv(testutil.PF_VLAN_KEY)

	hcl, resourceName := hclProvisionReqService(types, cloudProvider, vlan, vcRequestUuid)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_VC_REQUEST_UUID_KEY,
				testutil.PF_TYPE_KEY,
				testutil.PF_CLOUD_PROVIDER_KEY,
				testutil.PF_VLAN_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", types),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", cloudProvider),
					resource.TestCheckResourceAttr(resourceName, "vc_request_uuid", vcRequestUuid),
					resource.TestCheckResourceAttr(resourceName, "interface.0.vlan", vlan),
				),
			},
		},
	})

}
