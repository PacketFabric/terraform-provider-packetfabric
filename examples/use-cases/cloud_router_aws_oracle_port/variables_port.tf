
# Port
variable "pf_port_pop1" {
  type    = string
  default = "NYC4"
}
variable "pf_port_avzone1" {
  type    = string
  default = "A" # login to the portal https://portal.packetfabric.com and start a workflow to create a port (but don't create it, just note the pop/zone info to use in Terraform)
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
