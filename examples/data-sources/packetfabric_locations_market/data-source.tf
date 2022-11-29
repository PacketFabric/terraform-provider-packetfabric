data "packetfabric_locations_market" "locations_market_1" {
  provider = packetfabric
}

output "packetfabric_locations_market" {
  value = data.packetfabric_locations_market.locations_market_1
}