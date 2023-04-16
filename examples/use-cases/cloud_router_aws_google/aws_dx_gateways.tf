# From the AWS side: Create a gateway
resource "aws_dx_gateway" "direct_connect_gw_1" {
  provider        = aws
  name            = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  amazon_side_asn = var.amazon_side_asn1
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_1
  ]
}


# From the AWS side: Associate Transit GW to Direct Connect GW
resource "aws_dx_gateway_association" "transit_gw_to_direct_connect_1" {
  provider              = aws
  dx_gateway_id         = aws_dx_gateway.direct_connect_gw_1.id
  associated_gateway_id = aws_ec2_transit_gateway.transit_gw_1.id
  allowed_prefixes = [
    var.aws_vpc_cidr1
  ]
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_1
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}