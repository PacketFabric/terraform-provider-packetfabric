# Pre-req to enable AzureExpressRoute in the Azure Subscription
# az feature register --namespace Microsoft.Network --name AllowExpressRoutePorts
# az provider register -n Microsoft.Network

# From the Microsoft side: Create an ExpressRoute circuit in the Azure Console.
resource "azurerm_express_route_circuit" "azure_express_route_1" {
  provider              = azurerm
  name                  = "${var.resource_name}-${random_pet.name.id}"
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
    environment = "${var.resource_name}-${random_pet.name.id}"
  }
}

# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_azure" "crc_azure" {
  provider = packetfabric
  # using Azure Peering location for the name but removing the spaces with a regex
  description       = "${var.resource_name}-${random_pet.name.id}-${replace(var.azure_peering_location_1, "/\\s+/", "")}-primary"
  labels            = var.pf_labels
  circuit_id        = packetfabric_cloud_router.cr.id
  azure_service_key = azurerm_express_route_circuit.azure_express_route_1.service_key
  speed             = var.pf_crc_speed
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
  vlan_id                       = packetfabric_cloud_router_connection_azure.crc_azure.vlan_id_private
  shared_key                    = var.azure_bgp_shared_key
}

resource "packetfabric_cloud_router_bgp_session" "crbs_azure" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_azure.crc_azure.id
  address_family = var.pf_crbs_af
  remote_asn     = var.azure_side_asn1
  orlonger       = var.pf_crbs_orlonger
  # Only specify either the primary_subnet OR the secondary_subnet
  primary_subnet = var.azure_primary_peer_address_prefix
  prefixes {
    prefix = var.gcp_subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.azure_vnet_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
}
# output "packetfabric_cloud_router_bgp_session_crbs_azure" {
#   value = packetfabric_cloud_router_bgp_session.crbs_azure
# }

# # From the Microsoft side: Create a virtual network gateway for ExpressRoute.
# resource "azurerm_public_ip" "public_ip_vng_1" {
#   provider            = azurerm
#   name                = "${var.resource_name}-${random_pet.name.id}-public-ip-vng1"
#   location            = azurerm_resource_group.resource_group_1.location
#   resource_group_name = azurerm_resource_group.resource_group_1.name
#   allocation_method   = "Dynamic"
#   sku                 = "Standard"
#   tags = {
#     environment = "${var.resource_name}-${random_pet.name.id}"
#   }
# }

# # Please be aware that provisioning a Virtual Network Gateway takes a long time (between 30 minutes and 1 hour)
# # Deletion can take up to 15 minutes
# resource "azurerm_virtual_network_gateway" "vng_1" {
#   provider            = azurerm
#   name                = "${var.resource_name}-${random_pet.name.id}-vng1"
#   location            = azurerm_resource_group.resource_group_1.location
#   resource_group_name = azurerm_resource_group.resource_group_1.name
#   type                = "ExpressRoute"
#   sku                 = "Standard"
#   ip_configuration {
#     name                          = "vnetGatewayConfig"
#     public_ip_address_id          = azurerm_public_ip.public_ip_vng_1.id
#     private_ip_address_allocation = "Dynamic"
#     subnet_id                     = azurerm_subnet.subnet_gw.id
#   }
#   tags = {
#     environment = "${var.resource_name}-${random_pet.name.id}"
#   }
#   depends_on = [
#     azurerm_public_ip.public_ip_vng_1
#   ]
# }

# # From the Microsoft side: Link a virtual network gateway to the ExpressRoute circuit.
# resource "azurerm_virtual_network_gateway_connection" "vng_connection_1" {
#   provider                   = azurerm
#   name                       = "${var.resource_name}-${random_pet.name.id}-vng_connection_1"
#   location                   = azurerm_resource_group.resource_group_1.location
#   resource_group_name        = azurerm_resource_group.resource_group_1.name
#   type                       = "ExpressRoute"
#   express_route_circuit_id   = azurerm_express_route_circuit.azure_express_route_1.id
#   virtual_network_gateway_id = azurerm_virtual_network_gateway.vng_1.id
#   routing_weight             = 0
#   tags = {
#     environment = "${var.resource_name}-${random_pet.name.id}"
#   }
# }
