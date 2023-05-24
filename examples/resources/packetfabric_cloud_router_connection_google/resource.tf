resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US"]
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

# Google Cloud Router
resource "google_compute_router" "google_router" {
  provider = google
  name     = "myGoogleCloudRouter"
  region   = "us-west1"
  project  = "myGoogleProject"
  network  = "myNetwork"
  bgp {
    asn            = 16550
    advertise_mode = "CUSTOM"
  }
  lifecycle {
    # advertised_ip_ranges managed via BGP prefixes in configured in packetfabric_cloud_router_connection_google
    # asn could be change to a private ASN by PacketFabric in case of multiple google connection in the same cloud router
    ignore_changes = [
      bgp[0].advertised_ip_ranges,
      bgp[0].asn
    ]
  }
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
    google_cloud_router_name        = google_compute_router.google_router.name
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