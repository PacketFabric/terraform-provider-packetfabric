# From the PacketFabric side: Create a cloud router connection to AWS
resource "packetfabric_cloud_router_connection_aws" "crc_1" {
  provider = packetfabric
  # it is recommended to make sure the connection description is unique as this name will be used to search in AWS later with aws_dx_connection data source
  # vote for this issue https://github.com/hashicorp/terraform-provider-aws/issues/26919 if you want to get the filter added to the aws_dx_connection data source
  description = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  speed       = var.pf_crc_speed
  maybe_nat   = var.pf_crc_maybe_nat
  is_public   = var.pf_crc_is_public
}

# From the AWS side: Accept the connection
# Wait for the connection to show up in AWS
resource "time_sleep" "wait_aws_connection" {
  create_duration = "2m"
  depends_on = [
    packetfabric_cloud_router_connection_aws.crc_1
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
output "aws_dx_connection_1" {
  value = data.aws_dx_connection.current_1
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

# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc_1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.amazon_side_asn1
  orlonger       = var.pf_crbs_orlonger
  remote_address = aws_dx_transit_virtual_interface.direct_connect_vip_1.amazon_address   # AWS side
  l3_address     = aws_dx_transit_virtual_interface.direct_connect_vip_1.customer_address # PF side
  md5            = aws_dx_transit_virtual_interface.direct_connect_vip_1.bgp_auth_key
  prefixes {
    prefix = var.gcp_subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
    order  = 0
  }
  prefixes {
    prefix = var.aws_vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
    order  = 0
  }
}
output "packetfabric_cloud_router_bgp_session_crbs_1" {
  value = packetfabric_cloud_router_bgp_session.crbs_1
}