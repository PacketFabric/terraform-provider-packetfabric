resource "packetfabric_quick_connect_reject_request" "reject_request_quick_connect" {
  provider          = packetfabric
  import_circuit_id = "PF-L3-IMP-2896010"
}