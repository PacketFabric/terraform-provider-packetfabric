resource "packetfabric_cs_ibm_hosted_connection" "cs_conn1_hosted_ibm" {
  provider    = packetfabric
  ibm_bgp_asn = 64536
  description = "hello world"
  pop         = "BOS1"
  port        = packetfabric_port.port_1.id
  vlan        = 102
  speed       = "10Gbps"
  labels      = sort(["terraform", "dev"])
}