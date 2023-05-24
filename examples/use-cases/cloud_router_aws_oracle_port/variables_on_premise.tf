## On-premise VARs

variable "on_premise_cidr1" { # used in PF BGP prefix
  type        = string
  description = "CIDR for Network 1"
  default     = "10.10.1.0/24" # do not use 172.17.0.1/16, internal network used for docker containers used in the demo VMs
}
