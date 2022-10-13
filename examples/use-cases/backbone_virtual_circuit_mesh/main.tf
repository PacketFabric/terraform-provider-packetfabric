terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.3.2"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

# Create random name to use to name objects
resource "random_pet" "name" {}

# var.pf_port2 - var.pf_port3
resource "packetfabric_backbone_virtual_circuit" "vc1" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc1"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port2
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  interface_z {
    port_circuit_id = var.pf_port3
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port1 - var.pf_port2
resource "packetfabric_backbone_virtual_circuit" "vc2" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc2"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port1
    untagged        = false
    vlan            = var.pf_vc_vlan2
  }
  interface_z {
    port_circuit_id = var.pf_port2
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

# var.pf_port1 - var.pf_port5
resource "packetfabric_backbone_virtual_circuit" "vc3" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc3"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port1
    untagged        = false
    vlan            = var.pf_vc_vlan3
  }
  interface_z {
    port_circuit_id = var.pf_port5
    untagged        = false
    vlan            = var.pf_vc_vlan3
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port1 - var.pf_port4
resource "packetfabric_backbone_virtual_circuit" "vc4" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc4"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port1
    untagged        = false
    vlan            = var.pf_vc_vlan4
  }
  interface_z {
    port_circuit_id = var.pf_port4
    untagged        = false
    vlan            = var.pf_vc_vlan4
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port1 - var.pf_port6
resource "packetfabric_backbone_virtual_circuit" "vc5" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc5"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port1
    untagged        = false
    vlan            = var.pf_vc_vlan5
  }
  interface_z {
    port_circuit_id = var.pf_port6
    untagged        = false
    vlan            = var.pf_vc_vlan5
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port1 - var.pf_port3
resource "packetfabric_backbone_virtual_circuit" "vc6" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc6"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port1
    untagged        = false
    vlan            = var.pf_vc_vlan6
  }
  interface_z {
    port_circuit_id = var.pf_port3
    untagged        = false
    vlan            = var.pf_vc_vlan6
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port2 - var.pf_port5
resource "packetfabric_backbone_virtual_circuit" "vc7" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc7"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port2
    untagged        = false
    vlan            = var.pf_vc_vlan7
  }
  interface_z {
    port_circuit_id = var.pf_port5
    untagged        = false
    vlan            = var.pf_vc_vlan7
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port2 - var.pf_port4
resource "packetfabric_backbone_virtual_circuit" "vc8" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc8"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port2
    untagged        = false
    vlan            = var.pf_vc_vlan8
  }
  interface_z {
    port_circuit_id = var.pf_port4
    untagged        = false
    vlan            = var.pf_vc_vlan8
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port2 - var.pf_port6
resource "packetfabric_backbone_virtual_circuit" "vc9" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc9"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port2
    untagged        = false
    vlan            = var.pf_vc_vlan9
  }
  interface_z {
    port_circuit_id = var.pf_port6
    untagged        = false
    vlan            = var.pf_vc_vlan9
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port5 - var.pf_port4
resource "packetfabric_backbone_virtual_circuit" "vc10" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc10"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port5
    untagged        = false
    vlan            = var.pf_vc_vlan10
  }
  interface_z {
    port_circuit_id = var.pf_port4
    untagged        = false
    vlan            = var.pf_vc_vlan10
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port5 - var.pf_port6
resource "packetfabric_backbone_virtual_circuit" "vc11" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc11"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port5
    untagged        = false
    vlan            = var.pf_vc_vlan11
  }
  interface_z {
    port_circuit_id = var.pf_port6
    untagged        = false
    vlan            = var.pf_vc_vlan11
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port5 - var.pf_port3
resource "packetfabric_backbone_virtual_circuit" "vc12" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc12"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port5
    untagged        = false
    vlan            = var.pf_vc_vlan12
  }
  interface_z {
    port_circuit_id = var.pf_port3
    untagged        = false
    vlan            = var.pf_vc_vlan12
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port4 - var.pf_port6
resource "packetfabric_backbone_virtual_circuit" "vc13" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc13"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port4
    untagged        = false
    vlan            = var.pf_vc_vlan13
  }
  interface_z {
    port_circuit_id = var.pf_port6
    untagged        = false
    vlan            = var.pf_vc_vlan13
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port4 - var.pf_port3
resource "packetfabric_backbone_virtual_circuit" "vc14" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc14"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port4
    untagged        = false
    vlan            = var.pf_vc_vlan14
  }
  interface_z {
    port_circuit_id = var.pf_port3
    untagged        = false
    vlan            = var.pf_vc_vlan14
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}

# var.pf_port6 - var.pf_port3
resource "packetfabric_backbone_virtual_circuit" "vc15" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-vc15"
  epl         = false
  interface_a {
    port_circuit_id = var.pf_port6
    untagged        = false
    vlan            = var.pf_vc_vlan15
  }
  interface_z {
    port_circuit_id = var.pf_port3
    untagged        = false
    vlan            = var.pf_vc_vlan15
  }
  bandwidth {
    account_uuid      = var.pf_account_uuid
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}
