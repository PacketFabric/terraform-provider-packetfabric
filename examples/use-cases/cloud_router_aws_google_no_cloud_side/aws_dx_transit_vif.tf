# From the AWS side: Create and attach a VIF
resource "aws_dx_transit_virtual_interface" "direct_connect_vip_1" {
  provider       = aws
  connection_id  = data.aws_dx_connection.current.id
  dx_gateway_id  = aws_dx_gateway.direct_connect_gw_1.id
  name           = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  vlan           = packetfabric_cloud_router_connection_aws.crc_1.vlan_id_pf
  address_family = "ipv4"
  bgp_asn        = var.pf_cr_asn
  depends_on = [
    data.packetfabric_cloud_router_connections.current
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
  lifecycle {
    ignore_changes = [
      connection_id,
      vlan
    ]
  }
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
    aws_dx_transit_virtual_interface.direct_connect_vip_1
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}