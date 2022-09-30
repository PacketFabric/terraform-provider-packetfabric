data "packetfabric_cloud_router_connection_ipsec" "crc3" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr1.id
}
output "packetfabric_cloud_router_connection_ipsec" {
  value = data.packetfabric_cloud_router_connection_ipsec.crc3
}