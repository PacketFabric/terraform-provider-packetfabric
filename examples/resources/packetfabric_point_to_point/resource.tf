resource "packetfabric_point_to_point" "ptp1" {
  provider          = packetfabric
  description       = "hello world"
  speed             = "1Gbps"
  media             = "LX"
  subscription_term = 1
  endpoints {
    pop     = "SEA3"
    zone    = "A"
    autoneg = true
  }
  endpoints {
    pop     = "NYC5"
    zone    = "A"
    autoneg = true
  }
  labels = sort(["terraform", "dev"])
}
