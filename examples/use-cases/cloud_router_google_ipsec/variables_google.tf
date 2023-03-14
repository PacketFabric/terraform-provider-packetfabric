## Google VARs
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
variable "google_subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.5.1.0/24"
}
variable "public_key" {
  type        = string
  description = "Public Key used to access demo Virtual Machines."
  sensitive   = true
}
