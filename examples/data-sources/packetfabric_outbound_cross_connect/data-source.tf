data "packetfabric_outbound_cross_connect" "crossconnect_1" {
  provider = packetfabric
}
output "packetfabric_outbound_cross_connect" {
  value = data.packetfabric_outbound_cross_connect.crossconnect_1
}