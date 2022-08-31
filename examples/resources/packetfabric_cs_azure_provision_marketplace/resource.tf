resource "packetfabric_cs_azure_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider          = packetfabric
  description       = var.description
  account_uuid      = var.pf_account_uuid
  azure_service_key = var.azure_service_key
  routing_id        = var.routing_id
  market            = var.market
  speed             = var.pf_cs_speed # will be deprecated
}

output "packetfabric_cs_azure_hosted_marketplace_connection" {
  sensitive = true
  value     = packetfabric_cs_azure_hosted_marketplace_connection.cs_marketplace_conn1
}

resource "packetfabric_cs_azure_provision_marketplace" "accept_request_azure" {
  provider        = packetfabric
  description     = var.description
  port_circuit_id = var.port_circuit_id_marketplace
  vc_request_uuid = packetfabric_cs_azure_hosted_marketplace_connection.cs_marketplace_conn1.id
  vlan_private    = var.pf_cs_vlan_private
  vlan_microsoft  = var.pf_cs_vlan_microsoft
}
