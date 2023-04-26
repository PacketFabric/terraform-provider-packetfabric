# This block specifies that the PacketFabric provider is required and its version should be greater than or equal to 0.7.0
terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.4.0"
    }
  }
}

# This block creates a resource of type "packetfabric_cloud_router" with the provider set to "packetfabric"
# The asn, name, capacity and regions are passed in as variables
resource "packetfabric_cloud_router" "cr" {
  provider = packetfabric
  name     = var.name
  asn      = var.asn
  capacity = var.capacity
  regions  = var.regions
  labels   = var.labels
}

# This block creates a data source of type "packetfabric_locations_cloud" and is used to retrieve the zone information for a given pop.
# The data source is created for each item in the "aws_connections" variable.
data "packetfabric_locations_cloud" "locations_pop_zones_aws" {
  count                 = var.module_enabled ? 1 : 0
  provider              = packetfabric
  cloud_provider        = "aws"
  cloud_connection_type = "hosted"
  has_cloud_router      = true
  pop                   = each.value.pop
}

# This block creates a connection to an AWS Virtual Private Cloud (VPC) via the PacketFabric cloud router.
resource "packetfabric_cloud_router_connection_aws" "aws" {
  for_each    = var.aws_connections
  provider    = packetfabric
  circuit_id  = packetfabric_cloud_router.cr.id
  description = "aws-${each.key}-${lower(each.value.pop)}"
  labels      = var.labels
  pop         = each.value.pop
  speed       = each.value.speed
  zone        = data.packetfabric_locations_cloud.locations_pop_zones_aws[each.key].cloud_locations[0].zones[0]
  lifecycle {
    ignore_changes = [
      zone,
    ]
  }
}

# This block creates a connection to a GCP Virtual Private Cloud (VPC) via the PacketFabric cloud router.
resource "packetfabric_cloud_router_connection_google" "gcp" {
  for_each                    = var.gcp_connections
  provider                    = packetfabric
  circuit_id                  = packetfabric_cloud_router.cr.id
  description                 = "gcp-${each.key}-${lower(each.value.pop)}"
  labels                      = var.labels
  google_pairing_key          = each.value.pairing_key
  google_vlan_attachment_name = "gcp-${each.value.vlan_attachment_name}-${lower(each.value.pop)}"
  pop                         = each.value.pop
  speed                       = each.value.speed
}

# This block creates a Border Gateway Protocol (BGP) session for the AWS connection.
resource "packetfabric_cloud_router_bgp_session" "aws" {
  for_each       = var.aws_bgp_sessions
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_aws.aws[each.key].id
  address_family = "v4"
  remote_asn     = each.value.remote_asn
  remote_address = each.value.remote_address
  l3_address     = each.value.l3_address
  md5            = each.value.md5
  disabled       = each.value.disabled

  # This block creates a dynamic block of in and out prefixes for the BGP session.
  dynamic "prefixes" {
    for_each = { for i, v in var.gcp_outbound[each.key] : i => v }
    content {
      prefix     = prefixes.value.prefix
      type       = "in"
      match_type = prefixes.value.match_type
    }
  }

  dynamic "prefixes" {
    for_each = { for i, v in var.aws_outbound[each.key] : i => v }
    content {
      prefix     = prefixes.value.prefix
      type       = "out"
      match_type = prefixes.value.match_type
    }
  }
}

# This block creates a Border Gateway Protocol (BGP) session for the GCP connection.
resource "packetfabric_cloud_router_bgp_session" "gcp" {
  for_each       = var.gcp_bgp_sessions
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_google.gcp[each.key].id
  address_family = "v4"
  remote_asn     = each.value.remote_asn
  remote_address = each.value.remote_address
  l3_address     = each.value.l3_address
  disabled       = each.value.disabled

  # This block creates a dynamic block of in and out prefixes for the BGP session.
  dynamic "prefixes" {
    for_each = { for i, v in var.aws_outbound[each.key] : i => v }
    content {
      prefix     = prefixes.value.prefix
      type       = "in"
      match_type = prefixes.value.match_type
    }
  }

  dynamic "prefixes" {
    for_each = { for i, v in var.gcp_outbound[each.key] : i => v }
    content {
      prefix     = prefixes.value.prefix
      type       = "out"
      match_type = prefixes.value.match_type
      med        = lookup(prefixes.value, "med", 0)
    }
  }
}
