resource "packetfabric_cloud_router" "cr1" {
  provider     = packetfabric
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_google" "crc2" {
  provider                    = packetfabric
  description                 = var.pf_crc_description
  circuit_id                  = packetfabric_cloud_router.cr1.id
  account_uuid                = var.pf_account_uuid
  google_pairing_key          = var.pf_crc_google_pairing_key
  google_vlan_attachment_name = var.pf_crc_google_vlan_attachment_name
  pop                         = var.pf_crc_pop
  speed                       = var.pf_crc_speed
  maybe_nat                   = var.pf_crc_maybe_nat
}

output "packetfabric_cloud_router" {
  value = packetfabric_cloud_router.cr1
}

output "packetfabric_cloud_router_connection_google" {
  value = packetfabric_cloud_router_connection_google.crc2
}
