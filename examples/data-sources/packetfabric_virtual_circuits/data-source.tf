data "packetfabric_virtual_circuits" "vc" {
  provider = packetfabric
}
output "packetfabric_virtual_circuits" {
  value = data.packetfabric_virtual_circuits.vc
}