resource "packetfabric_cs_google_hosted_connection" "cs_conn1_hosted_google" {
  provider                    = packetfabric
  description                 = "hello world"
  port                        = packetfabric_port.port_1.id
  speed                       = "10Gbps"
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = var.google_vlan_attachment_name
  pop                         = "BOS1"
  vlan                        = 102
  labels                      = ["terraform", "dev"]
}

output "packetfabric_cs_google_hosted_connection" {
  value     = packetfabric_cs_google_hosted_connection.cs_conn1_hosted_google
  sensitive = true
}