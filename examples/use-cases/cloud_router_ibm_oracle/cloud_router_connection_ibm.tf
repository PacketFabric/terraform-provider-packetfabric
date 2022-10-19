# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_ibm" "crc_1" {
  provider         = packetfabric
  description      = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  circuit_id       = packetfabric_cloud_router.cr.id
  account_uuid     = var.pf_account_uuid
  ibm_account_id   = var.ibm_account_id
  ibm_bgp_asn      = var.ibm_bgp_asn
  ibm_bgp_cer_cidr = var.ibm_bgp_cer_cidr
  ibm_bgp_ibm_cidr = var.ibm_bgp_ibm_cidr
  pop              = var.pf_crc_pop1
  zone             = var.pf_crc_zone1
  maybe_nat        = var.pf_crc_maybe_nat
  speed            = var.pf_crc_speed
}

# From the IBM side: Accept the connection.
# Vote for 3978 New resource to accept a direct link creation request: ibm_dl_gateway_accept
# https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3978

# From the IBM side: Add a virtual connection to your IBM virtual private cloud (VPC)


# From the IBM side: Set up VRF


# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_ibm.crc_1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.ibm_bgp_asn
  orlonger       = var.pf_crbs_orlonger
  remote_address = "TBD" # IBM side
  l3_address     = "TBD" # PF side
}
output "packetfabric_cloud_router_bgp_session_crbs_1" {
  value = packetfabric_cloud_router_bgp_session.crbs_1
}

# Configure BGP Prefix is mandatory to setup the BGP session correctly
resource "packetfabric_cloud_router_bgp_prefixes" "crbp_1" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_1.id
  prefixes {
    prefix = var.oracle_subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
    order  = 0
  }
  prefixes {
    prefix = var.ibm_vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
    order  = 0
  }
}
data "packetfabric_cloud_router_bgp_prefixes" "bgp_prefix_crbp_1" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_1.id
}
output "packetfabric_bgp_prefix_crbp_1" {
  value = data.packetfabric_cloud_router_bgp_prefixes.bgp_prefix_crbp_1
}