# Create a PacketFabric interfaces
resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  autoneg           = var.pf_port_autoneg
  description       = var.description
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

### Get the site filtering on the pop using packetfabric_locations

# List PacketFabric locations
data "packetfabric_locations" "main" {
  provider = packetfabric
}

locals {
  all_locations = data.packetfabric_locations.main.locations[*]
  helper_map = { for val in local.all_locations :
  val["pop"] => val }
  pf_port_site1 = local.helper_map["${var.pf_port_pop1}"]["site_code"]
}
output "pf_port_site1" {
  value = local.pf_port_site1
}

# Create Cross Connect
resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
  provider      = packetfabric
  description   = var.description
  document_uuid = var.pf_document_uuid1
  port          = packetfabric_port.port_1.id
  site          = local.pf_port_site1
}
output "packetfabric_outbound_cross_connect1" {
  value = packetfabric_outbound_cross_connect.crossconnect_1
}