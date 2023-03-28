resource "packetfabric_cs_aws_hosted_connection" "cs_conn1_hosted_aws" {
  provider    = packetfabric
  description = "hello world"
  port        = packetfabric_port.port_1.id
  speed       = "10Gbps"
  pop         = "BOS1"
  vlan        = 102
  zone        = "A"
  labels      = ["terraform", "dev"]
}

output "packetfabric_cs_aws_hosted_connection" {
  value = packetfabric_cs_aws_hosted_connection.cs_conn1_hosted_aws
}