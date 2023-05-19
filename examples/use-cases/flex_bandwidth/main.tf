terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.5.0"
    }
  }
}

provider "packetfabric" {}

# Create random name to use to name objects
resource "random_pet" "name" {}

# Create a PacketFabric ports
resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  labels            = var.pf_labels
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
}
# output "packetfabric_port_1" {
#   value = packetfabric_port.port_1
# }
resource "packetfabric_port" "port_2" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  labels            = var.pf_labels
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop2
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone2
}
# output "packetfabric_port_2" {
#   value = packetfabric_port.port_2
# }

resource "packetfabric_flex_bandwidth" "flex1" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  subscription_term = var.pf_flex_subscription_term
  capacity          = var.pf_flex_capacity
}
# output "packetfabric_flex_bandwidth" {
#   value = packetfabric_flex_bandwidth.flex1
# }

# Create backbone Virtual Circuit 1
resource "packetfabric_backbone_virtual_circuit" "vc_1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  labels      = var.pf_labels
  epl         = false
  interface_a {
    port_circuit_id = packetfabric_port.port_1.id
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  interface_z {
    port_circuit_id = packetfabric_port.port_2.id
    untagged        = false
    vlan            = var.pf_vc_vlan2
  }
  bandwidth {
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
  flex_bandwidth_id = packetfabric_flex_bandwidth.flex1.id
}
# output "packetfabric_backbone_virtual_circuit_1" {
#   value = packetfabric_backbone_virtual_circuit.vc_1
# }

# Create backbone Virtual Circuit 2
resource "packetfabric_backbone_virtual_circuit" "vc_2" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  labels      = var.pf_labels
  epl         = false
  interface_a {
    port_circuit_id = packetfabric_port.port_1.id
    untagged        = false
    vlan            = var.pf_vc_vlan3
  }
  interface_z {
    port_circuit_id = packetfabric_port.port_2.id
    untagged        = false
    vlan            = var.pf_vc_vlan4
  }
  bandwidth {
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
  flex_bandwidth_id = packetfabric_flex_bandwidth.flex1.id
}
# output "packetfabric_backbone_virtual_circuit_2" {
#   value = packetfabric_backbone_virtual_circuit.vc_2
# }