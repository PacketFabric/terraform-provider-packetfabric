
# PacketFabric Cloud Router BGP Session - Google and IPsec
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

# PacketFabric Cloud Router BGP Session - IPsec
variable "vpn_side_asn2" {
  type        = number
  default     = 64534 # private (64512 to 65534)
  description = "VPN Side ASN"
}
variable "vpn_remote_address" {
  type        = string
  description = "The cloud-side router peer IP."
  default     = "169.254.51.1/29"
}
variable "vpn_l3_address" {
  type        = string
  description = "The L3 address of this instance."
  default     = "169.254.51.2/29"
}