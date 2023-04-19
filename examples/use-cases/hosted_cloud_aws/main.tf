terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.4.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.58.0"
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
    Name = "${var.resource_name}-${random_pet.name.id}"
  }
}

# Define the subnets
resource "aws_subnet" "subnet_1" {
  provider   = aws
  vpc_id     = aws_vpc.vpc_1.id
  cidr_block = var.aws_subnet_cidr1
  tags = {
    Name = "${var.resource_name}-${random_pet.name.id}"
  }
}

# Define the internet gateways
resource "aws_internet_gateway" "gw_1" {
  provider = aws
  vpc_id   = aws_vpc.vpc_1.id
  tags = {
    Name = "${var.resource_name}-${random_pet.name.id}"
  }
}

# Virtual Private Gateway (creation + attachement to the VPC)
resource "aws_vpn_gateway" "vpn_gw_1" {
  provider        = aws
  amazon_side_asn = var.amazon_side_asn1
  tags = {
    Name = "${var.resource_name}-${random_pet.name.id}"
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
    Name = "${var.resource_name}-${random_pet.name.id}"
  }
}

# Assign the route table to the subnet
resource "aws_route_table_association" "route_association_1" {
  provider       = aws
  subnet_id      = aws_subnet.subnet_1.id
  route_table_id = aws_route_table.route_table_1.id
}


# Create Direct Connect Gateway
resource "aws_dx_gateway" "direct_connect_gw_1" {
  provider        = aws
  name            = "${var.resource_name}-${random_pet.name.id}"
  amazon_side_asn = var.amazon_side_asn1
}

# Associate Virtual Private GW to Direct Connect GW
resource "aws_dx_gateway_association" "virtual_private_gw_to_direct_connect_1" {
  provider              = aws
  dx_gateway_id         = aws_dx_gateway.direct_connect_gw_1.id
  associated_gateway_id = aws_vpn_gateway.vpn_gw_1.id
  allowed_prefixes = [
    var.aws_vpc_cidr1
  ]
  depends_on = [
    aws_vpn_gateway_attachment.vpn_attachment_1
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}

# Create a PacketFabric port
resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  autoneg           = var.pf_port_autoneg
  description       = "${var.resource_name}-${random_pet.name.id}"
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
  labels            = var.pf_labels
}
# output "packetfabric_port_1" {
#   value = packetfabric_port.port_1
# }

resource "packetfabric_cloud_provider_credential_aws" "aws_creds1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-aws"
  # using env var PF_AWS_ACCESS_KEY_ID and PF_AWS_SECRET_ACCESS_KEY
}

# Create a AWS Hosted Connection + AWS VIF
resource "packetfabric_cs_aws_hosted_connection" "pf_cs_conn1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  labels      = var.pf_labels
  port        = packetfabric_port.port_1.id
  speed       = var.pf_cs_speed
  pop         = var.pf_cs_pop1
  vlan        = var.pf_cs_vlan1
  zone        = var.pf_cs_zone1
  cloud_settings {
    credentials_uuid = packetfabric_cloud_provider_credential_aws.aws_creds1.id
    aws_region       = var.aws_region1
    mtu              = var.pf_cs_mtu
    aws_vif_type     = "private"
    bgp_settings {
      customer_asn = var.customer_side_asn1
    }
    aws_gateways {
      type = "directconnect"
      id   = aws_dx_gateway.direct_connect_gw_1.id
    }
    aws_gateways {
      type   = "private"
      id     = aws_vpn_gateway.vpn_gw_1.id
      vpc_id = aws_vpc.vpc_1.id
    }
  }
  depends_on = [
    aws_dx_gateway.direct_connect_gw_1
  ]
}
# output "packetfabric_cs_aws_hosted_connection" {
#   value = packetfabric_cs_aws_hosted_connection.pf_cs_conn1
# }

data "packetfabric_cs_hosted_connection_router_config" "router_aws_cisco2900" {
  cloud_circuit_id = packetfabric_cs_aws_hosted_connection.pf_cs_conn1.id
  router_type      = "CiscoSystemsInc-2900SeriesRouters-IOS124"
}
resource "local_file" "router_aws_cisco2900_file" {
  filename = "router_config_aws_cisco2900.txt"
  content  = data.packetfabric_cs_hosted_connection_router_config.router_aws_cisco2900.router_config
}