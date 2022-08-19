## General VARs
variable "tag_name" {
  default = "demo-pf-aws"
}

## PacketFabic VARs
variable "pf_api_key" {
  type        = string
  description = "PacketFabric platform API access key"
  sensitive   = true
}
variable "pf_account_uuid" {
  type = string
}
variable "pf_api_server" {
  type        = string
  default     = "https://api.packetfabric.com"
  description = "PacketFabric API endpoint URL"
}
# AWS Hosted Connection
variable "pf_aws_account_id" {
  type = number
}
variable "pf_port_circuit_id" {
  type    = string
  default = "PF-AP-LAX2-1234567"
}
variable "pf_cs_pop1" {
  type    = string
  default = "SFO6"
}
variable "pf_cs_zone1" {
  type    = string
  default = "A"
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
  type     = number
  default  = 64538 # private (64512 to 65534)
  nullable = false
}
variable "aws_region1" {
  type        = string
  description = "AWS region"
  default     = "us-west-1"
}
variable "aws_region1_zone1" {
  type        = string
  description = "AWS Availability Zone"
  default     = "us-west-1b"
}
variable "vpc_cidr1" {
  type        = string
  description = "CIDR for the VPC"
  default     = "10.8.0.0/16"
}
# Subnet Variables
variable "subnet_cidr1" {
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