data "packetfabric_marketplace_service_port_requests" "marketplace_requests" {
  provider = packetfabric
  type     = "sent" # sent or received
}
output "packetfabric_marketplace_service_port_requests" {
  value = data.packetfabric_marketplace_service_port_requests.marketplace_requests
}