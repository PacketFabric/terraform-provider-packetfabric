resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  enabled           = true
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