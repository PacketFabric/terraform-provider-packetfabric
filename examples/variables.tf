## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo"
}

########################################
###### PORTS / CROSS CONNECT / VC
########################################

# Ports
variable "pf_port_media" {
  type    = string
  default = "LX" # LR
}
variable "pf_port_pop1" {
  type    = string
  default = "DEN2" # PDX1
}
variable "pf_port_avzone1" {
  type    = string
  default = "B" # check location /v2/locations/DEN2/port-availability
}
variable "pf_port_pop2" {
  type    = string
  default = "WDC1"
}
variable "pf_port_avzone2" {
  type    = string
  default = "E" # check location /v2/locations/WDC1/port-availability
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
  default = "1Gbps" # 100Mbps, 10Gbps, 5Gbps
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
  default = 40
}
variable "pf_vc_vlan2" {
  type    = number
  default = 50
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
  default = "A" # check location /v2/locations/SEA1/port-availability
}
variable "pf_ptp_pop2" {
  type    = string
  default = "CHI1"
}
variable "pf_ptp_zone2" {
  type    = string
  default = "A" # check location /v2/locations/CHI1/port-availability
}

########################################
###### HOSTED CLOUD CONNECTIONS
########################################

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
  default = "A" # check location /v2/locations/cloud?cloud_connection_type=hosted&cloud_provider=aws&pop=SFO6
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
  type    = string
  default = "A" # check location /v2/locations/cloud?cloud_connection_type=hosted&cloud_provider=oracle&pop=SFO6
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
  type      = string
  default   = "secret"
  sensitive = true
}

# IBM Hosted Connection
variable "pf_cs_pop7" {
  type    = string
  default = "SFO1"
}
variable "pf_cs_vlan7" {
  type    = number
  default = 108
}
variable "ibm_bgp_asn" {
  type    = string
  default = 64537 # private (64512 to 65534)
}

# Markeptlace
variable "pf_routing_id" {
  type    = string
  default = "PD-WUY-9VB0" # Demo A
}
variable "pf_market" {
  type    = string
  default = "HOU" # Demo A
}
variable "pf_market_port_circuit_id" {
  type        = string
  description = "Port Circuit ID used to provision a Marketplace request"
  default     = "PF-AP-HOU1-1751418" # Demo A
}
variable "pf_routing_id_ix" {
  type    = string
  default = "IXW-XRH-K2VX" # IX-Denver
}
variable "pf_market_ix" {
  type    = string
  default = "DEN" # IX-Denver
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
  default = "D" # check location /v2/locations/cloud?cloud_connection_type=dedicated&cloud_provider=aws&pop=NYC6
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
  default = "C" # check location /v2/locations/cloud?cloud_connection_type=dedicated&cloud_provider=google&pop=ATL1
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
  default = "E" # check location /v2/locations/cloud?cloud_connection_type=dedicated&cloud_provider=azure&pop=SFO6
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
  type        = number
  description = "The ASN of the cloud router"
  default     = 4556 # PacketFabric ASN
  nullable    = false
}
variable "pf_cr_capacity" {
  type        = string
  description = "The cloud router capacity"
  default     = "1Gbps" # 2Gbps
}
variable "pf_cr_regions" {
  type        = list(string)
  description = "The regions in which the Cloud Router connections will be located."
  default     = ["US"] # ["US"] or ["US", "UK"] or ["UK"]
}

# Cloud Router Connections Common
variable "pf_crc_maybe_nat" {
  type        = bool
  description = "Set this to true if you intend to use Source NAT on this connection"
  default     = false
}
variable "pf_crc_maybe_dnat" {
  type        = bool
  description = "Set this to true if you intend to use Destination NAT on this connection"
  default     = false
}
variable "pf_crc_is_public" {
  type        = bool
  description = "Whether PacketFabric should allocate a public IP address for this connection"
  default     = false
}


# Cloud Router BGP Session Common
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

# Cloud Router Connection AWS
variable "pf_crc_speed" {
  type        = string
  description = "The speed of the new connection"
  default     = "50Mbps"
}
variable "pf_crc_pop1" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "PDX2"
}
variable "pf_crc_zone1" {
  type    = string
  default = "A" # check location /v2/locations/cloud?cloud_connection_type=hosted&has_cloud_router: true=true&cloud_provider=aws&pop=PDX2
}

# Cloud Router BGP Session AWS
variable "aws_side_asn1" {
  type        = number
  default     = 64535 # private (64512 to 65534)
  description = "AWS Side ASN"
}
variable "aws_remote_address" {
  type        = string
  description = "The cloud-side router peer IP."
  default     = "169.254.52.1/29"
}
variable "aws_l3_address" {
  type        = string
  description = "The L3 address of this instance."
  default     = "169.254.52.2/29"
}

# Cloud Router Connection Google
variable "pf_crc_pop2" {
  type        = string
  description = "The POP in which you want to provision the connection"
  default     = "LAX1"
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

# Cloud Router Connection Azure
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
  default     = "3des-cbc"
}
variable "pf_crc_phase1_authentication_algo" {
  type        = string
  description = "The authentication algorithm to use during phase 1"
  default     = "sha-384"
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
  default     = "3des-cbc"
}
variable "pf_crc_phase2_authentication_algo" {
  type        = string
  description = "The authentication algorithm to use during phase 2"
  default     = "hmac-sha-256-128" # not needed to set pf_crc_phase2_authentication_algo if pf_crc_phase2_encryption_algo = aes-256-gcm
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

# Cloud Router BGP Session IPsec
variable "vpn_side_asn3" {
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

# Cloud Router Connection Port
variable "pf_crc_port_circuit_id" {
  type        = string
  description = "Port Circuit ID used as a source port to create a Port Cloud Router Connection"
  default     = "PF-AP-WDC1-1726464"
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
  default = "B"
}
variable "pf_crc_ibm_bgp_asn" {
  type    = number
  default = 64536 # private (64512 to 65534)
}

# Cloud Router Connection Oracle
variable "pf_crc_pop5" {
  type    = string
  default = "SFO1"
}
variable "pf_crc_zone5" {
  type    = string
  default = "A"
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