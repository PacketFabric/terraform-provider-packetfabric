# From the AWS side: Create a gateway
resource "aws_dx_gateway" "direct_connect_gw_1" {
  provider        = aws
  name            = "${var.resource_name}-${random_pet.name.id}"
  amazon_side_asn = var.amazon_side_asn1
}


# From the AWS side: Associate Transit GW to Direct Connect GW
resource "aws_dx_gateway_association" "transit_gw_to_direct_connect_1" {
  provider              = aws
  dx_gateway_id         = aws_dx_gateway.direct_connect_gw_1.id
  associated_gateway_id = aws_ec2_transit_gateway.transit_gw_1.id
  # needed for initial creation
  allowed_prefixes = [
    var.aws_vpc_cidr1
  ]
  # allowed_prefixes managed via BGP prefixes in configured in packetfabric_cloud_router_connection_aws
  lifecycle {
    ignore_changes = [
      allowed_prefixes
    ]
  }
  depends_on = [
    aws_ec2_transit_gateway_vpc_attachment.transit_attachment_1
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}