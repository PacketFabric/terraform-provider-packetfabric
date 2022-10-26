## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-mesh"
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

variable "pf_vc_longhaul_type" {
  type    = string
  default = "dedicated"
}
variable "pf_vc_speed" {
  type    = string
  default = "1Gbps"
}
variable "pf_vc_subterm" {
  type    = number
  default = 1 # default 1 month
}

variable "pf_port1" {
  type        = string
  description = "Port OpenColo-Santa Clara (SFO13)"
  default     = "PF-AP-SFO13-1748214"
}
variable "pf_port2" {
  type        = string
  description = "Port Coresite MI1 (MIA3)"
  default     = "PF-AP-MIA3-1715317"
}
variable "pf_port3" {
  type        = string
  description = "Port Switch Las Vegas 7 (LAS1)"
  default     = "PF-AP-LAS1-1715316"
}
variable "pf_port4" {
  type        = string
  description = "Port Coresite LA1 (LAX1)"
  default     = "PF-AP-LAX1-8620"
}
variable "pf_port5" {
  type        = string
  description = "Port 165 Halsey MMR (NYC1)"
  default     = "PF-AP-NYC1-824"
}
variable "pf_port6" {
  type        = string
  description = "Port Equinix DA2 (DAL1)"
  default     = "PF-AP-DAL1-823"
}

variable "pf_vc_vlan1" {
  type    = number
  default = 101
}
variable "pf_vc_vlan2" {
  type    = number
  default = 102
}
variable "pf_vc_vlan3" {
  type    = number
  default = 103
}
variable "pf_vc_vlan4" {
  type    = number
  default = 104
}
variable "pf_vc_vlan5" {
  type    = number
  default = 105
}
variable "pf_vc_vlan6" {
  type    = number
  default = 106
}
variable "pf_vc_vlan7" {
  type    = number
  default = 107
}
variable "pf_vc_vlan8" {
  type    = number
  default = 108
}
variable "pf_vc_vlan9" {
  type    = number
  default = 109
}
variable "pf_vc_vlan10" {
  type    = number
  default = 110
}
variable "pf_vc_vlan11" {
  type    = number
  default = 111
}
variable "pf_vc_vlan12" {
  type    = number
  default = 112
}
variable "pf_vc_vlan13" {
  type    = number
  default = 113
}
variable "pf_vc_vlan14" {
  type    = number
  default = 114
}
variable "pf_vc_vlan15" {
  type    = number
  default = 115
}
