terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.7.0"
    }
  }
}

provider "packetfabric" {}
