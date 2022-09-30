data "packetfabric_cloud_router_connection_google" "crc2" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr1.id
}
output "packetfabric_cloud_router_connection_google" {
  value = data.packetfabric_cloud_router_connection_google.crc2
}