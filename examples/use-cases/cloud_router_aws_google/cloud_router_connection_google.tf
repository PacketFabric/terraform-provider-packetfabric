resource "packetfabric_cloud_provider_credential_google" "google_creds1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-google"
  # using env var GOOGLE_CREDENTIALS
}

# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_google" "crc_google" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop2
  speed       = var.pf_crc_speed
  # Cloud side provisioning
  cloud_settings {
    credentials_uuid                = packetfabric_cloud_provider_credential_google.google_creds1.id
    google_region                   = var.gcp_region1
    google_vlan_attachment_name     = "${var.resource_name}-${random_pet.name.id}"
    google_cloud_router_name        = google_compute_router.google_router_1.name
    google_vpc_name                 = google_compute_network.vpc_1.name
    google_edge_availability_domain = 1
    bgp_settings {
      remote_asn = var.gcp_side_asn1
      orlonger   = var.pf_crbs_orlonger
      prefixes {
        prefix = var.aws_vpc_cidr1
        type   = "out" # Allowed Prefixes to Cloud
      }
      prefixes {
        prefix = var.gcp_subnet_cidr1
        type   = "in" # Allowed Prefixes from Cloud
      }
    }
  }
}