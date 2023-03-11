# Define the route table
resource "aws_route_table" "route_table_1" {
  provider = aws
  vpc_id   = aws_vpc.vpc_1.id
  # internet gw
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw_1.id
  }
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

