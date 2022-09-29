terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.3.0"
    }
    google = {
      source  = "hashicorp/google"
      version = "4.38.0"
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

resource "google_compute_firewall" "ssh-rule" {
  provider = google
  name     = "allow-icmp-ssh-http-locust-iperf"
  network  = google_compute_network.vpc_1.name
  allow {
    protocol = "icmp"
  }
  allow {
    protocol = "tcp"
    ports    = ["22", "80", "8089", "5001"]
  }
  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_instance" "vm_1" {
  provider     = google
  name         = "${var.tag_name}-${random_pet.name.id}-vm1"
  machine_type = "e2-micro"
  zone         = var.gcp_zone1
  tags         = ["${var.tag_name}-${random_pet.name.id}"]
  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2204-lts"
    }
  }
  network_interface {
    subnetwork = google_compute_subnetwork.subnet_1.name
    access_config {}
  }
  metadata_startup_script = file("./user-data-ubuntu.sh")
  metadata = {
    sshKeys = "ubuntu:${var.public_key}"
  }
}

data "google_compute_instance" "vm_1" {
  provider = google
  name     = "${var.tag_name}-${random_pet.name.id}-vm1"
  zone     = var.gcp_zone1
  depends_on = [
    google_compute_instance.vm_1
  ]
}

output "google_private_ip_vm_1" {
  description = "Private ip address for VM for Region 1"
  value       = data.google_compute_instance.vm_1.network_interface.0.network_ip
}

output "google_public_ip_vm_1" {
  description = "Public ip address for VM for Region 1 (ssh user: ubuntu)"
  value       = data.google_compute_instance.vm_1.network_interface.0.access_config.0.nat_ip
}

########################################
###### Google Connection
########################################

# From the Google side: Create a Google Cloud Router with ASN 16550.
resource "google_compute_router" "google_router_1" {
  provider = google
  name     = "${var.tag_name}-${random_pet.name.id}"
  network  = google_compute_network.vpc_1.id
  bgp {
    # You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
    asn               = var.gcp_side_asn1
    advertise_mode    = "CUSTOM"
    advertised_groups = ["ALL_SUBNETS"]
  }
}
# data "google_compute_router" "google_router_1" {
#   name    = google_compute_router.google_router_1.name
#   network = google_compute_network.vpc_1.id
# }
# output "google_cloud_router_1" {
#   value = data.google_compute_router.google_router_1
# }

# From the Google side: Create a VLAN attachment.
resource "google_compute_interconnect_attachment" "google_interconnect_1" {
  provider      = google
  name          = "${var.tag_name}-${random_pet.name.id}"
  region        = var.gcp_region1
  description   = "Interconnect to PacketFabric Network"
  type          = "PARTNER"
  admin_enabled = true # From the Google side: Accept (automatically) the connection.
  router        = google_compute_router.google_router_1.id
}
output "google_interconnect_1" {
  value = google_compute_interconnect_attachment.google_interconnect_1
}

# From the PacketFabric side: Create a cloud router
resource "packetfabric_cloud_router" "cr" {
  provider     = packetfabric
  name         = "${var.tag_name}-${random_pet.name.id}"
  account_uuid = var.pf_account_uuid
  asn          = var.pf_cr_asn
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

output "packetfabric_cloud_router" {
  value = packetfabric_cloud_router.cr
}

# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_google_cloud_router_connection" "crc_1" {
  provider                    = packetfabric
  description                 = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  circuit_id                  = packetfabric_cloud_router.cr.id
  account_uuid                = var.pf_account_uuid
  google_pairing_key          = google_compute_interconnect_attachment.google_interconnect_1.pairing_key
  google_vlan_attachment_name = google_compute_interconnect_attachment.google_interconnect_1.name
  pop                         = var.pf_crc_pop1
  speed                       = var.pf_crc_speed
  maybe_nat                   = var.pf_crc_maybe_nat
}

# From both sides: Configure BGP.

# Because the BGP session is created automatically, the only way to get the BGP Addresses it is to use gcloud
# To avoid using this workaround, vote for:
# https://github.com/hashicorp/terraform-provider-google/issues/11458
# https://github.com/hashicorp/terraform-provider-google/issues/12624

# Get the BGP Addresses using glcoud terraform module as a workaround
module "gcloud_bgp_addresses" {
  # https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest
  source                   = "terraform-google-modules/gcloud/google"
  version                  = "~> 2.0"
  service_account_key_file = var.gcp_credentials

  # https://cloud.google.com/sdk/gcloud/reference/compute/routers/update-bgp-peer
  create_cmd_entrypoint = "${path.module}/gcloud_bgp_addresses.sh"
  create_cmd_body       = "${var.gcp_project_id} ${var.gcp_region1} ${google_compute_router.google_router_1.name}"

  module_depends_on = [
    packetfabric_google_cloud_router_connection.crc_1
  ]
}
data "local_file" "cloud_router_ip_address" {
  filename = "${path.module}/cloud_router_ip_address.txt"
  depends_on = [
    module.gcloud_bgp_addresses
  ]
}
data "local_file" "customer_router_ip_address" {
  filename = "${path.module}/customer_router_ip_address.txt"
  depends_on = [
    module.gcloud_bgp_addresses
  ]
}

# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_google_cloud_router_connection.crc_1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.gcp_side_asn1
  orlonger       = var.pf_crbs_orlonger
  # when the google_compute_interconnect_attachment data source will exist, no need to use the gcloud terraform module
  # https://github.com/hashicorp/terraform-provider-google/issues/12624
  # remote_address = data.google_compute_interconnect_attachment.google_interconnect_1.cloud_router_ip_address    # Google side
  # l3_address     = data.google_compute_interconnect_attachment.google_interconnect_1.customer_router_ip_address # PF side
  remote_address = data.local_file.cloud_router_ip_address.content    # Google side
  l3_address     = data.local_file.customer_router_ip_address.content # PF side
}
output "packetfabric_cloud_router_bgp_session_crbs_1" {
  value = packetfabric_cloud_router_bgp_session.crbs_1
}

# Configure BGP Prefix is mandatory to setup the BGP session correctly
resource "packetfabric_cloud_router_bgp_prefixes" "crbp_1" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_1.id
  prefixes {
    prefix = var.subnet_cidr2
    type   = "out" # Allowed Prefixes to Cloud
    order  = 0
  }
  prefixes {
    prefix = var.subnet_cidr1
    type   = "in" # Allowed Prefixes from Cloud
    order  = 0
  }
}
data "packetfabric_cloud_router_bgp_prefixes" "bgp_prefix_crbp_1" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_1.id
}
output "packetfabric_bgp_prefix_crbp_1" {
  value = data.packetfabric_cloud_router_bgp_prefixes.bgp_prefix_crbp_1
}

data "packetfabric_google_cloud_router_connection" "current" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  depends_on = [
    packetfabric_cloud_router_bgp_session.crbs_1
  ]
}
output "packetfabric_google_cloud_router_connection" {
  value = data.packetfabric_google_cloud_router_connection.current
}

# Because the BGP session is created automatically, the only way to update it is to use gcloud
# To avoid using this workaround, vote for:
# https://github.com/hashicorp/terraform-provider-google/issues/12630

# Update BGP Peer in the BGP session's Google Cloud Router
module "gcloud_bgp_peer_update" {
  # https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest
  source                   = "terraform-google-modules/gcloud/google"
  version                  = "~> 2.0"
  service_account_key_file = var.gcp_credentials

  # https://cloud.google.com/sdk/gcloud/reference/compute/routers/update-bgp-peer
  create_cmd_entrypoint = "${path.module}/gcloud_bgp_peer_update.sh"
  create_cmd_body       = "${var.gcp_project_id} ${var.gcp_region1} ${google_compute_router.google_router_1.name} ${var.pf_cr_asn}"

  module_depends_on = [
    packetfabric_google_cloud_router_connection.crc_1
  ]
}

########################################
###### VPN Connection (IPsec)
########################################

resource "packetfabric_ipsec_cloud_router_connection" "crc_2" {
  provider                     = packetfabric
  description                  = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  circuit_id                   = packetfabric_cloud_router.cr.id
  account_uuid                 = var.pf_account_uuid
  pop                          = var.pf_crc_pop2
  speed                        = var.pf_crc_speed
  gateway_address              = var.pf_crc_gateway_address
  ike_version                  = var.pf_crc_ike_version
  phase1_authentication_method = var.pf_crc_phase1_authentication_method
  phase1_group                 = var.pf_crc_phase1_group
  phase1_encryption_algo       = var.pf_crc_phase1_encryption_algo
  phase1_authentication_algo   = var.pf_crc_phase1_authentication_algo
  phase1_lifetime              = var.pf_crc_phase1_lifetime
  phase2_pfs_group             = var.pf_crc_phase2_pfs_group
  phase2_encryption_algo       = var.pf_crc_phase2_encryption_algo
  phase2_authentication_algo   = var.pf_crc_phase2_authentication_algo
  phase2_lifetime              = var.pf_crc_phase2_lifetime
  shared_key                   = var.pf_crc_shared_key
}

resource "packetfabric_cloud_router_bgp_session" "crbs_2" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_ipsec_cloud_router_connection.crc_2.id
  address_family = var.pf_crbs_af
  remote_asn     = var.vpn_side_asn2
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.vpn_remote_address # On-Prem side
  l3_address     = var.vpn_l3_address     # PF side
}
output "packetfabric_cloud_router_bgp_session_crbs_2" {
  value = packetfabric_cloud_router_bgp_session.crbs_2
}

# Configure BGP Prefix is mandatory to setup the BGP session correctly
resource "packetfabric_cloud_router_bgp_prefixes" "crbp_2" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_2.id
  prefixes {
    prefix = var.subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
    order  = 0
  }
  prefixes {
    prefix = var.subnet_cidr2
    type   = "in" # Allowed Prefixes from Cloud
    order  = 0
  }
}

data "packetfabric_cloud_router_bgp_prefixes" "bgp_prefix_crbp_2" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_2.id
}
output "packetfabric_bgp_prefix_crbp_2" {
  value = data.packetfabric_cloud_router_bgp_prefixes.bgp_prefix_crbp_2
}

