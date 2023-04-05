resource "packetfabric_cs_azure_dedicated_connection" "pf_cs_conn1_dedicated_azure" {
  provider          = packetfabric
  description       = "hello world"
  zone              = "A"
  pop               = "BOS1"
  subscription_term = 1
  service_class     = "longhaul"
  encapsulation     = "dot1q"
  port_category     = "primary" # secondary
  speed             = "10Gbps"
  labels            = ["terraform", "dev"]
}