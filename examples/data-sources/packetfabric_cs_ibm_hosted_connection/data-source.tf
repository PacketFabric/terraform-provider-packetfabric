data "packetfabric_cs_ibm_hosted_connection" "ibm_hosted_conn_1" {
  provider = packetfabric
  cloud_circuit_id = "PF-AP-LAX1-1002"
}

output "packetfabric_cs_ibm_hosted_connection" {
  value = data.packetfabric_cs_ibm_hosted_connection.ibm_hosted_conn_1
}