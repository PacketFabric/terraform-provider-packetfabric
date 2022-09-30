resource "packetfabric_cloud_router_connection_ipsec" "crc_3" {
  provider                     = packetfabric
  description                  = var.pf_crc_description
  circuit_id                   = packetfabric_cloud_router.cr.id
  account_uuid                 = var.pf_account_uuid
  pop                          = var.pf_crc_pop3
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