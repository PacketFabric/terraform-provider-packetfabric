data "packetfabric_locations_cloud" "cloud_location_1" {
  provider              = packetfabric
  cloud_provider        = "aws"
  cloud_connection_type = "hosted"
}

output "packetfabric_locations_cloud" {
  value = data.packetfabric_locations_cloud.cloud_location_1
}
