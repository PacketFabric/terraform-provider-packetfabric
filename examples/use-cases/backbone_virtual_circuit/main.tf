terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.0.2"
    }
  }
}

provider "packetfabric" {}

# Create random name to use to name objects
resource "random_pet" "name" {}

# Create a PacketFabric ports
resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  autoneg           = var.pf_port_autoneg
  description       = "${var.tag_name}-${random_pet.name.id}"
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
}
output "packetfabric_port_1" {
  value = packetfabric_port.port_1
}
resource "packetfabric_port" "port_2" {
  provider          = packetfabric
  autoneg           = var.pf_port_autoneg
  description       = "${var.tag_name}-${random_pet.name.id}"
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop2
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone2
}
output "packetfabric_port_2" {
  value = packetfabric_port.port_2
}

# Get billing information related to the interface created
data "packetfabric_billing" "port_1" {
  provider   = packetfabric
  circuit_id = packetfabric_port.port_1.id
}
output "packetfabric_billing_port_1" {
  value = data.packetfabric_billing.port_1
}
data "packetfabric_billing" "port_2" {
  provider   = packetfabric
  circuit_id = packetfabric_port.port_2.id
}
output "packetfabric_billing_port_2" {
  value = data.packetfabric_billing.port_2
}

### Get the site filtering on the pop using packetfabric_locations

# # List PacketFabric locations
# data "packetfabric_locations" "locations_all" {
#   provider = packetfabric
# }
# # output "packetfabric_locations" {
# #   value = data.packetfabric_locations.locations_all
# # }

# locals {
#   all_locations = data.packetfabric_locations.locations_all.locations[*]
#   helper_map = { for val in local.all_locations :
#   val["pop"] => val }
#   pf_port_site1 = local.helper_map["${var.pf_port_pop1}"]["site_code"]
#   pf_port_site2 = local.helper_map["${var.pf_port_pop2}"]["site_code"]
# }
# output "pf_port_site1" {
#   value = local.pf_port_site1
# }
# output "pf_port_site2" {
#   value = local.pf_port_site2
# }

# # Create Cross Connect
# resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
#   provider      = packetfabric
#   description   = "${var.tag_name}-${random_pet.name.id}"
#   document_uuid = var.pf_document_uuid1
#   port          = packetfabric_port.port_1.id
#   site          = local.pf_port_site1
# }
# output "packetfabric_outbound_cross_connect1" {
#   value = packetfabric_outbound_cross_connect.crossconnect_1
# }
# resource "packetfabric_outbound_cross_connect" "crossconnect_2" {
#   provider      = packetfabric
#   description   = "${var.tag_name}-${random_pet.name.id}"
#   document_uuid = var.pf_document_uuid2
#   port          = packetfabric_port.port_2.id
#   site          = local.pf_port_site2
# }
# output "packetfabric_outbound_cross_connect2" {
#   value = packetfabric_outbound_cross_connect.crossconnect_2
# }

# Create backbone Virtual Circuit
resource "packetfabric_backbone_virtual_circuit" "vc_1" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}"
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
}
output "packetfabric_backbone_virtual_circuit_1" {
  value = packetfabric_backbone_virtual_circuit.vc_1
}
