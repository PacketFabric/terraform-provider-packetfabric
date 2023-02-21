package provider

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclIxVc(description, routingId, market, asn, vlan, subscriptionTerm, speed, longhaultype string, untagged bool) (hcl string, resourceName string) {

	portHcl, portResourceName := hclPort(
		os.Getenv(testutil.PF_PORT_DESCR),
		os.Getenv(testutil.PF_PORT_MEDIA_KEY),
		os.Getenv(testutil.PF_PORT_SPEED_KEY),
		os.Getenv(testutil.PF_PORT_POP1_KEY),
		os.Getenv(testutil.PF_PORT_SUBTERM_KEY),
	)

	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_ix_virtual_circuit_marketplace." + hclName
	ixVcHcl := fmt.Sprintf(testutil.RResourceIXVirtualCircuitMarketplace, hclName, description, routingId, market, asn,
		portResourceName, untagged, vlan, longhaultype, speed, subscriptionTerm)
	hcl = fmt.Sprintf("%s\n%s", portHcl, ixVcHcl)
	return
}

func TestAccIxVcRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	routingId := os.Getenv(testutil.PF_ROUTING_ID_KEY)
	description := os.Getenv(testutil.PF_PORT_DESCR)
	market := os.Getenv(testutil.PF_MARKET_IX_KEY)
	asn := os.Getenv(testutil.PF_ASN_IX_KEY)
	vlan := os.Getenv(testutil.PF_VC_VLAN1_KEY)
	subscriptionTerm := os.Getenv(testutil.PF_VC_SUBTERM_KEY)
	speed := os.Getenv(testutil.PF_VC_SPEED_KEY)
	longhaultype := os.Getenv(testutil.PF_VC_LONGHAUL_TYPE_KEY)
	untagged := testutil.GetEnvBool(testutil.PF_PORT_UNTAGGED_KEY)

	hcl, resourceName := hclIxVc(
		description,
		routingId,
		market,
		asn,
		vlan,
		subscriptionTerm,
		speed,
		longhaultype,
		untagged,
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_PORT_DESCR,
				testutil.PF_ROUTING_ID_IX_KEY,
				testutil.PF_ASN_IX_KEY,
				testutil.PF_MARKET_IX_KEY,
				testutil.PF_VC_SPEED_KEY,
				testutil.PF_VC_LONGHAUL_TYPE_KEY,
				testutil.PF_VC_SUBTERM_KEY,
				testutil.PF_VC_VLAN1_KEY,
				testutil.PF_PORT_UNTAGGED_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "routing_id", routingId),
					resource.TestCheckResourceAttr(resourceName, "asn", asn),
					resource.TestCheckResourceAttr(resourceName, "market", market),
					resource.TestCheckResourceAttr(resourceName, "interface.0.vlan", vlan),
					resource.TestCheckResourceAttr(resourceName, "interface.0.untagged", strconv.FormatBool(untagged)),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.longhaul_type", longhaultype),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.speed", speed),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.subscription_term", subscriptionTerm),
				),
			},
		},
	})
}
