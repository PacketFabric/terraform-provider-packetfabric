resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider    = packetfabric
  description = var.pf_description
  routing_id  = var.pf_routing_id
  market      = var.pf_market
  speed       = var.pf_cs_speed
  pop         = var.pf_cs_pop
  zone        = var.pf_cs_zone
}

resource "packetfabric_marketplace_service_port_accept_request" "accept_marketplace_request" {
  provider       = packetfabric
  type           = "cloud" # "backbone", "ix" or "cloud"
  cloud_provider = "aws"   # "aws, azure, google, oracle
  interface {
    port_circuit_id = var.pf_market_port_circuit_id
    vlan            = var.pf_cs_vlan2
  }
  vc_request_uuid = packetfabric_cs_aws_hosted_marketplace_connection.cs_marketplace_conn1.id
}

output "packetfabric_marketplace_service_port_accept_request" {
  value = packetfabric_marketplace_service_port_accept_request.accept_request_aws
}
