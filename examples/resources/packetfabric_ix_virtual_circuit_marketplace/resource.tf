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
  labels            = ["terraform", "dev"]
}

resource "packetfabric_ix_virtual_circuit_marketplace" "ix_marketplace_conn1" {
  provider    = packetfabric
  description = "hello world"
  routing_id  = "PD-WUY-9VB0"
  market      = "HOU"
  asn         = 64545
  interface {
    port_circuit_id = packetfabric_port.port_1.id
    untagged        = false
    vlan            = 100
  }
  bandwidth {
    longhaul_type     = "dedicated"
    speed             = "1Gbps"
    subscription_term = 1
  }
}