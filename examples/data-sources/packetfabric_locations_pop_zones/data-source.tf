data "packetfabric_locations_pop_zones" "locations_pop_zones_DAL_1" {
  provider = packetfabric
  pop      = "DAL1"
}
output "packetfabric_locations_pop_zones" {
  value = data.packetfabric_locations_pop_zones.locations_pop_zones_DAL_1
}