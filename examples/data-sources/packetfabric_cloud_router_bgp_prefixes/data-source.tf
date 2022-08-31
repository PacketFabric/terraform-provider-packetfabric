data "packetfabric_cloud_router_bgp_prefixes" "bgp_prefix_crbp" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs.id
}
output "packetfabric_cloud_router_bgp_prefixes" {
  value = data.packetfabric_cloud_router_bgp_prefixes.bgp_prefix_crbp
}
