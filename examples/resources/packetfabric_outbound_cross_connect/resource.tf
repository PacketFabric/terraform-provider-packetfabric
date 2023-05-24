# Create a PacketFabric interface
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

resource "packetfabric_document" "loa1" {
  provider        = packetfabric
  document        = "letter-of-authorization-PF-AP-LAB8-3339359.pdf"
  type            = "loa"
  description     = "My LOA"
  port_circuit_id = "PF-AP-LAB8-3339359"
}

# Create Cross Connect
resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
  provider      = packetfabric
  description   = "hello world"
  document_uuid = packetfabric_document.loa1.id
  port          = packetfabric_port.port_1.id
  site          = local.pf_port_site1
}