terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.5.0"
    }
  }
}

provider "packetfabric" {
  alias = "pf"
}

module "packetfabric-cloud-router" {
  source = "./modules/packetfabric"
  for_each = var.cloud_routers
  name = each.key 
  asn = each.value.asn 
  capacity = each.value.capacity 
  regions = each.value.regions
  aws_connections = var.aws_connections[each.key]
  aws_bgp_sessions = var.aws_bgp_sessions[each.key]
  aws_inbound = var.aws_inbound[regex("\\w+-\\w+$", each.key)]
  aws_outbound = var.aws_outbound[regex("\\w+-\\w+$", each.key)]
  gcp_connections = var.gcp_connections[each.key]
  gcp_bgp_sessions = var.gcp_bgp_sessions[each.key]
  gcp_inbound = var.gcp_inbound[regex("\\w+-\\w+$", each.key)]
  gcp_outbound = var.gcp_outbound[regex("\\w+-\\w+$", each.key)]
  providers = {
    packetfabric = packetfabric.pf
  }
}
