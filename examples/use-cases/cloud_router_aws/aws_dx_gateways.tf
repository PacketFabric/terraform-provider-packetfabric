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
