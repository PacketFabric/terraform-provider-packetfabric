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

# remove the resource to enable the port
resource "packetfabric_port_disable" "disable_port_1" {
  provider        = packetfabric
  port_circuit_id = packetfabric_port.port_1.id
}

data "packetfabric_port" "ports_all" {
  provider   = packetfabric
  depends_on = [packetfabric_port_disable.disable_port_1]
}

locals {
  port_1_admin_status = toset([for each in data.packetfabric_port.ports_all.interfaces[*] : each.admin_status if each.port_circuit_id == packetfabric_port.port_1.id])
  port_1_full_details = toset([for each in data.packetfabric_port.ports_all.interfaces[*] : each if each.port_circuit_id == packetfabric_port.port_1.id])
}
output "packetfabric_port_1_admin_status" {
  value = local.port_1_admin_status
}
output "packetfabric_port_1_full_details" {
  value = local.port_1_full_details
}
