resource "packetfabric_cloud_router" "cr1" {
  provider     = packetfabric
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_oracle" "crc6" {
  provider     = packetfabric
  description  = var.pf_crc_description
  circuit_id   = packetfabric_cloud_router.cr1.id
  account_uuid = var.pf_account_uuid
  region       = var.pf_crc_oracle_region
  vc_ocid      = var.pf_crc_oracle_vc_ocid
  pop          = var.pf_crc_pop
  zone         = var.pf_crc_zone
  maybe_nat    = var.pf_crc_maybe_nat
}

output "packetfabric_cloud_router" {
  value = packetfabric_cloud_router.cr1
}

output "packetfabric_cloud_router_connection_oracle" {
  value = packetfabric_cloud_router_connection_oracle.crc6
}
