## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf"
}

## PacketFabic VARs
variable "pf_api_key" {
  type        = string
  description = "PacketFabric platform API access key"
  sensitive   = true
  default     = "secret"
}
variable "pf_account_uuid" {
  type        = string
  description = "The UUID for the billing account (Find it under Billing > Accounts in the Portal)"
  default     = "34ff9ffb-9bbf-43b1-9cf8-6c8e62370597"
}
variable "pf_api_server" {
  type        = string
  description = "PacketFabric API endpoint URL"
  default     = "https://api.packetfabric.com"
}

# Port
variable "pf_port_pop1" {
  type    = string
  default = "PDX1"
}
variable "pf_port_avzone1" {
  type    = string
  default = "A"
}
variable "pf_port_pop2" {
  type    = string
  default = "NYC4"
}
variable "pf_port_avzone2" {
  type    = string
  default = "A"
}
variable "pf_port_media" {
  type    = string
  default = "LX"
}
variable "pf_port_subterm" {
  type    = number
  default = 1 # default 1 month
}
variable "pf_port_autoneg" {
  type    = bool
  default = false
}
variable "pf_port_speed" {
  type    = string
  default = "1Gbps"
}
variable "pf_port_nni" {
  type    = bool
  default = false
}

# Cross connect
variable "pf_document_uuid1" {
  type    = string
  default = "1d2fb159-b40e-4eda-8f63-1191a80a023e" # use API /v2/documents to get UUID
}
variable "pf_document_uuid2" {
  type    = string
  default = "1d2fb159-b40e-4eda-8f63-1191a80a023e"
}

# Virtual Circuit
variable "pf_vc_vlan1" {
  type    = number
  default = 145
}
variable "pf_vc_vlan2" {
  type    = number
  default = 146
}
variable "pf_vc_longhaul_type" {
  type    = string
  default = "dedicated"
}
variable "pf_vc_speed" {
  type    = string
  default = "200Mbps"
}
variable "pf_vc_subterm" {
  type    = number
  default = 1 # default 1 month
}
