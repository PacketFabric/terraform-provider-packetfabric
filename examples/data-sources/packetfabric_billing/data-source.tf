data "packetfabric_billing" "billing_port_1" {
  provider   = packetfabric
  circuit_id = packetfabric_port.port_1.id
}
output "packetfabric_billing_port_1" {
  value = data.packetfabric_billing.billing_port_1
}