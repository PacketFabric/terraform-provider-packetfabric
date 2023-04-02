package provider

import (
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAzureReqExpressConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	csAzureHostedConnectionResult := testutil.RHclCSAzureHostedConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_AZURE_SERVICE_KEY,
				testutil.PF_CS_VLAN_PRIVATE_KEY,
				testutil.PF_CS_VLAN_MICROSOFT_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             csAzureHostedConnectionResult.Hcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csAzureHostedConnectionResult.ResourceName, "description", csAzureHostedConnectionResult.Desc),
					resource.TestCheckResourceAttr(csAzureHostedConnectionResult.ResourceName, "azure_service_key", csAzureHostedConnectionResult.AzureServiceKey),
					resource.TestCheckResourceAttr(csAzureHostedConnectionResult.ResourceName, "speed", csAzureHostedConnectionResult.Speed),
					resource.TestCheckResourceAttr(csAzureHostedConnectionResult.ResourceName, "vlan_private", strconv.Itoa(csAzureHostedConnectionResult.VlanPrivate)),
					resource.TestCheckResourceAttr(csAzureHostedConnectionResult.ResourceName, "vlan_microsoft", strconv.Itoa(csAzureHostedConnectionResult.VlanMicrosoft)),
				),
			},
		},
	})
}
