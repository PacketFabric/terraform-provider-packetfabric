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
  type    = string
  default = "secret"
  #sensitive = true
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
# Backbone Virtual Circuit Speed Burst
variable "pf_vc_circuit_id" {
  type    = string
  default = "PF-BC-RNO-CHI-1729807-PF"
}
variable "pf_vc_speed_burst" {
  type    = string
  default = "400Mbps"
}

# Point to Point
variable "pf_ptp_speed" {
  type    = string
  default = "1Gbps" # 1Gbps 10Gbps 40Gbps 100Gbps
}
variable "pf_ptp_subterm" {
  type    = number
  default = 1 # default 1 month
}
variable "pf_ptp_autoneg" {
  type    = bool
  default = false
}
variable "pf_ptp_media" {
  type    = string
  default = "LX"
}
variable "pf_ptp_pop1" {
  type    = string
  default = "SEA1"
}
variable "pf_ptp_zone1" {
  type    = string
  default = "C" # A
}
variable "pf_ptp_pop2" {
  type    = string
  default = "CHI1"
}
variable "pf_ptp_zone2" {
  type    = string
  default = "A"
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
  default   = "secret"
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
  default   = "secret"
}
variable "google_vlan_attachment_name" {
  sensitive = true
  default   = "vlan_attachment_name"
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

# Oracle Hosted Connection
variable "pf_cs_pop6" {
  type    = string
  default = "SFO6"
}
variable "pf_cs_zone6" {
  type = string
  default = "C"
}
variable "pf_cs_vlan6" {
  type    = number
  default = 107
}
variable "pf_cs_oracle_region" {
  type        = string
  default     = "us-ashburn-1"
  description = "Oracle Cloud region"
}
variable "pf_cs_oracle_vc_ocid" {
  type = string
  default   = "secret"
  sensitive = true
}

# Markeptlace
variable "pf_routing_id" {
  type    = string
  default = "PDB-ROJ-9Y0K"
}
variable "pf_market" {
  type    = string
  default = "ATL"
}
variable "pf_routing_id_ix" {
  type    = string
  default = "PDB-ROJ-9Y0K"
}
variable "pf_market_ix" {
  type    = string
  default = "ATL"
}
variable "pf_port_circuit_id_marketplace" {
  type    = string
  default = "PF-AP-WDC1-1726464"
}
variable "pf_asn_ix" {
  type     = number
  default  = 64545
  nullable = false
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
  default = "PDX2"
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
variable "pf_crc_pop2" {
  type    = string
  default = "LAX1"
}
variable "pf_crc_google_pairing_key" {
  type      = string
  default   = "secret"
  sensitive = true
}
variable "pf_crc_google_vlan_attachment_name" {
  type    = string
  default = "vlan_attachement_name"
}
variable "pf_crc_pop3" {
  type    = string
  default = "SFO6"
}
variable "pf_crc_azure_service_key" {
  type      = string
  default   = "secret"
  sensitive = true
}

# Cloud Router Connection IPsec
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
  default = "3des-cbc"
}
variable "pf_crc_phase1_authentication_algo" {
  type    = string
  default = "sha-384"
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
  default = "3des-cbc"
}
variable "pf_crc_phase2_authentication_algo" {
  type    = string
  default = "hmac-sha-256-128"
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

# Cloud Router Connection Port
variable "pf_crc_port_circuit_id" {
  type    = string
  default = "PF-AP-WDC1-1726464"
}
variable "pf_crc_vlan" {
  type    = number
  default = 170
}

# Cloud Router Connection IBM
variable "pf_crc_pop4" {
  type    = string
  default = "SFO1"
}
variable "pf_crc_zone4" {
  type    = string
  default = "c"
}
variable "pf_crc_ibm_bgp_asn" {
  type    = number
  default = 64536 # private (64512 to 65534)
}
variable "pf_crc_ibm_bgp_cer_cidr" {
  type    = string
  default = "169.254.248.41/30"
}
variable "pf_crc_ibm_bgp_ibm_cidr" {
  type    = string
  default = "169.254.248.42/30"
}

# Cloud Router Connection Oracle
variable "pf_crc_pop5" {
  type    = string
  default = "WDC02"
}
variable "pf_crc_zone5" {
  type    = string
  default = "c"
}
variable "pf_crc_oracle_region" {
  type        = string
  default     = "us-ashburn-1"
  description = "Oracle Cloud region"
}
variable "pf_crc_oracle_vc_ocid" {
  type      = string
  default   = "secret"
  sensitive = true
}