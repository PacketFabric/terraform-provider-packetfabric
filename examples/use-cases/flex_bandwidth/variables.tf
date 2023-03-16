## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-flex"
}

## PacketFabic VARs

# Port
variable "pf_port_pop1" {
  type    = string
  default = "DEN1"
}
variable "pf_port_avzone1" {
  type    = string
  default = "F"
}
variable "pf_port_pop2" {
  type    = string
  default = "CHI6"
}
variable "pf_port_avzone2" {
  type    = string
  default = "B"
}
variable "pf_port_media" {
  type    = string
  default = "LR"
}
variable "pf_port_subterm" {
  type    = number
  default = 1 # default 1 month
}
variable "pf_port_speed" {
  type    = string
  default = "10Gbps"
}
variable "pf_port_nni" {
  type    = bool
  default = false
}

# Flex Bandwidth container
variable "pf_flex_subscription_term" {
  type    = number
  default = 1 # default 1 month
}
variable "pf_flex_capacity" {
  type    = string
  default = "100Gbps" # 50Gbps 100Gbps 150Gbps 200Gbps 250Gbps 300Gbps 350Gbps 400Gbps 450Gbps 500Gbps
}

# Virtual Circuits 
variable "pf_vc_longhaul_type" {
  type    = string
  default = "dedicated"
}
variable "pf_vc_speed" {
  type    = string
  default = "1Gbps" # 1Gbps - 2Gbps
}
variable "pf_vc_subterm" {
  type    = number
  default = 1 # default 1 month
}
# VLANs for VC1
variable "pf_vc_vlan1" {
  type    = number
  default = 145
}
variable "pf_vc_vlan2" {
  type    = number
  default = 146
}
# VLANs for VC2
variable "pf_vc_vlan3" {
  type    = number
  default = 147
}
variable "pf_vc_vlan4" {
  type    = number
  default = 148
}
