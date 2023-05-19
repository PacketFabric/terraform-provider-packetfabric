//go:build datasource || cloud_router || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCloudRouterConnectionComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, []string{"PF_AWS_ACCOUNT_ID"})

	datasourceCloudConnResult := testutil.DHclCloudConnection()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceCloudConnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "connection_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "port_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "connection_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "port_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "pending_delete"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "deleted"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "speed"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "state"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "account_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "service_class"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "service_provider"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "service_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "description"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_provider_connection_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.vlan_id_pf"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.vlan_id_cust"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.svlan_id_cust"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "user_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "customer_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_provider"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "pop"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "site"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "bgp_state_list"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_router_name"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_router_asn"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_router_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "nat_capable"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "dnat_capable"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "zone"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "vlan"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "desired_nat"),
				),
			},
		},
	})

}
