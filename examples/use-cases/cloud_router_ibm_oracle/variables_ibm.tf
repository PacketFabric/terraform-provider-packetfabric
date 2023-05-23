## IBM VARs
variable "ibm_resource_group" {
  type        = string
  default     = "My Resource Group"
  description = "IBM Resource Group"
}
variable "ibm_region1" {
  type        = string
  default     = "us-south"
  description = "IBM Cloud region"
}
variable "ibm_region1_zone1" {
  type        = string
  description = "IBM Availability Zone"
  default     = "us-south-1"
}
variable "ibm_vpc_cidr1" {
  type        = string
  description = "CIDR for the VPC"
  default     = "10.8.0.0/16" # do not use 172.17.0.1/16, internal network used for docker containers used in the demo VMs
}
variable "ibm_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.8.1.0/24"
}