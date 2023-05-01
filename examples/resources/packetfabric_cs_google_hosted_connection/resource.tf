# Example PacketFabric side provisioning only
resource "packetfabric_cs_google_hosted_connection" "cs_conn1_hosted_google" {
  provider                    = packetfabric
  description                 = "hello world"
  port                        = packetfabric_port.port_1.id
  speed                       = "10Gbps"
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = var.google_vlan_attachment_name
  pop                         = "BOS1"
  vlan                        = 102
  labels                      = ["terraform", "dev"]
}

# Example PacketFabric side + Google side provisioning
resource "packetfabric_cloud_provider_credential_google" "google_creds1" {
  provider               = packetfabric
  description            = "Google Staging Environement"
  google_service_account = var.google_service_account # or use env var GOOGLE_CREDENTIALS
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
    credentials_uuid                = packetfabric_cloud_provider_credential_google.google_creds1.id
    google_region                   = "us-west1"
    google_vlan_attachment_name     = "my-google-vlan-attachment-primary"
    google_cloud_router_name        = "my-google-cloud-router"
    google_vpc_name                 = "my-google-vpc"
    google_edge_availability_domain = 1 # primary
    bgp_settings {
      customer_asn = 65537
      md5          = "changeme"
    }
  }
  labels = ["terraform", "dev"]
}