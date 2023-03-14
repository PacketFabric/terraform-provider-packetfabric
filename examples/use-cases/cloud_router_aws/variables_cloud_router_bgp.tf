# PacketFabric Cloud Router BGP Session
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