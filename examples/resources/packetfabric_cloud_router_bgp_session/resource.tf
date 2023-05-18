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
  remote_address = "169.254.247.41/30" # AWS side
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
  remote_address = "169.254.247.41/30" # AWS side
  l3_address     = "169.254.247.42/30" # PF side
  nat {
    direction       = "input" # input=ingress output=egress
    nat_type        = "overload"
    pre_nat_sources = ["10.0.0.0/24"]
    pool_prefixes   = ["185.161.1.42/32"]
  }
  prefixes {
    prefix = "185.161.1.42/32"
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
  remote_address = "169.254.247.41/30" # AWS side
  l3_address     = "169.254.247.42/30" # PF side
  nat {
    nat_type = "inline_dnat"
    dnat_mappings {
      public_prefix  = "185.161.1.42/32" # Pre-translation IP prefix
      private_prefix = "10.1.1.123/32"   # Post-translation IP prefix
      # Note that the post-translation prefix must be equal to or included within the conditional IP prefix
      conditional_prefix = "10.1.1.0/24" # Conditional IP prefix
    }
  }
  prefixes {
    prefix = "185.161.1.42/32"
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = "10.1.1.0/24"
    type   = "in" # Allowed Prefixes from Cloud
  }
}