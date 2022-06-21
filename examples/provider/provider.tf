terraform {
  required_providers {
    packetfabric = {
      source  = var.pf_provider_source
      version = "~> 0.0.0"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}
