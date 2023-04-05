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

# Example PacketFabric side + AWS side provisioning
resource "packetfabric_cloud_provider_credential_aws" "aws_creds1" {
  provider       = packetfabric
  description    = "AWS Staging Environement"
  aws_access_key = var.pf_aws_key    # or use env var PF_AWS_ACCESS_KEY_ID
  aws_secret_key = var.pf_aws_secret # or use env var PF_AWS_SECRET_ACCESS_KEY
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
    aws_vif_type     = "private"
    bgp_settings {
      customer_asn   = 64513
      address_family = "ipv4"
    }
    aws_gateways {
      type = "directconnect"
      name = "${var.tag_name}-${random_pet.name.id}"
      asn  = 64513
    }
    aws_gateways {
      type   = "private"
      name   = "${var.tag_name}-${random_pet.name.id}"
      vpc_id = "vpc-bea401c4"
    }
  }
  labels = ["terraform", "dev"]
}