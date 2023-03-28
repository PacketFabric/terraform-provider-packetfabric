resource "packetfabric_backbone_virtual_circuit_speed_burst" "boost" {
  provider      = packetfabric
  vc_circuit_id = packetfabric_backbone_virtual_circuit.vc1.id
  speed         = "10Gbps"
}

output "packetfabric_backbone_virtual_circuit_speed_burst" {
  value = packetfabric_backbone_virtual_circuit_speed_burst.boost
}