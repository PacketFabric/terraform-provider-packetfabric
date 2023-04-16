# From the Google side: Create a Google Cloud Router with ASN 16550.
resource "google_compute_router" "google_router_1" {
  provider = google
  name     = "${var.resource_name}-${random_pet.name.id}"
  network  = google_compute_network.vpc_1.id
  bgp {
    # You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
    asn               = var.gcp_side_asn1
    advertise_mode    = "CUSTOM"
    advertised_groups = ["ALL_SUBNETS"]
  }
}

resource "packetfabric_cloud_provider_credential_google" "google_creds1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-google"
  # using env var GOOGLE_CREDENTIALS
}

# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_google" "crc_1" {
  provider                    = packetfabric
  description                 = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  labels                      = var.pf_labels
  circuit_id                  = packetfabric_cloud_router.cr.id
  pop                         = var.pf_crc_pop1
  speed                       = var.pf_crc_speed
  maybe_nat                   = var.pf_crc_maybe_nat
  # for cloud side provisioning - optional
  cloud_settings {
    credentials_uuid                = packetfabric_cloud_provider_credential_google.google_creds1.id
    google_region                   = var.pf_cs_google_region
    google_vlan_attachment_name     = "${var.resource_name}-${random_pet.name.id}"
    google_cloud_router_name        = "${var.resource_name}-${random_pet.name.id}"
    google_vpc_name                 = var.pf_cs_google_vpc_name
    google_edge_availability_domain = 1
    bgp_settings {
      multihop_ttl   = var.pf_crbs_mhttl
      remote_asn     = var.gcp_side_asn1
      orlonger       = var.pf_crbs_orlonger
      prefixes {
        prefix = var.azure_vnet_cidr1
        type   = "out" # Allowed Prefixes to Cloud
      }
      prefixes {
        prefix = var.gcp_subnet_cidr1
        type   = "in" # Allowed Prefixes from Cloud
      }
    }
  }
}
