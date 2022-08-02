variable "tag_name" {
  default = "demo-pf-aws"
}
variable "amazon_side_asn1" { # used in BGP session
  type     = number
  default  = 64532 # private (64512 to 65534)
  nullable = false
}
variable "amazon_side_asn2" { # used in BGP session
  type     = number
  default  = 64533 # private (64512 to 65534)
  nullable = false
}
# Make sure you set the correct AWS region based on the PacketFabric pop selected
# Find details on location https://packetfabric.com/locations/cloud-on-ramps and https://aws.amazon.com/directconnect/locations/)
# Essentially, select the PacketFabric pop the closest to the AWS region you want to connect to. 
# Example: AWS region us-west-1 is the closest to PacketFabric pop LAX1.
variable "aws_region1" {
  type        = string
  description = "AWS region"
  default     = "us-west-1" # aws_region1=us-west-1 when using pf_crc_pop1=LAX1
}
variable "aws_region1_zone1" {
  type        = string
  description = "AWS Availability Zone"
  default     = "us-west-1a"
}
variable "aws_region2" {
  type        = string
  description = "AWS region"
  default     = "us-east-1" # aws_region2=us-east-1 when using pf_crc_pop2=NYC1
}
variable "aws_region2_zone1" {
  type        = string
  description = "AWS Availability Zone"
  default     = "us-east-1a"
}
variable "vpc_cidr1" { # used in PF BGP prefix
  type        = string
  description = "CIDR for the VPC"
  default     = "10.1.0.0/16"
}
# Subnet Variables
variable "subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.1.1.0/24"
}
variable "vpc_cidr2" { # used in PF BGP prefix
  type        = string
  description = "CIDR for the VPC"
  default     = "10.2.0.0/16"
}
# Subnet Variables
variable "subnet_cidr2" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.2.1.0/24"
}
# Make sure you setup the correct AMI if you chance default AWS regions 1 and 2
variable "ec2_ami1" {
  default = "ami-085284d24fe829cd0" # Ubuntu 22.04 in aws_region1 (e.g. us-west-1)
}
variable "ec2_ami2" {
  default = "ami-052efd3df9dad4825" # Ubuntu 22.04 in aws_region2 (e.g. us-east-1)
}
variable "ec2_instance_type" {
  default = "t2.micro" # Free tier
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
variable "public_key" {
  sensitive = true
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
variable "pf_aws_account_id" {
  type = number
}
variable "pf_api_server" {
  type        = string
  default     = "https://api.packetfabric.com"
  description = "PacketFabric API endpoint URL"
}
variable "pf_provider_source" {
  type    = string
  default = "packetfabric/packetfabric"
}
# PacketFabric Cloud-Router Parameter configurations
variable "pf_cr_asn" {
  type     = number
  default  = 4556 # PacketFabric ASN
  nullable = false
}
# Parameter deprecated
variable "pf_cr_scope" {
  type    = string
  default = "private"
}
variable "pf_cr_capacity" {
  type    = string
  default = "1Gbps" # 100Mbps
}
variable "pf_cr_regions" {
  type    = list(string)
  default = ["US"] # ["UK"] ["US", "UK"]
}

# PacketFabric Cloud-Router-Connections Parameter configuration:
# Make sure you set the correct AWS region based on the PacketFabric pop selected
# Find details on location https://packetfabric.com/locations/cloud-on-ramps and https://aws.amazon.com/directconnect/locations/)
# Essentially, select the PacketFabric pop the closest to the AWS region you want to connect to. 
# Example: AWS region us-west-1 is the closest to PacketFabric pop LAX1.
variable "pf_crc_pop1" {
  type    = string
  default = "LAX1" # aws_region1=us-west-1 when using pf_crc_pop1=LAX1
}
variable "pf_crc_zone1" {
  type    = string
  default = "c" # you may look in the PacketFabric Web Portal to find out the zone available for the pop selected
}
variable "pf_crc_pop2" {
  type    = string
  default = "NYC1" # aws_region2=us-east-1 when using pf_crc_pop2=NYC1
}
variable "pf_crc_zone2" {
  type    = string
  default = "c" # you may look in the PacketFabric Web Portal to find out the zone available for the pop selected
}
variable "pf_crc_speed" {
  type    = string
  default = "100Mbps" # 50Mbps
}
variable "pf_crc_maybe_nat" {
  type    = bool
  default = false
}
variable "pf_crc_is_public" {
  type    = bool
  default = false
}

# PacketFabric Cloud-Router-BGP-Session Parameter configuration:
variable "pf_crbs_af" {
  type    = string
  default = "v4"
}
variable "pf_crbs_mhttl" {
  type    = number
  default = 1
}
variable "pf_crbs_orlonger" {
  type    = bool
  default = true # Allow longer prefixes
}