package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclCloudRouterConnectionAws(awsAccountID, accountUUID, description, pop, speed string) (hcl string, resourceName string) {

	hclCloudRouter, crResourceName := hclCloudRouter(
		os.Getenv(testutil.PF_CR_DESCR),
		os.Getenv(testutil.PF_ACCOUNT_ID),
		os.Getenv(testutil.PF_CR_CAPACITY_KEY),
	)

	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_cloud_router_connection_aws." + hclName
	crcHcl := fmt.Sprintf(testutil.RResourceCloudRouterConnectionAws, hclName, crResourceName, awsAccountID, accountUUID, description, pop, speed)

	hcl = fmt.Sprintf("%s\n%s", hclCloudRouter, crcHcl)
	return
}

func TestAccCloudRouterConnectionAwsRequiredFields(t *testing.T) {
	testutil.SkipIfEnvNotSet(t)

	awsAccountId := os.Getenv(testutil.PF_AWS_ACCOUNT_ID)
	accountUuid := os.Getenv(testutil.PF_AWS_ACCOUNT_ID)
	description := os.Getenv(testutil.PF_CRC_DESCR)
	pop := os.Getenv(testutil.PF_CRC_POP1_KEY)
	speed := os.Getenv(testutil.PF_CRC_SPEED_KEY)

	hcl, resourceName := hclCloudRouterConnectionAws(
		awsAccountId,
		accountUuid,
		description,
		pop,
		speed,
	)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_CRC_DESCR,
				testutil.PF_AWS_ACCOUNT_ID,
				testutil.PF_AWS_ACCOUNT_ID,
				testutil.PF_CRC_POP1_KEY,
				testutil.PF_CRC_SPEED_KEY,
			})
		},
		ProviderFactories: testAccProviderFactories,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "aws_account_id", awsAccountId),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "account_uuid", accountUuid),
					resource.TestCheckResourceAttr(resourceName, "pop", pop),
					resource.TestCheckResourceAttr(resourceName, "speed", speed),
				),
			},
		},
	})
}
