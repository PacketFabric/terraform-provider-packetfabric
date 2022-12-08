terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.5.1"
    }
  }
}

provider "packetfabric" {}
