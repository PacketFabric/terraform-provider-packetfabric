resource "packetfabric_cs_oracle_hosted_connection" "cs_conn1_hosted_oracle" {
  provider    = packetfabric
  description = "hello world"
  vc_ocid     = var.oracle_vc_ocid
  region      = "us-ashburn-1"
  port        = packetfabric_port.port_1.id
  pop         = "A"
  zone        = "BOS1"
  vlan        = 102
  labels      = ["terraform", "dev"]
}

output "packetfabric_cs_oracle_hosted_connection" {
  value     = packetfabric_cs_oracle_hosted_connection.cs_conn1_hosted_oracle
  sensitive = true
}