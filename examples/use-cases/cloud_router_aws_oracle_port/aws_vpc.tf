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