## General VARs
variable "resource_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-aws"
}
variable "pf_labels" {
  type        = list(string)
  description = "A list of labels to be applied to PacketFabric resources. These labels will be visible in the PacketFabric Portal and can be searched for easier resource identification."
  default     = ["terraform"] # Example: ["terraform", "dev"]
}
variable "public_key" {
  type        = string
  description = "Public Key used to access demo EC2 instances."
  sensitive   = true
}