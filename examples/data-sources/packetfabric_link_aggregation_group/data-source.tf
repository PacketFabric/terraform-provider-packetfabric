data "packetfabric_link_aggregation_group" "lag_1" {
  provider       = packetfabric
  lag_circuit_id = packetfabric_link_aggregation_group.lag_1.id
}
output "packetfabric_link_aggregation_group" {
  value = data.packetfabric_link_aggregation_group.lag_1
}
