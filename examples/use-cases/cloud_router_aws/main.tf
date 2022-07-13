terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = "0.1.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

# Define default profile
provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region1
  profile    = "region1-profile"
}
# Profile for Region1
provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region1
  profile    = "region1-profile"
  alias      = "region1"
}
# Profile for Region2
provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region2
  profile    = "region2-profile"
  alias      = "region2"
}

# Create random name to use to name objects
resource "random_pet" "name" {}

# Create the VPCs
resource "aws_vpc" "vpc_1" {
  cidr_block           = var.vpc_cidr1
  enable_dns_hostnames = true
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_vpc" "vpc_2" {
  provider             = aws.region2
  cidr_block           = var.vpc_cidr2
  enable_dns_hostnames = true
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Define the public subnets
resource "aws_subnet" "subnet_1" {
  vpc_id            = aws_vpc.vpc_1.id
  cidr_block        = var.public_subnet_cidr1
  availability_zone = var.aws_az1
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_subnet" "subnet_2" {
  provider          = aws.region2
  vpc_id            = aws_vpc.vpc_2.id
  cidr_block        = var.public_subnet_cidr2
  availability_zone = var.aws_az2
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Define the internet gateways
resource "aws_internet_gateway" "gw_1" {
  vpc_id = aws_vpc.vpc_1.id
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_internet_gateway" "gw_2" {
  provider = aws.region2
  vpc_id   = aws_vpc.vpc_2.id
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Virtual Private Gateway (creation + attachement to the VPC)
resource "aws_vpn_gateway" "vpn_gw_1" {
  vpc_id = aws_vpc.vpc_1.id
  amazon_side_asn = var.amazon_side_asn1
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_vpn_gateway" "vpn_gw_2" {
  provider = aws.region2
  vpc_id = aws_vpc.vpc_2.id
  amazon_side_asn = var.amazon_side_asn2
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Define the route tables
resource "aws_route_table" "route_table_1" {
  vpc_id = aws_vpc.vpc_1.id
  # internet gw
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw_1.id
  }
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_route_table" "route_table_2" {
  provider = aws.region2
  vpc_id   = aws_vpc.vpc_2.id
  # internet gw
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw_2.id
  }
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Assign the route table to the public subnet
resource "aws_route_table_association" "route_association_1" {
  subnet_id      = aws_subnet.subnet_1.id
  route_table_id = aws_route_table.route_table_1.id
}
resource "aws_route_table_association" "route_association_2" {
  provider       = aws.region2
  subnet_id      = aws_subnet.subnet_2.id
  route_table_id = aws_route_table.route_table_2.id
}

resource "aws_security_group" "ingress_all_1" {
  name   = "allow-icmp-ssh-http-locust-iperf-sg"
  vpc_id = aws_vpc.vpc_1.id
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 8089
    to_port     = 8089
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 5001
    to_port     = 5001
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = -1
    to_port = -1
    protocol = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  // Terraform removes the default rule
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_security_group" "ingress_all_2" {
  provider = aws.region2
  name   = "allow-icmp-ssh-http-locust-iperf-sg"
  vpc_id   = aws_vpc.vpc_2.id
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 8089
    to_port     = 8089
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 5001
    to_port     = 5001
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = -1
    to_port = -1
    protocol = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  // Terraform removes the default rule
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Create NIC for the EC2 instances
resource "aws_network_interface" "nic1" {
  subnet_id       = aws_subnet.subnet_1.id
  security_groups = ["${aws_security_group.ingress_all_1.id}"]
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_network_interface" "nic2" {
  provider        = aws.region2
  subnet_id       = aws_subnet.subnet_2.id
  security_groups = ["${aws_security_group.ingress_all_2.id}"]
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Create the Key Pair
resource "aws_key_pair" "ssh_key_1" {
  key_name   = "ssh_key-${random_pet.name.id}"
  public_key = var.public_key
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_key_pair" "ssh_key_2" {
  provider   = aws.region2
  key_name   = "ssh_key-${random_pet.name.id}"
  public_key = var.public_key
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Create the Ubuntu EC2 instances
resource "aws_instance" "ec2_instance_1" {
  ami           = var.ec2_ami1
  instance_type = var.ec2_instance_type
  network_interface {
    network_interface_id = aws_network_interface.nic1.id
    device_index         = 0
  }
  key_name  = aws_key_pair.ssh_key_1.id
  user_data = file("user-data-ubuntu.sh")
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_instance" "ec2_instance_2" {
  provider      = aws.region2
  ami           = var.ec2_ami2
  instance_type = var.ec2_instance_type
  network_interface {
    network_interface_id = aws_network_interface.nic2.id
    device_index         = 0
  }
  key_name  = aws_key_pair.ssh_key_2.id
  user_data = file("user-data-ubuntu.sh")
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# Assign a public IP to both EC2 instances
resource "aws_eip" "public_ip_1" {
  instance = aws_instance.ec2_instance_1.id
  vpc      = true
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}
resource "aws_eip" "public_ip_2" {
  provider = aws.region2
  instance = aws_instance.ec2_instance_2.id
  vpc      = true
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
}

# From the PacketFabric side: Create a cloud router
resource "cloud_router" "cr" {
  provider     = packetfabric
  scope        = var.pf_cr_scope
  asn          = var.pf_cr_asn
  name         = "${var.tag_name}-${random_pet.name.id}"
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

# From the PacketFabric side: Create a cloud router connection to AWS
resource "aws_cloud_router_connection" "crc_1" {
  provider       = packetfabric
  circuit_id     = cloud_router.cr.id
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  maybe_nat      = var.pf_crc_maybe_nat
  description    = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  pop            = var.pf_crc_pop1
  zone           = var.pf_crc_zone1
  is_public      = var.pf_crc_is_public
  speed          = var.pf_crc_speed
  depends_on = [
    cloud_router.cr
  ]
}
resource "aws_cloud_router_connection" "crc_2" {
  provider       = packetfabric
  circuit_id     = cloud_router.cr.id
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  maybe_nat      = var.pf_crc_maybe_nat
  description    = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  pop            = var.pf_crc_pop2
  zone           = var.pf_crc_zone2
  is_public      = var.pf_crc_is_public
  speed          = var.pf_crc_speed
  depends_on = [
    cloud_router.cr
  ]
}

# Wait 30s for the connection to show up in AWS
resource "null_resource" "previous" {}
resource "time_sleep" "wait_30_seconds" {
  depends_on = [null_resource.previous]
  create_duration = "30s"
}
# This resource will create (at least) 30 seconds after null_resource.previous
resource "null_resource" "next" {
  depends_on = [time_sleep.wait_30_seconds]
}

# Retrieve the Direct Connect connection in AWS
data "aws_dx_connection" "current_1" {
  provider = aws
  name     = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  depends_on = [
    null_resource.next
  ]
}
data "aws_dx_connection" "current_2" {
  provider = aws.region2
  name     = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  depends_on = [
    null_resource.next
  ]
}
# From the AWS side: Accept the connection
resource "aws_dx_connection_confirmation" "confirmation_1" {
  provider      = aws
  connection_id = data.aws_dx_connection.current_1.id
  depends_on = [
    data.aws_dx_connection.current_1
  ]
}
resource "aws_dx_connection_confirmation" "confirmation_2" {
  provider      = aws.region2
  connection_id = data.aws_dx_connection.current_2.id
  depends_on = [
    data.aws_dx_connection.current_2
  ]
}

# From the AWS side: Create a gateway
resource "aws_dx_gateway" "direct_connect_gw_1" {
  provider        = aws
  name            = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  amazon_side_asn = var.amazon_side_asn1
  depends_on = [
    aws_cloud_router_connection.crc_1
  ]
}
resource "aws_dx_gateway" "direct_connect_gw_2" {
  provider        = aws.region2
  name            = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  amazon_side_asn = var.amazon_side_asn2
  depends_on = [
    aws_cloud_router_connection.crc_2
  ]
}

# From the AWS side: Create and attach a VIF
data "aws_cloud_router_connection" "current" {
  provider   = packetfabric
  circuit_id = cloud_router.cr.id

  depends_on = [
    aws_dx_connection_confirmation.confirmation_1,
    aws_dx_connection_confirmation.confirmation_2,
  ]
}
locals {
  aws_cloud_connections = data.aws_cloud_router_connection.current.aws_cloud_connections[*]
  helper_map = {for val in local.aws_cloud_connections:
              val["pop"]=>val}
  cc1 = local.helper_map["${var.pf_crc_pop1}"]
  cc2 = local.helper_map["${var.pf_crc_pop2}"]
}
output "aws_cloud_router_connection" {
  value = data.aws_cloud_router_connection.current.aws_cloud_connections[*]
}
resource "aws_dx_private_virtual_interface" "direct_connect_vip_1" {
  provider       = aws
  connection_id  = data.aws_dx_connection.current_1.id
  dx_gateway_id  = aws_dx_gateway.direct_connect_gw_1.id
  name           = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  vlan           = one(local.cc1.cloud_settings[*].vlan_id_pf)
  address_family = "ipv4"
  bgp_asn        = var.pf_cr_asn
  depends_on = [
    data.aws_cloud_router_connection.current
  ]
}
resource "aws_dx_private_virtual_interface" "direct_connect_vip_2" {
  provider       = aws.region2
  connection_id  = data.aws_dx_connection.current_2.id
  dx_gateway_id  = aws_dx_gateway.direct_connect_gw_2.id
  name           = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  vlan           = one(local.cc2.cloud_settings[*].vlan_id_pf)
  address_family = "ipv4"
  bgp_asn        = var.pf_cr_asn
  depends_on = [
    data.aws_cloud_router_connection.current
  ]
}

# From the AWS side: Associate Virtual Private GW to Direct Connect GW
resource "aws_dx_gateway_association" "virtual_private_gw_to_direct_connect_1" {
  provider       = aws
  dx_gateway_id         = aws_dx_gateway.direct_connect_gw_1.id
  associated_gateway_id = aws_vpn_gateway.vpn_gw_1.id
  allowed_prefixes = [
    var.vpc_cidr1,
    var.vpc_cidr2
  ]
  depends_on = [
    aws_dx_private_virtual_interface.direct_connect_vip_1
  ]
}
resource "aws_dx_gateway_association" "virtual_private_gw_to_direct_connect_2" {
  provider       = aws.region2
  dx_gateway_id         = aws_dx_gateway.direct_connect_gw_2.id
  associated_gateway_id = aws_vpn_gateway.vpn_gw_2.id
  allowed_prefixes = [
    var.vpc_cidr1,
    var.vpc_cidr2
  ]
  depends_on = [
    aws_dx_private_virtual_interface.direct_connect_vip_2
  ]
}

# From the PacketFabric side: Configure BGP
resource "cloud_router_bgp_session" "crbs_1" {
  provider       = packetfabric
  circuit_id     = cloud_router.cr.id
  connection_id  = aws_cloud_router_connection.crc_1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.amazon_side_asn1
  orlonger       = var.pf_crbs_orlonger
  remote_address = aws_dx_private_virtual_interface.direct_connect_vip_1.amazon_address # AWS side
  l3_address     = aws_dx_private_virtual_interface.direct_connect_vip_1.customer_address # PF side
  md5            = aws_dx_private_virtual_interface.direct_connect_vip_1.bgp_auth_key
  depends_on = [
    aws_dx_private_virtual_interface.direct_connect_vip_1
  ]
}
resource "cloud_router_bgp_prefixes" "crbp_1" {
  provider = packetfabric
  bgp_settings_uuid = cloud_router_bgp_session.crbs_1.id
  prefixes {
    prefix = var.vpc_cidr1
    type = "in"
    order = 0
  }
  prefixes {
    prefix = var.vpc_cidr2
    type = "in"
    order = 0
  }
  prefixes {
    prefix = var.vpc_cidr1
    type = "out"
    order = 0
  }
  prefixes {
    prefix = var.vpc_cidr2
    type = "out"
    order = 0
  }
  depends_on = [
    cloud_router_bgp_session.crbs_1
  ]
}

resource "cloud_router_bgp_session" "crbs_2" {
  provider       = packetfabric
  circuit_id     = cloud_router.cr.id
  connection_id  = aws_cloud_router_connection.crc_2.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.amazon_side_asn2
  orlonger       = var.pf_crbs_orlonger
  remote_address = aws_dx_private_virtual_interface.direct_connect_vip_2.amazon_address # AWS side
  l3_address     = aws_dx_private_virtual_interface.direct_connect_vip_2.customer_address # PF side
  md5            = aws_dx_private_virtual_interface.direct_connect_vip_2.bgp_auth_key
  depends_on = [
    aws_dx_private_virtual_interface.direct_connect_vip_2
  ]
}
resource "cloud_router_bgp_prefixes" "crbp_2" {
  provider = packetfabric
  bgp_settings_uuid = cloud_router_bgp_session.crbs_2.id
  prefixes {
    prefix = var.vpc_cidr1
    type = "in"
    order = 0
  }
  prefixes {
    prefix = var.vpc_cidr2
    type = "in"
    order = 0
  }
  prefixes {
    prefix = var.vpc_cidr1
    type = "out"
    order = 0
  }
  prefixes {
    prefix = var.vpc_cidr2
    type = "out"
    order = 0
  }
  depends_on = [
    cloud_router_bgp_session.crbs_2
  ]
}

data "cloud_router" "current" {
  provider = packetfabric
  depends_on = [
    cloud_router.cr
  ]
}

output "cloud_router" {
  value = data.cloud_router.current
}


# Private IPs of the demo Ubuntu instances
output "ec2_private_ip_1" {
  description = "Private ip address for EC2 instance for Region 1"
  value = aws_instance.ec2_instance_1.private_ip
}

output "ec2_private_ip_2" {
  description = "Private ip address for EC2 instance for Region 2"
  value = aws_instance.ec2_instance_2.private_ip
}

# Public IPs of the demo Ubuntu instances
output "ec2_public_ip_1" {
  description = "Elastic ip address for EC2 instance for Region 1"
  value = aws_eip.public_ip_1.public_ip
}

output "ec2_public_ip_2" {
  description = "Elastic ip address for EC2 instance for Region 2"
  value = aws_eip.public_ip_2.public_ip
}
