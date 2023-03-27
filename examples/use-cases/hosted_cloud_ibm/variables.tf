## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-ibm"
}
variable "pf_labels" {
  type        = list(string)
  description = "A list of labels to be applied to PacketFabric resources. These labels will be visible in the PacketFabric Portal and can be searched for easier resource identification."
  default     = ["terraform"] # Example: ["terraform", "dev"]
}

## PacketFabic VARs
# Port
variable "pf_port_pop1" {
  type    = string
  default = "PDX1"
}
variable "pf_port_avzone1" {
  type    = string
  default = "A"
}
variable "pf_port_media" {
  type    = string
  default = "LX"
}
variable "pf_port_subterm" {
  type    = number
  default = 1 # default 1 month
}
variable "pf_port_autoneg" {
  type    = bool
  default = true # only for 1Gbps
}
variable "pf_port_speed" {
  type    = string
  default = "1Gbps"
}
variable "pf_port_nni" {
  type    = bool
  default = false
}

# IBM Hosted Connection
variable "pf_cs_pop1" {
  type    = string
  default = "LAB1" # SFO1
}
variable "pf_cs_zone1" {
  type    = string
  default = "B" # C
}
variable "pf_cs_speed" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_vlan1" {
  type    = number
  default = 110
}
variable "pf_cs_peer_asn" {
  type    = number
  default = 64535 # private (64512 to 65534)
}

# IBM VARs
variable "ibm_resource_group" {
  type        = string
  default     = "My Resource Group"
  description = "IBM Resource Group"
}
variable "ibm_region1" {
  type        = string
  default     = "us-east"
  description = "IBM Cloud region"
}
variable "ibm_region1_zone1" {
  type        = string
  description = "IBM Availability Zone"
  default     = "us-east-1" # "us-south-1"
}