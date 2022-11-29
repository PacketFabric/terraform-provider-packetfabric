data "packetfabric_locations_regions" "locations_regions_1" {
  provider = packetfabric
}
output "packetfabric_locations_regions" {
  value = data.packetfabric_locations_regions.locations_regions_1
}