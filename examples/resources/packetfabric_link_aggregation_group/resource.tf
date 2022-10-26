# Create a PacketFabric interfaces
resource "packetfabric_port" "port_1a" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  autoneg           = var.pf_port_autoneg
  description       = "${var.pf_description}-a"
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
}
output "packetfabric_port_1a" {
  value = packetfabric_port.port_1a
}

## 2nd port in the same location same zone to create a LAG
resource "packetfabric_port" "port_1b" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  autoneg           = var.pf_port_autoneg
  description       = "${var.pf_description}-b"
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
}
output "packetfabric_port_1b" {
  value = packetfabric_port.port_1b
}

resource "packetfabric_link_aggregation_group" "lag_1" {
  provider    = packetfabric
  description = var.pf_description
  interval    = "fast" # or slow
  members     = [packetfabric_port.port_1a.id, packetfabric_port.port_1b.id]
  #members = [packetfabric_port.port_1a.id]
  pop = var.pf_port_pop1
}

data "packetfabric_link_aggregation_group" "lag_1" {
  provider            = packetfabric
  lag_port_circuit_id = packetfabric_link_aggregation_group.lag_1.id
}

output "packetfabric_link_aggregation_group" {
  value = data.packetfabric_link_aggregation_group.lag_1
}