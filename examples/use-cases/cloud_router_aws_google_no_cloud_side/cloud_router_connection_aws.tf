# From the PacketFabric side: Create a cloud router connection to AWS
resource "packetfabric_cloud_router_connection_aws" "crc_aws" {
  provider = packetfabric
  # it is recommended to make sure the connection description is unique as this name will be used to search in AWS later with aws_dx_connection data source
  # vote for this issue https://github.com/hashicorp/terraform-provider-aws/issues/26919 if you want to get the filter added to the aws_dx_connection data source
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  speed       = var.pf_crc_speed
  maybe_nat   = var.pf_crc_maybe_nat
  is_public   = var.pf_crc_is_public
}
# output "packetfabric_cloud_router_connection_aws_crc_aws" {
#   value = packetfabric_cloud_router_connection_aws.crc_aws
# }

resource "aws_dx_connection_confirmation" "confirmation" {
  provider      = aws
  connection_id = packetfabric_cloud_router_connection_aws.crc_aws.cloud_provider_connection_id
}

# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_aws" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc_aws.id
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
  }
  prefixes {
    prefix = var.aws_vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
}
# output "packetfabric_cloud_router_bgp_session_crbs_aws" {
#   value = packetfabric_cloud_router_bgp_session.crbs_aws
# }