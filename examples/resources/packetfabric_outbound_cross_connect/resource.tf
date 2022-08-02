terraform {
  required_providers {
    packetfabric = {
      source  = "packetfabric/packetfabric"
      version = "~> 0.0.1"
    }
  }
}
provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

resource "packetfabric_outbound_cross_connect" "new" {
  provider = packetfabric
  description = "my-OCC"
  document_uuid = "55A7A654-4C3C-4C69-BCBE-755790F0417C"
  port = "PF-SO-ME-111"
  site = "DR-ATL1"

}

data "packetfabric_outbound_cross_connect" "current" {
  filter {
    id = packetfabric_outbound_cross_connect.new.id
  }
}

output "my-occ-info" {
  value = packetfabric_outbound_cross_connect.current
}