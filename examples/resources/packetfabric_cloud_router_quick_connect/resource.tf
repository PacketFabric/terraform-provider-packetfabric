resource "packetfabric_cloud_router_quick_connect" "cr_quick_connect" {
  provider              = packetfabric
  cr_circuit_id         = var.pf_cr_circuit_id
  connection_circuit_id = var.pf_connection_circuit_id
  service_uuid          = var.pf_service_uuid
  return_filters {
    prefix     = var.pf_return_filters_prefix1
    match_type = var.pf_return_filters_match_type1
  }
  return_filters {
    prefix     = var.pf_return_filters_prefix2
    match_type = var.pf_return_filters_match_type2
  }
}
output "packetfabric_cloud_router_quick_connect" {
  value = packetfabric_cloud_router_quick_connect.cr_quick_connect
}