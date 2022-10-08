resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_conn1_marketplace_aws" {
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

output "packetfabric_cs_aws_hosted_marketplace_connection" {
  value = packetfabric_cs_aws_hosted_marketplace_connection.cs_conn1_marketplace_aws
}