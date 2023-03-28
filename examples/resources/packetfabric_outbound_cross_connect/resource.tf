# Create a PacketFabric interfaces
resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  autoneg           = true
  description       = "hello world"
  media             = "LX"
  nni               = false
  pop               = "SEA2"
  speed             = "1Gbps"
  subscription_term = 1
  zone              = "A"
}
output "packetfabric_port_1" {
  value = packetfabric_port.port_1
}

### Get the site filtering on the pop using packetfabric_locations

# List PacketFabric locations
data "packetfabric_locations" "main" {
  provider = packetfabric
}

locals {
  all_locations = data.packetfabric_locations.main.locations[*]
  helper_map = { for val in local.all_locations :
  val["pop"] => val }
  pf_port_site1 = local.helper_map["SEA2"]["site_code"]
}
output "pf_port_site1" {
  value = local.pf_port_site1
}

# Create Cross Connect
resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
  provider      = packetfabric
  description   = "hello world"
  document_uuid = var.pf_document_uuid1
  port          = packetfabric_port.port_1.id
  site          = local.pf_port_site1
}
output "packetfabric_outbound_cross_connect1" {
  value = packetfabric_outbound_cross_connect.crossconnect_1
}