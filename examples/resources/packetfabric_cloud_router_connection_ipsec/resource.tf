resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = var.pf_cr_asn
  name     = var.pf_cr_name
  capacity = var.pf_cr_capacity
  regions  = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_ipsec" "crc3" {
  provider                     = packetfabric
  description                  = var.pf_crc_description
  circuit_id                   = packetfabric_cloud_router.cr1.id
  pop                          = var.pf_crc_pop
  speed                        = var.pf_crc_speed
  gateway_address              = var.pf_crc_gateway_address
  ike_version                  = var.pf_crc_ike_version
  phase1_authentication_method = var.pf_crc_phase1_authentication_method
  phase1_group                 = var.pf_crc_phase1_group
  phase1_encryption_algo       = var.pf_crc_phase1_encryption_algo
  phase1_authentication_algo   = var.pf_crc_phase1_authentication_algo
  phase1_lifetime              = var.pf_crc_phase1_lifetime
  phase2_pfs_group             = var.pf_crc_phase2_pfs_group
  phase2_encryption_algo       = var.pf_crc_phase2_encryption_algo
  phase2_authentication_algo   = var.pf_crc_phase2_authentication_algo
  phase2_lifetime              = var.pf_crc_phase2_lifetime
  shared_key                   = var.pf_crc_shared_key
}

output "packetfabric_cloud_router_connection_ipsec" {
  value = packetfabric_cloud_router_connection_ipsec.crc3
}
