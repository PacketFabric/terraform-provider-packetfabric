
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
  # Need to wait for the transit GW to be attached before adding it to the route table
  depends_on = [
    aws_vpn_gateway_attachment.vpn_attachment_1
  ]
  # # Workaround for https://github.com/hashicorp/terraform-provider-aws/issues/1426
  # lifecycle {
  #   ignore_changes = [
  #     route
  #   ]
  # }
}
resource "aws_route_table" "route_table_2" {
  provider = aws.region2
  vpc_id   = aws_vpc.vpc_2.id
  # internet gw
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw_2.id
  }
  propagating_vgws = ["${aws_vpn_gateway.vpn_gw_2.id}"]
  tags = {
    Name = "${var.tag_name}-${random_pet.name.id}"
  }
  # Need to wait for the transit GW to be attached before adding it to the route table
  depends_on = [
    aws_vpn_gateway_attachment.vpn_attachment_1
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
resource "aws_route_table_association" "route_association_2" {
  provider       = aws.region2
  subnet_id      = aws_subnet.subnet_2.id
  route_table_id = aws_route_table.route_table_2.id
}
