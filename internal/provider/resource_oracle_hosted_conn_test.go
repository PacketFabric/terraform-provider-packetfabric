package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclOracleHostedConn(description, vcOcid, region, pop, zone, vlan string) (hcl string, resourceName string) {
	portHcl, portResourceName := hclPort(
		os.Getenv(testutil.PF_PORT_DESCR),
		os.Getenv(testutil.PF_PORT_MEDIA_KEY),
		os.Getenv(testutil.PF_PORT_SPEED_KEY),
		os.Getenv(testutil.PF_PORT_POP1_KEY),
		os.Getenv(testutil.PF_PORT_SUBTERM_KEY),
	)
	hclName := testutil.GenerateUniqueResourceName()

	resourceName = "packetfabric_cs_oracle_hosted_connection." + hclName
	oracleHostedConnHcl := fmt.Sprintf(testutil.RResourceCSOracleHostedConnection, hclName, description, vcOcid, region, portResourceName, pop, zone, vlan)
	hcl = fmt.Sprintf("%s\n%s", portHcl, oracleHostedConnHcl)
	return
}

func TestOracleHostedConnectRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	description := os.Getenv(testutil.DESCR_HC_C)
	vcOcid := os.Getenv(testutil.PF_CS_ORACLE_VC_OCID_KEY)
	pop := os.Getenv(testutil.PF_CS_POP6_KEY)
	vlan := os.Getenv(testutil.PF_CS_VLAN6_KEY)
	region := os.Getenv(testutil.PF_CS_ORACLE_REGION_KEY)
	zone := os.Getenv(testutil.PF_CS_ZONE6_KEY)

	hcl, resourceName := hclOracleHostedConn(
		description,
		vcOcid,
		region,
		pop,
		zone,
		vlan,
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CS_ORACLE_VC_OCID_KEY,
				testutil.PF_CS_ORACLE_REGION_KEY,
				testutil.DESCR_HC_C,
				testutil.PF_CS_POP6_KEY,
				testutil.PF_CS_VLAN6_KEY,
				testutil.PF_CS_ZONE6_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "vc_ocid", vcOcid),
					resource.TestCheckResourceAttr(resourceName, "region", region),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "zone", zone),
					resource.TestCheckResourceAttr(resourceName, "pop", pop),
					resource.TestCheckResourceAttr(resourceName, "vlan", vlan),
				),
			},
		},
	})

}
