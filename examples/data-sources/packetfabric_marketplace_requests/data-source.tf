data "packetfabric_marketplace_requests" "marketplace_requests" {
  provider = packetfabric
}
output "packetfabric_marketplace_requests" {
  value = data.packetfabric_marketplace_requests.marketplace_requests
}