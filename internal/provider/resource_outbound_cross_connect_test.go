package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclOutboundCrossConnect(site, description, documentUuid string) (hcl string, resourceName string) {

	portHcl, portResourceName := hclPort(
		os.Getenv(testutil.PF_PORT_DESCR),
		os.Getenv(testutil.PF_PORT_MEDIA_KEY),
		os.Getenv(testutil.PF_PORT_SPEED_KEY),
		os.Getenv(testutil.PF_PORT_POP1_KEY),
		os.Getenv(testutil.PF_PORT_SUBTERM_KEY),
	)

	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_outbound_cross_connect." + hclName
	outboundCrossConnectHcl := fmt.Sprintf(testutil.RResourceOutboundCrossConnect, hclName, description, documentUuid, portResourceName, site)
	hcl = fmt.Sprintf("%s\n%s", portHcl, outboundCrossConnectHcl)
	return
}

func testOutboundCrossConnect(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	description := os.Getenv(testutil.PF_PORT_DESCR)
	site := os.Getenv(testutil.PF_SITE_KEY)
	documentUuid := os.Getenv(testutil.PF_DOCUMENT_UUID1_KEY)

	hcl, resourceName := hclOutboundCrossConnect(
		site,
		description,
		documentUuid,
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_PORT_DESCR,
				testutil.PF_SITE_KEY,
				testutil.PF_DOCUMENT_UUID1_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "document_uuid", documentUuid),
					resource.TestCheckResourceAttr(resourceName, "site", site),
				),
			},
		},
	})
}
