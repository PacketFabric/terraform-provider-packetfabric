terraform {
  required_providers {
    packetfabric = {
      source  = "packetfabric/packetfabric"
      version = ">= 0.2.1"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}
