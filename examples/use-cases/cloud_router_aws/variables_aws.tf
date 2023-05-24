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
  default     = "10.1.0.0/16" # do not use 172.17.0.1/16, internal network used for docker containers used in the demo VMs
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
  default     = "10.2.0.0/16" # do not use 172.17.0.1/16, internal network used for docker containers used in the demo VMs
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