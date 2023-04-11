package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAzureHostedMktConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	csAzureHostedMarketplaceConnectionResult := testutil.RHclCsAzureHostedMarketplaceConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_CRC_AZURE_SERVICE_KEY,
				testutil.PF_AZ_CS_HOSTED_MARKET_CONN_KEY,
				testutil.PF_ROUTING_ID_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: csAzureHostedMarketplaceConnectionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(csAzureHostedMarketplaceConnectionResult.ResourceName, "description", csAzureHostedMarketplaceConnectionResult.Desc),
					resource.TestCheckResourceAttr(csAzureHostedMarketplaceConnectionResult.ResourceName, "azure_service_key", csAzureHostedMarketplaceConnectionResult.AzureServiceKey),
					resource.TestCheckResourceAttr(csAzureHostedMarketplaceConnectionResult.ResourceName, "market", csAzureHostedMarketplaceConnectionResult.Market),
					resource.TestCheckResourceAttr(csAzureHostedMarketplaceConnectionResult.ResourceName, "speed", csAzureHostedMarketplaceConnectionResult.Speed),
				),
			},
			{
				ResourceName:      csAzureHostedMarketplaceConnectionResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
