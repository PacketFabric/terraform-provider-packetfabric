## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-z-side"
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


variable "pf_z_side_port_id" {
  type    = string
  default = "PF-AP-LAX-2741214"
}
variable "pf_z_side_vc_vlan2" {
  type    = number
  default = 50
}
variable "pf_a_side_vc_request_uuid" {
  type    = string
  description = "Update with the A side VC request UUID"
  default = "dedi1bf234df-7a9d-4b4e-8d3a-3c245d6049bfcated"
}
