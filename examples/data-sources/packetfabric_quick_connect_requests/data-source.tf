data "packetfabric_quick_connect_requests" "quick_connect_requests" {
  provider = packetfabric
  type     = "received" # sent or received
}
output "packetfabric_quick_connect_requests" {
  value = data.packetfabric_quick_connect_requests.quick_connect_requests
}
