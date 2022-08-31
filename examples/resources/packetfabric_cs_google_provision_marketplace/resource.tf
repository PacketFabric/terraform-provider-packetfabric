resource "packetfabric_cs_google_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider                    = packetfabric
  description                 = var.description
  account_uuid                = var.pf_account_uuid
  routing_id                  = var.routing_id
  market                      = var.market
  speed                       = var.pf_cs_speed
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = var.google_vlan_attachment_name
  pop                         = var.pf_cs_pop

}

resource "packetfabric_cs_google_provision_marketplace" "accept_request_google" {
  provider        = packetfabric
  description     = var.description
  port_circuit_id = var.port_circuit_id_marketplace
  vc_request_uuid = packetfabric_cs_google_hosted_marketplace_connection.cs_marketplace_conn1.id
  vlan            = var.pf_cs_vlan
}