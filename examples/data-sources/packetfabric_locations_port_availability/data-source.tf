data "packetfabric_locations_port_availability" "port_availabilty_dal_1" {
  provider = packetfabric
  pop = "DAL"
}

output "packetfabric_locations_port_availability" {
  value = data.packetfabric_locations_port_availability.port_availabilty_dal_1
}
