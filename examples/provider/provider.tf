terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.8.0"
    }
  }
}

provider "packetfabric" {}
