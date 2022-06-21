# General PacketFabic API VARs
variable "pf_api_key" {
  type = string
  default = "api-abcde-change-me"
  description = "PacketFabric platform API access key"
  sensitive = true
}

variable "pf_api_server" {
  type    = string
  default = "https://api.packetfabric.com"
  description = "PacketFabric API endpoint URL"
}

variable "pf_provider_source" {
  type = string
  default = packetfabric/packetfabric
}

variable "pf_account_uuid" {
  type = string
  default = "change me"
}

variable "pf_aws_account_id" {
  type = number
  default = 123456789
}



# PacketFabric Cloud-Router Parameter configurations

variable "pf_cr_asn" {
  type = number
  nullable = false
}

variable "pf_cr_scope" {
  type = string
  default = "private"
}

variable "pf_cr_name" {
  type = string
  nullable = false
  default = "PF New Router"
}

variable "pf_cr_capacity" {
  type = string
  default = "1Gps"
}
variable "pf_cr_regions" {
  type = list(string)
  default = ["US"]
}



# PacketFabric Cloud-Router-Connections Parameter configuration:

variable "pf_crc_maybe_nat" {
  type = bool
  default = false
}

variable "pf_crc_description" {
  type = string
  default = "PF Cloud-Router-Connection Description here"
}

variable "pf_crc_pop" {
  type = string
  default = "LAX1"
}

variable "pf_crc_zone" {
  type = string
  default = "a"
}

variable "pf_crc_is_public" {
  type = bool
  default = true
}

variable "pf_crc_speed" {
  type = string
  default = "50Mbps"
}

# PacketFabric Cloud-Router-BGP-Session Parameter configuration:

variable "pf_crbs_af" {
  type = string
  default = "v4"
}

variable "pf_crbs_mhttl" {
  type = number
  default = 1
}

variable "pf_crbs_rasn" {
  type = number
  default = 64515
}

variable "pf_crbs_orlonger" {
  type = bool
  default = false
}

variable "pf_crbs_remoteaddr" {
  type = string
  default = "169.254.253.5/30"
}
 variable "pf_crbs_l3addr" {
  type = string
  default = "169.254.253.6/30"
 }

 variable "pf_crbs_md5" {
  type = string
  default = "123456789abcdef123456789abcdef12"
 }

# PacketFabric Cloud-Router-BGP-Prefixes Parameter configuration:

variable "pf_crbp_pfx00" {
  type = string
  default = "169.254.253.4/30"
}
variable "pf_crbp_pfx00_type" {
  type = string
  default = "in"
}
variable "pf_crbp_pfx00_order" {
  type = number
  default = 0
}
variable "pf_crbp_pfx01" {
  type = string
  default = "169.254.253.4/30"
}
variable "pf_crbp_pfx01_type" {
  type = string
  default = "out"
}
variable "pf_crbp_pfx01_order" {
  type = number
  default = 0
}

# PacketFabric Cloud-Services-AWS-Dedicated Parameter configuration:
variable "pf_cs_aws_d_region" {
  type = string
  default = "us-west-2"
}
variable "pf_cs_aws_d_descr" {
  type = string
  default = "CloudService AWS Dedicated Description"
}
variable "pf_cs_aws_d_avzone" {
  type = string
  default = "A"
}
variable "pf_cs_aws_d_pop" {
  type = string
  default = "PDX1"
}
variable "pf_cs_aws_d_subterm" {
  type = number
  default = 1
}
variable "pf_cs_aws_d_srvclass" {
  type = string
  default = "metro"
}
variable "pf_cs_aws_d_autoneg" {
  type = bool
  default = false
}
variable "pf_cs_aws_d_speed" {
  type = string
  default = "10Gbps"
}
variable "pf_cs_aws_d_createlag" {
  type = bool
  default = true
}

# PacketFabric Cloud-Service-AWS-Hosted-Marketplace parameter configuration:
variable "pf_cs_aws_hm_descr" {
  type = string
  default = "CloudService AWS Hosted Marketplace Description"
}
variable "pf_cs_aws_hm_avzone" {
  type = string
  default = "A"
}
variable "pf_cs_aws_hm_pop" {
  type = string
  default = "PDX1"
}
variable "pf_cs_aws_hm_rid" {
  type = string
  default = "CLO-DWI-UDUC"
}
variable "pf_cs_aws_hm_market" {
  type = string
  default = "DAL"
}
variable "pf_cs_aws_hm_svcuuid" {
  type = string
  default = "0d5225a8-f000-47ee-b3ea-64888f6d5da9"
}
variable "pf_cs_aws_hm_speed" {
  type = string
  default = "10Gbps"
}

# PacketFabric interface Parameter configuration:
variable "pf_cs_interface_media" {
  type = string
  default = "LX"
}
variable "pf_cs_interface_descr" {
  type = string
  default = "CloudService AWS Dedicated Description"
}
variable "pf_cs_interface_avzone" {
  type = string
  default = "A"
}
variable "pf_cs_interface_pop" {
  type = string
  default = "SFO1"
}
variable "pf_cs_interface_subterm" {
  type = number
  default = 1
}
variable "pf_cs_interface_srvclass" {
  type = string
  default = "metro"
}
variable "pf_cs_interface_autoneg" {
  type = bool
  default = false
}
variable "pf_cs_interface_speed" {
  type = string
  default = "1Gbps"
}
variable "pf_cs_interface_nni" {
  type = bool
  default = false
}
