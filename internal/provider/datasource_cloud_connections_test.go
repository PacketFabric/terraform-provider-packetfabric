//go:build datasource || cloud_router || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCloudRouterConnectionsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_AWS_ACCOUNT_ID"})

	datasourceCloudConnsResult := testutil.DHclCloudConnections()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceCloudConnsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.connection_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.port_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.connection_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.pending_delete"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.deleted"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.speed"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.state"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.account_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.service_class"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.service_provider"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.service_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.description"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_provider_connection_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_connections.0.cloud_settings.0.vlan_id_pf"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_connections.0.cloud_settings.0.vlan_id_cust"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_connections.0.cloud_settings.0.svlan_id_cust"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.user_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.time_updated"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_provider"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.pop"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.site"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.bgp_state_list"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_router_name"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_router_asn"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_router_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.nat_capable"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.dnat_capable"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.zone"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.vlan"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.desired_nat"),
				),
			},
		},
	})

}
