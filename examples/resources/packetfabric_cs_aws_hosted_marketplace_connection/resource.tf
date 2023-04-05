resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_conn1_marketplace_aws" {
  provider    = packetfabric
  description = "hello world"
  routing_id  = "PD-WUY-9VB0"
  market      = "HOU"
  speed       = "10Gbps"
  pop         = "BOS1"
  zone        = "A"
}