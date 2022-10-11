## General VARs
variable "tag_name" {
  default = "demo-pf-ibm-oracle"
}
variable "public_key" {
  sensitive = true
}

## IBM VARs
variable "ibm_account_id" {
  type        = string
  sensitive   = true
  description = "IBM Account ID"
}
variable "ibmcloud_api_key" {
  type        = string
  sensitive   = true
  description = "IBM API key"
}
variable "iaas_classic_username" {
  type        = string
  sensitive   = true
  description = "IBM Classic Username"
}
variable "iaas_classic_api_key" {
  type        = string
  sensitive   = true
  description = "IBM Classic API key"
}
variable "ibm_region1" {
  type        = string
  default     = "us-south"
  description = "IBM Cloud region"
}
variable "ibm_region1_zone1" {
  type        = string
  description = "IBM Availability Zone"
  default     = "us-south-1"
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

## Oracle VARs
variable "tenancy_ocid" {
  type        = string
  sensitive   = true
  description = "Oracle Tenancy OCID"
}
variable "user_ocid" {
  type        = string
  sensitive   = true
  description = "Oracle User OCID"
}
variable "private_key" {
  type        = string
  sensitive   = true
  description = "Oracle Private Key"
}
# variable "private_key_password" {
#   type        = string
#   sensitive   = true
#   description = "Oracle Private Key Password"
# }
variable "fingerprint" {
  type        = string
  sensitive   = true
  description = "Oracle Public Key fingerprint"
}
variable "parent_compartment_id" {
  type        = string
  description = "Oracle Parent Compartment OCID"
}

variable "oracle_region1" {
  type        = string
  default     = "us-sanjose-1"
  description = "Oracle Cloud region"
}

variable "oracle_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.9.1.0/24"
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

# PacketFabric Cloud-Router
variable "pf_cr_asn" {
  type     = number
  default  = 4556 # PacketFabric ASN
  nullable = false
}
variable "pf_cr_capacity" {
  type    = string
  default = "1Gbps" # 100Mbps
}
variable "pf_cr_regions" {
  type    = list(string)
  default = ["US"] # ["UK"] ["US", "UK"]
}

# PacketFabric Cloud Router Connection - IBM and Oracle
variable "pf_crc_maybe_nat" {
  type    = bool
  default = false
}

# PacketFabric Cloud Router Connection & BGP Session - IBM
variable "pf_crc_pop1" {
  type    = string
  default = "SFO1"
}
variable "pf_crc_zone1" {
  type    = string
  default = "c"
}
variable "pf_crc_speed" {
  type    = string
  default = "50Mbps" # must match bandwidth_in_mbps
}
variable "ibm_bgp_asn" {
  type    = number
  default = 64536 # private (64512 to 65534)
}
variable "ibm_bgp_cer_cidr" {
  type    = string
  default = "169.254.248.41/30"
}
variable "ibm_bgp_ibm_cidr" {
  type    = string
  default = "169.254.248.42/30"
}

# PacketFabric Cloud Router Connection - Oracle
variable "pf_crc_pop2" {
  type    = string
  default = "WDC02"
}
variable "pf_crc_zone2" {
  type    = string
  default = "c"
}

# PacketFabric Cloud Router BGP Session - IBM and Oracle
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

# PacketFabric Cloud Router BGP Session - Oracle
variable "oracle_bandwidth_shape_name" {
  type    = string
  default = "1 Gbps"
}
variable "oracle_peer_asn" {
  type    = number
  default = 64537 # private (64512 to 65534)
}
variable "oracle_primary_peer_address_prefix" {
  type    = string
  default = "169.254.247.41/30"
}
variable "oracle_secondary_peer_address_prefix" {
  type    = string
  default = "169.254.247.42/30"
}
variable "oracle_bgp_shared_key" {
  type      = string
  default   = "dd02c7c2232759874e1c20558" # echo "secret" | md5sum | cut -c1-25
  sensitive = true
}