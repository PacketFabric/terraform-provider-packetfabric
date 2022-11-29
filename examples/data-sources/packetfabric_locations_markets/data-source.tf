data "packetfabric_locations_markets" "locations_market_1" {
  provider = packetfabric
}

output "packetfabric_locations_markets" {
  value = data.packetfabric_locations_markets.locations_market_1
}