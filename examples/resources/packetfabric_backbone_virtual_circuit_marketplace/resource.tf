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

resource "packetfabric_backbone_virtual_circuit_marketplace" "vc_marketplace_conn1" {
  provider    = packetfabric
  description = "hello world"
  routing_id  = "PD-WUY-9VB0"
  market      = "HOU"
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