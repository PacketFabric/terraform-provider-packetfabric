data "packetfabric_cloud_routers" "cr1" {
  provider = packetfabric
}
output "packetfabric_cloud_routers" {
  value = data.packetfabric_cloud_routers.cr1
}