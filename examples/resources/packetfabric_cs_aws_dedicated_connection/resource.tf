resource "packetfabric_cs_aws_dedicated_connection" "pf_cs_conn1_dedicated_aws" {
  provider          = packetfabric
  aws_region        = "us-east-1"
  description       = "hello world"
  zone              = "A"
  pop               = "BOS1"
  subscription_term = 1
  service_class     = "longhaul"
  autoneg           = false
  speed             = "10Gbps"
  should_create_lag = true
  labels            = sort(["terraform", "dev"])
}