## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-aws"
}
variable "public_key" {
  type        = string
  description = "Public Key used to access demo EC2 instances."
  sensitive   = true
}

## AWS VARs
variable "amazon_side_asn1" {
  type        = number
  description = "Direct Connect Gateway ASN Region 1. Used in BGP session. Also set in Virtual Private Gateway 1."
  default     = 64532 # private (64512 to 65534)
}
variable "amazon_side_asn2" { # used in BGP session
  type        = number
  description = "Direct Connect Gateway ASN Region 2. Used in BGP session. Also set in Virtual Private Gateway 2."
  default     = 64533 # private (64512 to 65534)
}
variable "aws_region1" {
  type        = string
  description = "AWS region 1"
  default     = "us-west-1" # aws_region1=us-west-1 when using pf_crc_pop1=LAX1
}
variable "aws_region2" {
  type        = string
  description = "AWS region 2"
  default     = "us-east-1" # aws_region2=us-east-1 when using pf_crc_pop2=NYC1
}
variable "aws_vpc_cidr1" { # used in PF BGP prefix
  type        = string
  description = "CIDR for the VPC in AWS Region 1"
  default     = "10.1.0.0/16"
}
# Subnet Variables
variable "aws_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet in AWS Region 1"
  default     = "10.1.1.0/24"
}
variable "aws_vpc_cidr2" { # used in PF BGP prefix
  type        = string
  description = "CIDR for the VPC in AWS Region 2"
  default     = "10.2.0.0/16"
}
# Subnet Variables
variable "aws_subnet_cidr2" {
  type        = string
  description = "CIDR for the subnet in AWS Region 2"
  default     = "10.2.1.0/24"
}
# Make sure you setup the correct AMI if you chance default AWS regions 1 and 2
variable "ec2_ami1" {
  description = "Ubuntu 22.04 in aws_region1 (e.g. us-west-1)"
  default     = "ami-085284d24fe829cd0"
}
variable "ec2_ami2" {
  description = "Ubuntu 22.04 in aws_region2 (e.g. us-east-1)"
  default     = "ami-052efd3df9dad4825"
}
variable "ec2_instance_type" {
  description = "Instance Type/Size"
  default     = "t2.micro" # Free tier
}

## PacketFabic VARs
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
  default     = "1Gbps" # 100Mbps
}
variable "pf_cr_regions" {
  type        = list(string)
  description = "The regions in which the Cloud Router connections will be located."
  default     = ["US"] # ["UK"] ["US", "UK"]
}

# PacketFabric Cloud Router Connections
variable "pf_crc_pop1" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "SFO6" # aws_region1=us-west-1 when using pf_crc_pop1=LAX1
}
variable "pf_crc_zone1" {
  type    = string
  default = "B"
}
variable "pf_crc_pop2" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "NYC1" # aws_region2=us-east-1 when using pf_crc_pop2=NYC1
}
variable "pf_crc_zone2" {
  type    = string
  default = "C"
}
variable "pf_crc_speed" {
  type        = string
  description = "The speed of the new connection"
  default     = "50Mbps" # 1Gbps
}
variable "pf_crc_maybe_nat" {
  type        = bool
  description = "Set this to true if you intend to use NAT on this connection"
  default     = false
}
variable "pf_crc_is_public" {
  type        = bool
  description = "Whether PacketFabric should allocate a public IP address for this connection"
  default     = false
}

# PacketFabric Cloud Router BGP Session
variable "pf_crbs_af" {
  type        = string
  description = "Whether this instance is IPv4 or IPv6. At this time, only IPv4 is supported"
  default     = "v4"
}
variable "pf_crbs_mhttl" {
  type        = number
  description = "The TTL of this session. The default is 1."
  default     = 1
}
variable "pf_crbs_orlonger" {
  type        = bool
  description = "Whether to use exact match or longer for all prefixes"
  default     = true # Allow longer prefixes
}