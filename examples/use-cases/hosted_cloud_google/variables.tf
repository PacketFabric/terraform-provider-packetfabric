## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-google"
}

## PacketFabic VARs
# Port
variable "pf_port_pop1" {
  type    = string
  default = "PDX1"
}
variable "pf_port_avzone1" {
  type    = string
  default = "A" # check availability /v2/locations/PDX1/port-availability or login to the portal https://portal.packetfabric.com and start a workflow to create a port (but don't create it, just note the pop/zone info to use in Terraform)
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

# GCP Hosted Connection
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