data "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
}
output "packetfabric_cloud_router" {
  value = data.packetfabric_cloud_router.cr1
}