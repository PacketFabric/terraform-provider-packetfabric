resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  autoneg           = var.pf_port_autoneg
  description       = var.pf_description
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
}
output "packetfabric_port_1" {
  value = packetfabric_port.port_1
}