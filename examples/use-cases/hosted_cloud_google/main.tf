terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.4.0"
    }
    google = {
      source  = "hashicorp/google"
      version = ">= 4.56.0"
    }
  }
}

provider "packetfabric" {}

# Make sure you enabled Compute Engine API
provider "google" {
  project = var.gcp_project_id
  # use GOOGLE_CREDENTIALS environment variable
  region = var.gcp_region1
  zone   = var.gcp_zone1
}

# create random name to use to name objects
resource "random_pet" "name" {}

resource "google_compute_network" "vpc_1" {
  provider                = google
  name                    = "${var.resource_name}-${random_pet.name.id}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet_1" {
  provider      = google
  name          = "${var.resource_name}-${random_pet.name.id}"
  ip_cidr_range = var.subnet_cidr1
  region        = var.gcp_region1
  network       = google_compute_network.vpc_1.id
}

# From the Google side: Create a Google Cloud Router with ASN 16550.
resource "google_compute_router" "router_1" {
  provider = google
  name     = "${var.resource_name}-${random_pet.name.id}"
  network  = google_compute_network.vpc_1.id
  bgp {
    # You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
    asn            = 16550
    advertise_mode = "DEFAULT"
  }
}

# # From the Google side: Create a VLAN attachment.
# resource "google_compute_interconnect_attachment" "interconnect_1" {
#   provider                 = google
#   name                     = "${var.resource_name}-${random_pet.name.id}-${lower(var.pf_cs_pop1)}-primary"
#   region                   = var.gcp_region1
#   description              = "Interconnect to PacketFabric Network"
#   type                     = "PARTNER"
#   edge_availability_domain = "AVAILABILITY_DOMAIN_1"
#   admin_enabled            = true # From the Google side: Accept (automatically) the connection.
#   router                   = google_compute_router.router_1.id
# }

# Create a PacketFabric port
resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  autoneg           = var.pf_port_autoneg
  description       = "${var.resource_name}-${random_pet.name.id}"
  labels            = var.pf_labels
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
}

resource "packetfabric_cloud_provider_credential_google" "google_creds1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-google"
  # using env var GOOGLE_CREDENTIALS
}

# From the PacketFabric side: Create a GCP Hosted Connection 
resource "packetfabric_cs_google_hosted_connection" "pf_cs_conn1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_cs_pop1}"
  labels      = var.pf_labels
  port        = packetfabric_port.port_1.id
  speed       = var.pf_cs_speed
  # set if cloud_settings not used
  # google_pairing_key          = google_compute_interconnect_attachment.interconnect_1.pairing_key
  google_vlan_attachment_name = "${var.resource_name}-${random_pet.name.id}-${lower(var.pf_cs_pop1)}-primary"
  pop                         = var.pf_cs_pop1
  vlan                        = var.pf_cs_vlan1
  # for cloud side provisioning - optional
  cloud_settings {
    credentials_uuid                = packetfabric_cloud_provider_credential_google.google_creds1.id
    google_region                   = var.gcp_region1
    google_project_id               = var.gcp_project_id
    google_vlan_attachment_name     = "${var.resource_name}-${random_pet.name.id}-${lower(var.pf_cs_pop1)}-primary"
    google_cloud_router_name        = "${var.resource_name}-${random_pet.name.id}"
    google_vpc_name                 = "${var.resource_name}-${random_pet.name.id}"
    google_edge_availability_domain = 1 # primary
    bgp_settings {
      customer_asn = var.pf_cs_google_customer_asn
      md5          = var.pf_cs_google_bgp_md5
    }
  }
  depends_on = [
    google_compute_router.router_1
  ]
}
# output "packetfabric_cs_google_hosted_connection" {
#   value = packetfabric_cs_google_hosted_connection.pf_cs_conn1
# }

##########################################################################################
#### Here you would need to setup BGP in your Router
##########################################################################################

# # Vote for
# # https://github.com/hashicorp/terraform-provider-google/issues/11458
# # https://github.com/hashicorp/terraform-provider-google/issues/12624

# data "google_compute_router" "router_1" {
#   provider = google
#   name     = "${var.resource_name}-${random_pet.name.id}"
#   network  = google_compute_network.vpc_1.id
# }
# # output "google_compute_router" {
# #   value = data.google_compute_router.router_1
# # }

# data "packetfabric_cs_hosted_connection_router_config" "router_google_cisco" {
#   cloud_circuit_id = packetfabric_cs_google_hosted_connection.pf_cs_conn1.id
#   router_type      = "CiscoSystemsInc-Routers-Generic"
# }
# resource "local_file" "router_google_cisco_file" {
#   filename = "router_config_google_cisco.txt"
#   content  = data.packetfabric_cs_hosted_connection_router_config.router_google_cisco.router_config
# }
