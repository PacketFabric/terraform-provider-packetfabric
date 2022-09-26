## General VARs
variable "tag_name" {
  default = "demo-pf-gcp-vpn"
}

# GCP VARs
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
# You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
variable "gcp_side_asn1" {
  type        = number
  default     = 16550
  description = "Google Cloud ASN"
}
variable "gcp_zone1" {
  type        = string
  default     = "us-west1-a"
  description = "Google Cloud zone"
}
variable "subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.5.1.0/24"
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
# To avoid manually setting the BGP session name, vote for:
# https://github.com/hashicorp/terraform-provider-google/issues/11458
# https://github.com/hashicorp/terraform-provider-google/issues/12624
variable "google_cloud_router_bgp_peer_name" {
  type    = string
  default = "auto-ia-bgp-demo-pf-aws-rel-5d2456c8e1b2fb4"
}

# PacketFabric Google Cloud Router Connections
variable "pf_crc_pop1" {
  type    = string
  default = "LAX2"
}
variable "pf_crc_speed" {
  type    = string
  default = "100Mbps" # 50Mbps
}
variable "pf_crc_maybe_nat" {
  type    = bool
  default = false
}

# PacketFabric Cloud Router BGP Session
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
variable "vpn_remote_address" {
  type    = string
  default = "169.254.51.1/29"
}
variable "vpn_l3_address" {
  type    = string
  default = "169.254.51.2/29"
}

# PacketFabric IPsec Cloud Router Connections
variable "vpn_side_asn2" {
  type        = number
  default     = 64534 # private (64512 to 65534)
  description = "VPN Side ASN"
}
variable "subnet_cidr2" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.6.1.0/24"
}
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
  type    = string
  default = "superCoolKey"
}