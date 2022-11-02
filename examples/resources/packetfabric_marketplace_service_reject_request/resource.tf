resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider       = packetfabric
  description    = var.pf_description
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  routing_id     = var.pf_routing_id
  market         = var.pf_market
  speed          = var.pf_cs_speed
  pop            = var.pf_cs_pop
  zone           = var.pf_cs_zone
}

resource "packetfabric_marketplace_service_reject_request" "reject_request_aws" {
  provider        = packetfabric
  vc_request_uuid = packetfabric_cs_aws_hosted_marketplace_connection.cs_marketplace_conn1.id
}

output "packetfabric_marketplace_service_reject_request" {
  value = packetfabric_marketplace_service_reject_request.reject_request_aws
}
