terraform {
  required_providers {
    packetfabric = {
      source  = "packetfabric/packetfabric"
      version = "~> 0.0.1"
    }
  }
}

provider "packetfabric" {
  host = "https://api.packetfabric.com"
  token = "api-ddae87ba-f88d-4d1f-8c84-xyz"
}

data "cloud_routers" "new" {

}

resource "aws_cloud_router_connection" "new" {
  provider = packetfabric
  account_uuid      = "6e5a143c-6ab6-4dfd-a064-xxxx"
  aws_account_id    = "xxxxx"
  maybe_nat         = false
  description       = "New AWS Cloud Router Connection"
  pop               = "DALI1"
  zone              = "A"
  is_public         = true
  speed             = "100Mbps"
}

output "circuit_id" {
  value = cloud_router.new.circuit_id
}
