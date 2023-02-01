data "packetfabric_cloud_router_bgp_session" "bgp_session_crbs" {
  provider = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_ipsec.crc_1.id
}
output "packetfabric_cloud_router_bgp_session_crbs" {
  value = data.packetfabric_cloud_router_bgp_session.bgp_session_crbs
}