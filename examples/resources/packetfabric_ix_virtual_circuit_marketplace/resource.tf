resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  autoneg           = var.pf_port_autoneg
  description       = var.pf_description
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
}

resource "packetfabric_ix_virtual_circuit_marketplace" "ix_marketplace_conn1" {
  provider    = packetfabric
  description = var.pf_description
  routing_id  = var.pf_routing_id
  market      = var.pf_market
  asn         = var.pf_asn_ix
  interface {
    port_circuit_id = packetfabric_port.port_1.id
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  bandwidth {
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

output "packetfabric_ix_virtual_circuit_marketplace" {
  value = packetfabric_ix_virtual_circuit_marketplace.ix_marketplace_conn1
}