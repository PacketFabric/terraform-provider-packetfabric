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

# Define the route table
resource "aws_route_table" "route_table_1" {
  provider = aws
  vpc_id   = aws_vpc.vpc_1.id
  # internet gw
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw_1.id
  }
  # Comment when using Transit Gateway
  # propagating_vgws = ["${aws_vpn_gateway.vpn_gw_1.id}"]
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

