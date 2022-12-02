data "packetfabric_port_router_logs" "port_logs_1" {
  provider        = packetfabric
  port_circuit_id = var.pf_port_circuit_id
  time_from       = "2022-11-30 00:00:00"
  time_to         = "2022-12-01 00:00:00"
}
output "packetfabric_port_router_logs" {
  value = data.packetfabric_port_router_logs.port_logs_1
}
