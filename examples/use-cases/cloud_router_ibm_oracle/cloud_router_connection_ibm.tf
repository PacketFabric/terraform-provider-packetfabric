# From the PacketFabric side: Create a Cloud Router connection.
resource "packetfabric_cloud_router_connection_ibm" "crc_ibm" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  labels      = var.pf_labels
  circuit_id  = packetfabric_cloud_router.cr.id
  ibm_bgp_asn = packetfabric_cloud_router.cr.asn
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  maybe_nat   = var.pf_crc_maybe_nat
  speed       = var.pf_crc_speed
  
  depends_on = [
    ibm_resource_group.resource_group_1
  ]
}

# From the IBM side: Accept the connection
resource "time_sleep" "wait_ibm_connection" {
  create_duration = "1m"
}
# Retrieve the Direct Connect connections in IBM
data "ibm_dl_gateway" "current" {
  provider   = ibm
  name       = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  depends_on = [time_sleep.wait_ibm_connection]
}

# Used in case you are using an existing resource group and you don't create a new one
# data "ibm_resource_group" "existing_rg" {
#   provider   = ibm
#   name       = "Packet Fabric"
# }

resource "ibm_dl_gateway_action" "confirmation" {
  provider = ibm
  gateway  = data.ibm_dl_gateway.current.id
  # resource_group = data.ibm_resource_group.existing_rg.id # used for existing resource group
  resource_group = ibm_resource_group.resource_group_1.id # used for new resource group
  action         = "create_gateway_approve"
  global         = true
  metered        = true # If set true gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway

  provisioner "local-exec" {
    when    = destroy
    command = "sleep 30"
  }
}
# output "ibm_dl_gateway_action" {
#   value = data.ibm_dl_gateway.current
# }

# data "ibm_dl_gateway" "after_approved" {
#   provider   = ibm
#   name       = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
#   depends_on = [ibm_dl_gateway_action.confirmation]
# }
# output "ibm_dl_gateway_after" {
#   value = data.ibm_dl_gateway.after_approved
# }

# From the IBM side: Set up VRF
# TBD

# From the IBM side: Add a virtual connection to your IBM virtual private cloud (VPC)
resource "ibm_dl_virtual_connection" "dl_gateway_vc" {
  provider   = ibm
  gateway    = data.ibm_dl_gateway.direct_link_gw.id
  name       = "${var.resource_name}-${random_pet.name.id}"
  type       = "vpc"
  network_id = ibm_is_vpc.vpc_1.id
}

# From the PacketFabric side: Configure BGP
resource "packetfabric_cloud_router_bgp_session" "crbs_ibm" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_ibm.crc_ibm.id
  address_family = var.pf_crbs_af
  remote_asn     = ibm_dl_gateway_action.confirmation.bgp_ibm_asn # IBM Direct Link ASN (13884) and cannot be modified
  orlonger       = var.pf_crbs_orlonger
  remote_address = ibm_dl_gateway_action.confirmation.bgp_base_cidr # IBM side
  l3_address     = ibm_dl_gateway_action.confirmation.bgp_cer_cidr  # PF side
  prefixes {
    prefix = var.oracle_subnet_cidr1
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = var.ibm_vpc_cidr1
    type   = "in" # Allowed Prefixes from Cloud
  }
}
# output "packetfabric_cloud_router_bgp_session_crbs_ibm" {
#   value = packetfabric_cloud_router_bgp_session.crbs_ibm
# }

# Create a Direct Link Gateway Virtual Connection
resource "ibm_dl_virtual_connection" "connect_gw_vpc"{
        gateway = data.ibm_dl_gateway.current.id
        name = "${var.resource_name}-${random_pet.name.id}"
        type = "vpc"
        network_id = ibm_is_vpc.vpc_1.resource_crn   
}

# Missing part: routing