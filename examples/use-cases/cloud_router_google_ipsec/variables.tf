## General VARs
variable "tag_name" {
  default = "demo-pf-gcp-vpn"
}

## Google VARs
variable "gcp_project_id" {
  type = string
  # sensitive   = true
  description = "Google Cloud project ID"
}

variable "gcp_credentials" {
  type        = string
  description = "Google Cloud service account credentials"
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
variable "google_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.5.1.0/24"
}
variable "public_key" {
  sensitive = true
}

## IPsec VAR
variable "ipsec_subnet_cidr2" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.6.1.0/24"
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

# PacketFabric Cloud Router Connection - Google and IPsec
variable "pf_crc_speed" {
  type    = string
  default = "50Mbps" # 1Gbps
}

# PacketFabric Cloud Router Connection - Google 
variable "pf_crc_pop1" {
  type    = string
  default = "SFO1"
}
variable "pf_crc_maybe_nat" {
  type    = bool
  default = false
}

# PacketFabric Cloud Router Connections - IPsec
variable "pf_crc_pop2" {
  type    = string
  default = "CHI1"
}
variable "pf_crc_ike_version" {
  type    = number
  default = 1
}
variable "pf_crc_phase1_authentication_method" {
  type    = string
  default = "pre-shared-key"
}
variable "pf_crc_phase1_group" {
  type    = string
  default = "group14"
}
variable "pf_crc_phase1_encryption_algo" {
  type    = string
  default = "aes-128-cbc"
}
variable "pf_crc_phase1_authentication_algo" {
  type    = string
  default = "sha1"
}
variable "pf_crc_phase1_lifetime" {
  type    = number
  default = 10800
}
variable "pf_crc_phase2_pfs_group" {
  type    = string
  default = "group14"
}
variable "pf_crc_phase2_encryption_algo" {
  type    = string
  default = "aes-128-cbc"
}
variable "pf_crc_phase2_authentication_algo" {
  type    = string
  default = "hmac-sha1-96"
}
variable "pf_crc_phase2_lifetime" {
  type    = number
  default = 28800
}
variable "pf_crc_gateway_address" {
  type    = string
  default = "127.0.0.1"
}
variable "pf_crc_shared_key" {
  type      = string
  default   = "superCoolKey"
  sensitive = true
}

# PacketFabric Cloud Router BGP Session - Google and IPsec
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

# PacketFabric Cloud Router BGP Session - Google
# You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
variable "gcp_side_asn1" {
  type        = number
  default     = 16550
  description = "Google Cloud ASN"
}

# PacketFabric Cloud Router BGP Session - IPsec
variable "vpn_side_asn2" {
  type        = number
  default     = 64534 # private (64512 to 65534)
  description = "VPN Side ASN"
}
variable "vpn_remote_address" {
  type    = string
  default = "169.254.51.1/29"
}
variable "vpn_l3_address" {
  type    = string
  default = "169.254.51.2/29"
}