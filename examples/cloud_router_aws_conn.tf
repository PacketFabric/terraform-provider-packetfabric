terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric/packetfabric"
      version = "~> 0.0.1"
    }
  }
}

provider "packetfabric" {
  host = "https://api.packetfabric.com"
  token = "api-ddae87ba-f88d-4d1f-8c84-8e6225bdcc42-0ce71a7b-4cf3-4c95-ac5e-d74a03448c51"
}

data "cloud_routers" "new" {
  
}

resource "aws_cloud_router_connection" "new" {
  provider = packetfabric
  account_uuid      = "6e5a143c-6ab6-4dfd-a064-462ccdfa9a8a"
  aws_account_id    = "023456789102"
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
