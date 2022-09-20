resource "packetfabric_cloud_router" "cr1" {
  provider     = packetfabric
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "packetfabric_aws_cloud_router_connection" "crc1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  maybe_nat      = var.pf_crc_maybe_nat
  description    = var.pf_crc_description
  pop            = var.pf_crc_pop
  zone           = var.pf_crc_zone
  is_public      = var.pf_crc_is_public
  speed          = var.pf_crc_speed
}

output "packetfabric_cloud_router" {
  value = packetfabric_cloud_router.cr1
}

output "packetfabric_cloud_router_conn" {
  value = packetfabric_aws_cloud_router_connection.crc1
}
