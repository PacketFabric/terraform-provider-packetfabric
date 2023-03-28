## General VARs
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
# AWS Dedicated Connection
variable "pf_cs_pop1" {
  type    = string
  default = "SFO6" # needs to match with aws_dx_location
}
variable "pf_cs_zone1" {
  type    = string
  default = "A"
}
variable "pf_cs_speed" {
  type    = string
  default = "10Gbps"
}
variable "pf_cs_subterm" {
  type    = number
  default = 1 # default 1 month
}
variable "pf_cs_srvclass" {
  type    = string
  default = "longhaul" # longhaul or metro
}
variable "pf_cs_aws_d_autoneg" {
  type    = bool
  default = false
}
variable "should_create_lag" {
  type    = bool
  default = false
}

# Cross connect
variable "pf_document_uuid1" {
  type    = string
  default = "1d2fb159-b40e-4eda-8f63-1191a80a023e" # use API /v2/documents to get UUID
}
variable "pf_cs_site1" {
  type    = string
  default = "CS-SV4" # realted to pf_cs_pop1, use API /v2/locations to get site_code
}

# Virtual Circuit
variable "pf_interface_a_circuit_id" {
  type    = string
  default = "PF-AP-NYC6-1743863" # AWS dedicated cloud 
}
variable "pf_interface_b_circuit_id" {
  type    = string
  default = "PF-AP-WDC1-1726464" # existing PF port
}
variable "pf_vc_vlan1" {
  type    = number
  default = 145
}
variable "pf_vc_vlan2" {
  type    = number
  default = 146
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

# AWS VARs
variable "amazon_side_asn1" { # used in BGP session
  type     = number
  default  = 64538 # private (64512 to 65534)
  nullable = false
}
variable "customer_side_asn1" { # used in BGP session
  type    = number
  default = 64539 # private (64512 to 65534)
}
variable "aws_region1" {
  type        = string
  description = "AWS region"
  default     = "us-west-1"
}
# https://aws.amazon.com/directconnect/locations/
variable "aws_dx_location" {
  type    = string
  default = "CSSV4" # need to match with pf_cs_pop1 and aws_region1
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

