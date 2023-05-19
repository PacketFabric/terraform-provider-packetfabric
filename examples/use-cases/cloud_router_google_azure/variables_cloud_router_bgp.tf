# PacketFabric Cloud Router BGP Session - Google and Azure
variable "pf_crbs_af" {
  type        = string
  description = "Whether this instance is IPv4 or IPv6. At this time, only IPv4 is supported"
  default     = "v4"
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
