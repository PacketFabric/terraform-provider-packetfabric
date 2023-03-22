package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAwsHostedMktRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	awsHostedMktResult := testutil.RHclCSAwsHostedMktConnection()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: awsHostedMktResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(awsHostedMktResult.ResourceName, "pop", awsHostedMktResult.Pop),
					resource.TestCheckResourceAttr(awsHostedMktResult.ResourceName, "market", awsHostedMktResult.Market),
					resource.TestCheckResourceAttr(awsHostedMktResult.ResourceName, "speed", awsHostedMktResult.Speed),
				),
			},
			{
				ResourceName:      awsHostedMktResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
