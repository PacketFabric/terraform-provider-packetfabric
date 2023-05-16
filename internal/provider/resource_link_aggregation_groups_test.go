//go:build resource || core || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccLinkAggregGroupsRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	linkAggregationGroupResult := testutil.RHclLinkAggregationGroup()

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: linkAggregationGroupResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "description", linkAggregationGroupResult.Desc),
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "interval", linkAggregationGroupResult.Interval),
					resource.TestCheckResourceAttr(linkAggregationGroupResult.ResourceName, "pop", linkAggregationGroupResult.Pop),
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
