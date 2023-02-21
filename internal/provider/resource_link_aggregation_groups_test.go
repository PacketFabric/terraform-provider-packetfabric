package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func hclLinkAggregGroups(description, interval, pop, member1, member2 string) (hcl string, resourceName string) {

	hclName := testutil.GenerateUniqueResourceName()
	resourceName = "packetfabric_link_aggregation_group." + hclName
	hcl = fmt.Sprintf(testutil.RResourceLinkAggregationGroup, hclName, description, interval, member1, member2, pop)
	return
}

func TestAccLinkAggregGroupsRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	description := os.Getenv(testutil.DESCR_HC_C)
	interval := os.Getenv(testutil.PF_INTERVAL_KEY)
	pop := os.Getenv(testutil.PF_CS_POP7_KEY)
	member1 := os.Getenv(testutil.PF_MEMBER1_KEY)
	member2 := os.Getenv(testutil.PF_MEMBER2_KEY)

	hcl, resourceName := hclLinkAggregGroups(
		description,
		interval,
		pop,
		member1,
		member2,
	)
	members := []string{member1, member2}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_INTERVAL_KEY,
				testutil.PF_CS_POP7_KEY,
				testutil.PF_MEMBER1_KEY,
				testutil.PF_MEMBER2_KEY,
				testutil.DESCR_HC_C,
			})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "interval", interval),
					resource.TestCheckResourceAttr(resourceName, "pop", pop),
					resource.TestCheckResourceAttr(resourceName, "members", members[0]),
				),
			},
		},
	})

}
