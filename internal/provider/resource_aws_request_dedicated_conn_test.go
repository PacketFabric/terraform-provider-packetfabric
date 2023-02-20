package provider

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclAwsDedicatedConnection(awsRegion, description, pop, subscriptionTerm, serviceClass, speed string, autoneg bool) (hcl string, resourceName string) {

	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_cs_aws_dedicated_connection." + hclName
	hcl = fmt.Sprintf(testutil.RResourceCSAwsDedicatedConnection, hclName, awsRegion, description, pop, subscriptionTerm, serviceClass, speed, autoneg)
	return
}

func TestAccAwsDedicatedConnectionRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	awsRegion := os.Getenv(testutil.AWS_REGION3_KEY)
	description := os.Getenv(testutil.PF_CS_DESCR)
	pop := os.Getenv(testutil.PF_CS_POP3_KEY)
	speed := os.Getenv(testutil.PF_CS_SPEED3_KEY)
	subscriptionTerm := os.Getenv(testutil.PF_CS_SUBTERM_KEY)
	serviceClass := os.Getenv(testutil.PF_CS_SRVCLASS_KEY)
	autoneg := testutil.GetEnvBool(testutil.PF_CS_AUTONEG_KEY)

	hcl, resourceName := hclAwsDedicatedConnection(
		awsRegion,
		description,
		pop,
		subscriptionTerm,
		serviceClass,
		speed,
		autoneg,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.AWS_REGION3_KEY,
				testutil.PF_CS_SPEED3_KEY,
				testutil.PF_CS_DESCR,
				testutil.PF_CS_POP3_KEY,
				testutil.PF_CS_SUBTERM_KEY,
				testutil.PF_CS_SRVCLASS_KEY,
				testutil.PF_CS_AUTONEG_KEY,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "aws_region", awsRegion),
					resource.TestCheckResourceAttr(resourceName, "speed", speed),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "pop", pop),
					resource.TestCheckResourceAttr(resourceName, "subscription_term", subscriptionTerm),
					resource.TestCheckResourceAttr(resourceName, "service_class", serviceClass),
					resource.TestCheckResourceAttr(resourceName, "autoneg", strconv.FormatBool(autoneg)),
				),
			},
		},
	})
}
