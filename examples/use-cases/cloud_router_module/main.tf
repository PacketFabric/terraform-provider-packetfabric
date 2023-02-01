# This code specifies the PacketFabric provider and sets its version to be greater than or equal to 0.7.0
terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.0.0"
    }
  }
}

# This block configures the PacketFabric provider and assigns it an alias of "pf"
provider "packetfabric" {
  alias = "pf"
}

# This block sets up the "packetfabric-cloud-router" module, using variables and the PacketFabric provider
module "packetfabric-cloud-router" {
  source = "./modules/packetfabric"
  # This variable is used to iterate over multiple cloud routers
  for_each         = var.cloud_routers
  name             = each.key
  asn              = each.value.asn
  capacity         = each.value.capacity
  regions          = each.value.regions
  aws_connections  = var.aws_connections
  aws_bgp_sessions = var.aws_bgp_sessions
  aws_outbound     = var.aws_outbound
  gcp_connections  = var.gcp_connections
  gcp_bgp_sessions = var.gcp_bgp_sessions
  gcp_outbound     = var.gcp_outbound
  # This block sets the provider to use the "pf" alias created earlier
  providers = {
    packetfabric = packetfabric.pf
  }
}
