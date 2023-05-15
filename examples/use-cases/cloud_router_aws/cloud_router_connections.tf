resource "packetfabric_cloud_provider_credential_aws" "aws_creds1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-aws"
  # using env var PF_AWS_ACCESS_KEY_ID and PF_AWS_SECRET_ACCESS_KEY
}

# From the PacketFabric side: Create a cloud router connection to AWS
resource "packetfabric_cloud_router_connection_aws" "crc_aws1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  speed       = var.pf_crc_speed
  # Cloud side provisioning
  cloud_settings {
    credentials_uuid = packetfabric_cloud_provider_credential_aws.aws_creds1.id
    aws_region       = var.aws_region1
    aws_vif_type     = "private"
    aws_gateways {
      type = "directconnect"
      id   = aws_dx_gateway.direct_connect_gw_1.id
      asn  = var.amazon_side_asn1
    }
    aws_gateways {
      type   = "private"
      id     = aws_vpn_gateway.vpn_gw_1.id
      vpc_id = aws_vpc.vpc_1.id
    }
    bgp_settings {
      orlonger = var.pf_crbs_orlonger
      prefixes {
        prefix = var.aws_vpc_cidr2
        type   = "out" # Allowed Prefixes to Cloud
      }
      prefixes {
        prefix = var.aws_vpc_cidr1
        type   = "in" # Allowed Prefixes from Cloud
      }
    }
  }
  depends_on = [
    aws_dx_gateway_association.virtual_private_gw_to_direct_connect_1
  ]
}

resource "packetfabric_cloud_router_connection_aws" "crc_aws2" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop2
  zone        = var.pf_crc_zone2
  speed       = var.pf_crc_speed
  # Cloud side provisioning
  cloud_settings {
    credentials_uuid = packetfabric_cloud_provider_credential_aws.aws_creds1.id
    aws_region       = var.aws_region2
    aws_vif_type     = "private"
    aws_gateways {
      type = "directconnect"
      id   = aws_dx_gateway.direct_connect_gw_2.id
    }
    aws_gateways {
      type   = "private"
      id     = aws_vpn_gateway.vpn_gw_2.id
      vpc_id = aws_vpc.vpc_2.id
    }
    bgp_settings {
      orlonger = var.pf_crbs_orlonger
      prefixes {
        prefix = var.aws_vpc_cidr1
        type   = "out" # Allowed Prefixes to Cloud
      }
      prefixes {
        prefix = var.aws_vpc_cidr2
        type   = "in" # Allowed Prefixes from Cloud
      }
    }
  }
  depends_on = [
    aws_dx_gateway_association.virtual_private_gw_to_direct_connect_2
  ]
}
