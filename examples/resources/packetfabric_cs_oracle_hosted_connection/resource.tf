resource "packetfabric_cs_oracle_hosted_connection" "cs_conn1_hosted_oracle" {
  provider    = packetfabric
  description = var.pf_description
  vc_ocid     = var.oracle_vc_ocid
  region      = var.oracle_region
  port        = var.pf_port
  pop         = var.pf_cs_zone
  zone        = var.pf_cs_pop
  vlan        = var.pf_cs_vlan
  src_svlan   = var.pf_cs_src_svlan
}

output "packetfabric_cs_oracle_hosted_connection" {
  value     = packetfabric_cs_oracle_hosted_connection.cs_conn1_hosted_oracle
  sensitive = true
}