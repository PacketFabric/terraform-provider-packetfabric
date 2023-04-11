terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.3.0"
    }
  }
}

provider "packetfabric" {}
