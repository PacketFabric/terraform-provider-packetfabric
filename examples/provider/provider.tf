terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.3.1"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}
