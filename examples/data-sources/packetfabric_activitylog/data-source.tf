data "packetfabric_activitylog" "current" {
  provider = packetfabric
  filter {
    user = "alice"
  }
}
output "my-activity-logs" {
  value = data.packetfabric_activitylog.current
}