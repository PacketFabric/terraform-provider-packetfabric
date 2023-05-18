resource "packetfabric_document" "loa1" {
  provider        = packetfabric
  document        = "letter-of-authorization-PF-AP-LAB8-3339359.pdf"
  type            = "loa"
  description     = "My LOA"
  port_circuit_id = "PF-AP-LAB8-3339359"
}