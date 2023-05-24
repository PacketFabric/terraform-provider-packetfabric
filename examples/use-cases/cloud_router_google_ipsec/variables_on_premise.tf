## IPsec VAR
variable "ipsec_subnet_cidr2" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.6.1.0/24" # do not use 172.17.0.1/16, internal network used for docker containers used in the demo VMs
}