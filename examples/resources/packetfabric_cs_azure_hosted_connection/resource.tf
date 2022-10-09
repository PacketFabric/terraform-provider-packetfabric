resource "packetfabric_cs_azure_hosted_connection" "cs_conn1_hosted_azure" {
  provider          = packetfabric
  description       = var.pf_description
  account_uuid      = var.pf_account_uuid
  azure_service_key = var.azure_service_key
  port              = var.pf_port
  speed             = var.pf_cs_speed # will be deprecated
  vlan_private      = var.pf_cs_vlan_private
  vlan_microsoft    = var.pf_cs_vlan_microsoft
}

output "packetfabric_cs_azure_hosted_connection" {
  sensitive = true
  value     = packetfabric_cs_azure_hosted_connection.cs_conn1_hosted_azure
}