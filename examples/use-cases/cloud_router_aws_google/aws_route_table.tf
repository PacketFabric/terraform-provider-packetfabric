# Define the route table
resource "aws_route_table" "route_table_1" {
  provider = aws
  vpc_id   = aws_vpc.vpc_1.id
  # internet gw
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw_1.id
  }
  route {
    cidr_block = var.gcp_subnet_cidr1
    gateway_id = aws_ec2_transit_gateway.transit_gw_1.id
  }
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
  # Need to wait for the transit GW to be attached before adding it to the route table
  depends_on = [
    aws_ec2_transit_gateway_vpc_attachment.transit_attachment_1
  ]
  # # Workaround for https://github.com/hashicorp/terraform-provider-aws/issues/1426
  # lifecycle {
  #   ignore_changes = [
  #     route
  #   ]
  # }
}

# Assign the route table to the subnet
resource "aws_route_table_association" "route_association_1" {
  provider       = aws
  subnet_id      = aws_subnet.subnet_1.id
  route_table_id = aws_route_table.route_table_1.id
}

