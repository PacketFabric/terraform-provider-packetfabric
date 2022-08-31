data "packetfabric_port" "ports_all" {
  provider = packetfabric
}

output "packetfabric_ports_all" {
  value = data.packetfabric_port.ports_all
}