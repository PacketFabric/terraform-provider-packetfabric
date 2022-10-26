resource "ibm_resource_group" "resource_group_1" {
  name = "${var.tag_name}-${random_pet.name.id}"
}

resource "ibm_is_vpc" "vpc_1" {
  name                      = "${var.tag_name}-${random_pet.name.id}"
  resource_group            = ibm_resource_group.resource_group_1.id
  address_prefix_management = "manual" # no default prefix will be created for each zone in this VPC.
}

resource "ibm_is_vpc_address_prefix" "vpc_prefix_1" {
  provider = ibm
  name     = "${var.tag_name}-${random_pet.name.id}"
  zone     = var.ibm_region1_zone1
  vpc      = ibm_is_vpc.vpc_1.id
  cidr     = var.ibm_vpc_cidr1
}

resource "ibm_is_subnet" "subnet_1" {
  provider        = ibm
  name            = "${var.tag_name}-${random_pet.name.id}"
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

output "ibm_is_subnet" {
  value = data.ibm_is_subnet.subnet_1
}