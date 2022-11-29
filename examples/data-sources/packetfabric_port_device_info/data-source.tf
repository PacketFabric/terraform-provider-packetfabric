data "packetfabric_port_device_info" "port_device_info_1" {
  provider = packetfabric
  port_circuit_id = var.pf_port_circuit_id
}

output "packetfabric_port_device_info" {
  value = data.packetfabric_port_device_info.port_device_info_1
}
