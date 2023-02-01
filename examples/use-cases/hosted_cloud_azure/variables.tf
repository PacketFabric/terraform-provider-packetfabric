## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-azure"
}

## PacketFabic VARs
# Port
variable "pf_port_pop1" {
  type    = string
  default = "PDX1"
}
variable "pf_port_avzone1" {
  type    = string
  default = "A"
}
variable "pf_port_media" {
  type    = string
  default = "LX"
}
variable "pf_port_subterm" {
  type    = number
  default = 1 # default 1 month
}
variable "pf_port_autoneg" {
  type    = bool
  default = false
}
variable "pf_port_speed" {
  type    = string
  default = "1Gbps"
}
variable "pf_port_nni" {
  type    = bool
  default = false
}

# Azure Hosted Connection
variable "pf_cs_speed" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_vlan_private" {
  type    = number
  default = 100
}
variable "pf_cs_vlan_microsoft" {
  type    = number
  default = 101
}

## Azure VARs
# https://docs.microsoft.com/en-us/azure/availability-zones/az-overview
variable "azure_region1" {
  type        = string
  description = "Azure region"
  default     = "East US"
}

# https://docs.microsoft.com/en-us/azure/expressroute/expressroute-locations-providers
# West US (Silicon Valley)
# West Central US (Denver)
# North Central US (Chicago)
# East US, East US2 (New York, Washington DC)
# South Central US (Dallas)
# Las Vegas
variable "peering_location_1" {
  type        = string
  description = "Azure Peering Location"
  default     = "New York"
}
variable "bandwidth_in_mbps" {
  type        = string
  description = "Azure Bandwidth"
  default     = 50
}
variable "service_provider_name" {
  type    = string
  default = "PacketFabric"
}
variable "sku_tier" {
  type    = string
  default = "Standard" # Standard or Premium
}
variable "sku_family" {
  type    = string
  default = "MeteredData"
}

# Express Route GW SKUs ErGw1AZ, ErGw2AZ, ErGw3AZ
variable "vnet_cidr1" {
  type        = string
  description = "CIDR for the VNET"
  default     = "10.3.0.0/16"
}
variable "subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.3.1.0/24"
}
variable "subnet_cidr1gw" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.3.2.0/24"
}

# BGP peering
variable "peer_asn" {
  type    = number
  default = 64535 # private (64512 to 65534)
}
variable "primary_peer_address_prefix" {
  type    = string
  default = "169.254.247.40/30"
}
variable "secondary_peer_address_prefix" {
  type    = string
  default = "169.254.248.40/30"
}
variable "shared_key" {
  type      = string
  default   = "dd02c7c2232759874e1c20558" # echo "secret" | md5sum | cut -c1-25
  sensitive = true
}

