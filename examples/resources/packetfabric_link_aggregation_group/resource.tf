# Create a PacketFabric interfaces
resource "packetfabric_port" "port_1a" {
  provider          = packetfabric
  autoneg           = true
  description       = "hello world-a"
  media             = "LX"
  nni               = false
  pop               = "SEA2"
  speed             = "1Gbps"
  subscription_term = 1
  zone              = "A"
  labels            = ["terraform", "dev"]
}
output "packetfabric_port_1a" {
  value = packetfabric_port.port_1a
}

## 2nd port in the same location same zone to create a LAG
resource "packetfabric_port" "port_1b" {
  provider          = packetfabric
  autoneg           = true
  description       = "hello world-b"
  media             = "LX"
  nni               = false
  pop               = "SEA2"
  speed             = "1Gbps"
  subscription_term = 1
  zone              = "B"
  labels            = ["terraform", "dev"]
}
output "packetfabric_port_1b" {
  value = packetfabric_port.port_1b
}

resource "packetfabric_link_aggregation_group" "lag_1" {
  provider    = packetfabric
  description = "hello world"
  interval    = "fast" # or slow
  members     = [packetfabric_port.port_1a.id, packetfabric_port.port_1b.id]
  pop         = "SEA2"
  labels      = ["terraform", "dev"]
}
output "packetfabric_link_aggregation_group" {
  value = packetfabric_link_aggregation_group.lag_1
}