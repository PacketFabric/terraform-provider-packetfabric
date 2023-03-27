## General VARs
variable "tag_name" {
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
