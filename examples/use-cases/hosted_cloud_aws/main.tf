terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.3.2"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.37.0"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region1
}

# create random name to use to name objects
resource "random_pet" "name" {}

# Create the VPCs
resource "aws_vpc" "vpc_1" {
  provider             = aws
  cidr_block           = var.vpc_cidr1
  enable_dns_hostnames = true
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Define the subnets
resource "aws_subnet" "subnet_1" {
  provider   = aws
  vpc_id     = aws_vpc.vpc_1.id
  cidr_block = var.subnet_cidr1
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Define the internet gateways
resource "aws_internet_gateway" "gw_1" {
  provider = aws
  vpc_id   = aws_vpc.vpc_1.id
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Virtual Private Gateway (creation + attachement to the VPC)
resource "aws_vpn_gateway" "vpn_gw_1" {
  provider        = aws
  amazon_side_asn = var.amazon_side_asn1
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
  depends_on = [
    aws_vpc.vpc_1
  ]
}
resource "aws_vpn_gateway_attachment" "vpn_attachment_1" {
  provider       = aws
  vpc_id         = aws_vpc.vpc_1.id
  vpn_gateway_id = aws_vpn_gateway.vpn_gw_1.id
}

# Define the route tables
resource "aws_route_table" "route_table_1" {
  provider = aws
  vpc_id   = aws_vpc.vpc_1.id
  # internet gw
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw_1.id
  }
  propagating_vgws = ["${aws_vpn_gateway.vpn_gw_1.id}"]
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Assign the route table to the subnet
resource "aws_route_table_association" "route_association_1" {
  provider       = aws
  subnet_id      = aws_subnet.subnet_1.id
  route_table_id = aws_route_table.route_table_1.id
}

# From the PacketFabric side: Create a AWS Hosted Connection 
resource "packetfabric_cs_aws_hosted_connection" "pf_cs_conn1" {
  provider       = packetfabric
  description    = "${var.tag_name}-${random_pet.name.id}"
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  port           = var.pf_port_circuit_id
  speed          = var.pf_cs_speed
  pop            = var.pf_cs_pop1
  vlan           = var.pf_cs_vlan1
  zone           = var.pf_cs_zone1
}

output "packetfabric_cs_aws_hosted_connection" {
  value = packetfabric_cs_aws_hosted_connection.pf_cs_conn1
}

# From the AWS side: Accept the connection
# Wait at least 90s for the connection to show up in AWS
resource "null_resource" "previous" {}
resource "time_sleep" "wait_90_seconds" {
  depends_on      = [null_resource.previous]
  create_duration = "90s"
}
# This resource will create (at least) 90 seconds after null_resource.previous
resource "null_resource" "next" {
  depends_on = [time_sleep.wait_90_seconds]
}

# Retrieve the Direct Connect connections in AWS
data "aws_dx_connection" "current_1" {
  provider = aws
  name     = "${var.tag_name}-${random_pet.name.id}"
  depends_on = [
    null_resource.next,
    packetfabric_cs_aws_hosted_connection.pf_cs_conn1
  ]
}
output "aws_dx_connection_1" {
  value = data.aws_dx_connection.current_1
}

# Vote for 26335 aws_dx_connection_confirmation add timeout and do not fail when state is available
# https://github.com/hashicorp/terraform-provider-aws/issues/26335
resource "aws_dx_connection_confirmation" "confirmation_1" {
  provider      = aws
  connection_id = data.aws_dx_connection.current_1.id
}

# From the AWS side: Create a gateway
resource "aws_dx_gateway" "direct_connect_gw_1" {
  provider        = aws
  name            = "${var.tag_name}-${random_pet.name.id}"
  amazon_side_asn = var.amazon_side_asn1
  depends_on = [
    packetfabric_cs_aws_hosted_connection.pf_cs_conn1
  ]
}

# From the AWS side: Create and attach a VIF
data "aws_dx_gateway" "direct_connect_gw_1" {
  provider = aws
  name     = "${var.tag_name}-${random_pet.name.id}"

  depends_on = [
    aws_dx_gateway.direct_connect_gw_1
  ]
}

resource "aws_dx_private_virtual_interface" "direct_connect_vip_1" {
  provider       = aws
  connection_id  = data.aws_dx_connection.current_1.id
  dx_gateway_id  = aws_dx_gateway.direct_connect_gw_1.id
  name           = "${var.tag_name}-${random_pet.name.id}"
  vlan           = data.aws_dx_gateway.direct_connect_gw_1.vlan_id # provider version >= 4.35.0 https://github.com/hashicorp/terraform-provider-aws/issues/26461
  address_family = "ipv4"
  bgp_asn        = var.customer_side_asn1
  depends_on = [
    aws_dx_connection_confirmation.confirmation_1
  ]
}

 # provider version >= 4.37.0
data "aws_dx_router_configuration" "router_config" {
  provider               = aws
  virtual_interface_id   = aws_dx_private_virtual_interface.direct_connect_vip_1.id
  router_type_identifier = "CiscoSystemsInc-2900SeriesRouters-IOS124"
}

##########################################################################################
#### Here you would need to setup BGP in your Router
##########################################################################################

# # From the AWS side: Associate Virtual Private GW to Direct Connect GW
# resource "aws_dx_gateway_association" "virtual_private_gw_to_direct_connect_1" {
#   provider              = aws
#   dx_gateway_id         = aws_dx_gateway.direct_connect_gw_1.id
#   associated_gateway_id = aws_vpn_gateway.vpn_gw_1.id
#   allowed_prefixes = [
#     var.vpc_cidr1
#   ]
#   depends_on = [
#     aws_dx_private_virtual_interface.direct_connect_vip_1
#   ]
#   timeouts {
#     create = "2h"
#     delete = "2h"
#   }
# }
