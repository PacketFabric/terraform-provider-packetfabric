# From the PacketFabric side: Create a cloud router connection to AWS
resource "packetfabric_cloud_router_connection_aws" "crc_1" {
  provider = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  speed       = var.pf_crc_speed
  maybe_nat   = var.pf_crc_maybe_nat
  is_public   = var.pf_crc_is_public
  cloud_settings {
    credentials_uuid = packetfabric_cloud_provider_credential_aws.aws_creds1.id
    aws_region       = var.aws_region1
    aws_vif_type     = "transit"
    aws_gateways {
      type = "directconnect"
      id   = aws_dx_gateway.direct_connect_gw_1.id
    }
    aws_gateways {
      type   = "transit"
      id     = aws_ec2_transit_gateway.transit_gw_1.id
      vpc_id = aws_vpc.vpc_1.id
    }
    bgp_settings {
      multihop_ttl   = var.pf_crbs_mhttl
      remote_asn     = var.amazon_side_asn1
      orlonger       = var.pf_crbs_orlonger
      prefixes {
        prefix = var.gcp_subnet_cidr1
        type   = "out" # Allowed Prefixes to Cloud
      }
      prefixes {
        prefix = var.aws_vpc_cidr1
        type   = "in" # Allowed Prefixes from Cloud
      }
    }
  }
}
