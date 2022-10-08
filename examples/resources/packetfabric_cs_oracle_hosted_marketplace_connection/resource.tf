
resource "packetfabric_cs_oracle_hosted_marketplace_connection" "cs_conn1_marketplace_oracle" {
  provider     = packetfabric
  description  = var.description
  account_uuid = var.pf_account_uuid
  vc_ocid      = var.oracle_vc_ocid
  region       = var.oracle_region
  routing_id   = var.routing_id
  market       = var.market
  pop          = var.pf_cs_pop
}

output "packetfabric_cs_oracle_hosted_marketplace_connection" {
  value     = packetfabric_cs_oracle_hosted_marketplace_connection.cs_conn1_marketplace_oracle
  sensitive = true
}