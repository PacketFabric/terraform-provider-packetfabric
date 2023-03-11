# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_port" "crc_3" {
  provider        = packetfabric
  description     = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_port_circuit_id}"
  circuit_id      = packetfabric_cloud_router.cr.id
  # port_circuit_id = var.pf_crc_port_circuit_id # in case you prefer to use an existing port
  port_circuit_id = packetfabric_port.port_1.id
  vlan            = var.pf_crc_vlan3
  untagged        = var.pf_crc_untagged
  speed           = var.pf_crc_speed3
}

# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_3" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_port.crc_3.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.oracle_peer_asn
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.on_premise_bgp_peering_prefix # Customer On-premise Router Peer IP side
  l3_address     = var.pf_side_bgp_peering_prefix    # PF side
  prefixes {
    prefix = var.on_premise_cidr1
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.oracle_subnet_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
  prefixes {
    prefix = var.aws_vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
}
# output "packetfabric_cloud_router_bgp_session_crbs_3" {
#   value = packetfabric_cloud_router_bgp_session.crbs_3
# }
