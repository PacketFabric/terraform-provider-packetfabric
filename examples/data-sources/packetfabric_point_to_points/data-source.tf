data "packetfabric_point_to_points" "ptp" {
  provider = packetfabric
}
output "packetfabric_point_to_points" {
  value = data.packetfabric_point_to_points.ptp
}