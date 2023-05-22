data "packetfabric_activitylogs" "current" {
  provider = packetfabric
}
output "packetfabric_activitylogs" {
  value = data.packetfabric_activitylogs.current
}