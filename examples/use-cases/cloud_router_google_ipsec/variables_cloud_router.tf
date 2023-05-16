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

# PacketFabric Cloud Router Connection - Common
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

# PacketFabric Cloud Router Connections - IPsec
variable "pf_crc_pop2" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "SFO1" # CHI1
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