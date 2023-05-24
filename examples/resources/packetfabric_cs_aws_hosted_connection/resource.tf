# Example PacketFabric side provisioning only
resource "packetfabric_cs_aws_hosted_connection" "cs_conn1_hosted_aws" {
  provider    = packetfabric
  description = "hello world"
  port        = packetfabric_port.port_1.id
  speed       = "10Gbps"
  pop         = "BOS1"
  vlan        = 102
  zone        = "A"
  labels      = ["terraform", "dev"]
}
resource "aws_dx_connection_confirmation" "confirmation" {
  provider      = aws
  connection_id = packetfabric_cs_aws_hosted_connection.cs_conn1_hosted_aws.cloud_provider_connection_id
}

# Example PacketFabric side + AWS side provisioning
resource "packetfabric_cloud_provider_credential_aws" "aws_creds1" {
  provider       = packetfabric
  description    = "AWS Staging Environement"
  aws_access_key = var.pf_aws_key    # or use env var AWS_ACCESS_KEY_ID
  aws_secret_key = var.pf_aws_secret # or use env var AWS_SECRET_ACCESS_KEY
}

resource "packetfabric_cs_aws_hosted_connection" "cs_conn1_hosted_aws_cloud_side" {
  provider    = packetfabric
  description = "hello world"
  port        = packetfabric_port.port_1.id
  speed       = "10Gbps"
  pop         = "BOS1"
  vlan        = 102
  zone        = "A"
  cloud_settings {
    credentials_uuid = packetfabric_cloud_provider_credential_aws.aws_creds1.id
    aws_region       = "us-east-1"
    mtu              = 1500
    aws_vif_type     = "private" # or transit
    bgp_settings {
      customer_asn   = 64513
      address_family = "ipv4"
    }
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