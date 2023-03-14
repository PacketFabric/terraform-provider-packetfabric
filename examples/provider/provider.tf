terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.2.0"
    }
  }
}

provider "packetfabric" {}
