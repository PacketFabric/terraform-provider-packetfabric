terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.0.2"
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

# From the PacketFabric side: Create a AWS Dedicated Connection 
resource "packetfabric_cs_aws_dedicated_connection" "pf_cs_conn1" {
  provider          = packetfabric
  aws_region        = var.aws_region1
  description       = "${var.tag_name}-${random_pet.name.id}"
  zone              = var.pf_cs_zone1
  pop               = var.pf_cs_pop1
  subscription_term = var.pf_cs_subterm
  service_class     = var.pf_cs_srvclass
  autoneg           = var.pf_cs_aws_d_autoneg
  speed             = var.pf_cs_speed
  should_create_lag = var.should_create_lag
}

# data "packetfabric_cs_aws_dedicated_connection" "pf_cs_conn1" {
#   provider = packetfabric
# }
# output "packetfabric_cs_aws_dedicated_connection" {
#   value = data.packetfabric_cs_aws_dedicated_connection.pf_cs_conn1
# }

# # Get PacketFabric locations
# data "packetfabric_locations" "location_1" {
#   provider = packetfabric
# }
# output "packetfabric_locations" {
#   value = data.packetfabric_locations.location_1
# }

# # Get AWS locations
# data "aws_dx_locations" "aws_locations" {
#     provider = aws
# }

# output "aws_dx_locations" {
#   value = data.aws_dx_locations.aws_locations
# }

# Vote for 26438 aws_dx_locations: add Direct Connect Locations & Speed + filter capability
# https://github.com/hashicorp/terraform-provider-aws/issues/26438

# From the AWS side: Create the Direct Connect Connection 
resource "aws_dx_connection" "current_1" {
  provider      = aws
  name          = "${var.tag_name}-${random_pet.name.id}-${var.pf_cs_pop1}"
  bandwidth     = var.pf_cs_speed
  location      = var.aws_dx_location
  provider_name = "PacketFabric"
}

# # Wait at least 60s for the connection to show up in AWS
# resource "null_resource" "previous" {}
# resource "time_sleep" "wait_60_seconds" {
#   depends_on      = [null_resource.previous]
#   create_duration = "60s"
# }
# # This resource will create (at least) 60 seconds after null_resource.previous
# resource "null_resource" "next" {
#   depends_on = [time_sleep.wait_60_seconds]
# }

# # Retrieve the Direct Connect connections in AWS
# data "aws_dx_connection" "current_1" {
#   provider = aws
#   name     = "${var.tag_name}-${random_pet.name.id}-${var.pf_cs_pop1}"
#   depends_on = [
#     null_resource.next,
#     packetfabric_cs_aws_dedicated_connection.pf_cs_conn1
#   ]
# }
# output "aws_dx_connection_1" {
#   value = data.aws_dx_connection.current_1
# }

##########################################################################################
#### From the AWS portal, download the LOA-CFA and use it to order a cross connect in the PacketFabric portal.
##########################################################################################

# Vote for 26436 aws_dx_connection data source: add PDF LOA in base64 encoded 
# https://github.com/hashicorp/terraform-provider-aws/issues/26436

# # Create Cross Connect
# resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
#   provider      = packetfabric
#   description   = "${var.tag_name}-${random_pet.name.id}"
#   document_uuid = var.pf_document_uuid1
#   port          = var.pf_interface_a_circuit_id
#   #port          = data.packetfabric_cs_aws_dedicated_connection.pf_cs_conn1.cloud_circuit_id
#   site          = var.pf_cs_site1
#   # https://github.com/PacketFabric/terraform-provider-packetfabric/issues/63
#   #site = data.packetfabric_locations.location_1.site_code
# }
# output "packetfabric_outbound_cross_connect1" {
#   value = packetfabric_outbound_cross_connect.crossconnect_1
# }

# From the PacketFabric portal, create a virtual circuit from your source interface to the PacketFabric-AWS Direct Connection connection.

# # Create backbone Virtual Circuit
# resource "packetfabric_backbone_virtual_circuit" "vc_1" {
#   provider    = packetfabric
#   description = "${var.tag_name}-${random_pet.name.id}"
#   epl         = false
#   interface_a {
#     port_circuit_id = var.pf_interface_a_circuit_id # AWS dedicated cloud
#     #port_circuit_id = data.packetfabric_cs_aws_dedicated_connection.pf_cs_conn1.cloud_circuit_id
#     untagged        = false
#     vlan            = var.pf_vc_vlan1
#   }
#   interface_z {
#     port_circuit_id = var.pf_interface_b_circuit_id # existing PF port
#     untagged        = false
#     vlan            = var.pf_vc_vlan2
#   }
#   bandwidth {
#     longhaul_type     = var.pf_vc_longhaul_type
#     speed             = var.pf_vc_speed
#     subscription_term = var.pf_vc_subterm
#   }
# }

# From the AWS portal, create a virtual interface to complete your Direct Connect connection.

# # From the AWS side: Create a gateway
# resource "aws_dx_gateway" "direct_connect_gw_1" {
#   provider        = aws
#   name            = "${var.tag_name}-${random_pet.name.id}-${var.pf_cs_pop1}"
#   amazon_side_asn = var.amazon_side_asn1
#   depends_on = [
#     packetfabric_cloud_router_connection_aws.crc_1
#   ]
# }

# # From the AWS side: Create and attach a VIF
# resource "aws_dx_private_virtual_interface" "direct_connect_vip_1" {
#   provider       = aws
#   connection_id  = data.aws_dx_connection.current_1.id
#   dx_gateway_id  = aws_dx_gateway.direct_connect_gw_1.id
#   name           = "${var.tag_name}-${random_pet.name.id}-${var.pf_cs_pop1}"
#   vlan           = var.pf_cs_vlan1
#   address_family = "ipv4"
#   bgp_asn        = var.customer_side_asn1
#   depends_on = [
#     data.packetfabric_cloud_router_connection.current
#   ]
# }

# ##########################################################################################
# #### Here you would need to setup BGP in your Router
# ##########################################################################################

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
