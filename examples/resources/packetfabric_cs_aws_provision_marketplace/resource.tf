resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider       = packetfabric
  description    = var.description
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  routing_id     = var.routing_id
  market         = var.market
  speed          = var.pf_cs_speed
  pop            = var.pf_cs_pop
  zone           = var.pf_cs_zone
}

resource "packetfabric_cs_aws_provision_marketplace" "accept_request_aws" {
  provider        = packetfabric
  description     = var.description
  port_circuit_id = var.port_circuit_id_marketplace
  vc_request_uuid = packetfabric_cs_aws_hosted_marketplace_connection.cs_marketplace_conn1.id
  vlan            = var.pf_cs_vlan
}

