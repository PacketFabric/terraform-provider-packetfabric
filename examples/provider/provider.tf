terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.4.0"
    }
  }
}

provider "packetfabric" {}
