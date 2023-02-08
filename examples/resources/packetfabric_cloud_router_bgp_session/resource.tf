resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = var.pf_cr_asn
  name     = var.pf_cr_name
  capacity = var.pf_cr_capacity
  regions  = var.pf_cr_regions
}

resource "packetfabric_cloud_router_connection_aws" "crc1" {
  provider    = packetfabric
  circuit_id  = packetfabric_cloud_router.cr1.id
  maybe_nat   = var.pf_crc_maybe_nat
  maybe_dnat  = var.pf_crc_maybe_dnat
  description = var.pf_crc_description
  pop         = var.pf_crc_pop
  zone        = var.pf_crc_zone
  is_public   = var.pf_crc_is_public
  speed       = var.pf_crc_speed
}

# Example maybe_nat set to false
resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.pf_crbs_rasn
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.pf_crbs_remoteaddr
  l3_address     = var.pf_crbs_l3addr
  md5            = var.pf_crbs_md5
  prefixes {
    prefix = var.pf_crbp_pfx00
    type   = "out" # Allowed Prefixes to Cloud
    order  = var.pf_crbp_pfx00_order
  }
  prefixes {
    prefix = var.pf_crbp_pfx01
    type   = "out" # Allowed Prefixes to Cloud
    order  = var.pf_crbp_pfx01_order
  }
  prefixes {
    prefix = var.pf_crbp_pfx02
    type   = "in" # Allowed Prefixes from Cloud
    order  = var.pf_crbp_pfx02_order
  }
  prefixes {
    prefix = var.pf_crbp_pfx03
    type   = "in" # Allowed Prefixes from Cloud
    order  = var.pf_crbp_pfx03_order
  }
}

# Example maybe_nat set to true (source NAT)
resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.pf_crbs_rasn
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.pf_crbs_remoteaddr
  l3_address     = var.pf_crbs_l3addr
  md5            = var.pf_crbs_md5
  nat {
    direction       = "input" # or output
    nat_type        = "overload"
    pre_nat_sources = var.pf_crbs_pre_nat_sources # e.g. ["10.1.1.0/24", "10.1.2.0/24", "10.1.3.0/24"]
    pool_prefixes   = var.pf_crbs_pool_prefixes   # e.g. ["192.168.1.50/32", "192.168.1.51/32"]
  }
  prefixes {
    prefix = var.pf_crbp_pfx00
    type   = "out" # Allowed Prefixes to Cloud
    order  = var.pf_crbp_pfx00_order
  }
  prefixes {
    prefix = var.pf_crbp_pfx01
    type   = "in"  # Allowed Prefixes from Cloud
    order  = var.pf_crbp_pfx01_order
  }
}

# Example maybe_dnat set to true (destination NAT)
resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.pf_crbs_rasn
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.pf_crbs_remoteaddr
  l3_address     = var.pf_crbs_l3addr
  md5            = var.pf_crbs_md5
  nat {
    nat_type = "inline_dnat"
    dnat_mappings {
      private_prefix = var.pf_crbs_private_prefix1
      public_prefix  = var.pf_crbs_public_prefix1
    }
    dnat_mappings {
      private_prefix     = var.pf_crbs_private_prefix2
      public_prefix      = var.pf_crbs_public_prefix2
      conditional_prefix = var.pf_crbs_conditional_prefix2 # must be a subnet of private_prefix
    }
  }
  prefixes {
    prefix = var.pf_crbp_pfx00
    type   = "out" # Allowed Prefixes to Cloud
    order  = var.pf_crbp_pfx00_order
  }
  prefixes {
    prefix = var.pf_crbp_pfx01
    type   = "in"  # Allowed Prefixes from Cloud
    order  = var.pf_crbp_pfx01_order
  }
}