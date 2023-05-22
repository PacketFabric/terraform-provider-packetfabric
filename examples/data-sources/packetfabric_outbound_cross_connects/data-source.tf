data "packetfabric_outbound_cross_connects" "crossconnects" {
  provider = packetfabric
}
output "packetfabric_outbound_cross_connects" {
  value = data.packetfabric_outbound_cross_connects.crossconnects
}