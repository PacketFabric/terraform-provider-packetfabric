## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-aws"
}

## PacketFabic VARs
# AWS Hosted Connection
variable "pf_aws_account_id" {
  type        = number
  description = "The AWS account ID to connect with. Must be 12 characters long"
}
variable "pf_port_circuit_id" {
  type    = string
  default = "PF-AP-WDC1-1726464"
}
variable "pf_cs_pop1" {
  type    = string
  default = "SFO6"
}
variable "pf_cs_zone1" {
  type    = string
  default = "A" # check availability /v2/locations/cloud?cloud_connection_type=hosted&cloud_provider=aws&pop=SFO6
}
variable "pf_cs_speed" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_vlan1" {
  type    = number
  default = 110
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

variable "aws_access_key" {
  type        = string
  description = "AWS access key"
  sensitive   = true
}
variable "aws_secret_key" {
  type        = string
  description = "AWS secret key"
  sensitive   = true
}