# PacketFabric Cloud-Router
variable "pf_cr_asn" {
  type        = number
  description = "The ASN of the cloud router"
  default     = 4556 # PacketFabric ASN
  nullable    = false
}
variable "pf_cr_capacity" {
  type        = string
  description = "The cloud router capacity"
  default     = "1Gbps" # 100Mbps
}
variable "pf_cr_regions" {
  type        = list(string)
  description = "The regions in which the Cloud Router connections will be located."
  default     = ["US"] # ["UK"] ["US", "UK"]
}

# PacketFabric Google Cloud Router Connection - Google and Azure
variable "pf_crc_speed" {
  type        = string
  description = "The speed of the new connection"
  default     = "50Mbps" # must match bandwidth_in_mbps for Azure Express
}
variable "pf_crc_maybe_nat" {
  type        = bool
  description = "Set this to true if you intend to use NAT on this connection"
  default     = false
}

# PacketFabric Google Cloud Router Connection - Google
variable "pf_crc_pop1" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "SFO1"
}

# PacketFabric Google Cloud Router Connection - Azure ExpressRoute Circuit
variable "pf_crc_is_public" {
  type        = bool
  description = "Whether PacketFabric should allocate a public IP address for this connection"
  default     = false # set to true if peering_type = MicrosoftPeering
}
# https://docs.microsoft.com/en-us/azure/expressroute/expressroute-locations-providers
# West US (Silicon Valley)
# West Central US (Denver)
# North Central US (Chicago)
# East US, East US2 (New York, Washington DC)
# South Central US (Dallas)
# Las Vegas
# The pop is defined on the Azure side
variable "azure_peering_location_1" {
  type        = string
  description = "Azure Peering Location"
  default     = "New York"
}
variable "azure_bandwidth_in_mbps" {
  type        = string
  description = "Azure Bandwidth"
  default     = 50 # must match pf_crc_speed for Azure Cloud Router Connection
}
variable "azure_service_provider_name" {
  type    = string
  default = "PacketFabric"
}
variable "azure_sku_tier" {
  type    = string
  default = "Standard" # Standard or Premium
}
variable "azure_sku_family" {
  type    = string
  default = "MeteredData"
}