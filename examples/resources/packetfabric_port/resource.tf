resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  enabled           = true
  autoneg           = var.pf_port_autoneg
  description       = var.pf_description
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
  labels            = ["terraform", "dev"]
}
output "packetfabric_port_1" {
  value = packetfabric_port.port_1
}
data "packetfabric_port" "ports_all" {
  provider   = packetfabric
  depends_on = [packetfabric_port.port_1]
}
locals {
  port_1_details = toset([for each in data.packetfabric_port.ports_all.interfaces[*] : each if each.port_circuit_id == packetfabric_port.port_1.id])
}
output "packetfabric_port_1_details" {
  value = local.port_1_details
}