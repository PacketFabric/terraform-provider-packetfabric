package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAzureCloudRouterConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	cloudRouterConnectionAzureResult := testutil.RHclCloudRouterConnectionAzure()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_ACCOUNT_ID_KEY,
				testutil.PF_CRC_AZURE_SERVICE_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: cloudRouterConnectionAzureResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(cloudRouterConnectionAzureResult.ResourceName, "azure_service_key", cloudRouterConnectionAzureResult.AzureServiceKey),
					resource.TestCheckResourceAttr(cloudRouterConnectionAzureResult.ResourceName, "description", cloudRouterConnectionAzureResult.Desc),
					resource.TestCheckResourceAttr(cloudRouterConnectionAzureResult.ResourceName, "speed", cloudRouterConnectionAzureResult.Speed),
				),
			},
			{
				ResourceName:      cloudRouterConnectionAzureResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}
