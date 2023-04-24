resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = ["terraform", "dev"]
}

# Example PacketFabric side provisioning only
resource "packetfabric_cloud_router_connection_google" "crc2" {
  provider                    = packetfabric
  description                 = "hello world"
  circuit_id                  = packetfabric_cloud_router.cr1.id
  google_pairing_key          = var.pf_crc_google_pairing_key
  google_vlan_attachment_name = var.pf_crc_google_vlan_attachment_name
  pop                         = "PDX2"
  speed                       = "1Gbps"
  maybe_nat                   = false
  labels                      = ["terraform", "dev"]
}


# Example PacketFabric side + Google side provisioning
resource "packetfabric_cloud_provider_credential_google" "google_creds1" {
  provider               = packetfabric
  description            = "Google Staging Environement"
  google_service_account = var.google_service_account # or use env var GOOGLE_CREDENTIALS
}

resource "packetfabric_cloud_router_connection_google" "crc2" {
  provider    = packetfabric
  description = "hello world"
  circuit_id  = packetfabric_cloud_router.cr1.id
  pop         = "PDX2"
  speed       = "1Gbps"
  maybe_nat   = false
  cloud_settings {
    credentials_uuid                = packetfabric_cloud_provider_credential_google.google_creds1.id
    google_region                   = "us-west1"
    google_vlan_attachment_name     = "my-google-vlan-attachment-primary"
    google_cloud_router_name        = "my-google-cloud-router"
    google_vpc_name                 = "my-google-vpc"
    google_edge_availability_domain = 1 # primary
    bgp_settings {
      md5 = "changeme"
      prefixes {
        prefix = var.pf_crbp_pfx00
        type   = "out" # Allowed Prefixes to Cloud
      }
      prefixes {
        prefix = var.pf_crbp_pfx03
        type   = "in" # Allowed Prefixes from Cloud
      }
    }
  }
  labels = ["terraform", "dev"]
}