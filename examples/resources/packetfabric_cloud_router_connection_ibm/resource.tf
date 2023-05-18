resource "packetfabric_cloud_router" "cr1" {
  provider = packetfabric
  asn      = 4556
  name     = "hello world"
  capacity = "10Gbps"
  regions  = ["US"]
  labels   = ["terraform", "dev"]
}

resource "packetfabric_cloud_router_connection_ibm" "crc5" {
  provider    = packetfabric
  description = "hello world"
  circuit_id  = packetfabric_cloud_router.cr1.id
  ibm_bgp_asn = packetfabric_cloud_router.cr1.asn
  pop         = "DAL2"
  zone        = "A"
  maybe_nat   = false
  speed       = "1Gbps"
  labels      = ["terraform", "dev"]
}

resource "time_sleep" "wait_ibm_connection" {
  create_duration = "1m"
  depends_on = [packetfabric_cloud_router_connection_ibm.crc5]
}
data "ibm_dl_gateway" "current" {
  provider   = ibm
  name       = "hello world" # same as the PacketFabric IBM Cloud Router Connection description
  depends_on = [time_sleep.wait_ibm_connection]
}
data "ibm_resource_group" "existing_rg" {
  provider = ibm
  name     = "My Resource Group"
}

resource "ibm_dl_gateway_action" "confirmation" {
  provider       = ibm
  gateway        = data.ibm_dl_gateway.current.id
  resource_group = data.ibm_resource_group.existing_rg.id
  action         = "create_gateway_approve"
  global         = true
  metered        = true # If set true gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway

  provisioner "local-exec" {
    when    = destroy
    command = "sleep 30"
  }
}