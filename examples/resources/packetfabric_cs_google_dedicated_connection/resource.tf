
resource "packetfabric_cs_google_dedicated_connection" "pf_cs_conn1_dedicated_google" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  description       = var.pf_description
  zone              = var.pf_cs_zone
  pop               = var.pf_cs_pop
  subscription_term = var.pf_cs_subterm
  service_class     = var.pf_cs_srvclass
  autoneg           = var.pf_cs_autoneg
  speed             = var.pf_cs_speed4
}

output "packetfabric_cs_google_dedicated_connection" {
  value = data.packetfabric_cs_google_dedicated_connection.pf_cs_conn1_dedicated_google
}
