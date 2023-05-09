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
}

# To avoid the error conflicting pending workflow when deleting EC2 VPN Gateway Attachment during the destroy
resource "time_sleep" "delay1" {
  create_duration  = "0s"
  destroy_duration = "2m"

  depends_on = [
    aws_vpn_gateway.vpn_gw_1,
    aws_dx_gateway.direct_connect_gw_1
  ]
}
resource "time_sleep" "delay2" {
  create_duration  = "0s"
  destroy_duration = "2m"

  depends_on = [
    aws_vpn_gateway.vpn_gw_2,
    aws_dx_gateway.direct_connect_gw_2
  ]
}