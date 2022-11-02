resource "packetfabric_cloud_router" "cr1" {
  provider     = packetfabric
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_azure" "crc4" {
  provider          = packetfabric
  description       = var.pf_crc_description
  circuit_id        = packetfabric_cloud_router.cr1.id
  account_uuid      = var.pf_account_uuid
  azure_service_key = var.pf_crc_azure_service_key
  speed             = var.pf_crc_speed
  maybe_nat         = var.pf_crc_maybe_nat
  is_public         = var.pf_crc_is_public
}

output "packetfabric_cloud_router_connection_azure" {
  value = packetfabric_cloud_router_connection_azure.crc4
}
