data "packetfabric_cloud_router" "cr1" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr1.id
}
output "packetfabric_cloud_router" {
  value = data.packetfabric_cloud_router.cr1
}