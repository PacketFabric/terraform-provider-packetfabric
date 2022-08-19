terraform {
  required_providers {
    packetfabric = {
      # source  = "PacketFabric/packetfabric"
      # version = "0.2.0"
      source  = "terraform.local/PacketFabric/packetfabric"
      version = "~> 0.0.0"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

# Create random name to use to name objects
resource "random_pet" "name" {}

# Create a PacketFabric ports/interfaces
resource "packetfabric_interface" "port_1" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  autoneg           = var.pf_interface_autoneg
  description       = "${var.tag_name}-${random_pet.name.id}"
  media             = var.pf_interface_media
  nni               = var.pf_interface_nni
  pop               = var.pf_interface_pop1
  speed             = var.pf_interface_speed
  subscription_term = var.pf_interface_subterm
  zone              = var.pf_interface_avzone1
}
output "packetfabric_port_1" {
  value = packetfabric_interface.port_1
}
resource "packetfabric_interface" "port_2" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  autoneg           = var.pf_interface_autoneg
  description       = "${var.tag_name}-${random_pet.name.id}"
  media             = var.pf_interface_media
  nni               = var.pf_interface_nni
  pop               = var.pf_interface_pop2
  speed             = var.pf_interface_speed
  subscription_term = var.pf_interface_subterm
  zone              = var.pf_interface_avzone2
}
output "packetfabric_port_2" {
  value = packetfabric_interface.port_2
}

# Get billing information related to the interface created
data "packetfabric_billing" "port_1" {
  provider   = packetfabric
  circuit_id = packetfabric_interface.port_1.id
  depends_on = [
    packetfabric_interface.port_1
  ]
}
output "packetfabric_billing_port_1" {
  value = data.packetfabric_billing.port_1
}
data "packetfabric_billing" "port_2" {
  provider   = packetfabric
  circuit_id = packetfabric_interface.port_2.id
  depends_on = [
    packetfabric_interface.port_2
  ]
}
output "packetfabric_billing_port_2" {
  value = data.packetfabric_billing.port_2
}

# Get PacketFabric locations
data "packetfabric_locations" "location_1" {
  provider = packetfabric
  # filter {
  #   pop = var.pf_interface_pop1
  # }
}
data "packetfabric_locations" "location_2" {
  provider = packetfabric
  # filter {
  #   pop = var.pf_interface_pop2
  # }
}

# Create Cross Connect
resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
  provider      = packetfabric
  description   = "${var.tag_name}-${random_pet.name.id}"
  document_uuid = var.pf_document_uuid1
  port          = packetfabric_interface.port_1.id
  site          = var.pf_interface_site1
  # https://github.com/PacketFabric/terraform-provider-packetfabric/issues/63
  #site = data.packetfabric_locations.location_1.site_code
}
output "packetfabric_outbound_cross_connect1" {
  value = packetfabric_outbound_cross_connect.crossconnect_1
}
resource "packetfabric_outbound_cross_connect" "crossconnect_2" {
  provider      = packetfabric
  description   = "${var.tag_name}-${random_pet.name.id}"
  document_uuid = var.pf_document_uuid2
  port          = packetfabric_interface.port_1.id
  site          = var.pf_interface_site2
  # https://github.com/PacketFabric/terraform-provider-packetfabric/issues/63
  #site = data.packetfabric_locations.location_2.site_code
}
output "packetfabric_outbound_cross_connect2" {
  value = packetfabric_outbound_cross_connect.crossconnect_2
}

# Create backbone Virtual Circuit
resource "packetfabric_backbone_virtual_circuit" "vc_1" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}"
  epl         = false
  interface_a {
    port_circuit_id = packetfabric_interface.port_1.id
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  interface_z {
    port_circuit_id = packetfabric_interface.port_2.id
    untagged        = false
    vlan            = var.pf_vc_vlan2
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}
