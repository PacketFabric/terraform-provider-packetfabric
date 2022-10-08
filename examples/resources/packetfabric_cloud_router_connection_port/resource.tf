resource "packetfabric_cloud_router" "cr1" {
  provider     = packetfabric
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_port" "crc7" {
  provider        = packetfabric
  description     = var.pf_crc_description
  circuit_id      = packetfabric_cloud_router.cr1.id
  account_uuid    = var.pf_account_uuid
  port_circuit_id = var.port_circuit_id
  vlan            = var.vlan
  untagged        = var.untagged
  speed           = var.pf_crspeedc_zone
  is_public       = var.is_public
  maybe_nat       = var.pf_crc_maybe_nat
}

output "packetfabric_cloud_router" {
  value = packetfabric_cloud_router.cr1
}

output "packetfabric_cloud_router_connection_port" {
  value = packetfabric_cloud_router_connection_port.crc7
}
