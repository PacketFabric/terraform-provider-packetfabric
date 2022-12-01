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

terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.5.0"
    }
  }
}

resource "packetfabric_cloud_router" "cr" {
  provider     = packetfabric
  asn          = var.asn
  name         = var.name
  capacity     = var.capacity
  regions      = var.regions
  timeouts {
    create = "20m"
  }
}

resource "packetfabric_cloud_router_connection_aws" "aws" {
  for_each = var.aws_connections
  provider = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id
  description = "aws-${each.key}"
  pop = each.value.pop
  speed = each.value.speed
  zone = each.value.zone
  timeouts {
    create = "20m"
  }
}

resource "packetfabric_cloud_router_connection_google" "gcp" {
  for_each = var.gcp_connections
  provider                    = packetfabric
  circuit_id                  = packetfabric_cloud_router.cr.id
  google_pairing_key          = each.value.pairing_key
  google_vlan_attachment_name = each.value.vlan_attachment_name
  pop                         = each.value.pop
  speed                       = each.value.speed
  description                 = "gcp-${each.key}"
  timeouts {
    create = "20m"
  }
}


resource "packetfabric_cloud_router_bgp_session" "gcp" {
  for_each = var.gcp_bgp_sessions
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_google.gcp[each.key].id
  address_family = "v4"
  remote_asn     = each.value.asn
  remote_address = each.value.remote_address
  l3_address     = each.value.l3_address
  bfd_interval   = 1000
  bfd_multiplier = 5
  multihop_ttl   = 1
  disabled       = each.value.disabled

  dynamic "prefixes" {
    for_each   = { for i,v in var.gcp_inbound : i => v if contains(v.env, each.key)}
    content {
      prefix     = prefixes.value.prefix
      type       = "in"
      match_type = prefixes.value.match_type
      order      = prefixes.key
    }
  }

  dynamic "prefixes" {
    for_each   = { for i,v in var.gcp_outbound : i => v }
    content {
      prefix     = prefixes.value.prefix
      type       = "out"
      match_type = prefixes.value.match_type
      med        = lookup(prefixes.value, "med", 0)
      order      = prefixes.key
    }
  }
}


resource "packetfabric_cloud_router_bgp_session" "aws" {
  for_each = var.aws_bgp_sessions
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_aws.aws[each.key].id
  address_family = "v4"
  remote_asn     = each.value.asn
  remote_address = each.value.remote_address
  l3_address     = each.value.l3_address
  bfd_interval   = 300
  bfd_multiplier = 3
  multihop_ttl   = 1
  md5            = each.value.md5
  disabled       = each.value.disabled

  dynamic "prefixes" {
    for_each   = { for i,v in var.aws_inbound : i => v }
    content {
      prefix     = prefixes.value.prefix
      type       = "in"
      match_type = prefixes.value.match_type
      order      = prefixes.key
    }
  }

  dynamic "prefixes" {
    for_each   = { for i,v in var.aws_outbound : i => v }
    content {
      prefix     = prefixes.value.prefix
      type       = "out"
      match_type = prefixes.value.match_type
      order      = prefixes.key
    }
  }
}
