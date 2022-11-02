## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-gcp-azure"
}
variable "public_key" {
  type        = string
  description = "Public Key used to access demo Virtual Machines."
  sensitive   = true
}

## Google VARs
variable "gcp_project_id" {
  type = string
  # sensitive   = true
  description = "Google Cloud project ID"
}
variable "gcp_credentials_path" {
  type        = string
  sensitive   = true
  description = "Google Cloud service account credentials (path to GCP json file)"
}
# https://cloud.google.com/compute/docs/regions-zones
variable "gcp_region1" {
  type        = string
  default     = "us-west1"
  description = "Google Cloud region"
}
variable "gcp_zone1" {
  type        = string
  default     = "us-west1-a"
  description = "Google Cloud zone"
}
variable "gcp_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.5.1.0/24"
}

## Azure VARs
variable "subscription_id" {
  type        = string
  description = "Azure Subscription ID"
  sensitive   = true
}
variable "client_id" {
  type        = string
  description = "Azure Client ID"
  sensitive   = true
}
variable "client_secret" {
  type        = string
  description = "Azure Client Secret ID"
  sensitive   = true
}
variable "tenant_id" {
  type        = string
  description = "Azure Tenant ID"
  sensitive   = true
}
# https://docs.microsoft.com/en-us/azure/availability-zones/az-overview
variable "azure_region1" {
  type        = string
  description = "Azure region"
  default     = "East US"
}
variable "azure_vnet_cidr1" {
  type        = string
  description = "CIDR for the VNET"
  default     = "10.7.0.0/16"
}
variable "azure_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.7.1.0/24"
}
variable "azure_subnet_cidr2" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.7.2.0/24"
}

## PacketFabic VARs
variable "pf_api_key" {
  type        = string
  description = "PacketFabric platform API access key"
  sensitive   = true
  default     = "secret"
}
variable "pf_account_uuid" {
  type        = string
  description = "The UUID for the billing account (Find it under Billing > Accounts in the Portal)"
  default     = "34ff9ffb-9bbf-43b1-9cf8-6c8e62370597"
}
variable "pf_api_server" {
  type        = string
  description = "PacketFabric API endpoint URL"
  default     = "https://api.packetfabric.com"
}

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
  default     = "50Mbps" # must match bandwidth_in_mbps
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
  default     = 50 # must match pf_crc_speed
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

# PacketFabric Cloud Router BGP Session - Google and Azure
variable "pf_crbs_af" {
  type        = string
  description = "Whether this instance is IPv4 or IPv6. At this time, only IPv4 is supported"
  default     = "v4"
}
variable "pf_crbs_mhttl" {
  type        = number
  description = "The TTL of this session. The default is 1."
  default     = 1
}
variable "pf_crbs_orlonger" {
  type        = bool
  description = "Whether to use exact match or longer for all prefixes"
  default     = true # Allow longer prefixes
}

# PacketFabric Cloud Router BGP Session - Google
# You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
variable "gcp_side_asn1" {
  type        = number
  default     = 16550
  description = "Google Cloud ASN"
}

# PacketFabric Cloud Router BGP Session - Azure
variable "azure_side_asn1" {
  type        = number
  default     = 12076 # cannot be changed
  description = "Azure Cloud ASN"
}
variable "azure_peering_type" {
  type    = string
  default = "AzurePrivatePeering" # MicrosoftPeering 
}
variable "azure_primary_peer_address_prefix" {
  type        = string
  description = "Provide this as the primary subnet when creating an Azure cloud router connection"
  default     = "169.254.247.40/30" # Use Public IP if using MicrosoftPeering
}
variable "azure_secondary_peer_address_prefix" {
  type        = string
  description = "Provide this as the secondary subnet when creating an Azure cloud router connection"
  default     = "169.254.247.44/30" # Use Public IP if using MicrosoftPeering
}
variable "azure_bgp_shared_key" {
  type        = string
  description = "The MD5 value of the authenticated BGP sessions."
  default     = "secret"
  sensitive   = true
}
