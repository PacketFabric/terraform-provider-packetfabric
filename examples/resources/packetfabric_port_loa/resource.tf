resource "packetfabric_port_loa" "new" {
  provider          = packetfabric
  port_circuit_id   = var.pf_port_circuit_id
  loa_customer_name = "MyName"
  destination_email = "email@mydomain.com"
}
