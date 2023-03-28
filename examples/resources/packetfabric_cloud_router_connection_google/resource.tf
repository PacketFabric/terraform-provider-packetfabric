resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = ["terraform", "dev"]
}

resource "packetfabric_cloud_router_connection_google" "crc2" {
  provider                    = packetfabric
  description                 = "hello world"
  circuit_id                  = packetfabric_cloud_router.cr1.id
  google_pairing_key          = var.pf_crc_google_pairing_key
  google_vlan_attachment_name = var.pf_crc_google_vlan_attachment_name
  pop                         = "PDX2"
  speed                       = "1Gbps"
  maybe_nat                   = false
  labels                      = ["terraform", "dev"]
}

output "packetfabric_cloud_router_connection_google" {
  value = packetfabric_cloud_router_connection_google.crc2
}
