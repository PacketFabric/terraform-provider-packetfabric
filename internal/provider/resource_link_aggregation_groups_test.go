<<<<<<< HEAD
//go:build resource || core || all

=======
>>>>>>> 10162f3 (Acc test structure: Resource packetfabric_link_aggregation_group)
package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccLinkAggregGroupsRequiredFields(t *testing.T) {
<<<<<<< HEAD
	testutil.PreCheck(t, nil)
=======

	testutil.SkipIfEnvNotSet(t)
>>>>>>> 10162f3 (Acc test structure: Resource packetfabric_link_aggregation_group)

	linkAggregationGroupResult := testutil.RHclLinkAggregationGroup()

	resource.ParallelTest(t, resource.TestCase{
<<<<<<< HEAD
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: linkAggregationGroupResult.Hcl,
=======
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             linkAggregationGroupResult.Hcl,
				ExpectNonEmptyPlan: true,
>>>>>>> 10162f3 (Acc test structure: Resource packetfabric_link_aggregation_group)
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "description", linkAggregationGroupResult.Desc),
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "interval", linkAggregationGroupResult.Interval),
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "pop", linkAggregationGroupResult.Pop),
<<<<<<< HEAD
				),
			},
			{
				ResourceName:      linkAggregationGroupResult.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
=======
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "members", linkAggregationGroupResult.Members[0]),
				),
			},
>>>>>>> 10162f3 (Acc test structure: Resource packetfabric_link_aggregation_group)
		},
	})

}
