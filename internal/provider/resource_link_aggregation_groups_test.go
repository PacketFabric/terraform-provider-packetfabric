package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccLinkAggregGroupsRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	linkAggregationGroupResult := testutil.RHclLinkAggregationGroup()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: linkAggregationGroupResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "description", linkAggregationGroupResult.Desc),
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "interval", linkAggregationGroupResult.Interval),
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "pop", linkAggregationGroupResult.Pop),
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "members", linkAggregationGroupResult.Members[0]),
				),
			},
			{
				ResourceName:      linkAggregationGroupResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}
