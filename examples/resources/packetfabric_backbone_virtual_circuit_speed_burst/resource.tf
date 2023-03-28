resource "packetfabric_backbone_virtual_circuit_speed_burst" "boost" {
  provider      = packetfabric
  vc_circuit_id = var.pf_vc_circuit_id
  speed         = "10Gbps"
}

output "packetfabric_backbone_virtual_circuit_speed_burst" {
  value = packetfabric_backbone_virtual_circuit_speed_burst.boost
}