data "packetfabric_cloud_router_connection_aws" "crc1" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr1.id
}
output "packetfabric_cloud_router_connection_aws" {
  value = data.packetfabric_cloud_router_connection_aws.crc1
}