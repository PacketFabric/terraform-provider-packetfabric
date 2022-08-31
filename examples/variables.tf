## General VARs
variable "tag_name" {
  default = "demo"
}

## PacketFabic VARs
variable "pf_api_key" {
  type        = string
  description = "PacketFabric platform API access key"
  default     = "secret"
  sensitive   = true
}
variable "pf_account_uuid" {
  type      = string
  default   = "secret"
  sensitive = true
}
variable "pf_api_server" {
  type        = string
  default     = "https://api.packetfabric.com"
  description = "PacketFabric API endpoint URL"
}

########################################
###### PORTS / CROSS CONNECT / VC
########################################

# Ports
variable "pf_port_media" {
  type    = string
  default = "LR" # LX
}
variable "pf_port_pop1" {
  type    = string
  default = "DEN2" # PDX1
}
variable "pf_port_avzone1" {
  type    = string
  default = "A" # A
}
variable "pf_port_pop2" {
  type    = string
  default = "WDC1"
}
variable "pf_port_avzone2" {
  type    = string
  default = "E"
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
  default = "10Gbps" # 100Mbps, 10Gbps, 5Gbps
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
  default = 4
}
variable "pf_vc_vlan2" {
  type    = number
  default = 5
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

########################################
###### HOSTED CLOUD CONNECTIONS
########################################
# Demo Port used for all
variable "pf_demo_port" {
  type    = string
  default = "PF-AP-WDC1-1726464"
}

# Azure Hosted Connection
variable "azure_service_key" {
  sensitive = true
  default   = "123456-abcdef-123456-abcdef-123456"
}
variable "pf_cs_src_svlan" {
  type    = number
  default = 100
}
variable "pf_cs_vlan_private" {
  type    = number
  default = 166
}
variable "pf_cs_vlan_microsoft" {
  type    = number
  default = 167
}

# GCP Hosted Connection
variable "google_pairing_key" {
  sensitive = true
  default   = "123456-abcdef-123456-abcdef-123456/us-west1/any"
}
variable "google_vlan_attachment_name" {
  sensitive = true
  default   = "demo-darling-albacore"
}
variable "pf_cs_pop1" {
  type    = string
  default = "SFO1"
}
variable "pf_cs_speed1" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_vlan1" {
  type    = number
  default = 106
}

# AWS Hosted Connection
variable "pf_cs_pop2" {
  type    = string
  default = "SFO6"
}
variable "pf_cs_zone2" {
  type    = string
  default = "A"
}
variable "pf_cs_speed2" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_vlan2" {
  type    = number
  default = 107
}

# Markeptlace
variable "routing_id" {
  type    = string
  default = "PDB-ROJ-9Y0K" # DEMO B
}
variable "market" {
  type    = string
  default = "ATL" # DEMO B
}
variable "port_circuit_id_marketplace" {
  type    = string
  default = "PF-AP-ATL1-1744189" # DEMO B in the market
}

########################################
###### DEDICATED CLOUD CONNECTIONS
########################################

# AWS Dedicated Connection
variable "pf_cs_pop3" {
  type    = string
  default = "NYC6"
}
variable "pf_cs_zone3" {
  type    = string
  default = "D" #A
}
variable "pf_cs_speed3" {
  type    = string
  default = "10Gbps"
}
variable "aws_region3" {
  type        = string
  description = "AWS region"
  default     = "us-east-1"
}

# Google Dedicated Connection
variable "pf_cs_pop4" {
  type    = string
  default = "ATL1"
}
variable "pf_cs_zone4" {
  type    = string
  default = "C"
}
variable "pf_cs_speed4" {
  type    = string
  default = "10Gbps"
}

# Azure Dedicated Connection
variable "pf_cs_pop5" {
  type    = string
  default = "DEN1"
}
variable "pf_cs_zone5" {
  type    = string
  default = "E"
}
variable "pf_cs_speed5" {
  type    = string
  default = "10Gbps"
}
variable "encapsulation" {
  type    = string
  default = "dot1q"
}
variable "port_category" {
  type    = string
  default = "primary"
}

# Dedicated all clouds
variable "pf_cs_srvclass" {
  type    = string
  default = "longhaul" # longhaul or metro
}
variable "pf_cs_autoneg" {
  type    = bool
  default = false
}
variable "should_create_lag" {
  type    = bool
  default = false
}
variable "pf_cs_subterm" {
  type    = number
  default = 1 # default 1 month
}


########################################
###### CLOUD ROUTER
########################################
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
  default = "1Gbps" # 2Gbps
}
variable "pf_cr_regions" {
  type    = list(string)
  default = ["US"] # ["US"] or ["US", "UK"] or ["UK"]
}
variable "pf_aws_account_id" {
  type    = string
  default = "123456789"
}
variable "pf_crc_speed" {
  type    = string
  default = "50Mbps"
}
variable "pf_crc_pop1" {
  type    = string
  default = "PDX2" # PDX2/a LAX1/c SF06/a LON1/a
}
variable "pf_crc_zone1" {
  type    = string
  default = "a"
}
variable "pf_crc_maybe_nat" {
  type    = bool
  default = false
}
variable "pf_crc_is_public" {
  type    = bool
  default = false
}
