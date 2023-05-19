data "packetfabric_cloud_router_connection" "crc1" {
  provider      = packetfabric
  circuit_id    = packetfabric_cloud_router.cr1.id
  connection_id = packetfabric_cloud_router_connection_aws.crc1.id
}
output "packetfabric_cloud_router_connection_crc1" {
  value = data.packetfabric_cloud_router_connection.crc1
}