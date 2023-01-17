terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.6.0"
    }
  }
}

provider "packetfabric" {}
