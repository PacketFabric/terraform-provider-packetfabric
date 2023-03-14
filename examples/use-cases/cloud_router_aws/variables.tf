## General VARs
variable "tag_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-aws"
}
variable "public_key" {
  type        = string
  description = "Public Key used to access demo EC2 instances."
  sensitive   = true
}