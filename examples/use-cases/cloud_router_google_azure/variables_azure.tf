## Azure VARs
# https://docs.microsoft.com/en-us/azure/availability-zones/az-overview
variable "azure_region1" {
  type        = string
  description = "Azure region"
  default     = "East US" # East US, East US 2, West US 2, West US 3
}
variable "azure_vnet_cidr1" {
  type        = string
  description = "CIDR for the VNET"
  default     = "10.7.0.0/16"
}
variable "azure_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.7.1.0/24"
}
variable "azure_subnet_cidr2" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.7.2.0/24"
}