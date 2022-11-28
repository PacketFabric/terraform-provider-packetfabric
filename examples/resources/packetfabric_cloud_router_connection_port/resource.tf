resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = var.pf_cr_asn
  name     = var.pf_cr_name
  capacity = var.pf_cr_capacity
  regions  = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_port" "crc7" {
  provider        = packetfabric
  description     = var.pf_crc_description
  circuit_id      = packetfabric_cloud_router.cr1.id
  port_circuit_id = var.pf_crc_port_circuit_id
  vlan            = var.pf_crc_vlan
  untagged        = var.pf_crc_untagged
  speed           = var.pf_crc_speed
  is_public       = var.pf_crc_is_public
  maybe_nat       = var.pf_crc_maybe_nat
}

output "packetfabric_cloud_router_connection_port" {
  value = packetfabric_cloud_router_connection_port.crc7
}
