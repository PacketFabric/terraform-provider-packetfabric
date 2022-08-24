terraform {
  required_providers {
    packetfabric = {
      source  = "packetfabric/packetfabric"
      version = "~> 0.0.1"
    }
  }
}

provider "packetfabric" {
  host = var.pf_api_server
  token = var.pf_api_key
}

data "packetfabric_cloud_router" "current" {
  provider = packetfabric
}

data "packetfabric_cloud_router_bgp_session" "current" {
  provider = packetfabric
}

data "packetfabric_cs_aws_dedicated_connection" "current" {
  provider = packetfabric
}

data "packetfabric_cloud_router_bgp_prefixes" "current" {
  provider = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.new.id
  depends_on = [
    packetfabric_cloud_router_bgp_session.new
  ]
}

resource "packetfabric_cloud_router" "new" {
  provider = packetfabric
  scope        = var.pf_cr_scope
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "packetfabric_aws_cloud_router_connection" "new" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.new.id
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  maybe_nat      = var.pf_crc_maybe_nat
  description    = var.pf_crc_description
  pop            = var.pf_crc_pop
  zone           = var.pf_crc_zone
  is_public      = var.pf_crc_is_public
  speed          = var.pf_crc_speed
  depends_on = [
    packetfabric_cloud_router.new,
    data.packetfabric_cloud_router.current
  ]
}

resource "packetfabric_cloud_router_bgp_session" "new" {
  provider         = packetfabric
  circuit_id       = packetfabric_cloud_router.new.id
  connection_id    = packetfabric_aws_cloud_router_connection.new.id
  address_family   = var.pf_crbs_af
  multihop_ttl     = var.pf_crbs_mhttl
  remote_asn       = var.pf_crbs_rasn
  orlonger         = var.pf_crbs_orlonger
  remote_address   = var.pf_crbs_remoteaddr
  l3_address       = var.pf_crbs_l3addr
  md5              = var.pf_crbs_md5
  depends_on = [
    packetfabric_aws_cloud_router_connection.new,
    data.packetfabric_cs_aws_dedicated_connection.new,
    data.packetfabric_cloud_router.current
  ]
}

output "packetfabric_cloud_router" {
  value = data.packetfabric_cloud_router.current
}

output "packetfabric_cs_aws_dedicated_connection" {
  value = data.packetfabric_cs_aws_dedicated_connection.new
}

output "packetfabric_cloud_router_bgp_session" {
  value = data.packetfabric_cloud_router_bgp_session.current
}
