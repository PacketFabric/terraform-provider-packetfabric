terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.6.0"
    }
  }
}

provider "packetfabric" {}

# Create random name to use to name objects
resource "random_pet" "name" {}


# Accept the Request Backbone VC
resource "packetfabric_marketplace_service_port_accept_request" "accept_marketplace_request" {
  provider    = packetfabric
  type        = "backbone"
  description = "${var.resource_name}-${random_pet.name.id}"
  labels      = var.pf_labels
  interface {
    port_circuit_id = var.pf_z_side_port_id
    vlan            = var.pf_z_side_vc_vlan2
  }
  vc_request_uuid = var.pf_a_side_vc_request_uuid
}

# # Reject the Request
# resource "packetfabric_marketplace_service_port_reject_request" "reject_marketplace_request" {
#   provider        = packetfabric
#   vc_request_uuid = var.pf_a_side_vc_request_uuid
# }