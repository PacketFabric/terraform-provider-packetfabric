resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = ["terraform", "dev"]
}

resource "packetfabric_cloud_router_connection_oracle" "crc6" {
  provider    = packetfabric
  description = "hello world"
  circuit_id  = packetfabric_cloud_router.cr1.id
  region      = "us-ashburn-1"
  vc_ocid     = var.pf_crc_oracle_vc_ocid
  pop         = "SFO1"
  zone        = "A"
  maybe_nat   = false
  labels      = ["terraform", "dev"]
}

output "packetfabric_cloud_router_connection_oracle" {
  value = packetfabric_cloud_router_connection_oracle.crc6
}
