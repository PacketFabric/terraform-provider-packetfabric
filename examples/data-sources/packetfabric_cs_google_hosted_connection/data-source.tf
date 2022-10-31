data "packetfabric_cs_google_hosted_connection" "current" {
  provider         = packetfabric
  cloud_circuit_id = "PF-AP-LAX1-1002"
}

output "packetfabric_cs_google_hosted_connection_data" {
  value = data.packetfabric_cs_azure_hosted_connection.current
}