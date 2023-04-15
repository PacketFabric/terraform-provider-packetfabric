package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceCloudConnComputedRequiredFields(t *testing.T) {

	testutil.SkipIfEnvNotSet(t)

	datasourceCloudConnResult := testutil.DHclDatasourceCloudConn()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testutil.PreCheck(t, nil)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceCloudConnResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.vlan_id_pf"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.vlan_id_cust"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.svlan_id_cust"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.aws_region"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.aws_hosted_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.aws_connection_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.aws_account_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.google_pairing_key"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.google_vlan_attachment_name"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.vlan_id_private"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.vlan_id_microsoft"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.azure_service_key"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.azure_service_tag"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.azure_connection_type"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.oracle_region"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.vc_ocid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.port_cross_connect_ocid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.port_compartment_ocid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.account_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.gateway_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.port_id"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.name"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.bgp_asn"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.bgp_cer_cidr"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.bgp_ibm_cidr"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.public_ip"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.cloud_settings.0.nat_public_ip"),

					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.bgp_state_list.0.bgp_settings_uuid"),
					resource.TestCheckResourceAttrSet(datasourceCloudConnResult.ResourceName, "cloud_connections.0.bgp_state_list.0.bgp_state"),
				),
			},
		},
	})

}
