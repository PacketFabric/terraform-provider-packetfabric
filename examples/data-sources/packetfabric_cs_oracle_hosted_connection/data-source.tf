data "packetfabric_cs_oracle_hosted_connection" "current" {
  provider         = packetfabric
  cloud_circuit_id = "PF-AP-LAX1-1002"
}
output "packetfabric_cs_oracle_hosted_connection" {
  value = data.packetfabric_cs_oracle_hosted_connection.current
}