//go:build datasource || cloud_router || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCloudRouterConnectionsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_AWS_ACCOUNT_ID"})

	datasourceCloudConnsResult := testutil.DHclCloudRouterConns()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceCloudConnsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.port_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.connection_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.pending_delete"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.deleted"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.speed"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.state"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.account_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.service_class"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.service_provider"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.service_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.description"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_provider_connection_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_settings.0.vlan_id_pf"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_settings.0.vlan_id_cust"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_settings.0.svlan_id_cust"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.user_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.time_updated"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.pop"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.site"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_router_name"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_router_asn"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.cloud_router_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.nat_capable"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.dnat_capable"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.zone"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.vlan"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnsResult.ResourceName, "cloud_connections.0.subscription_term"),
				),
			},
		},
	})

}
