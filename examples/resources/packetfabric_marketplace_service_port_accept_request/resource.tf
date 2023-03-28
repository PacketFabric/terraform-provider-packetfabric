resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider    = packetfabric
  description = "hello world"
  routing_id  = "PD-WUY-9VB0"
  market      = "HOU"
  speed       = "10Gbps"
  pop         = "BOS1"
  zone        = "A"
}

resource "packetfabric_marketplace_service_port_accept_request" "accept_marketplace_request" {
  provider       = packetfabric
  type           = "cloud" # "backbone", "ix" or "cloud"
  cloud_provider = "aws"   # "aws, azure, google, oracle
  interface {
    port_circuit_id = var.pf_market_port_circuit_id
    vlan            = 1022
  }
  vc_request_uuid = packetfabric_cs_aws_hosted_marketplace_connection.cs_marketplace_conn1.id
}

output "packetfabric_marketplace_service_port_accept_request" {
  value = packetfabric_marketplace_service_port_accept_request.accept_request_aws
}
