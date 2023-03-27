# From the PacketFabric side: Create a cloud router connection to AWS
resource "packetfabric_cloud_router_connection_aws" "crc_1" {
  provider = packetfabric
  # it is recommended to make sure the connection description is unique as this name will be used to search in AWS later with aws_dx_connection data source
  # vote for this issue https://github.com/hashicorp/terraform-provider-aws/issues/26919 if you want to get the filter added to the aws_dx_connection data source
  description = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  speed       = var.pf_crc_speed
  maybe_nat   = var.pf_crc_maybe_nat
  is_public   = var.pf_crc_is_public
}
resource "packetfabric_cloud_router_connection_aws" "crc_2" {
  provider    = packetfabric
  description = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop2
  zone        = var.pf_crc_zone2
  speed       = var.pf_crc_speed
  maybe_nat   = var.pf_crc_maybe_nat
  is_public   = var.pf_crc_is_public
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

# Retrieve the Direct Connect connections in AWS
data "aws_dx_connection" "current_1" {
  provider   = aws
  name       = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  depends_on = [time_sleep.wait_aws_connection]
}
data "aws_dx_connection" "current_2" {
  provider   = aws.region2
  name       = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  depends_on = [time_sleep.wait_aws_connection]
}
# output "aws_dx_connection_1" {
#   value = data.aws_dx_connection.current_1
# }
# output "aws_dx_connection_2" {
#   value = data.aws_dx_connection.current_2
# }

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
  prefixes {
    prefix = var.aws_vpc_cidr2
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.aws_vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
}
# output "packetfabric_cloud_router_bgp_session_crbs_1" {
#   value = packetfabric_cloud_router_bgp_session.crbs_1
# }

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
  prefixes {
    prefix = var.aws_vpc_cidr1
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.aws_vpc_cidr2
    type   = "in" # Allowed Prefixes from Cloud
  }
}
# output "packetfabric_cloud_router_bgp_session_crbs_2" {
#   value = packetfabric_cloud_router_bgp_session.crbs_2
# }

# # just informative
# data "packetfabric_cloud_router_connections" "all_crc" {
#   provider   = packetfabric
#   circuit_id = packetfabric_cloud_router.cr.id

#   depends_on = [
#     packetfabric_cloud_router_bgp_session.crbs_1,
#     packetfabric_cloud_router_bgp_session.crbs_2
#   ]
# }
# output "packetfabric_cloud_router_connections" {
#   value = data.packetfabric_cloud_router_connections.all_crc
# }