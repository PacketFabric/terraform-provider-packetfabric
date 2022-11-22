resource "packetfabric_cs_aws_hosted_connection" "cs_conn1_hosted_aws" {
  provider    = packetfabric
  description = var.pf_description
  port        = var.pf_port
  speed       = var.pf_cs_speed
  pop         = var.pf_cs_pop
  vlan        = var.pf_cs_vlan
  zone        = var.pf_cs_zone
}

output "packetfabric_cs_aws_hosted_connection" {
  value = packetfabric_cs_aws_hosted_connection.cs_conn1_hosted_aws
}