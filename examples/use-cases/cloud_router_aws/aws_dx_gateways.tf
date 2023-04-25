# From the AWS side: Create a gateway
resource "aws_dx_gateway" "direct_connect_gw_1" {
  provider        = aws
  name            = "${var.resource_name}-${random_pet.name.id}-1"
  amazon_side_asn = var.amazon_side_asn1
  depends_on = [
    aws_vpn_gateway_attachment.vpn_attachment_1
  ]
}
resource "aws_dx_gateway" "direct_connect_gw_2" {
  provider        = aws.region2
  name            = "${var.resource_name}-${random_pet.name.id}-2"
  amazon_side_asn = var.amazon_side_asn2
  depends_on = [
    aws_vpn_gateway_attachment.vpn_attachment_2
  ]
}

# From the AWS side: Associate Virtual Private GW to Direct Connect GW
resource "aws_dx_gateway_association" "virtual_private_gw_to_direct_connect_1" {
  provider              = aws
  dx_gateway_id         = aws_dx_gateway.direct_connect_gw_1.id
  associated_gateway_id = aws_vpn_gateway.vpn_gw_1.id
  # allowed_prefixes managed via BGP prefixes in configured in packetfabric_cloud_router_connection_aws
  depends_on = [
    aws_vpn_gateway_attachment.vpn_attachment_1
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
  # allowed_prefixes managed via BGP prefixes in configured in packetfabric_cloud_router_connection_aws
  depends_on = [
    aws_vpn_gateway_attachment.vpn_attachment_2
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}