terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.4.2"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.14.0"
    }
  }
}

provider "packetfabric" {}

provider "azurerm" {
  features {
    resource_group {
      prevent_deletion_if_contains_resources = false
    }
  }
}

# create random name to use to name objects
resource "random_pet" "name" {}

# From the Microsoft side: Create a Microsoft Azure account and set up a virtual network (VNet)
resource "azurerm_resource_group" "resource_group_1" {
  provider = azurerm
  name     = "${var.tag_name}-${random_pet.name.id}"
  location = var.azure_region1
}

resource "azurerm_virtual_network" "virtual_network_1" {
  provider            = azurerm
  name                = "${var.tag_name}-${random_pet.name.id}-vnet1"
  location            = azurerm_resource_group.resource_group_1.location
  resource_group_name = azurerm_resource_group.resource_group_1.name
  address_space       = ["${var.vnet_cidr1}"]
  tags = {
    environment = "${var.tag_name}-${random_pet.name.id}"
  }
}

resource "azurerm_subnet" "subnet_1" {
  provider             = azurerm
  name                 = "${var.tag_name}-${random_pet.name.id}-subnet1"
  address_prefixes     = ["${var.subnet_cidr1}"]
  resource_group_name  = azurerm_resource_group.resource_group_1.name
  virtual_network_name = azurerm_virtual_network.virtual_network_1.name
}

# Subnet used for the azurerm_virtual_network_gateway only
resource "azurerm_subnet" "subnet_gw" {
  provider             = azurerm
  name                 = "GatewaySubnet"
  address_prefixes     = ["${var.subnet_cidr1gw}"]
  resource_group_name  = azurerm_resource_group.resource_group_1.name
  virtual_network_name = azurerm_virtual_network.virtual_network_1.name
}

# From the Microsoft side: Create an ExpressRoute circuit from the Azure portal.
resource "azurerm_express_route_circuit" "azure_express_route_1" {
  provider              = azurerm
  name                  = "${var.tag_name}-${random_pet.name.id}"
  resource_group_name   = azurerm_resource_group.resource_group_1.name
  location              = azurerm_resource_group.resource_group_1.location
  peering_location      = var.peering_location_1
  service_provider_name = var.service_provider_name
  bandwidth_in_mbps     = var.bandwidth_in_mbps
  sku {
    tier   = var.sku_tier
    family = var.sku_family
  }
  tags = {
    environment = "${var.tag_name}-${random_pet.name.id}"
  }
}

# From the PacketFabric side: Create a PacketFabric Hosted Cloud Connection.
resource "packetfabric_cs_azure_hosted_connection" "pf_cs_conn1" {
  provider          = packetfabric
  description       = "${var.tag_name}-${random_pet.name.id}"
  azure_service_key = azurerm_express_route_circuit.azure_express_route_1.service_key
  port              = var.pf_port_circuit_id
  speed             = var.pf_cs_speed # will be deprecated
  vlan_private      = var.pf_cs_vlan_private
  #vlan_microsoft = var.pf_cs_vlan_microsoft
}

output "packetfabric_cs_azure_hosted_connection" {
  value     = packetfabric_cs_azure_hosted_connection.pf_cs_conn1
  sensitive = true
}

# From the Microsoft side: Configure peering in the Azure portal.
resource "azurerm_express_route_circuit_peering" "private_circuit_1" {
  provider                      = azurerm
  peering_type                  = "AzurePrivatePeering"
  express_route_circuit_name    = azurerm_express_route_circuit.azure_express_route_1.name
  resource_group_name           = azurerm_resource_group.resource_group_1.name
  peer_asn                      = var.peer_asn
  primary_peer_address_prefix   = var.primary_peer_address_prefix
  secondary_peer_address_prefix = var.secondary_peer_address_prefix
  vlan_id                       = var.pf_cs_vlan_private
  shared_key                    = var.shared_key
  depends_on = [
    packetfabric_cs_azure_hosted_connection.pf_cs_conn1
  ]
}

data "azurerm_express_route_circuit" "azure_express_route_1" {
  resource_group_name = azurerm_resource_group.resource_group_1.name
  name                = azurerm_express_route_circuit.azure_express_route_1.name
  depends_on = [
    azurerm_express_route_circuit_peering.private_circuit_1
  ]
}

output "express_route_circuit" {
  value = data.azurerm_express_route_circuit.azure_express_route_1
}

##########################################################################################
#### Here you would need to setup BGP in your Router
##########################################################################################

### Below resources can take a while to create/delete, manually uncomment below code

# # From the Microsoft side: Create a virtual network gateway for ExpressRoute.
# resource "azurerm_public_ip" "public_ip_vng_1" {
#   provider            = azurerm
#   name                = "${var.tag_name}-${random_pet.name.id}-public-ip-vng1"
#   location            = azurerm_resource_group.resource_group_1.location
#   resource_group_name = azurerm_resource_group.resource_group_1.name
#   allocation_method   = "Dynamic"
#   tags = {
#     environment = "${var.tag_name}-${random_pet.name.id}"
#   }
# }
# # This resource creation can take up to 50min - deletion up to 12min
# resource "azurerm_virtual_network_gateway" "vng_1" {
#   provider            = azurerm
#   name                = "${var.tag_name}-${random_pet.name.id}-vng1"
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
#     environment = "${var.tag_name}-${random_pet.name.id}"
#   }
# }

# # From the Microsoft side: Link a virtual network gateway to the ExpressRoute circuit.
# resource "azurerm_virtual_network_gateway_connection" "vng_connection_1" {
#   provider                   = azurerm
#   name                       = "${var.tag_name}-${random_pet.name.id}-vng_connection_1"
#   location                   = azurerm_resource_group.resource_group_1.location
#   resource_group_name        = azurerm_resource_group.resource_group_1.name
#   type                       = "ExpressRoute"
#   express_route_circuit_id   = azurerm_express_route_circuit.azure_express_route_1.id
#   virtual_network_gateway_id = azurerm_virtual_network_gateway.vng_1.id
#   routing_weight             = 0
#   tags = {
#     environment = "${var.tag_name}-${random_pet.name.id}"
#   }
#   depends_on = [
#     packetfabric_cs_azure_hosted_connection.pf_cs_conn1
#   ]
# }
