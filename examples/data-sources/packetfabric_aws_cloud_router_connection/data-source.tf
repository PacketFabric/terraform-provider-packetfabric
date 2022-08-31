data "packetfabric_aws_cloud_router_connection" "crc1" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr1.id
}
output "packetfabric_aws_cloud_router_connection" {
  value = data.packetfabric_aws_cloud_router_connection.crc1
}