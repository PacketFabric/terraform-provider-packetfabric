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

data "cloud_router" "current" {
  provider = packetfabric
}

resource "cloud_router" "new" {
  provider = packetfabric
  scope        = var.pf_cr_scope
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

output "cloud_router" {
  value = data.cloud_router.current
}
