data "packetfabric_cs_google_dedicated_connection" "current" {
  provider = packetfabric
}

output "packetfabric_cs_google_dedicated_connection" {
  value = data.packetfabric_cs_aws_dedicated_connection.current
}