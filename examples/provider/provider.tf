terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.1.0"
    }
  }
}

provider "packetfabric" {}
