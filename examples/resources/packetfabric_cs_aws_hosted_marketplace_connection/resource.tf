resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_conn1_marketplace_aws" {
  provider       = packetfabric
  description    = var.pf_description
  aws_account_id = var.pf_aws_account_id
  routing_id     = var.pf_routing_id
  market         = var.pf_market
  speed          = var.pf_cs_speed
  pop            = var.pf_cs_pop
  zone           = var.pf_cs_zone
}

output "packetfabric_cs_aws_hosted_marketplace_connection" {
  value = packetfabric_cs_aws_hosted_marketplace_connection.cs_conn1_marketplace_aws
}