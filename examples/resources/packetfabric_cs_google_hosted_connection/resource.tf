resource "packetfabric_cs_google_hosted_connection" "cs_conn1_hosted_google" {
  provider                    = packetfabric
  description                 = var.pf_description
  account_uuid                = var.pf_account_uuid
  port                        = var.pf_port
  speed                       = var.pf_cs_speed
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = var.google_vlan_attachment_name
  pop                         = var.pf_cs_pop
  vlan                        = var.pf_cs_vlan
}

output "packetfabric_cs_google_hosted_connection" {
  value     = packetfabric_cs_google_hosted_connection.cs_conn1_hosted_google
  sensitive = true
}