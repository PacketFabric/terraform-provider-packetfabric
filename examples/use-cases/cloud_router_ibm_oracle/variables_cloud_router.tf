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
  default     = "2Gbps"
}
variable "pf_cr_regions" {
  type        = list(string)
  description = "The regions in which the Cloud Router connections will be located."
  default     = ["US"] # ["UK"] ["US", "UK"]
}

# PacketFabric Cloud Router Connection - Common
variable "pf_crc_maybe_nat" {
  type        = bool
  description = "Set this to true if you intend to use NAT on this connection"
  default     = false
}

# PacketFabric Cloud Router Connection - IBM
variable "pf_crc_pop1" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "SFO1" # WDC1, SFO1, DAL2
}
variable "pf_crc_zone1" {
  type    = string
  default = "C"
}
variable "pf_crc_speed" {
  type        = string
  description = "The speed of the new connection"
  default     = "1Gbps"
}
variable "pf_crc_speed_ibm" {
  type        = number
  description = "The speed of the new connection"
  default     = 1000
}

# PacketFabric Cloud Router Connection - Oracle
variable "pf_crc_pop2" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "WDC2"
}
variable "pf_crc_zone2" {
  type    = string
  default = "F" # login to the portal https://portal.packetfabric.com and start a workflow to create a connection (but don't create it, just note the pop/zone info to use in Terraform)
}
variable "oracle_bandwidth_shape_name" {
  type    = string
  default = "1 Gbps" # 1 Gbps, 10 Gbps, or 100 Gbps increments
}