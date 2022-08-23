terraform {
  required_providers {
    packetfabric = {
      source  = "packetfabric/packetfabric"
      version = "~> 0.0.0"
    }
  }
}

provider "packetfabric" {
  host = var.pf_api_server
  token = var.pf_api_key
}


resource "packetfabric_cs_aws_hosted_marketplace_connection" "new" {
  provider = packetfabric
  routing_id = var.pf_cs_aws_hm_rid
  market = var.pf_cs_aws_hm_market
  description = var.pf_cs_aws_hm_descr
  aws_account_id = var.pf_aws_account_id
  account_uuid = var.pf_account_uuid
  service_uuid = var.pf_cs_aws_hm_svcuuid
  pop = var.pf_cs_aws_hm_pop
  zone = var.pf_cs_aws_hm_avzone
  speed = var.pf_cs_aws_hm_speed
}
