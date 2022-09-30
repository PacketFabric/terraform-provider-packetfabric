data "packetfabric_cloud_router_connections" "crc" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr1.id
}
output "packetfabric_cloud_router_connections" {
  value = data.packetfabric_cloud_router_connection.crc
}