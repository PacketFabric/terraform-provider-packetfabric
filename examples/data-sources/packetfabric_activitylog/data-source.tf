data "packetfabric_activitylog" "current" {
  provider = packetfabric
}
output "my-activity-logs" {
  value = data.packetfabric_activitylog.current
}