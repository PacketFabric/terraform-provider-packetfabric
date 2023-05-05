terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.5.0"
    }
    ibm = {
      source  = "IBM-Cloud/ibm"
      version = ">= 1.53.0"
    }
  }
}

provider "packetfabric" {}

provider "ibm" {
  region = var.ibm_region1
  # use PF_IBM_ACCOUNT_ID, IC_API_KEY, IAAS_CLASSIC_USERNAME, IAAS_CLASSIC_API_KEY environment variables
}

# create random name to use to name objects
resource "random_pet" "name" {}

resource "ibm_resource_group" "resource_group_1" {
  name = "${var.resource_name}-${random_pet.name.id}"
}

resource "ibm_is_vpc" "vpc_1" {
  name                      = "${var.resource_name}-${random_pet.name.id}"
  resource_group            = ibm_resource_group.resource_group_1.id
  address_prefix_management = "manual" # no default prefix will be created for each zone in this VPC.
}

resource "ibm_is_vpc_address_prefix" "vpc_prefix_1" {
  provider = ibm
  name     = "${var.resource_name}-${random_pet.name.id}"
  zone     = var.ibm_region1_zone1
  vpc      = ibm_is_vpc.vpc_1.id
  cidr     = var.ibm_vpc_cidr1
}

resource "ibm_is_subnet" "subnet_1" {
  provider        = ibm
  name            = "${var.resource_name}-${random_pet.name.id}"
  resource_group  = ibm_resource_group.resource_group_1.id
  vpc             = ibm_is_vpc.vpc_1.id
  zone            = var.ibm_region1_zone1
  ipv4_cidr_block = var.ibm_subnet_cidr1
  #routing_table   = ibm_is_vpc_routing_table.example.routing_table
  depends_on = [
    ibm_is_vpc_address_prefix.vpc_prefix_1
  ]
}

data "ibm_is_subnet" "subnet_1" {
  provider   = ibm
  identifier = ibm_is_subnet.subnet_1.id
}
# output "ibm_is_subnet" {
#   value = data.ibm_is_subnet.subnet_1
# }

# Create a PacketFabric port
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

# From the PacketFabric side: Create a IBM Hosted Connection 
resource "packetfabric_cs_ibm_hosted_connection" "pf_cs_conn1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  labels      = var.pf_labels
  ibm_bgp_asn = var.pf_cs_peer_asn
  port        = packetfabric_port.port_1.id
  speed       = var.pf_cs_speed
  pop         = var.pf_cs_pop1
  vlan        = var.pf_cs_vlan1
  zone        = var.pf_cs_zone1

  depends_on = [
    ibm_resource_group.resource_group_1
  ]
}

# From the IBM side: Accept the connection
resource "time_sleep" "wait_ibm_connection" {
  create_duration = "1m"
}
# Retrieve the Direct Connect connections in IBM
data "ibm_dl_gateway" "current" {
  provider   = ibm
  name       = "${var.resource_name}-${random_pet.name.id}"
  depends_on = [time_sleep.wait_ibm_connection]
}

# Used in case you are using an existing resource group and you don't create a new one
# data "ibm_resource_group" "existing_rg" {
#   provider   = ibm
#   name       = "Packet Fabric"
# }

resource "ibm_dl_gateway_action" "confirmation" {
  provider = ibm
  gateway  = data.ibm_dl_gateway.current.id
  # resource_group = data.ibm_resource_group.existing_rg.id # used for existing resource group
  resource_group = ibm_resource_group.resource_group_1.id # used for new resource group
  action         = "create_gateway_approve"
  global         = true
  metered        = true # If set true gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway

  provisioner "local-exec" {
    when    = destroy
    command = "sleep 30"
  }
}
# output "ibm_dl_gateway_action" {
#   value = data.ibm_dl_gateway.current
# }

# data "ibm_dl_gateway" "after_approved" {
#   provider   = ibm
#   name       = "${var.resource_name}-${random_pet.name.id}"
#   depends_on = [ibm_dl_gateway_action.confirmation]
# }
# output "ibm_dl_gateway_after" {
#   value = data.ibm_dl_gateway.after_approved
# }

##########################################################################################
#### Here you would need to setup BGP in your Router
##########################################################################################

# use ibm_dl_gateway_action.confirmation.bgp_base_cidr # IBM side
# use ibm_dl_gateway_action.confirmation.bgp_cer_cidr  # PF side
