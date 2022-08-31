data "packetfabric_locations" "locations_all" {
  provider = packetfabric
}
output "packetfabric_locations" {
  value = data.packetfabric_locations.locations_all
}