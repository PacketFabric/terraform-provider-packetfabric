resource "packetfabric_cs_azure_hosted_marketplace_connection" "cs_conn1_marketplace_azure" {
  provider          = packetfabric
  description       = var.pf_description
  account_uuid      = var.pf_account_uuid
  azure_service_key = var.azure_service_key
  routing_id        = var.pf_routing_id
  market            = var.pf_market
  speed             = var.pf_cs_speed # will be deprecated
}

output "packetfabric_cs_azure_hosted_marketplace_connection" {
  sensitive = true
  value     = packetfabric_cs_azure_hosted_marketplace_connection.cs_conn1_marketplace_azure
}