resource "packetfabric_cloud_router" "cr1" {
  provider     = packetfabric
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_ibm" "crc5" {
  provider       = packetfabric
  description    = var.pf_crc_description
  circuit_id     = packetfabric_cloud_router.cr1.id
  account_uuid   = var.pf_account_uuid
  ibm_account_id = var.pf_crc_ibm_account_id
  ibm_bgp_asn    = packetfabric_cloud_router.cr1.asn
  pop            = var.pf_crc_pop
  zone           = var.pf_crc_zone
  maybe_nat      = var.pf_crc_maybe_nat
  speed          = var.pf_crc_speed
}

output "packetfabric_cloud_router_connection_ibm" {
  value = packetfabric_cloud_router_connection_ibm.crc5
}
