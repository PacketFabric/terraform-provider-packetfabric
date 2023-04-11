resource "packetfabric_cs_google_hosted_marketplace_connection" "cs_conn1_marketplace_google" {
  provider                    = packetfabric
  description                 = "hello world"
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = var.google_vlan_attachment_name
  routing_id                  = "PD-WUY-9VB0"
  market                      = "HOU"
  speed                       = "10Gbps"
  pop                         = "BOS1"

}