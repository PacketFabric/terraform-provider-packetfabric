## General VARs
# Must follow ^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$
# Any lowercase ASCII letter or digit, and possibly hyphen, which should start with a letter and end with a letter or digit, 
# and have at most 63 characters (1 for the starting letter + up to 61 characters in the middle + 1 for the ending letter/digit).
variable "resource_name" {
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
  default = "WDC1"
}
variable "pf_port_avzone1" {
  type    = string
  default = "E"
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
  default = "WDC1" # WDC1, SFO1, DAL2
}
variable "pf_cs_zone1" {
  type    = string
  default = "E" # login to the portal https://portal.packetfabric.com and start a workflow to create a connection (but don't create it, just note the pop/zone info to use in Terraform)
}
variable "pf_cs_speed" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_speed_ibm" {
  type    = number
  default = 50
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
variable "ibm_vpc_cidr1" {
  type        = string
  description = "CIDR for the VPC"
  default     = "10.8.0.0/16"
}
variable "ibm_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.8.1.0/24"
}