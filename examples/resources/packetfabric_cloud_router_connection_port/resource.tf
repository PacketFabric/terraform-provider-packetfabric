resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = ["terraform", "dev"]
}

resource "packetfabric_cloud_router_connection_port" "crc7" {
  provider        = packetfabric
  description     = "hello world"
  circuit_id      = packetfabric_cloud_router.cr1.id
  port_circuit_id = packetfabric_port.port_1.id
  vlan            = 104
  untagged        = false
  speed           = "1Gbps"
  is_public       = false
  maybe_nat       = false
  labels          = ["terraform", "dev"]
}

output "packetfabric_cloud_router_connection_port" {
  value = packetfabric_cloud_router_connection_port.crc7
}
