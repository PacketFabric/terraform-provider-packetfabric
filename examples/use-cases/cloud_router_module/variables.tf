# declare variables in the root module
variable "cloud_routers" {
  type = map(any)
}

variable "aws_connections" {
  type = map(any)
}

variable "gcp_connections" {
  type = map(any)
}

variable "gcp_bgp_sessions" {
  type = map(any)
}

variable "aws_bgp_sessions" {
  type = map(any)
}

variable "gcp_inbound" {
  type = list(any)
}

variable "gcp_outbound" {
  type = list(any)
}

variable "aws_inbound" {
  type = list(any)
}

variable "aws_outbound" {
  type = list(any)
}