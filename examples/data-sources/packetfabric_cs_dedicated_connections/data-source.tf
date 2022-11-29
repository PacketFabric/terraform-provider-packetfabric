data "packetfabric_cs_dedicated_connections" "current" {
  provider = packetfabric
}
output "packetfabric_cs_dedicated_connections" {
  value = data.packetfabric_cs_dedicated_connections.current
}