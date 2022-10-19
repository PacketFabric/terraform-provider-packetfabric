# create random name to use to name objects
resource "random_pet" "name" {}

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

###################################################
###### PacketFabric Google Cloud Router Connection
###################################################

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

# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_google" "crc_1" {
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
  source  = "terraform-google-modules/gcloud/google"
  version = "~> 2.0"
  # when running locally with gcloud already installed
  service_account_key_file = var.GOOGLE_CREDENTIALS
  skip_download            = true
  # when running in a CI/CD pipeline without glcoud installed
  # use_tf_google_credentials_env_var = true
  # skip_download                     = false

  # https://cloud.google.com/sdk/gcloud/reference/compute/routers/update-bgp-peer
  create_cmd_entrypoint = "${path.module}/gcloud_bgp_addresses.sh"
  create_cmd_body       = "${var.gcp_project_id} ${var.gcp_region1} ${google_compute_router.google_router_1.name}"

  # no destroy needed
  destroy_cmd_entrypoint = "echo"
  destroy_cmd_body       = "skip"

  module_depends_on = [
    packetfabric_cloud_router_connection_google.crc_1
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
resource "packetfabric_cloud_router_bgp_session" "crbs_1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_google.crc_1.id
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

  # # workaround until we can use lifecycle into Terraform gcloud Module
  # # https://github.com/hashicorp/terraform/issues/27360
  # lifecycle {
  #   ignore_changes = [
  #     remote_address,
  #     l3_address
  #   ]
  # }
}
output "packetfabric_cloud_router_bgp_session_crbs_1" {
  value = packetfabric_cloud_router_bgp_session.crbs_1
}

# Configure BGP Prefix is mandatory to setup the BGP session correctly
resource "packetfabric_cloud_router_bgp_prefixes" "crbp_1" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_1.id
  prefixes {
    prefix = var.azure_vnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
    order  = 0
  }
  prefixes {
    prefix = var.gcp_subnet_cidr1
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

# Because the BGP session is created automatically, the only way to update it is to use gcloud
# To avoid using this workaround, vote for:
# https://github.com/hashicorp/terraform-provider-google/issues/12630

# Update BGP Peer in the BGP session's Google Cloud Router
module "gcloud_bgp_peer_update" {
  # https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest
  source  = "terraform-google-modules/gcloud/google"
  version = "~> 2.0"
  # when running locally with gcloud already installed
  service_account_key_file = var.GOOGLE_CREDENTIALS
  skip_download            = true
  # when running in a CI/CD pipeline without glcoud installed
  # use_tf_google_credentials_env_var = true
  # skip_download                     = false

  # https://cloud.google.com/sdk/gcloud/reference/compute/routers/update-bgp-peer
  create_cmd_entrypoint = "${path.module}/gcloud_bgp_peer_update.sh"
  create_cmd_body       = "${var.gcp_project_id} ${var.gcp_region1} ${google_compute_router.google_router_1.name} ${var.pf_cr_asn}"

  # no destroy needed
  destroy_cmd_entrypoint = "echo"
  destroy_cmd_body       = "skip"

  module_depends_on = [
    packetfabric_cloud_router_connection_google.crc_1
  ]

  # When "gcloud_bin_abs_path" changes, it should not trigger a replacement
  # https://github.com/hashicorp/terraform/issues/27360
}

###################################################
###### PacketFabric Azure Cloud Router Connection
###################################################

# Pre-req to enable AzureExpressRoute in the Azure Subscription
# az feature register --namespace Microsoft.Network --name AllowExpressRoutePorts
# az provider register -n Microsoft.Network

# From the Microsoft side: Create an ExpressRoute circuit in the Azure Console.
resource "azurerm_express_route_circuit" "azure_express_route_1" {
  provider              = azurerm
  name                  = "${var.tag_name}-${random_pet.name.id}"
  resource_group_name   = azurerm_resource_group.resource_group_1.name
  location              = azurerm_resource_group.resource_group_1.location
  peering_location      = var.azure_peering_location_1
  service_provider_name = var.azure_service_provider_name
  bandwidth_in_mbps     = var.azure_bandwidth_in_mbps
  sku {
    tier   = var.azure_sku_tier
    family = var.azure_sku_family
  }
  tags = {
    environment = "${var.tag_name}-${random_pet.name.id}"
  }
}

# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_azure" "crc_2" {
  provider = packetfabric
  # using Azure Peering location for the name but removing the spaces with a regex
  description       = "${var.tag_name}-${random_pet.name.id}-${replace(var.azure_peering_location_1, "/\\s+/", "")}-primary"
  circuit_id        = packetfabric_cloud_router.cr.id
  account_uuid      = var.pf_account_uuid
  azure_service_key = azurerm_express_route_circuit.azure_express_route_1.service_key
  speed             = var.pf_crc_speed
  maybe_nat         = var.pf_crc_maybe_nat
  is_public         = var.pf_crc_is_public
}

