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

resource "packetfabric_port_status" "change_port_1_status" {
  provider        = packetfabric
  port_circuit_id = packetfabric_port.port_1.id
  enabled         = false # disabling port
}
output "packetfabric_port_status" {
  value = packetfabric_port_status.change_port_1_status
}