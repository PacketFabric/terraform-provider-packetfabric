data "packetfabric_point_to_point" "ptp" {
  provider   = packetfabric
}
output "packetfabric_point_to_point" {
  value = data.packetfabric_point_to_point.ptp
}