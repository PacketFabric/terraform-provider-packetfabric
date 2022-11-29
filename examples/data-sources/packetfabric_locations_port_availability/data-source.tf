data "packetfabric_locations_port_availability" "port_availabilty_dal_1" {
  provider = packetfabric
  pop      = "DAL1"
}
output "packetfabric_locations_port_availability_dal_1" {
  value = data.packetfabric_locations_port_availability.port_availabilty_dal_1
}

# retreive the zone automatically for packetfabric_port
locals {
  zones_pop_dal = toset([for each in data.packetfabric_locations_port_availability.port_availabilty_dal_1.ports_available[*] : each.zone if each.media == "LX"])
}
output "packetfabric_locations_port_availability_dal_1_single_zone" {
  value = tolist(local.zones_pop_dal)[0]
}