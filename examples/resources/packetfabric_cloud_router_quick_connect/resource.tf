resource "packetfabric_cloud_router_quick_connect" "cr_quick_connect" {
  provider              = packetfabric
  cr_circuit_id         = packetfabric_cloud_router.cr1.id
  connection_circuit_id = packetfabric_cloud_router_connection_aws.crc1.id
  service_uuid          = var.pf_service_uuid
  return_filters {
    prefix     = "185.56.153.165/32"
    match_type = "orlonger"
  }
  return_filters {
    prefix     = "185.56.153.166/32"
    match_type = "orlonger"
  }
}