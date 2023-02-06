terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.9.0"
    }
  }
}

provider "packetfabric" {}
