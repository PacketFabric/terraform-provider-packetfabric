resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US"]
  labels   = ["terraform", "dev"]
}

resource "packetfabric_cloud_router_connection_aws" "crc1" {
  provider    = packetfabric
  circuit_id  = packetfabric_cloud_router.cr1.id
  maybe_nat   = false
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
  remote_asn     = 64535
  remote_address = "169.254.247.41/30"  # AWS side
  l3_address     = "169.254.247.42/30" # PF side
  prefixes {
    prefix = "10.1.1.0/24"
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = "10.1.2.0/24"
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = "172.16.1.0/24"
    type   = "in" # Allowed Prefixes from Cloud
  }
  prefixes {
    prefix = "172.16.2.0/24"
    type   = "in" # Allowed Prefixes from Cloud
  }
}

# Example maybe_nat set to true (source NAT)
resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc1.id
  remote_asn     = 64535
  remote_address = "169.254.247.41/30"  # AWS side
  l3_address     = "169.254.247.42/30" # PF side
  nat {
    direction       = "input" # or output
    nat_type        = "overload"
    pre_nat_sources = ["10.1.1.0/24", "10.1.2.0/24", "10.1.3.0/24"]
    pool_prefixes   = ["192.168.1.50/32", "192.168.1.51/32"]
  }
  prefixes {
    prefix = "172.16.1.0/24"
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = "10.1.1.0/24"
    type   = "in" # Allowed Prefixes from Cloud
  }
}

# Example maybe_dnat set to true (destination NAT)
resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc1.id
  remote_asn     = 64535
  remote_address = "169.254.247.41/30"  # AWS side
  l3_address     = "169.254.247.42/30" # PF side
  nat {
    nat_type = "inline_dnat"
    dnat_mappings {
      private_prefix = "10.1.1.0/24"
      public_prefix  = "1.2.3.4/32"
    }
    dnat_mappings {
      private_prefix     = "172.16.0.0/16"
      public_prefix      = "1.2.3.5/32"
      conditional_prefix = "172.16.1.0/24" # must be a subnet of private_prefix
    }
  }
  prefixes {
    prefix = "172.16.1.0/24"
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = "10.1.1.0/24"
    type   = "in" # Allowed Prefixes from Cloud
  }
}