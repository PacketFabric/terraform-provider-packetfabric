
resource "packetfabric_cs_oracle_hosted_marketplace_connection" "cs_conn1_marketplace_oracle" {
  provider    = packetfabric
  description = var.pf_description
  vc_ocid     = var.oracle_vc_ocid
  region      = var.oracle_region
  routing_id  = var.pf_routing_id
  market      = var.pf_market
  pop         = var.pf_cs_pop
}

output "packetfabric_cs_oracle_hosted_marketplace_connection" {
  value     = packetfabric_cs_oracle_hosted_marketplace_connection.cs_conn1_marketplace_oracle
  sensitive = true
}