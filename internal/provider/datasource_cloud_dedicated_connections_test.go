//go:build datasource || dedicated_cloud || all

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDedicatedConnsComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	dedicatedConnectionsResult := testutil.DHclDedicatedConnections()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dedicatedConnectionsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.uuid"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.customer_uuid"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.user_uuid"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.service_provider"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.port_type"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.deleted"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.time_updated"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.time_created"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.cloud_circuit_id"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.account_uuid"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.pop"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.site"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.service_class"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.description"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.state"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.speed"),
					resource.TestCheckResourceAttrSet(dedicatedConnectionsResult.ResourceName, "dedicated_connections.0.is_lag"),
				),
			},
		},
	})

}
