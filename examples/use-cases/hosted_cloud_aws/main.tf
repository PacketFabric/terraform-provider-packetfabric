terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.7.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.52.0"
    }
  }
}

provider "packetfabric" {}

provider "aws" {
  region = var.aws_region1
  # use AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables
}

# create random name to use to name objects
resource "random_pet" "name" {}

# Create the VPCs
resource "aws_vpc" "vpc_1" {
  provider             = aws
  cidr_block           = var.aws_vpc_cidr1
  enable_dns_hostnames = true
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Define the subnets
resource "aws_subnet" "subnet_1" {
  provider   = aws
  vpc_id     = aws_vpc.vpc_1.id
  cidr_block = var.aws_subnet_cidr1
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

# Create a PacketFabric port
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

# From the PacketFabric side: Create a AWS Hosted Connection 
resource "packetfabric_cs_aws_hosted_connection" "pf_cs_conn1" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}"
  port        = "PF-AP-WDC1-1726464" # packetfabric_port.port_1.id
  speed       = var.pf_cs_speed
  pop         = var.pf_cs_pop1
  vlan        = var.pf_cs_vlan1
  zone        = var.pf_cs_zone1
}

output "packetfabric_cs_aws_hosted_connection" {
  value = packetfabric_cs_aws_hosted_connection.pf_cs_conn1
}

# From the AWS side: Accept the connection
# Wait at least 90s for the connection to show up in AWS
resource "time_sleep" "wait_90_seconds" {
  create_duration = "90s"
  depends_on = [
    packetfabric_cs_aws_hosted_connection.pf_cs_conn1
  ]
}
# This resource will create (at least) 90 seconds after null_resource.previous
resource "null_resource" "next" {
  depends_on = [time_sleep.wait_90_seconds]
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
# Retrieve the Direct Connect connections in AWS
data "aws_dx_connection" "current_1" {
  provider = aws
  name     = "${var.tag_name}-${random_pet.name.id}"
  depends_on = [
    null_resource.next,
    packetfabric_cs_aws_hosted_connection.pf_cs_conn1
  ]
}
output "aws_dx_connection_current" {
  value = data.aws_dx_connection.current_1
}

# Vote for 29165 [Bug]: aws_dx_connection data-source reports vlan_id null
# https://github.com/hashicorp/terraform-provider-aws/issues/29165
# resource "aws_dx_private_virtual_interface" "direct_connect_vip_1" {
#   provider       = aws
#   connection_id  = data.aws_dx_connection.current_1.id
#   dx_gateway_id  = aws_dx_gateway.direct_connect_gw_1.id
#   name           = "${var.tag_name}-${random_pet.name.id}"
#   vlan           = data.aws_dx_connection.current_1.vlan_id # bug #29165
#   address_family = "ipv4"
#   bgp_asn        = var.customer_side_asn1
#   depends_on = [
#     aws_dx_connection_confirmation.confirmation_1
#   ]
# }

# # provider version >= 4.37.0
# data "aws_dx_router_configuration" "router_config" {
#   provider               = aws
#   virtual_interface_id   = aws_dx_private_virtual_interface.direct_connect_vip_1.id
#   router_type_identifier = "CiscoSystemsInc-2900SeriesRouters-IOS124"
# }

##########################################################################################
#### Here you would need to setup BGP in your Router
##########################################################################################

# # From the AWS side: Associate Virtual Private GW to Direct Connect GW
# resource "aws_dx_gateway_association" "virtual_private_gw_to_direct_connect_1" {
#   provider              = aws
#   dx_gateway_id         = aws_dx_gateway.direct_connect_gw_1.id
#   associated_gateway_id = aws_vpn_gateway.vpn_gw_1.id
#   allowed_prefixes = [
#     var.aws_vpc_cidr1
#   ]
#   depends_on = [
#     aws_dx_private_virtual_interface.direct_connect_vip_1
#   ]
#   timeouts {
#     create = "2h"
#     delete = "2h"
#   }
# }
