//go:build datasource || cloud_router || all || smoke

package provider

import (
	"testing"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/testutil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceBgpSessionComputedRequiredFields(t *testing.T) {
	testutil.PreCheck(t, nil)

	datasourceBgpSessionResult := testutil.DHclDatasourceBgpSession()

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: datasourceBgpSessionResult.Hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "circuit_id"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "connection_id"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.bgp_settings_uuid"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.address_family"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.remote_address"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.remote_asn"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.multihop_ttl"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.local_preference"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.community"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.as_prepend"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.med"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.bfd_interval"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.bfd_multiplier"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.disabled"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.prefixes.0.bgp_prefix_uuid"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.prefixes.0.prefix"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.prefixes.0.match_type"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.prefixes.0.as_prepend"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.prefixes.0.med"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.prefixes.0.local_preference"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.prefixes.0.type"),
					resource.TestCheckResourceAttrSet(datasourceBgpSessionResult.ResourceName, "bgp_sessions.0.prefixes.0.order"),
				),
			},
		},
	})

}
