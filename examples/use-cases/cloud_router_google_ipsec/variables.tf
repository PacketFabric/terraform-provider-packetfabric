## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-gcp-vpn"
}

## Google VARs
variable "gcp_project_id" {
  type = string
  # sensitive   = true
  description = "Google Cloud project ID"
}

variable "gcp_credentials" {
  type        = string
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
variable "google_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.5.1.0/24"
}
variable "public_key" {
  type        = string
  description = "Public Key used to access demo Virtual Machines."
  sensitive   = true
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
  type        = string
  description = "The UUID for the billing account (Find it under Billing > Accounts in the Portal)"
  sensitive   = true
}
variable "pf_api_server" {
  type        = string
  description = "PacketFabric API endpoint URL"
  default     = "https://api.packetfabric.com"
}

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

# PacketFabric Cloud Router Connection - Google and IPsec
variable "pf_crc_speed" {
  type        = string
  description = "The speed of the new connection"
  default     = "50Mbps" # 1Gbps
}

# PacketFabric Cloud Router Connection - Google 
variable "pf_crc_pop1" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "SFO1"
}
variable "pf_crc_maybe_nat" {
  type        = bool
  description = "Set this to true if you intend to use NAT on this connection"
  default     = false
}

# PacketFabric Cloud Router Connections - IPsec
variable "pf_crc_pop2" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "CHI1"
}
variable "pf_crc_ike_version" {
  type        = number
  description = "The Internet Key Exchange (IKE) version supported by your device"
  default     = 1
}
variable "pf_crc_phase1_authentication_method" {
  type        = string
  description = "The authentication method to use during phase 1"
  default     = "pre-shared-key"
}
variable "pf_crc_phase1_group" {
  type        = string
  description = "Phase 1 is when the VPN peers are authenticated and we establish security associations"
  default     = "group14"
}
variable "pf_crc_phase1_encryption_algo" {
  type        = string
  description = "The encryption algorithm to use during phase 1"
  default     = "aes-128-cbc"
}
variable "pf_crc_phase1_authentication_algo" {
  type        = string
  description = "The authentication algorithm to use during phase 1"
  default     = "sha1"
}
variable "pf_crc_phase1_lifetime" {
  type        = number
  description = "The time in seconds before a tunnel will need to re-authenticate"
  default     = 10800
}
variable "pf_crc_phase2_pfs_group" {
  type        = string
  description = "Phase 2 is when SAs are further established to protect and encrypt IP traffic within the tunnel"
  default     = "group14"
}
variable "pf_crc_phase2_encryption_algo" {
  type        = string
  description = "The encryption algorithm to use during phase 2"
  default     = "aes-128-cbc"
}
variable "pf_crc_phase2_authentication_algo" {
  type        = string
  description = "The authentication algorithm to use during phase 2"
  default     = "hmac-sha1-96"
}
variable "pf_crc_phase2_lifetime" {
  type        = number
  description = "The time in seconds before phase 2 expires and needs to reauthenticate"
  default     = 28800
}
variable "pf_crc_gateway_address" {
  type        = string
  description = "The gateway address of your VPN device. Because VPNs traverse the public internet, this must be a public IP address owned by you."
  default     = "127.0.0.1"
}
variable "pf_crc_shared_key" {
  type        = string
  description = "The pre-shared-key to use for authentication."
  default     = "superCoolKey"
  sensitive   = true
}

# PacketFabric Cloud Router BGP Session - Google and IPsec
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

# PacketFabric Cloud Router BGP Session - IPsec
variable "vpn_side_asn2" {
  type        = number
  default     = 64534 # private (64512 to 65534)
  description = "VPN Side ASN"
}
variable "vpn_remote_address" {
  type        = string
  description = "The cloud-side router peer IP."
  default     = "169.254.51.1/29"
}
variable "vpn_l3_address" {
  type        = string
  description = "The L3 address of this instance."
  default     = "169.254.51.2/29"
}