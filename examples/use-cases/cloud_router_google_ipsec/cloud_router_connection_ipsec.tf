resource "packetfabric_cloud_router_connection_ipsec" "crc_ipsec" {
  provider                     = packetfabric
  description                  = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  labels                       = var.pf_labels
  circuit_id                   = packetfabric_cloud_router.cr.id
  pop                          = var.pf_crc_pop2
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

resource "packetfabric_cloud_router_bgp_session" "crbs_ipsec" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_ipsec.crc_ipsec.id
  address_family = var.pf_crbs_af
  remote_asn     = var.vpn_side_asn2
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.vpn_remote_address # On-Prem side
  l3_address     = var.vpn_l3_address     # PF side
  prefixes {
    prefix = var.google_subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.ipsec_subnet_cidr2
    type   = "in" # Allowed Prefixes from Cloud
  }
}
# output "packetfabric_cloud_router_bgp_session_crbs_ipsec" {
#   value = packetfabric_cloud_router_bgp_session.crbs_ipsec
# }
