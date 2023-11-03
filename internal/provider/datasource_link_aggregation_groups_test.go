//go:build datasource || core || all

package provider

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func updateLagHclMembers(lagHcl string, portHclResult testutil.RHclPortResult) string {
	r := regexp.MustCompile(`members.+`)
	matches := r.FindAllString(lagHcl, -1)
	newMembers := strings.Replace(matches[0], "]", fmt.Sprintf(", %s.id]", portHclResult.ResourceName), -1)
	return r.ReplaceAllString(lagHcl, newMembers)
}

func TestAccDataSourceLinkAggregationGroupsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourceLinkAggregationGroupsResult := testutil.DHclLinkAggregationGroups()
	portHclResult := testutil.LinkAggregationGroupPort()
	singleLagMemberHcl := fmt.Sprintf("%s\n%s", portHclResult.Hcl, datasourceLinkAggregationGroupsResult.Hcl)
	doubleLagMembersHcl := updateLagHclMembers(singleLagMemberHcl, portHclResult)

	resource.ParallelTest(t, resource.TestCase{
		Providers:         testAccProviders,
		ExternalProviders: testAccExternalProviders,
		Steps: []resource.TestStep{
			{
				Config: doubleLagMembersHcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "lag_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.autoneg"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.state"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.speed"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.media"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.zone"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.region"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.market"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.market_description"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.pop"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.site"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.site_code"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.operational_status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.admin_status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.mtu"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.description"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_lag"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_lag_member"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.member_count"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.account_uuid"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.autoneg"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.state"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.speed"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.media"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.zone"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.region"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.market"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.market_description"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.pop"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.site"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.site_code"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.operational_status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.admin_status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.mtu"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.description"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.is_lag"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.is_lag_member"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.member_count"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.account_uuid"),
				),
			},
			{
				Config:             singleLagMemberHcl,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "lag_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.autoneg"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.state"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.speed"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.media"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.zone"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.region"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.market"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.market_description"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.pop"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.site"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.site_code"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.operational_status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.admin_status"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.mtu"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.description"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_lag"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_lag_member"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.member_count"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.account_uuid"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.autoneg"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.port_circuit_id"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.state"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.status"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.speed"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.media"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.zone"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.region"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.market"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.market_description"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.pop"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.site"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.site_code"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.operational_status"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.admin_status"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.mtu"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.description"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.is_lag"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.is_lag_member"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.member_count"),
					resource.TestCheckNoResourceAttr(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.1.account_uuid"),
				),
			},
		},
	})

}
