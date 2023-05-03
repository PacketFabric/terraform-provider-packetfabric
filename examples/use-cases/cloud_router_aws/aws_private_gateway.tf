# Virtual Private Gateway (creation + attachement to the VPC)
resource "aws_vpn_gateway" "vpn_gw_1" {
  provider        = aws
  amazon_side_asn = var.amazon_side_asn1
  vpc_id          = aws_vpc.vpc_1.id
  tags = {
    Name = "${var.resource_name}-${random_pet.name.id}"
  }
}
resource "aws_vpn_gateway" "vpn_gw_2" {
  provider        = aws.region2
  amazon_side_asn = var.amazon_side_asn2
  vpc_id          = aws_vpc.vpc_2.id
  tags = {
    Name = "${var.resource_name}-${random_pet.name.id}"
  }
  depends_on = [
    aws_vpc.vpc_2
  ]
}