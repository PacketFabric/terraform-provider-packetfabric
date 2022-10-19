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
