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
  autoneg           = var.pf_port_autoneg
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
  autoneg           = var.pf_port_autoneg
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

# OPTION 1
# Generate a LOA for both ports (outbound cross connect)
resource "packetfabric_document" "loa1_outbound" {
  provider        = packetfabric
  document        = "demo_port_loa.pdf" # LOA port 1
  type            = "loa"
  description     = "${var.resource_name}-${random_pet.name.id}-port1"
  port_circuit_id = packetfabric_port.port_1.id
}
resource "packetfabric_document" "loa2_outbound" {
  provider        = packetfabric
  document        = "demo_port_loa.pdf" # LOA port 2
  type            = "loa"
  description     = "${var.resource_name}-${random_pet.name.id}-port2"
  port_circuit_id = packetfabric_port.port_2.id
}

# # OPTION 2
# # Generate a LOA for a port (inbound cross connect)
# resource "packetfabric_port_loa" "loa1_inbound" {
#   provider          = packetfabric
#   port_circuit_id   = packetfabric_port.port_1.id
#   loa_customer_name = "My Awesome Company"
#   destination_email = "email@mydomain.com"
# }
# resource "packetfabric_port_loa" "loa2_inbound" {
#   provider          = packetfabric
#   port_circuit_id   = packetfabric_port.port_2.id
#   loa_customer_name = "My Awesome Company"
#   destination_email = "email@mydomain.com"
# }

### Get the site filtering on the pop using packetfabric_locations
data "packetfabric_locations" "locations_all" {
  provider = packetfabric
}
locals {
  all_locations = data.packetfabric_locations.locations_all.locations[*]
  helper_map = { for val in local.all_locations :
  val["pop"] => val }
  pf_port_site1 = local.helper_map["${var.pf_port_pop1}"]["site_code"]
  pf_port_site2 = local.helper_map["${var.pf_port_pop2}"]["site_code"]
}

# Create Cross Connect
resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
  provider      = packetfabric
  description   = "${var.resource_name}-${random_pet.name.id}-port1"
  document_uuid = packetfabric_document.loa1_outbound.id
  port          = packetfabric_port.port_1.id
  site          = local.pf_port_site1
}
# output "packetfabric_outbound_cross_connect1" {
#   value = packetfabric_outbound_cross_connect.crossconnect_1
# }
resource "packetfabric_outbound_cross_connect" "crossconnect_2" {
  provider      = packetfabric
  description   = "${var.resource_name}-${random_pet.name.id}-port2"
  document_uuid = packetfabric_document.loa2_outbound.id
  port          = packetfabric_port.port_2.id
  site          = local.pf_port_site2
}
# output "packetfabric_outbound_cross_connect2" {
#   value = packetfabric_outbound_cross_connect.crossconnect_2
# }

# Create backbone Virtual Circuit
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
}
# output "packetfabric_backbone_virtual_circuit_1" {
#   value = packetfabric_backbone_virtual_circuit.vc_1
# }
