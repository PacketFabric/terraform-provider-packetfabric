resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider    = packetfabric
  description = "hello world"
  routing_id  = "PD-WUY-9VB0"
  market      = "HOU"
  speed       = "10Gbps"
  pop         = "BOS1"
  zone        = "A"
}

resource "packetfabric_marketplace_service_port_reject_request" "reject_request_aws" {
  provider        = packetfabric
  vc_request_uuid = packetfabric_cs_aws_hosted_marketplace_connection.cs_marketplace_conn1.id
}