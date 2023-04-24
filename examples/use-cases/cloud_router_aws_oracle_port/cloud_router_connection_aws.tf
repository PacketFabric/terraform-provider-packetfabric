resource "packetfabric_cloud_provider_credential_aws" "aws_creds1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-aws"
  # using env var PF_AWS_ACCESS_KEY_ID and PF_AWS_SECRET_ACCESS_KEY
}

# From the PacketFabric side: Create a cloud router connection to AWS
resource "packetfabric_cloud_router_connection_aws" "crc_1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  speed       = var.pf_crc_speed1
  # Cloud side provisioning
  cloud_settings {
    credentials_uuid = packetfabric_cloud_provider_credential_aws.aws_creds1.id
    aws_region       = var.aws_region1
    aws_vif_type     = "transit"
    aws_gateways {
      type = "directconnect"
      id   = aws_dx_gateway.direct_connect_gw_1.id
      asn  = var.amazon_side_asn1
    }
    aws_gateways {
      type   = "transit"
      id     = aws_ec2_transit_gateway.transit_gw_1.id
      vpc_id = aws_vpc.vpc_1.id
      subnet_ids = [
        aws_subnet.subnet_1.id
      ]
    }
    bgp_settings {
      orlonger = var.pf_crbs_orlonger
      prefixes {
        prefix = var.aws_vpc_cidr1
        type   = "out" # Allowed Prefixes to Cloud
      }
      prefixes {
        prefix = var.oracle_subnet_cidr1
        type   = "in" # Allowed Prefixes from Cloud
      }
      prefixes {
        prefix = var.on_premise_cidr1
        type   = "in" # Allowed Prefixes from Cloud
      }
    }
  }
}
