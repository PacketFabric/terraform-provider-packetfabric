resource "packetfabric_cs_azure_hosted_marketplace_connection" "cs_conn1_marketplace_azure" {
  provider          = packetfabric
  description       = "hello world"
  azure_service_key = var.azure_service_key
  routing_id        = "PD-WUY-9VB0"
  market            = "HOU"
  speed             = "10Gbps" # will be deprecated
}

output "packetfabric_cs_azure_hosted_marketplace_connection" {
  sensitive = true
  value     = packetfabric_cs_azure_hosted_marketplace_connection.cs_conn1_marketplace_azure
}