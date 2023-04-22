package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceLinkAggregationGroupsComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceLinkAggregationGroupsResult := testutil.DHclDatasourceLinkAggregationGroups()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceLinkAggregationGroupsResult.Hcl,
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
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.vc_mode"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_lag"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_lag_member"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_cloud"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_ptp"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.is_nni"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.lag_interval"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.member_count"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.parent_lag_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.account_uuid"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.subscription_term"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.disabled"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.customer_name"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceLinkAggregationGroupsResult.ResourceName, "interfaces.0.time_updated"),
				)},
		},
	})

}
