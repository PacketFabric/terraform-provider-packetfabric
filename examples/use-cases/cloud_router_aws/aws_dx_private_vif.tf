# From the AWS side: Create and attach a VIF
data "packetfabric_cloud_router_connections" "current" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  depends_on = [
    aws_dx_connection_confirmation.confirmation_1,
    aws_dx_connection_confirmation.confirmation_2,
  ]
}
locals {
  # below may need to be updated
  # check https://github.com/PacketFabric/terraform-provider-packetfabric/issues/23
  cloud_connections = data.packetfabric_cloud_router_connections.current.cloud_connections[*]
  helper_map = { for val in local.cloud_connections :
  val["description"] => val }
  cc1 = local.helper_map["${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"]
  cc2 = local.helper_map["${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"]
}
output "cc1_vlan_id_pf" {
  value = one(local.cc1.cloud_settings[*].vlan_id_pf)
}
output "cc2_vlan_id_pf" {
  value = one(local.cc2.cloud_settings[*].vlan_id_pf)
}
output "packetfabric_cloud_router_connection_aws" {
  value = data.packetfabric_cloud_router_connections.current.cloud_connections[*]
}
resource "aws_dx_private_virtual_interface" "direct_connect_vip_1" {
  provider      = aws
  connection_id = data.aws_dx_connection.current_1.id
  dx_gateway_id = aws_dx_gateway.direct_connect_gw_1.id
  name          = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  # The VLAN is automatically assigned by PacketFabric and available in the packetfabric_cloud_router_connection_aws data source. 
  # We use local in order to parse the data source output and get the VLAN ID assigned by PacketFabric so we can use it to create the VIF in AWS
  vlan           = one(local.cc1.cloud_settings[*].vlan_id_pf)
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
      connection_id
    ]
  }
}
resource "aws_dx_private_virtual_interface" "direct_connect_vip_2" {
  provider       = aws.region2
  connection_id  = data.aws_dx_connection.current_2.id
  dx_gateway_id  = aws_dx_gateway.direct_connect_gw_2.id
  name           = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  vlan           = one(local.cc2.cloud_settings[*].vlan_id_pf)
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
      connection_id
    ]
  }
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
    aws_dx_private_virtual_interface.direct_connect_vip_1
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
    aws_dx_private_virtual_interface.direct_connect_vip_2
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}