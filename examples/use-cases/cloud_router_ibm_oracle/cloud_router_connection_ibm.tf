# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_ibm" "crc_1" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  circuit_id  = packetfabric_cloud_router.cr.id
  ibm_bgp_asn = packetfabric_cloud_router.cr.asn
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  maybe_nat   = var.pf_crc_maybe_nat
  speed       = var.pf_crc_speed
}

# From the IBM side: Accept the connection.
# Vote for 3978 New resource to accept a direct link creation request: ibm_dl_gateway_accept
# https://github.com/IBM-Cloud/terraform-provider-ibm/issues/3978

data "ibm_dl_gateway" "direct_link_gw" {
  provider = ibm
  name     = "${var.tag_name}-${random_pet.name.id}"
}

output "ibm_dl_gateway" {
  value = data.ibm_dl_gateway.direct_link_gw
}

# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_ibm.crc_1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.ibm_bgp_asn
  orlonger       = var.pf_crbs_orlonger
  remote_address = data.ibm_dl_gateway.direct_link_gw.bgp_base_cidr # IBM side
  l3_address     = data.ibm_dl_gateway.direct_link_gw.bgp_cer_cidr  # PF side
  prefixes {
    prefix = var.oracle_subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.ibm_vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
}
output "packetfabric_cloud_router_bgp_session_crbs_1" {
  value = packetfabric_cloud_router_bgp_session.crbs_1
}

# From the IBM side: Add a virtual connection to your IBM virtual private cloud (VPC)
resource "ibm_dl_virtual_connection" "dl_gateway_vc" {
  provider   = ibm
  gateway    = data.ibm_dl_gateway.direct_link_gw.id
  name       = "${var.tag_name}-${random_pet.name.id}"
  type       = "vpc"
  network_id = ibm_is_vpc.vpc_1.id
}

# From the IBM side: Set up VRF
# TBD