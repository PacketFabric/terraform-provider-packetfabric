# Transit Gateway (creation + attachement to the VPC)
resource "aws_ec2_transit_gateway" "transit_gw_1" {
  provider        = aws
  description     = "${var.resource_name}-${random_pet.name.id}"
  amazon_side_asn = var.amazon_side_asn2
  tags = {
    Name = "${var.resource_name}-${random_pet.name.id}"
  }
  depends_on = [
    aws_vpc.vpc_1
  ]
}

resource "aws_ec2_transit_gateway_vpc_attachment" "transit_attachment_1" {
  provider           = aws
  vpc_id             = aws_vpc.vpc_1.id
  transit_gateway_id = aws_ec2_transit_gateway.transit_gw_1.id
  subnet_ids = [
    aws_subnet.subnet_1.id
  ]
}

# The transit gateway to VPC attachment will propogate the route from the TGW to 
# the DX gateway all the way to the VPC and the subnets