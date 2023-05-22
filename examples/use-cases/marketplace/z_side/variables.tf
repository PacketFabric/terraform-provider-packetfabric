## General VARs
# Must follow ^(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?)$
# Any lowercase ASCII letter or digit, and possibly hyphen, which should start with a letter and end with a letter or digit, 
# and have at most 63 characters (1 for the starting letter + up to 61 characters in the middle + 1 for the ending letter/digit).
variable "resource_name" {
  type        = string
  description = "Used to name all resources created in this example"
  default     = "demo-pf-z-side"
}
variable "pf_labels" {
  type        = list(string)
  description = "A list of labels to be applied to PacketFabric resources. These labels will be visible in the PacketFabric Portal and can be searched for easier resource identification."
  default     = ["terraform"] # Example: ["terraform", "dev"]
}

## PacketFabic VARs

variable "pf_z_side_port_id" {
  type    = string
  default = "PF-AP-NYC10-1739866"
}
variable "pf_z_side_vc_vlan2" {
  type    = number
  default = 50
}
variable "pf_a_side_vc_request_uuid" {
  type        = string
  description = "Update with the A side VC request UUID (use the id in the response of the packetfabric_backbone_virtual_circuit_marketplace resource)"
  default     = "f9654a38-0722-4cc7-9aa5-7ad9af691fef"
}
