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

resource "packetfabric_port" "port_2" {
  provider          = packetfabric
  autoneg           = var.pf_port_autoneg
  description       = var.pf_description
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop2
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone2
}

resource "packetfabric_backbone_virtual_circuit" "vc1" {
  provider    = packetfabric
  description = var.pf_description
  epl         = false
  interface_a {
    port_circuit_id = packetfabric_port.port_1.id
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  interface_z {
    port_circuit_id = packetfabric_port.port_2.id
    untagged        = false
    vlan            = var.pf_vc_vlan2
  }
  bandwidth {
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}
