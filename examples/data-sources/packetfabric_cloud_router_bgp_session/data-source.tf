data "packetfabric_cloud_router_bgp_session" "all_cr_bgp_sessions" {
  provider = packetfabric
}
output "packetfabric_cloud_router_bgp_session" {
  value = data.packetfabric_cloud_router_bgp_session.all_cr_bgp_sessions
}