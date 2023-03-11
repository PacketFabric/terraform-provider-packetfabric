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

# From the Google side: Create a VLAN attachment.
resource "google_compute_interconnect_attachment" "google_interconnect_1" {
  provider                 = google
  name                     = "${var.tag_name}-${random_pet.name.id}"
  region                   = var.gcp_region1
  description              = "Interconnect to PacketFabric Network"
  type                     = "PARTNER"
  edge_availability_domain = "AVAILABILITY_DOMAIN_1"
  admin_enabled            = true # From the Google side: Accept (automatically) the connection.
  router                   = google_compute_router.google_router_1.id
}
# output "google_interconnect_1" {
#   value = google_compute_interconnect_attachment.google_interconnect_1
# }

# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_google" "crc_2" {
  provider                    = packetfabric
  description                 = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  circuit_id                  = packetfabric_cloud_router.cr.id
  google_pairing_key          = google_compute_interconnect_attachment.google_interconnect_1.pairing_key
  google_vlan_attachment_name = google_compute_interconnect_attachment.google_interconnect_1.name
  pop                         = var.pf_crc_pop2
  speed                       = var.pf_crc_speed
  maybe_nat                   = var.pf_crc_maybe_nat
}

# # Verify Terraform gcloud module works in your environment
# module "gcloud_version" {
#   # https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest
#   source                            = "terraform-google-modules/gcloud/google"
#   version                           = "~> 2.0"

#   use_tf_google_credentials_env_var = true
#   # when running locally with gcloud already installed
#   skip_download = true
#   # when running in a CI/CD pipeline without glcoud installed
#   # skip_download                     = false

#   # https://cloud.google.com/sdk/gcloud/reference/compute/routers/update-bgp-peer
#   create_cmd_entrypoint = "gcloud"
#   create_cmd_body       = "version"

#   # no destroy needed
#   destroy_cmd_entrypoint = "echo"
#   destroy_cmd_body       = "skip"
# }

# From both sides: Configure BGP.

# Because the BGP session is created automatically, the only way to get the BGP Addresses it is to use gcloud
# To avoid using this workaround, vote for:
# https://github.com/hashicorp/terraform-provider-google/issues/11458
# https://github.com/hashicorp/terraform-provider-google/issues/12624

# Get the BGP Addresses using glcoud terraform module as a workaround
module "gcloud_bgp_addresses" {
  # https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest
  source  = "terraform-google-modules/gcloud/google"
  version = "~> 2.0"

  use_tf_google_credentials_env_var = true
  # when running locally with gcloud already installed
  skip_download = true
  # when running in a CI/CD pipeline without glcoud installed
  # skip_download                     = false

  # https://cloud.google.com/sdk/gcloud/reference/compute/routers/update-bgp-peer
  create_cmd_entrypoint = "${path.module}/gcloud_bgp_addresses.sh"
  create_cmd_body       = "${var.gcp_project_id} ${var.gcp_region1} ${google_compute_router.google_router_1.name}"

  # no destroy needed
  destroy_cmd_entrypoint = "echo"
  destroy_cmd_body       = "skip"

  module_depends_on = [
    packetfabric_cloud_router_connection_google.crc_2
  ]

  # When "gcloud_bin_abs_path" changes, it should not trigger a replacement
  # https://github.com/hashicorp/terraform/issues/27360
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
resource "packetfabric_cloud_router_bgp_session" "crbs_2" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_google.crc_2.id
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
  prefixes {
    prefix = var.aws_vpc_cidr1
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.gcp_subnet_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }

  # # workaround until we can use lifecycle into Terraform gcloud Module
  # # https://github.com/hashicorp/terraform/issues/27360
  # lifecycle {
  #   ignore_changes = [
  #     remote_address,
  #     l3_address
  #   ]
  # }
}
# output "packetfabric_cloud_router_bgp_session_crbs_2" {
#   value = packetfabric_cloud_router_bgp_session.crbs_2
# }

# Because the BGP session is created automatically, the only way to update it is to use gcloud
# To avoid using this workaround, vote for:
# https://github.com/hashicorp/terraform-provider-google/issues/12630

# Update BGP Peer in the BGP session's Google Cloud Router
module "gcloud_bgp_peer_update" {
  # https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest
  source  = "terraform-google-modules/gcloud/google"
  version = "~> 2.0"

  use_tf_google_credentials_env_var = true
  # when running locally with gcloud already installed
  skip_download = true
  # when running in a CI/CD pipeline without glcoud installed
  # skip_download                     = false

  # https://cloud.google.com/sdk/gcloud/reference/compute/routers/update-bgp-peer
  create_cmd_entrypoint = "${path.module}/gcloud_bgp_peer_update.sh"
  create_cmd_body       = "${var.gcp_project_id} ${var.gcp_region1} ${google_compute_router.google_router_1.name} ${var.pf_cr_asn}"

  # no destroy needed
  destroy_cmd_entrypoint = "echo"
  destroy_cmd_body       = "skip"

  module_depends_on = [
    packetfabric_cloud_router_connection_google.crc_2
  ]

  # When "gcloud_bin_abs_path" changes, it should not trigger a replacement
  # https://github.com/hashicorp/terraform/issues/27360
}