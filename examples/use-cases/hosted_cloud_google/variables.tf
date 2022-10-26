## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-google"
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
# GCP Hosted Connection
variable "pf_port_circuit_id" {
  type    = string
  default = "PF-AP-WDC1-1726464"
}
variable "pf_cs_pop1" {
  type    = string
  default = "SFO1"
}
variable "pf_cs_speed" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_vlan1" {
  type    = number
  default = 105
}

# GCP VARs
variable "gcp_project_id" {
  type = string
  # sensitive   = true
  description = "Google Cloud project ID"
}

variable "gcp_credentials_path" {
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
variable "subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.5.1.0/24"
}
# You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
variable "gcp_side_asn1" {
  type        = number
  default     = 16550
  description = "Google Cloud ASN"
}