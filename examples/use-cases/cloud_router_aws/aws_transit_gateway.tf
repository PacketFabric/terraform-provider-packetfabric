# # Transit Gateway (creation + attachement to the VPC)
# resource "aws_ec2_transit_gateway" "transit_gw_1" {
#   provider        = aws
#   description = "${var.tag_name}-${random_pet.name.id}"
#   amazon_side_asn = var.amazon_side_asn1
#   tags = {
#     Name = "${var.tag_name}-${random_pet.name.id}"
#   }
#   depends_on = [
#     aws_vpc.vpc_1
#   ]
# }

# resource "aws_ec2_transit_gateway_vpc_attachment" "transit_attachment_1" {
#   provider        = aws
#   vpc_id             = aws_vpc.vpc_1.id
#   transit_gateway_id = aws_ec2_transit_gateway.transit_gw_1.id
#   subnet_ids         = [
#     aws_subnet.subnet_1.id
#   ]
# }

# resource "aws_ec2_transit_gateway" "transit_gw_2" {
#   provider        = aws.region2
#   description = "${var.tag_name}-${random_pet.name.id}"
#   amazon_side_asn = var.amazon_side_asn2
#   tags = {
#     Name = "${var.tag_name}-${random_pet.name.id}"
#   }
#   depends_on = [
#     aws_vpc.vpc_2
#   ]
# }
# resource "aws_ec2_transit_gateway_vpc_attachment" "transit_attachment_2" {
#   provider        = aws.region2
#   vpc_id             = aws_vpc.vpc_2.id
#   transit_gateway_id = aws_ec2_transit_gateway.transit_gw_2.id
#   subnet_ids         = [
#     aws_subnet.subnet_2.id
#   ]
# }

# # Define the route tables
# resource "aws_route_table" "route_table_1" {
#   provider = aws
#   vpc_id   = aws_vpc.vpc_1.id
#   # internet gw
#   route {
#     cidr_block = "0.0.0.0/0"
#     gateway_id = aws_internet_gateway.gw_1.id
#   }
#   # transit gw
#   route {
#     cidr_block = var.vpc_cidr2
#     transit_gateway_id = aws_ec2_transit_gateway.transit_gw_1.id
#   }
#   tags = {
#     Name = "${var.tag_name}-${random_pet.name.id}"
#   }
# }
# resource "aws_route_table" "route_table_2" {
#   provider = aws.region2
#   vpc_id   = aws_vpc.vpc_2.id
#   # internet gw
#   route {
#     cidr_block = "0.0.0.0/0"
#     gateway_id = aws_internet_gateway.gw_2.id
#   }
#   # transit gw
#   route {
#     cidr_block = var.vpc_cidr1
#     transit_gateway_id = aws_ec2_transit_gateway.transit_gw_2.id
#   }
#   tags = {
#     Name = "${var.tag_name}-${random_pet.name.id}"
#   }
# }
