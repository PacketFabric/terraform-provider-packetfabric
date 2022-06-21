terraform {
  required_providers {
    packetfabric = {
      source  = var.provider_source
      version = "~> 0.0.0"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

data "cloud_router" "current" {
  provider = packetfabric
}

data "aws_cloud_router_connection" "new" {
  provider = packetfabric
}

resource "cloud_router" "cr1" {
  provider     = packetfabric
  scope        = var.pf_cr_scope
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "aws_cloud_router_connection" "crc1" {
  provider       = packetfabric
  circuit_id     = cloud_router.cr1.id
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  maybe_nat      = var.pf_crc_maybe_nat
  description    = var.pf_crc_description
  pop            = var.pf_crc_pop
  zone           = var.pf_crc_zone
  is_public      = var.pf_crc_is_public
  speed          = var.pf_crc_speed
  depends_on = [
    cloud_router.new,
    data.cloud_router.current
  ]
}

output "cloud_router" {
  value = data.cloud_router.current
}

output "cloud_router_conn" {
  value = data.aws_cloud_router_connection.new
}
