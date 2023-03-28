
resource "packetfabric_cs_google_dedicated_connection" "pf_cs_conn1_dedicated_google" {
  provider          = packetfabric
  description       = "hello world"
  zone              = "A"
  pop               = "BOS1"
  subscription_term = 1
  service_class     = "longhaul"
  autoneg           = false
  speed             = "10Gbps"
  labels            = ["terraform", "dev"]
}

output "packetfabric_cs_google_dedicated_connection" {
  value = data.packetfabric_cs_google_dedicated_connection.pf_cs_conn1_dedicated_google
}
