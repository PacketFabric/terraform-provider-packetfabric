data "packetfabric_cs_oracle_hosted_connection" "current" {
  provider = packetfabric
}

output "packetfabric_cs_oracle_hosted_connection" {
  value = data.packetfabric_cs_oracle_hosted_connection.current
}