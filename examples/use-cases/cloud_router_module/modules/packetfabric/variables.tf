variable "asn" {
  type = number
}

variable "name" {
  type = string
}

variable "capacity" {
  type = string
}

variable "regions" {
  type = list(string)
}

variable "labels" {
  type = list(string)
}

variable "aws_connections" {
  type = map(map(string))
}

variable "aws_bgp_sessions" {
  type = map(map(string))
}

variable "aws_outbound" {
  type = map(list(any))
}

variable "gcp_connections" {
  type = map(map(string))
}

variable "gcp_bgp_sessions" {
  type = map(map(string))
}

variable "gcp_outbound" {
  type = map(list(any))
}


