resource "packetfabric_cloud_router_connection_google" "crc_2" {
  provider                    = packetfabric
  description                 = var.pf_crc_description
  circuit_id                  = packetfabric_cloud_router.cr.id
  account_uuid                = var.pf_account_uuid
  google_pairing_key          = var.pf_crc_google_pairing_key
  google_vlan_attachment_name = var.pf_crc_google_vlan_attachment_name
  pop                         = var.pf_crc_pop2
  speed                       = var.pf_crc_speed
  maybe_nat                   = var.pf_crc_maybe_nat
}