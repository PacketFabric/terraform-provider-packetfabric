# Define the route table
resource "aws_route_table" "route_table_1" {
  provider = aws
  vpc_id   = aws_vpc.vpc_1.id
  tags = {
    Name = "${var.resource_name}-${random_pet.name.id}"
  }
  # Need to wait for the transit GW to be attached before adding it to the route table
  depends_on = [
    aws_ec2_transit_gateway_vpc_attachment.transit_attachment_1
  ]
}

resource "aws_route" "route_table_1_to_internet_gateway" {
  route_table_id         = aws_route_table.route_table_1.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.gw_1.id
}

resource "aws_route" "route_table_1_to_transit_gateway" {
  route_table_id         = aws_route_table.route_table_1.id
  destination_cidr_block = var.gcp_subnet_cidr1
  # destination_cidr_block = "0.0.0.0/0"
  transit_gateway_id     = aws_ec2_transit_gateway.transit_gw_1.id
}

# Assign the route table to the subnet
resource "aws_route_table_association" "route_association_1" {
  provider       = aws
  subnet_id      = aws_subnet.subnet_1.id
  route_table_id = aws_route_table.route_table_1.id
}

