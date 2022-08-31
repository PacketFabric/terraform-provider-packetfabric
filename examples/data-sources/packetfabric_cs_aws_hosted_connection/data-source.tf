data "packetfabric_cs_aws_hosted_connection" "current" {
  provider = packetfabric
}

output "packetfabric_cs_aws_hosted_connection_data" {
  value = data.packetfabric_cs_aws_hosted_connection.current
}