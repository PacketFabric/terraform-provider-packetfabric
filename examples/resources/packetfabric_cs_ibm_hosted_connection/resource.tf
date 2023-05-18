resource "packetfabric_cs_ibm_hosted_connection" "cs_conn1_hosted_ibm" {
  provider    = packetfabric
  ibm_bgp_asn = 64536
  description = "hello world"
  pop         = "WDC1"
  port        = packetfabric_port.port_1.id
  vlan        = 102
  speed       = "10Gbps"
  labels      = ["terraform", "dev"]
}

resource "time_sleep" "wait_ibm_connection" {
  create_duration = "60s"
}
data "ibm_dl_gateway" "current" {
  provider   = ibm
  name       = "hello world" # same as the PacketFabric IBM Hosted Cloud description
  depends_on = [time_sleep.wait_ibm_connection]
}
data "ibm_resource_group" "existing_rg" {
  provider = ibm
  name     = "My Resource Group"
}

resource "ibm_dl_gateway_action" "confirmation" {
  provider                    = ibm
  gateway                     = data.ibm_dl_gateway.current.id
  resource_group              = data.ibm_resource_group.existing_rg.id
  action                      = "create_gateway_approve"
  global                      = true
  metered                     = true # If set true gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway
  bgp_asn                     = 64536
  default_export_route_filter = "permit"
  default_import_route_filter = "permit"
  speed_mbps                  = 10000 # must match PacketFabric speed

  provisioner "local-exec" {
    when    = destroy
    command = "sleep 30"
  }
}