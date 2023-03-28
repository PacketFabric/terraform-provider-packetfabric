resource "packetfabric_cs_azure_hosted_connection" "cs_conn1_hosted_azure" {
  provider          = packetfabric
  description       = "hello world"
  azure_service_key = var.azure_service_key
  port              = packetfabric_port.port_1.id
  speed             = "10Gbps" # will be deprecated
  vlan_private      = 102
  vlan_microsoft    = 103
  labels            = ["terraform", "dev"]
}

output "packetfabric_cs_azure_hosted_connection" {
  sensitive = true
  value     = packetfabric_cs_azure_hosted_connection.cs_conn1_hosted_azure
}