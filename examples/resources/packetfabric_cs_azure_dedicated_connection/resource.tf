resource "packetfabric_cs_azure_dedicated_connection" "pf_cs_conn1_dedicated_azure" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  description       = var.description
  zone              = var.pf_cs_zone
  pop               = var.pf_cs_pop
  subscription_term = var.pf_cs_subterm
  service_class     = var.pf_cs_srvclass
  encapsulation     = var.encapsulation
  port_category     = var.port_category
  speed             = var.pf_cs_speed
}

output "packetfabric_cs_azure_dedicated_connection" {
  value = packetfabric_cs_azure_dedicated_connection.pf_cs_conn1_dedicated_azure
}