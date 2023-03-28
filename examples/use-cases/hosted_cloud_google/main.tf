terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.3.0"
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
  # https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#credentials-1
  credentials = file(var.gcp_credentials_path) # or use environment variable called GOOGLE_CREDENTIALS
  region      = var.gcp_region1
  zone        = var.gcp_zone1
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
# output "google_compute_network" {
#   value = google_compute_network.vpc_1
# }

# From the Google side: Create a Google Cloud Router with ASN 16550.
resource "google_compute_router" "router_1" {
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

# From the Google side: Create a VLAN attachment.
resource "google_compute_interconnect_attachment" "interconnect_1" {
  provider                 = google
  name                     = "${var.resource_name}-${random_pet.name.id}"
  region                   = var.gcp_region1
  description              = "Interconnect to PacketFabric Network"
  type                     = "PARTNER"
  edge_availability_domain = "AVAILABILITY_DOMAIN_1"
  admin_enabled            = true # From the Google side: Accept (automatically) the connection.
  router                   = google_compute_router.router_1.id
}

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
# output "packetfabric_port_1" {
#   value = packetfabric_port.port_1
# }

# From the PacketFabric side: Create a GCP Hosted Connection 
resource "packetfabric_cs_google_hosted_connection" "pf_cs_conn1" {
  provider                    = packetfabric
  description                 = "${var.resource_name}-${random_pet.name.id}-${var.pf_cs_pop1}"
  labels                      = var.pf_labels
  port                        = packetfabric_port.port_1.id
  speed                       = var.pf_cs_speed
  google_pairing_key          = google_compute_interconnect_attachment.interconnect_1.pairing_key
  google_vlan_attachment_name = "${var.resource_name}-${random_pet.name.id}"
  pop                         = var.pf_cs_pop1
  vlan                        = var.pf_cs_vlan1
}
# output "packetfabric_cs_google_hosted_connection" {
#   value = packetfabric_cs_google_hosted_connection.pf_cs_conn1
# }

# data "packetfabric_cs_google_hosted_connection" "current" {
#   provider         = packetfabric
#   cloud_circuit_id = packetfabric_cs_google_hosted_connection.pf_cs_conn1.id
# }

# output "data_packetfabric_cs_google_hosted_connection" {
#   value = data.packetfabric_cs_google_hosted_connection.current
# }

##########################################################################################
#### Here you would need to setup BGP in your Router
##########################################################################################

# Vote for
# https://github.com/hashicorp/terraform-provider-google/issues/11458
# https://github.com/hashicorp/terraform-provider-google/issues/12624

data "google_compute_router" "router_1" {
  provider = google
  name     = "${var.resource_name}-${random_pet.name.id}"
  network  = google_compute_network.vpc_1.id
}
# output "google_compute_router" {
#   value = data.google_compute_router.router_1
# }
