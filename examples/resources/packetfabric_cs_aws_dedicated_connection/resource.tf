
resource "packetfabric_cs_aws_dedicated_connection" "pf_cs_conn1_dedicated_aws" {
  provider          = packetfabric
  aws_region        = var.aws_region
  account_uuid      = var.pf_account_uuid
  description       = var.description
  zone              = var.pf_cs_zone
  pop               = var.pf_cs_pop
  subscription_term = var.pf_cs_subterm
  service_class     = var.pf_cs_srvclass
  autoneg           = var.pf_cs_autoneg
  speed             = var.pf_cs_speed
  should_create_lag = var.should_create_lag
}

output "packetfabric_cs_aws_dedicated_connection" {
  value = packetfabric_cs_aws_dedicated_connection.pf_cs_conn1_dedicated_aws
}