# Get the VLAN ID from PacketFabric
data "packetfabric_cloud_router_connections" "current" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  depends_on = [
    packetfabric_cloud_router_connection_azure.crc_2
  ]
}
locals {
  # below may need to be updated
  # check https://github.com/PacketFabric/terraform-provider-packetfabric/issues/23
  cloud_connections = data.packetfabric_cloud_router_connections.current.cloud_connections[*]
  helper_map = { for val in local.cloud_connections :
  val["description"] => val }
  cc1 = local.helper_map["${var.tag_name}-${random_pet.name.id}-${replace(var.azure_peering_location_1, "/\\s+/", "")}-primary"]
}
output "cc1_vlan_private" {
  value = one(local.cc1.cloud_settings[*].vlan_id_private)
}
output "packetfabric_cloud_router_connection_azure" {
  value = data.packetfabric_cloud_router_connections.current.cloud_connections[*]
}

# From both sides: Configure BGP.
resource "azurerm_express_route_circuit_peering" "private_circuit_1" {
  provider                      = azurerm
  peering_type                  = var.azure_peering_type
  express_route_circuit_name    = azurerm_express_route_circuit.azure_express_route_1.name
  resource_group_name           = azurerm_resource_group.resource_group_1.name
  peer_asn                      = var.pf_cr_asn
  primary_peer_address_prefix   = var.azure_primary_peer_address_prefix
  secondary_peer_address_prefix = var.azure_secondary_peer_address_prefix
  # The VLAN is automatically assigned by PacketFabric and available in the packetfabric_cloud_router_connection_aws data source. 
  # We use local in order to parse the data source output and get the VLAN ID assigned by PacketFabric so we can use it to setup the ExpressRoute Peering
  vlan_id    = one(local.cc1.cloud_settings[*].vlan_id_private)
  shared_key = var.azure_bgp_shared_key
}

resource "packetfabric_cloud_router_bgp_session" "crbs_2" {
  provider         = packetfabric
  circuit_id       = packetfabric_cloud_router.cr.id
  connection_id    = packetfabric_cloud_router_connection_azure.crc_2.id
  address_family   = var.pf_crbs_af
  multihop_ttl     = var.pf_crbs_mhttl
  remote_asn       = var.azure_side_asn1
  orlonger         = var.pf_crbs_orlonger
  primary_subnet   = var.azure_primary_peer_address_prefix
  secondary_subnet = var.azure_secondary_peer_address_prefix
}
output "packetfabric_cloud_router_bgp_session_crbs_2" {
  value = packetfabric_cloud_router_bgp_session.crbs_2
}

# Configure BGP Prefix is mandatory to setup the BGP session correctly
resource "packetfabric_cloud_router_bgp_prefixes" "crbp_2" {
  provider          = packetfabric
  bgp_settings_uuid = packetfabric_cloud_router_bgp_session.crbs_2.id
  prefixes {
    prefix = var.gcp_subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
    order  = 0
  }
  prefixes {
    prefix = var.azure_vnet_cidr1
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

# From the Microsoft side: Create a virtual network gateway for ExpressRoute.
resource "azurerm_public_ip" "public_ip_vng_1" {
  provider            = azurerm
  name                = "${var.tag_name}-${random_pet.name.id}-public-ip-vng1"
  location            = azurerm_resource_group.resource_group_1.location
  resource_group_name = azurerm_resource_group.resource_group_1.name
  allocation_method   = "Dynamic"
  tags = {
    environment = "${var.tag_name}-${random_pet.name.id}"
  }
}
# Please be aware that provisioning a Virtual Network Gateway takes a long time (between 30 minutes and 1 hour)
# Deletion can take up to 15 minutes
resource "azurerm_virtual_network_gateway" "vng_1" {
  provider            = azurerm
  name                = "${var.tag_name}-${random_pet.name.id}-vng1"
  location            = azurerm_resource_group.resource_group_1.location
  resource_group_name = azurerm_resource_group.resource_group_1.name
  type                = "ExpressRoute"
  sku                 = "Standard"
  ip_configuration {
    name                          = "vnetGatewayConfig"
    public_ip_address_id          = azurerm_public_ip.public_ip_vng_1.id
    private_ip_address_allocation = "Dynamic"
    subnet_id                     = azurerm_subnet.subnet_gw.id
  }
  tags = {
    environment = "${var.tag_name}-${random_pet.name.id}"
  }
}

# From the Microsoft side: Link a virtual network gateway to the ExpressRoute circuit.
resource "azurerm_virtual_network_gateway_connection" "vng_connection_1" {
  provider                   = azurerm
  name                       = "${var.tag_name}-${random_pet.name.id}-vng_connection_1"
  location                   = azurerm_resource_group.resource_group_1.location
  resource_group_name        = azurerm_resource_group.resource_group_1.name
  type                       = "ExpressRoute"
  express_route_circuit_id   = azurerm_express_route_circuit.azure_express_route_1.id
  virtual_network_gateway_id = azurerm_virtual_network_gateway.vng_1.id
  routing_weight             = 0
  tags = {
    environment = "${var.tag_name}-${random_pet.name.id}"
  }
}

data "packetfabric_cloud_router_connections" "all_crc" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  depends_on = [
    packetfabric_cloud_router_bgp_session.crbs_1,
    packetfabric_cloud_router_bgp_session.crbs_2
  ]
}
output "packetfabric_cloud_router_connections" {
  value = data.packetfabric_cloud_router_connections.all_crc
}