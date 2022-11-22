## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-a-side"
}

## PacketFabic VARs

# Marketplace Service (go to the PacketFabric Portal to get the list of services available)
variable "pf_z_side_routing_id" {
  type    = string
  default = "PD-WUY-9VB0" # Demo A
}
variable "pf_z_side_market" {
  type    = string
  default = "NYC" # Demo A
}

# Virtual Circuit
variable "pf_a_side_port_id" {
  type    = string
  default = "PF-AP-PHX2-1748037"
}
variable "pf_z_side_port_id" {
  type    = string
  default = "PF-AP-NYC10-1739866"
}
variable "pf_a_side_vc_vlan1" {
  type    = number
  default = 40
}
variable "pf_z_side_vc_vlan2" {
  type    = number
  default = 50
}
variable "pf_vc_longhaul_type" {
  type    = string
  default = "dedicated"
}
variable "pf_vc_speed" {
  type    = string
  default = "200Mbps"
}
variable "pf_vc_subterm" {
  type    = number
  default = 1 # default 1 month
}
