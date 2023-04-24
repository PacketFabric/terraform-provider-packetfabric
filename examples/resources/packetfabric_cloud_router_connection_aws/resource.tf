resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
}

# Example PacketFabric side provisioning only
resource "packetfabric_cloud_router_connection_aws" "crc1" {
  provider    = packetfabric
  circuit_id  = packetfabric_cloud_router.cr1.id
  maybe_nat   = false
  description = "hello world"
  pop         = "PDX2"
  zone        = "A"
  is_public   = false
  speed       = "1Gbps"
  labels      = ["terraform", "dev"]
}

# Example PacketFabric side + AWS side provisioning
resource "packetfabric_cloud_provider_credential_aws" "aws_creds1" {
  provider       = packetfabric
  description    = "AWS Staging Environement"
  aws_access_key = var.pf_aws_key    # or use env var PF_AWS_ACCESS_KEY_ID
  aws_secret_key = var.pf_aws_secret # or use env var PF_AWS_SECRET_ACCESS_KEY
}

resource "packetfabric_cloud_router_connection_aws" "crc1" {
  provider    = packetfabric
  circuit_id  = packetfabric_cloud_router.cr1.id
  maybe_nat   = false
  description = "hello world"
  pop         = "PDX2"
  zone        = "A"
  is_public   = false
  speed       = "1Gbps"
  cloud_settings {
    credentials_uuid = packetfabric_cloud_provider_credential_aws.aws_creds1.id
    aws_region       = "us-west-1"
    mtu              = 1500
    aws_vif_type     = "private" # or transit
    aws_gateways {
      type = "directconnect"
      id   = "760f047b-53ce-4a9d-9ed6-6fac5ca2fa81"
    }
    aws_gateways { #  Private VIF
      type   = "private"
      id     = "vgw-066eb6dcd07dcbb65"
      vpc_id = "vpc-bea401c4"
    }
    # aws_gateways { # Transit VIF
    #   type = "transit"
    #   id   = "tgw-0b7a1390af74b9728"
    #   vpc_id = "vpc-bea401c4"
    #   subnet_ids = [
    #     "subnet-0c222c8047660ca13",
    #     "subnet-03838a8ea2270c40a"
    #   ]
    # }
  }
  labels = ["terraform", "dev"]
}
