resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US", "UK"]
  labels   = ["terraform", "dev"]
}

resource "packetfabric_cloud_router_connection_ipsec" "crc3" {
  provider                     = packetfabric
  description                  = "hello world"
  circuit_id                   = packetfabric_cloud_router.cr1.id
  pop                          = "SFO6"
  speed                        = "1Gbps"
  gateway_address              = "127.0.0.1"
  ike_version                  = 1
  phase1_authentication_method = "pre-shared-key"
  phase1_group                 = "group14"
  phase1_encryption_algo       = "3des-cbc"
  phase1_authentication_algo   = "sha-384"
  phase1_lifetime              = 10800
  phase2_pfs_group             = "group14"
  phase2_encryption_algo       = "3des-cbc"
  phase2_authentication_algo   = "hmac-sha-256-128"
  phase2_lifetime              = 28800
  shared_key                   = "superCoolKey"
  labels                       = ["terraform", "dev"]
}