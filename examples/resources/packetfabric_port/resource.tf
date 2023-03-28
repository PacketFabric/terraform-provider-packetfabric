resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  enabled           = true
  autoneg           = true
  description       = "hello world"
  media             = "LX"
  nni               = false
  pop               = "SEA2"
  speed             = "1Gbps"
  subscription_term = 1
  zone              = "A"
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