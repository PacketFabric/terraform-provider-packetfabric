# # Virtual Private Gateway (creation + attachement to the VPC)
# resource "aws_vpn_gateway" "vpn_gw_1" {
#   provider        = aws
#   amazon_side_asn = var.amazon_side_asn1
#   tags = {
#     Name = "${var.tag_name}-${random_pet.name.id}"
#   }
#   depends_on = [
#     aws_vpc.vpc_1
#   ]
# }
# resource "aws_vpn_gateway_attachment" "vpn_attachment_1" {
#   provider       = aws
#   vpc_id         = aws_vpc.vpc_1.id
#   vpn_gateway_id = aws_vpn_gateway.vpn_gw_1.id
# }