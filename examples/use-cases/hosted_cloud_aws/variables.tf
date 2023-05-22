## General VARs
# Must follow ^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$
# Any lowercase ASCII letter or digit, and possibly hyphen, which should start with a letter and end with a letter or digit, 
# and have at most 63 characters (1 for the starting letter + up to 61 characters in the middle + 1 for the ending letter/digit).
variable "resource_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-aws"
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

# AWS Hosted Connection
variable "pf_cs_pop1" {
  type    = string
  default = "SFO6"
}
variable "pf_cs_zone1" {
  type    = string
  default = "A" # login to the portal https://portal.packetfabric.com and start a workflow to create a connection (but don't create it, just note the pop/zone info to use in Terraform)
}
variable "pf_cs_speed" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_vlan1" {
  type    = number
  default = 110
}
# AWS Cloud Side Provisioning
variable "pf_cs_mtu" {
  description = "Maximum Transmission Unit this port supports (size of the largest supported PDU)"
  type        = number
  default     = 1500
}

# AWS VARs
variable "amazon_side_asn1" { # used in BGP session
  type    = number
  default = 64538 # private (64512 to 65534)
}
variable "customer_side_asn1" { # used in BGP session
  type    = number
  default = 64539 # private (64512 to 65534)
}
variable "aws_region1" {
  type        = string
  description = "AWS region"
  default     = "us-west-1" # has to be in the same region as aws_region1 var
}
variable "aws_vpc_cidr1" {
  type        = string
  description = "CIDR for the VPC"
  default     = "10.8.0.0/16"
}
# Subnet Variables
variable "aws_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.8.1.0/24"
}

