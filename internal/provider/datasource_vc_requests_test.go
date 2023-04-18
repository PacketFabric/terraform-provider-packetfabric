package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceVcRequestsComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceVcRequestsResult := testutil.DHclDataSourceVcRequests()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, []string{
				testutil.PF_DTS_TYPE_KEY})
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceVcRequestsResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "type"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.vc_request_uuid"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.vc_circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.status"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.request_type"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.text"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.rate_limit_in"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.rate_limit_out"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.service_name"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.allow_untagged_z"),
					resource.TestCheckResourceAttrSet(datasourceVcRequestsResult.ResourceName, "vc_requests.0.flex_bandwidth_id"),
				),
			},
		},
	})

}
