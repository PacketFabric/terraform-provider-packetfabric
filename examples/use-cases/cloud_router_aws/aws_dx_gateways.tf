# From the AWS side: Create a gateway
resource "aws_dx_gateway" "direct_connect_gw_1" {
  provider        = aws
  name            = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  amazon_side_asn = var.amazon_side_asn1
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_1
  ]
}
resource "aws_dx_gateway" "direct_connect_gw_2" {
  provider        = aws.region2
  name            = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  amazon_side_asn = var.amazon_side_asn2
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_2
  ]
}

# From the AWS side: Associate Virtual Private GW to Direct Connect GW
resource "aws_dx_gateway_association" "virtual_private_gw_to_direct_connect_1" {
  provider              = aws
  dx_gateway_id         = aws_dx_gateway.direct_connect_gw_1.id
  associated_gateway_id = aws_vpn_gateway.vpn_gw_1.id
  allowed_prefixes = [
    var.aws_vpc_cidr1,
    var.aws_vpc_cidr2
  ]
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_1
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}
resource "aws_dx_gateway_association" "virtual_private_gw_to_direct_connect_2" {
  provider              = aws.region2
  dx_gateway_id         = aws_dx_gateway.direct_connect_gw_2.id
  associated_gateway_id = aws_vpn_gateway.vpn_gw_2.id
  allowed_prefixes = [
    var.aws_vpc_cidr1,
    var.aws_vpc_cidr2
  ]
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_2
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}