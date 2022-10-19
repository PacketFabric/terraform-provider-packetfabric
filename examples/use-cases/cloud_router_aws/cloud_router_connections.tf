# From the PacketFabric side: Create a cloud router connection to AWS
resource "packetfabric_cloud_router_connection_aws" "crc_1" {
  provider = packetfabric
  # it is recommended to make sure the connection description is unique as this name will be used to search in AWS later with aws_dx_connection data source
  # vote for this issue https://github.com/hashicorp/terraform-provider-aws/issues/26919 if you want to get the filter added to the aws_dx_connection data source
  description    = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  circuit_id     = packetfabric_cloud_router.cr.id
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  pop            = var.pf_crc_pop1
  zone           = var.pf_crc_zone1
  speed          = var.pf_crc_speed
  maybe_nat      = var.pf_crc_maybe_nat
  is_public      = var.pf_crc_is_public
}
resource "packetfabric_cloud_router_connection_aws" "crc_2" {
  provider       = packetfabric
  description    = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  circuit_id     = packetfabric_cloud_router.cr.id
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  pop            = var.pf_crc_pop2
  zone           = var.pf_crc_zone2
  speed          = var.pf_crc_speed
  maybe_nat      = var.pf_crc_maybe_nat
  is_public      = var.pf_crc_is_public
}

# From the AWS side: Accept the connection
# Wait for the connection to show up in AWS
resource "time_sleep" "wait_aws_connection" {
  create_duration = "2m"
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_1,
    packetfabric_cloud_router_connection_aws.crc_2
  ]
}
resource "null_resource" "next_aws_connection" {
  depends_on = [time_sleep.wait_aws_connection]
}

# Retrieve the Direct Connect connections in AWS
data "aws_dx_connection" "current_1" {
  provider = aws
  name     = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  depends_on = [
    null_resource.next_aws_connection,
    packetfabric_cloud_router_connection_aws.crc_1
  ]
}
data "aws_dx_connection" "current_2" {
  provider = aws.region2
  name     = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  depends_on = [
    null_resource.next_aws_connection,
    packetfabric_cloud_router_connection_aws.crc_2
  ]
}
output "aws_dx_connection_1" {
  value = data.aws_dx_connection.current_1
}
output "aws_dx_connection_2" {
  value = data.aws_dx_connection.current_2
}

# Sometimes, it takes more than 10min for the connection to become available
# Vote for below issue to get a timeout attribute added to the aws_dx_connection_confirmation resource
# https://github.com/hashicorp/terraform-provider-aws/issues/26335

resource "aws_dx_connection_confirmation" "confirmation_1" {
  provider      = aws
  connection_id = data.aws_dx_connection.current_1.id

  lifecycle {
    ignore_changes = [
      connection_id
    ]
  }
}
resource "aws_dx_connection_confirmation" "confirmation_2" {
  provider      = aws.region2
  connection_id = data.aws_dx_connection.current_2.id

  lifecycle {
    ignore_changes = [
      connection_id
    ]
  }
}

# From the AWS side: Create a gateway
resource "aws_dx_gateway" "direct_connect_gw_1" {
  provider        = aws
  name            = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  amazon_side_asn = var.amazon_side_asn1
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_1
  ]
}
resource "aws_dx_gateway" "direct_connect_gw_2" {
  provider        = aws.region2
  name            = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  amazon_side_asn = var.amazon_side_asn2
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_2
  ]
}

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
    var.vpc_cidr1,
    var.vpc_cidr2
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
    var.vpc_cidr1,
    var.vpc_cidr2
  ]
  depends_on = [
    aws_dx_private_virtual_interface.direct_connect_vip_2
  ]
  timeouts {
    create = "2h"
    delete = "2h"
  }
}

# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc_1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.amazon_side_asn1
  orlonger       = var.pf_crbs_orlonger
  remote_address = aws_dx_private_virtual_interface.direct_connect_vip_1.amazon_address   # AWS side
  l3_address     = aws_dx_private_virtual_interface.direct_connect_vip_1.customer_address # PF side
  md5            = aws_dx_private_virtual_interface.direct_connect_vip_1.bgp_auth_key
}
output "packetfabric_cloud_router_bgp_session_crbs_1" {
  value = packetfabric_cloud_router_bgp_session.crbs_1
}

# Configure BGP Prefix is mandatory to setup the BGP session correctly
resource "packetfabric_cloud_router_bgp_prefixes" "crbp_1" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_1.id
  prefixes {
    prefix = var.vpc_cidr2
    type   = "out" # Allowed Prefixes to Cloud
    order  = 0
  }
  prefixes {
    prefix = var.vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
    order  = 0
  }
}
data "packetfabric_cloud_router_bgp_prefixes" "bgp_prefix_crbp_1" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_1.id
}
output "packetfabric_bgp_prefix_crbp_1" {
  value = data.packetfabric_cloud_router_bgp_prefixes.bgp_prefix_crbp_1
}

resource "packetfabric_cloud_router_bgp_session" "crbs_2" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc_2.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.amazon_side_asn2
  orlonger       = var.pf_crbs_orlonger
  remote_address = aws_dx_private_virtual_interface.direct_connect_vip_2.amazon_address   # AWS side
  l3_address     = aws_dx_private_virtual_interface.direct_connect_vip_2.customer_address # PF side
  md5            = aws_dx_private_virtual_interface.direct_connect_vip_2.bgp_auth_key
}
output "packetfabric_cloud_router_bgp_session_crbs_2" {
  value = packetfabric_cloud_router_bgp_session.crbs_2
}

# Configure BGP Prefix is mandatory to setup the BGP session correctly
resource "packetfabric_cloud_router_bgp_prefixes" "crbp_2" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_2.id
  prefixes {
    prefix = var.vpc_cidr1
    type   = "out" # Allowed Prefixes to Cloud
    order  = 0
  }
  prefixes {
    prefix = var.vpc_cidr2
    type   = "in" # Allowed Prefixes from Cloud
    order  = 0
  }
}
data "packetfabric_cloud_router_bgp_prefixes" "bgp_prefix_crbp_2" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_2.id
}
output "packetfabric_bgp_prefix_crbp_2" {
  value = data.packetfabric_cloud_router_bgp_prefixes.bgp_prefix_crbp_2
}

data "packetfabric_cloud_router_connections" "all_crc" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  depends_on = [
    packetfabric_cloud_router_bgp_session.crbs_1,
    packetfabric_cloud_router_bgp_session.crbs_2
  ]
}
output "packetfabric_cloud_router_connections" {
  value = data.packetfabric_cloud_router_connections.all_crc
}