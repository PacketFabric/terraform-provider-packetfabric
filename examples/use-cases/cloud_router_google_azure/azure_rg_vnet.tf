resource "azurerm_resource_group" "resource_group_1" {
  provider = azurerm
  name     = "${var.resource_name}-${random_pet.name.id}"
  location = var.azure_region1
}

resource "azurerm_virtual_network" "virtual_network_1" {
  provider            = azurerm
  name                = "${var.resource_name}-${random_pet.name.id}-vnet1"
  location            = azurerm_resource_group.resource_group_1.location
  resource_group_name = azurerm_resource_group.resource_group_1.name
  address_space       = ["${var.azure_vnet_cidr1}"]
  tags = {
    environment = "${var.resource_name}-${random_pet.name.id}"
  }
}

resource "azurerm_subnet" "subnet_1" {
  provider             = azurerm
  name                 = "${var.resource_name}-${random_pet.name.id}-subnet1"
  address_prefixes     = ["${var.azure_subnet_cidr1}"]
  resource_group_name  = azurerm_resource_group.resource_group_1.name
  virtual_network_name = azurerm_virtual_network.virtual_network_1.name
}

# Subnet used for the azurerm_virtual_network_gateway only
# https://learn.microsoft.com/en-us/azure/expressroute/expressroute-about-virtual-network-gateways#gwsub
resource "azurerm_subnet" "subnet_gw" {
  provider             = azurerm
  name                 = "GatewaySubnet"
  address_prefixes     = ["${var.azure_subnet_cidr2}"]
  resource_group_name  = azurerm_resource_group.resource_group_1.name
  virtual_network_name = azurerm_virtual_network.virtual_network_1.name
}