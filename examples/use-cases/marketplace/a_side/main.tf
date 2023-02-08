terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.0.0"
    }
  }
}

provider "packetfabric" {}

# Create random name to use to name objects
resource "random_pet" "name" {}

# Create a VC Marketplace Connection 
resource "packetfabric_backbone_virtual_circuit_marketplace" "vc_marketplace_conn1" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}"
  routing_id  = var.pf_z_side_routing_id
  market      = var.pf_z_side_market
  interface {
    port_circuit_id = var.pf_a_side_port_id
    untagged        = false
    vlan            = var.pf_a_side_vc_vlan1
  }
  bandwidth {
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}
output "packetfabric_backbone_virtual_circuit_marketplace" {
  value = packetfabric_backbone_virtual_circuit_marketplace.vc_marketplace_conn1
}

# Once the request has been accepted:
# 1. Comment above packetfabric_backbone_virtual_circuit_marketplace resource
# 2. Import the new Marketplace backbone Virtual Circuit (replace with correct VC ID)
#    $ terraform import packetfabric_backbone_virtual_circuit.vc_marketplace PF-BC-RNO-CHI-1729807-PF
# 3. Comment out below packetfabric_backbone_virtual_circuit resource
# 4. Apply the plan to confirm the resource is correctly imported and managed by Terraform

# resource "packetfabric_backbone_virtual_circuit" "vc_marketplace" {
#   provider    = packetfabric
#   description = "${var.tag_name}-${random_pet.name.id}"
#   epl         = false
#   interface_a {
#     port_circuit_id = var.pf_a_side_port_id
#     untagged        = false
#     vlan            = var.pf_a_side_vc_vlan1
#   }
#   interface_z {
#     port_circuit_id = var.pf_z_side_port_id
#     untagged        = false
#     vlan            = var.pf_z_side_vc_vlan2
#   }
#   bandwidth {
#     longhaul_type     = var.pf_vc_longhaul_type
#     speed             = var.pf_vc_speed
#     subscription_term = var.pf_vc_subterm
#   }
# }
# output "packetfabric_backbone_virtual_circuit" {
#   value = packetfabric_backbone_virtual_circuit.vc_marketplace
# }