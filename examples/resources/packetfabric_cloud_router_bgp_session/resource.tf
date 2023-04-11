resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = ["terraform", "dev"]
}

resource "packetfabric_cloud_router_connection_aws" "crc1" {
  provider    = packetfabric
  circuit_id  = packetfabric_cloud_router.cr1.id
  maybe_nat   = false
  maybe_dnat  = false
  description = "hello world"
  pop         = "PDX2"
  zone        = "A"
  is_public   = false
  speed       = "1Gbps"
  labels      = ["terraform", "dev"]
}

# Example maybe_nat set to false
resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc1.id
  address_family = "v4"
  multihop_ttl   = 1
  remote_asn     = 64535
  prefixes {
    prefix = var.pf_crbp_pfx00
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.pf_crbp_pfx01
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.pf_crbp_pfx02
    type   = "in" # Allowed Prefixes from Cloud
  }
  prefixes {
    prefix = var.pf_crbp_pfx03
    type   = "in" # Allowed Prefixes from Cloud
  }
}

# Example maybe_nat set to true (source NAT)
resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc1.id
  address_family = "v4"
  multihop_ttl   = 1
  remote_asn     = 64535
  nat {
    direction       = "input" # or output
    nat_type        = "overload"
    pre_nat_sources = ["10.1.1.0/24", "10.1.2.0/24", "10.1.3.0/24"]
    pool_prefixes   = ["192.168.1.50/32", "192.168.1.51/32"]
  }
  prefixes {
    prefix = var.pf_crbp_pfx00
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.pf_crbp_pfx01
    type   = "in" # Allowed Prefixes from Cloud
  }
}

# Example maybe_dnat set to true (destination NAT)
resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc1.id
  address_family = "v4"
  multihop_ttl   = 1
  remote_asn     = 64535
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
  }
  prefixes {
    prefix = var.pf_crbp_pfx01
    type   = "in" # Allowed Prefixes from Cloud
  }
}