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
  labels            = sort(["terraform", "dev"])
}

resource "packetfabric_port" "port_2" {
  provider          = packetfabric
  autoneg           = true
  description       = "hello world"
  media             = "LX"
  nni               = false
  pop               = "NYC5"
  speed             = "1Gbps"
  subscription_term = 1
  zone              = "A"
  labels            = sort(["terraform", "dev"])
}

resource "packetfabric_backbone_virtual_circuit" "vc1" {
  provider    = packetfabric
  description = "hello world"
  epl         = false
  interface_a {
    port_circuit_id = packetfabric_port.port_1.id
    untagged        = false
    vlan            = 100
  }
  interface_z {
    port_circuit_id = packetfabric_port.port_2.id
    untagged        = false
    vlan            = 101
  }
  bandwidth {
    longhaul_type     = "dedicated"
    speed             = "1Gbps"
    subscription_term = 1
  }
  labels = sort(["terraform", "dev"])
}
