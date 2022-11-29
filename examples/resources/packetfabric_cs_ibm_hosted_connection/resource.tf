resource "packetfabric_cs_ibm_hosted_connection" "new" {
  provider = packetfabric
  ibm_account_id = (environment variable PF_IBM_ACCOUNT_ID)
  ibm_bgp_asn = var.ibm_bgp_asn
  description = var.pf_description
  pop = var.pf_cs_pop
  port = var.pf_port
  vlan = var.pf_cs_vlan
  speed = var.pf_cs_speed
 }

output "packetfabric_cs_ibm_hosted_connection" {
  value     = packetfabric_cs_ibm_hosted_connection.new
  sensitive = true
}