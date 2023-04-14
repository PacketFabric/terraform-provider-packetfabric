resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = sort(["terraform", "dev"])
}

resource "packetfabric_cloud_router_connection_ibm" "crc5" {
  provider    = packetfabric
  description = "hello world"
  circuit_id  = packetfabric_cloud_router.cr1.id
  ibm_bgp_asn = packetfabric_cloud_router.cr1.asn
  pop         = "DAL2"
  zone        = "A"
  maybe_nat   = false
  speed       = "1Gbps"
  labels      = sort(["terraform", "dev"])
}