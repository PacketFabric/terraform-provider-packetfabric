terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.0.0"
    }
  }
}

provider "packetfabric" {}
