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

resource "packetfabric_port" "new" {
  provider = packetfabric
  account_uuid = var.pf_account_uuid
  autoneg = var.pf_cs_interface_autoneg
  description = var.pf_cs_interface_descr
  media = var.pf_cs_interface_media
  nni = var.pf_cs_interface_nni
  pop = var.pf_cs_interface_pop
  speed = var.pf_cs_interface_speed
  subscription_term = var.pf_cs_interface_subterm
  zone = var.pf_cs_interface_avzone
}
