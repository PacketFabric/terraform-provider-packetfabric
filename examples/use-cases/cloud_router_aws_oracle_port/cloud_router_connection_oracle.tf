# From the Oracle side: Create a dynamic routing gateway
resource "oci_core_drg" "dyn_routing_gw_1" {
  provider       = oci
  # compartment_id = oci_identity_compartment.compartment_1.id
  compartment_id = var.parent_compartment_id
  display_name   = "${var.tag_name}-${random_pet.name.id}"
}
# output "oci_core_drg" {
#   value = oci_core_drg.dyn_routing_gw_1
# }

data "oci_core_fast_connect_provider_services" "packetfabric_provider" {
  provider       = oci
  # compartment_id = oci_identity_compartment.compartment_1.id
  compartment_id = var.parent_compartment_id
  filter {
    name   = "provider_name"
    values = ["PacketFabric"]
  }
}
# output "oci_core_fast_connect_provider_services" {
#   value = data.oci_core_fast_connect_provider_services.packetfabric_provider
# }

# From the Oracle side: Create a FastConnect connection 
resource "oci_core_virtual_circuit" "fast_connect_1" {
  provider             = oci
  # compartment_id       = oci_identity_compartment.compartment_1.id
  compartment_id       = var.parent_compartment_id
  display_name         = "${var.tag_name}-${random_pet.name.id}"
  region               = var.oracle_region1
  type                 = "PRIVATE"
  gateway_id           = oci_core_drg.dyn_routing_gw_1.id
  bandwidth_shape_name = var.oracle_bandwidth_shape_name
  customer_asn         = var.oracle_peer_asn
  ip_mtu               = "MTU_1500" # or "MTU_9000"
  is_bfd_enabled       = false
  cross_connect_mappings {
    bgp_md5auth_key         = var.oracle_bgp_shared_key
    customer_bgp_peering_ip = var.customer_bgp_peering_prefix
    oracle_bgp_peering_ip   = var.oracle_bgp_peering_prefix
  }
  provider_service_id = data.oci_core_fast_connect_provider_services.packetfabric_provider.fast_connect_provider_services.0.id
  # public_prefixes {
  #     cidr_block = var.virtual_circuit_public_prefixes_cidr_block
  # }
  # routing_policy = var.virtual_circuit_routing_policy
}

# data "oci_core_virtual_circuit" "fast_connect_1" {
#   virtual_circuit_id = oci_core_virtual_circuit.fast_connect_1.id
# }

# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_oracle" "crc_2" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  circuit_id  = packetfabric_cloud_router.cr.id
  region      = var.oracle_region1
  vc_ocid     = oci_core_virtual_circuit.fast_connect_1.id
  pop         = var.pf_crc_pop2
  zone        = var.pf_crc_zone2
}

# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_2" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_oracle.crc_2.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.oracle_peer_asn
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.oracle_bgp_peering_prefix   # Oracle side
  l3_address     = var.customer_bgp_peering_prefix # PF side
  md5            = var.oracle_bgp_shared_key
  prefixes {
    prefix = var.oracle_subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.aws_vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
  prefixes {
    prefix = var.on_premise_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
}
# output "packetfabric_cloud_router_bgp_session_crbs_2" {
#   value     = packetfabric_cloud_router_bgp_session.crbs_2
#   sensitive = true
# }
