# From the AWS side: Create and attach a VIF
data "packetfabric_cloud_router_connections" "current" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  depends_on = [
    aws_dx_connection_confirmation.confirmation
  ]
}
locals {
  cloud_connections = data.packetfabric_cloud_router_connections.current.cloud_connections[*]
  helper_map = { for val in local.cloud_connections :
  val["description"] => val }
  cc1 = local.helper_map["${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"]
}
# output "cc1_vlan_id_pf" {
#   value = one(local.cc1.cloud_settings[*].vlan_id_pf)
# }
# output "packetfabric_cloud_router_connection_aws" {
#   value = data.packetfabric_cloud_router_connections.current.cloud_connections[*]
# }
resource "aws_dx_transit_virtual_interface" "direct_connect_vip_1" {
  provider      = aws
  connection_id = data.aws_dx_connection.current.id
  dx_gateway_id = aws_dx_gateway.direct_connect_gw_1.id
  name          = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  # The VLAN is automatically assigned by PacketFabric and available in the packetfabric_cloud_router_connection data source. 
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