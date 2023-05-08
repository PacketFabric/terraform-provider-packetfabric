## PacketFabic VARs
# PacketFabric Cloud Router BGP Session - AWS, Oracle and Port
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

# PacketFabric Cloud Router BGP Session - Oracle
variable "oracle_peer_asn" {
  type    = number
  default = 64537 # private (64512 to 65534)
}
variable "oracle_bgp_peering_prefix" {
  type    = string
  default = "169.254.247.41/30"
}
variable "customer_bgp_peering_prefix" {
  type    = string
  default = "169.254.247.42/30"
}
variable "oracle_bgp_shared_key" {
  type      = string
  default   = "dd02c7c2232759874e1c20558" # echo "secret" | md5sum | cut -c1-25
  sensitive = true
}

# PacketFabric Cloud Router BGP Session - PacketFabric Port / On-Premise
variable "on_premise_peer_asn" {
  type    = number
  default = 64538 # private (64512 to 65534)
}
variable "on_premise_bgp_peering_prefix" {
  type    = string
  default = "169.254.248.41/30"
}
variable "pf_side_bgp_peering_prefix" {
  type    = string
  default = "169.254.248.42/30"
}
variable "on_premise_bgp_shared_key" {
  type      = string
  default   = "555f136061769b05aa8fcdc86" # echo "secret2" | md5sum | cut -c1-25
  sensitive = true
}
