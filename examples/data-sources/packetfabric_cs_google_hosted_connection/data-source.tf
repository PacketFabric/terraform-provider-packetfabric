data "packetfabric_cs_google_hosted_connection" "current" {
  provider = packetfabric
}

output "packetfabric_cs_google_hosted_connection_data" {
  value = data.packetfabric_cs_azure_hosted_connection.current
}