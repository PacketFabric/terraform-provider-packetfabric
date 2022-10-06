resource "packetfabric_cloud_router_connection_azure" "crc_4" {
  provider          = packetfabric
  description       = var.pf_crc_description
  circuit_id        = packetfabric_cloud_router.cr.id
  account_uuid      = var.pf_account_uuid
  azure_service_key = var.pf_crc_azure_service_key
  speed             = var.pf_crc_speed
  maybe_nat         = var.pf_crc_maybe_nat
  is_public         = var.pf_crc_is_public
}