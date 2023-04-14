resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = sort(["terraform", "dev"])
}

resource "packetfabric_cloud_router_connection_azure" "crc4" {
  provider          = packetfabric
  description       = "hello world"
  circuit_id        = packetfabric_cloud_router.cr1.id
  azure_service_key = var.pf_crc_azure_service_key
  speed             = "1Gbps"
  maybe_nat         = false
  is_public         = false
  labels            = sort(["terraform", "dev"])
}