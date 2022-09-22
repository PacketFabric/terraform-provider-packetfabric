terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.2.1"
    }
    google = {
      source  = "hashicorp/google"
      version = "4.30.0"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

# Make sure you enabled Compute Engine API
provider "google" {
  project     = var.gcp_project_id
  credentials = file(var.gcp_credentials)
  region      = var.gcp_region1
  zone        = var.gcp_zone1
}

# create random name to use to name objects
resource "random_pet" "name" {}

resource "google_compute_network" "vpc_1" {
  provider                = google
  name                    = "${var.tag_name}-${random_pet.name.id}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet_1" {
  provider      = google
  name          = "${var.tag_name}-${random_pet.name.id}"
  ip_cidr_range = var.subnet_cidr1
  region        = var.gcp_region1
  network       = google_compute_network.vpc_1.id
}

output "google_compute_network" {
  value = google_compute_network.vpc_1
}

# From the Google side: Create a Google Cloud Router with ASN 16550.
resource "google_compute_router" "router_1" {
  provider = google
  name     = "${var.tag_name}-${random_pet.name.id}"
  network  = google_compute_network.vpc_1.id
  bgp {
    # You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
    asn               = var.peer_asn
    advertise_mode    = "CUSTOM"
    advertised_groups = ["ALL_SUBNETS"]
  }
}

# From the Google side: Create a VLAN attachment.
resource "google_compute_interconnect_attachment" "interconnect_1" {
  provider      = google
  name          = "${var.tag_name}-${random_pet.name.id}"
  region        = var.gcp_region1
  description   = "Interconnect to PacketFabric Network"
  type          = "PARTNER"
  admin_enabled = true # From the Google side: Accept (automatically) the connection.
  router        = google_compute_router.router_1.id
}

# From the PacketFabric side: Create a GCP Hosted Connection 
resource "packetfabric_cs_google_hosted_connection" "pf_cs_conn1" {
  provider                    = packetfabric
  description                 = "${var.tag_name}-${random_pet.name.id}-${var.pf_cs_pop1}"
  account_uuid                = var.pf_account_uuid
  port                        = var.pf_port_circuit_id
  speed                       = var.pf_cs_speed
  google_pairing_key          = google_compute_interconnect_attachment.interconnect_1.pairing_key
  google_vlan_attachment_name = "${var.tag_name}-${random_pet.name.id}"
  pop                         = var.pf_cs_pop1
  vlan                        = var.pf_cs_vlan1
}

output "packetfabric_cs_google_hosted_connection" {
  value = packetfabric_cs_google_hosted_connection.pf_cs_conn1
}

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
  name     = "${var.tag_name}-${random_pet.name.id}"
  network  = google_compute_network.vpc_1.id
}

output "google_compute_router" {
  value = data.google_compute_router.router_1
}
