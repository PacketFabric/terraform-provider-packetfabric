## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-aws-gcp"
}
variable "public_key" {
  type        = string
  description = "Public Key used to access demo Virtual Machines."
  sensitive   = true
}

## Google VARs
variable "gcp_project_id" {
  type = string
  # sensitive   = true
  description = "Google Cloud project ID"
}
variable "gcp_credentials_path" {
  type        = string
  sensitive   = true
  description = "Google Cloud service account credentials (path to GCP json file)"
}
# https://cloud.google.com/compute/docs/regions-zones
variable "gcp_region1" {
  type        = string
  default     = "us-west1"
  description = "Google Cloud region"
}
variable "gcp_zone1" {
  type        = string
  default     = "us-west1-a"
  description = "Google Cloud zone"
}
variable "gcp_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.5.1.0/24"
}

## AWS VARs
variable "amazon_side_asn1" {
  type        = number
  description = "Direct Connect Gateway ASN. Used in BGP session"
  default     = 64532 # private (64512 to 65534)
}
variable "amazon_side_asn2" {
  type        = number
  description = "Transit Gateway ASN. (must be different than Direct Connect Gateway)"
  default     = 64533 # private (64512 to 65534)
}
# Make sure you set the correct AWS region based on the PacketFabric pop selected
# Find details on location https://packetfabric.com/locations/cloud-on-ramps and https://aws.amazon.com/directconnect/locations/)
# Essentially, select the PacketFabric pop the closest to the AWS region you want to connect to. 
# Example: AWS region us-west-1 is the closest to PacketFabric pop LAX1.
variable "aws_region1" {
  type        = string
  description = "AWS region 1"
  default     = "us-east-1" # aws_region2=us-east-1 when using pf_crc_pop2=NYC1
}
variable "aws_vpc_cidr1" { # used in PF BGP prefix
  type        = string
  description = "CIDR for the VPC in AWS Region 2"
  default     = "10.2.0.0/16"
}
# Subnet Variables
variable "aws_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet in AWS Region 2"
  default     = "10.2.1.0/24"
}
# Make sure you setup the correct AMI if you chance default AWS region1
variable "ec2_ami1" {
  description = "Ubuntu 22.04 in aws_region2 (e.g. us-east-1)"
  default     = "ami-052efd3df9dad4825"
}
variable "ec2_instance_type" {
  description = "Instance Type/Size"
  default     = "t2.micro" # Free tier
}
variable "pf_aws_account_id" {
  type        = number
  description = "The AWS account ID to connect with. Must be 12 characters long"
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

# PacketFabric Google Cloud Router Connection - AWS and Google
variable "pf_crc_speed" {
  type        = string
  description = "The speed of the new connection"
  default     = "50Mbps" # must match bandwidth_in_mbps
}
variable "pf_crc_maybe_nat" {
  type        = bool
  description = "Set this to true if you intend to use NAT on this connection"
  default     = false
}

# PacketFabric AWS Cloud Router Connection - AWS
variable "pf_crc_pop1" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "NYC1" # aws_region=us-east-1 when using pf_crc_pop1=NYC1
}
variable "pf_crc_zone1" {
  type    = string
  default = "C" # check availability /v2/locations/cloud?cloud_connection_type=hosted&has_cloud_router: true=true&cloud_provider=aws&pop=PDX2
}

# PacketFabric Google Cloud Router Connection - Google
variable "pf_crc_pop2" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "SFO1"
}

# PacketFabric Google Cloud Router Connection - Azure ExpressRoute Circuit
variable "pf_crc_is_public" {
  type        = bool
  description = "Whether PacketFabric should allocate a public IP address for this connection"
  default     = false # set to true if peering_type = MicrosoftPeering
}


# PacketFabric Cloud Router BGP Session - Google and Azure
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

# PacketFabric Cloud Router BGP Session - Google
# You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
variable "gcp_side_asn1" {
  type        = number
  default     = 16550
  description = "Google Cloud ASN"
}