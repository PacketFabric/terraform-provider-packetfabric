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
  default     = "5Gbps"
}
variable "pf_cr_regions" {
  type        = list(string)
  description = "The regions in which the Cloud Router connections will be located."
  default     = ["US"] # ["UK"] ["US", "UK"]
}

# PacketFabric Cloud Router Connection - AWS
variable "pf_crc_pop1" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "NYC1" # aws_region=us-east-1 when using pf_crc_pop1=NYC1
}
variable "pf_crc_zone1" {
  type    = string
  default = "C" # login to the portal https://portal.packetfabric.com and start a workflow to create a connection (but don't create it, just note the pop/zone info to use in Terraform)
}
variable "pf_crc_speed1" {
  type        = string
  description = "The speed of the new connection AWS"
  default     = "1Gbps"
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

# PacketFabric Cloud Router Connection - Port
variable "pf_crc_port_circuit_id" {
  type        = string
  description = "Port Circuit ID used as a source port to create a Port Cloud Router Connection"
  default     = "PF-AP-ATL3-1751473"
}
variable "pf_crc_speed3" {
  type        = string
  description = "The speed of the new connection Port"
  default     = "1Gbps"
}
variable "pf_crc_vlan3" {
  type    = number
  default = 170
}
variable "pf_crc_untagged" {
  type    = bool
  default = false
}