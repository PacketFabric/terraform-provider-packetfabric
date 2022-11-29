data "packetfabric_port_vlans" "port_vlans_1" {
  provider        = packetfabric
  port_circuit_id = var.pf_port_circuit_id
}
output "packetfabric_port_vlans" {
  value = data.packetfabric_port_vlans.port_vlans_1
}
