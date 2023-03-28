##############################################################
# THIS TF FILE IS USED TO TEST PACKETFABRIC TERRAFORM PROVIDER
# LOOK FOR SPECIFIC USE CASES UNDER THE USE-CASES FOLDER
##############################################################

terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.2.0"
    }
  }
}

provider "packetfabric" {}

# Create random name to use to name objects
resource "random_pet" "name" {}

#######################################
##### Ports/Interfaces
######################################

# Get the zone from the pop automatically
data "packetfabric_locations_port_availability" "port_availabilty_pop1" {
  provider = packetfabric
  pop      = var.pf_port_pop1
}
output "packetfabric_locations_port_availability_pop1" {
  value = data.packetfabric_locations_port_availability.port_availabilty_pop1
}
locals {
  zones_pop1 = toset([for each in data.packetfabric_locations_port_availability.port_availabilty_pop1.ports_available[*] : each.zone if each.media == var.pf_port_media])
}
output "packetfabric_locations_port_availability_pop1_single_zone" {
  value = tolist(local.zones_pop1)[0]
}

# Create a PacketFabric Ports
resource "packetfabric_port" "port_1a" {
  provider          = packetfabric
  autoneg           = var.pf_port_autoneg
  description       = "${var.resource_name}-${random_pet.name.id}-a"
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = tolist(local.zones_pop1)[0] # var.pf_port_avzone1
  labels            = var.pf_labels
  po_number         = var.pf_po_number
  lifecycle {
    ignore_changes = [
      # Ignore changes to zone because zone cannot be modified 
      # after the port is created but can change as we get it 
      # from packetfabric_locations_port_availability data source
      zone,
    ]
  }
}
output "packetfabric_port_1a" {
  value = packetfabric_port.port_1a
}

## 2nd port in the same location same zone to create a LAG
resource "packetfabric_port" "port_1b" {
  provider          = packetfabric
  autoneg           = var.pf_port_autoneg
  description       = "${var.resource_name}-${random_pet.name.id}-b"
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = tolist(local.zones_pop1)[0] # var.pf_port_avzone1
  labels            = var.pf_labels
  po_number         = var.pf_po_number
  lifecycle {
    ignore_changes = [
      # Ignore changes to zone because zone cannot be modified 
      # after the port is created but can change as we get it 
      # from packetfabric_locations_port_availability data source
      zone,
    ]
  }
}
output "packetfabric_port_1b" {
  value = packetfabric_port.port_1b
}

resource "packetfabric_link_aggregation_group" "lag_1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  interval    = "fast" # or slow
  members     = [packetfabric_port.port_1a.id, packetfabric_port.port_1b.id]
  #members = [packetfabric_port.port_1a.id]
  pop       = var.pf_port_pop1
  labels    = var.pf_labels
  po_number = var.pf_po_number
}

data "packetfabric_link_aggregation_group" "lag_1" {
  provider       = packetfabric
  lag_circuit_id = packetfabric_link_aggregation_group.lag_1.id
}
output "packetfabric_link_aggregation_group" {
  value = data.packetfabric_link_aggregation_group.lag_1
}

resource "packetfabric_port" "port_2" {
  provider          = packetfabric
  enabled           = true # set to false to disable the port
  autoneg           = var.pf_port_autoneg
  description       = "${var.resource_name}-${random_pet.name.id}"
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop2
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone2
  labels            = var.pf_labels
  po_number         = var.pf_po_number
}
output "packetfabric_port_2" {
  value = packetfabric_port.port_2
}

data "packetfabric_port" "ports_all" {
  provider   = packetfabric
  depends_on = [packetfabric_port.port_2]
}
locals {
  port_2_details = toset([for each in data.packetfabric_port.ports_all.interfaces[*] : each if each.port_circuit_id == packetfabric_port.port_2.id])
}
output "packetfabric_port_2_details" {
  value = local.port_2_details
}

data "packetfabric_port_router_logs" "port_1a_logs" {
  provider        = packetfabric
  port_circuit_id = packetfabric_port.port_1a.id
  time_from       = "2022-11-30 00:00:00"
  time_to         = "2022-12-01 00:00:00"
  depends_on      = [packetfabric_port.port_1]
}
output "packetfabric_port_router_logs" {
  value = data.packetfabric_port_router_logs.port_1a_logs
}

data "packetfabric_port_device_info" "port_1a_device_info" {
  provider        = packetfabric
  port_circuit_id = packetfabric_port.port_1a.id
  depends_on      = [packetfabric_port.port_1]
}
output "packetfabric_port_device_info" {
  value = data.packetfabric_port_device_info.port_1a_device_info
}

#######################################
##### Virtual Circuit (VC)
#######################################

# Create backbone Virtual Circuit
resource "packetfabric_backbone_virtual_circuit" "vc1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  epl         = false
  interface_a {
    port_circuit_id = packetfabric_port.port_1a.id
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  interface_z {
    port_circuit_id = packetfabric_port.port_2.id
    untagged        = false
    vlan            = var.pf_vc_vlan2
  }
  bandwidth {
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
  labels    = var.pf_labels
  po_number = var.pf_po_number
}
output "packetfabric_backbone_virtual_circuit" {
  value = packetfabric_backbone_virtual_circuit.vc1
}

# Show all Virtual Circuits
data "packetfabric_virtual_circuits" "all_vcs" {
  provider = packetfabric
}
output "packetfabric_virtual_circuits" {
  value = data.packetfabric_virtual_circuits.all_vcs
}

#######################################
##### Virtual Circuit Speed Burst
#######################################

resource "packetfabric_backbone_virtual_circuit_speed_burst" "boost" {
  provider      = packetfabric
  vc_circuit_id = var.pf_vc_circuit_id
  speed         = var.pf_vc_speed_burst
}
output "packetfabric_backbone_virtual_circuit_speed_burst" {
  value = packetfabric_backbone_virtual_circuit_speed_burst.boost
}

#######################################
##### Point to Point
#######################################

resource "packetfabric_point_to_point" "ptp1" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  speed             = var.pf_ptp_speed
  media             = var.pf_ptp_media
  subscription_term = var.pf_ptp_subterm
  endpoints {
    pop     = var.pf_ptp_pop1
    zone    = var.pf_ptp_zone1
    autoneg = var.pf_ptp_autoneg
  }
  endpoints {
    pop     = var.pf_ptp_pop2
    zone    = var.pf_ptp_zone2
    autoneg = var.pf_ptp_autoneg
  }
  labels    = var.pf_labels
  po_number = var.pf_po_number
}
output "packetfabric_point_to_point" {
  value = packetfabric_point_to_point.ptp1
}

#######################################
##### Cross Connect
#######################################

## Get the site filtering on the pop using packetfabric_locations

# List PacketFabric locations
data "packetfabric_locations" "locations_all" {
  provider = packetfabric
}
# output "packetfabric_locations" {
#   value = data.packetfabric_locations.locations_all
# }

locals {
  all_locations = data.packetfabric_locations.locations_all.locations[*]
  helper_map = { for val in local.all_locations :
  val["pop"] => val }
  pf_port_site1 = local.helper_map["${var.pf_port_pop1}"]["site_code"]
  pf_port_site2 = local.helper_map["${var.pf_port_pop2}"]["site_code"]

  pop_in_market = toset([for each in data.packetfabric_locations.locations_all.locations[*] : each.pop if each.market == "WDC"])
}
output "pf_port_site1" {
  value = local.pf_port_site1
}
output "pf_port_site2" {
  value = local.pf_port_site2
}
output "packetfabric_location_pop_in_market" {
  value = local.pop_in_market
}

# Generate a LOA for a port (inbound cross connect)
resource "packetfabric_port_loa" "inbound_crossconnect_1" {
  provider          = packetfabric
  port_circuit_id   = packetfabric_port.port_1a.id
  loa_customer_name = "My Awesome Company"
  destination_email = "email@mydomain.com"
}
output "packetfabric_port_loa" {
  value = packetfabric_port_loa.inbound_crossconnect_1
}

# Create Cross Connect
resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
  provider      = packetfabric
  description   = "${var.resource_name}-${random_pet.name.id}"
  document_uuid = var.pf_document_uuid1
  port          = packetfabric_port.port_1a.id
  site          = local.pf_port_site1
}
output "packetfabric_outbound_cross_connect1" {
  value = packetfabric_outbound_cross_connect.crossconnect_1
}

resource "packetfabric_outbound_cross_connect" "crossconnect_2" {
  provider      = packetfabric
  description   = "${var.resource_name}-${random_pet.name.id}"
  document_uuid = var.pf_document_uuid2
  port          = packetfabric_port.port_1a.id
  site          = local.pf_port_site2
}
output "packetfabric_outbound_cross_connect2" {
  value = packetfabric_outbound_cross_connect.crossconnect_2
}

data "packetfabric_outbound_cross_connect" "crossconnect_1" {
  provider = packetfabric
}
output "packetfabric_outbound_cross_connect" {
  value = data.packetfabric_outbound_cross_connect.crossconnect_1
}

#######################################
##### ACTIVITY LOG
#######################################

data "packetfabric_activitylog" "current" {
  provider = packetfabric
}
output "my-activity-logs" {
  value = data.packetfabric_activitylog.current
}

#######################################
##### Locations
#######################################

data "packetfabric_locations_cloud" "cloud_location_aws" {
  provider              = packetfabric
  cloud_provider        = "aws"
  cloud_connection_type = "hosted"
  # has_cloud_router      = true
  # nat_capable           = true
  # pop                   = var.pf_crc_pop1
}
output "packetfabric_locations_cloud" {
  value = data.packetfabric_locations_cloud.cloud_location_aws
}

data "packetfabric_locations_pop_zones" "locations_pop_zones_DAL_1" {
  provider = packetfabric
  pop      = "DAL1"
}
output "packetfabric_locations_pop_zones" {
  value = data.packetfabric_locations_pop_zones.locations_pop_zones_DAL_1
}

#######################################
##### Hosted Cloud Connections
#######################################

# Create a AWS Hosted Connection
resource "packetfabric_cs_aws_hosted_connection" "cs_conn1_hosted_aws" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  port        = packetfabric_port.port_1a.id
  speed       = var.pf_cs_speed2
  pop         = var.pf_cs_pop2
  vlan        = var.pf_cs_vlan2
  zone        = var.pf_cs_zone2
  # # for cloud side provisioning - optional
  # cloud_settings {
  #   credentials_uuid = packetfabric_cloud_provider_credential_aws.aws_creds1.id
  #   aws_region       = var.pf_cs_aws_region
  #   mtu              = var.pf_cs_mtu
  #   aws_vif_type     = var.pf_cs_aws_vif_type
  #   bgp_settings {
  #     customer_asn   = var.pf_cs_customer_asn
  #     address_family = var.pf_cs_address_family
  #   }
  #   aws_gateways {
  #     type = "directconnect"
  #     name = "${var.resource_name}-${random_pet.name.id}"
  #     asn  = var.pf_cs_directconnect_gw_asn
  #   }
  #   aws_gateways {
  #     type   = "private"
  #     name   = "${var.resource_name}-${random_pet.name.id}"
  #     vpc_id = var.pf_cs_aws_vpc_id
  #   }
  # }
  labels    = var.pf_labels
  po_number = var.pf_po_number
}
output "packetfabric_cs_aws_hosted_connection" {
  value = packetfabric_cs_aws_hosted_connection.cs_conn1_hosted_aws
}

data "packetfabric_cs_aws_hosted_connection" "current" {
  provider         = packetfabric
  cloud_circuit_id = packetfabric_cs_aws_hosted_connection.cs_conn1_hosted_aws.id
}
output "packetfabric_cs_aws_hosted_connection_data" {
  value = data.packetfabric_cs_aws_hosted_connection.current
}

# Create a Azure Hosted Connection 
resource "packetfabric_cs_azure_hosted_connection" "cs_conn1_hosted_azure" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  azure_service_key = var.azure_service_key
  port              = packetfabric_port.port_1a.id
  speed             = var.pf_cs_speed1 # will be deprecated
  vlan_private      = var.pf_cs_vlan_private
  #vlan_microsoft = var.pf_cs_vlan_microsoft
  labels    = var.pf_labels
  po_number = var.pf_po_number
}
output "packetfabric_cs_azure_hosted_connection" {
  sensitive = true
  value     = packetfabric_cs_azure_hosted_connection.cs_conn1_hosted_azure
}

data "packetfabric_cs_azure_hosted_connection" "current" {
  provider         = packetfabric
  cloud_circuit_id = packetfabric_cs_azure_hosted_connection.cs_conn1_hosted_azure.id
}
output "packetfabric_cs_azure_hosted_connection_data" {
  value = data.packetfabric_cs_azure_hosted_connection.current
}

# Create a GCP Hosted Connection 
resource "packetfabric_cs_google_hosted_connection" "cs_conn1_hosted_google" {
  provider                    = packetfabric
  description                 = "${var.resource_name}-${random_pet.name.id}"
  port                        = packetfabric_port.port_1a.id
  speed                       = var.pf_cs_speed1
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = "${var.resource_name}-${random_pet.name.id}"
  pop                         = var.pf_cs_pop1
  vlan                        = var.pf_cs_vlan1
  labels                      = var.pf_labels
  po_number                   = var.pf_po_number
}
output "packetfabric_cs_google_hosted_connection" {
  value     = packetfabric_cs_google_hosted_connection.cs_conn1_hosted_google
  sensitive = true
}

data "packetfabric_cs_google_hosted_connection" "current" {
  provider         = packetfabric
  cloud_circuit_id = packetfabric_cs_google_hosted_connection.cs_conn1_hosted_google.id
}
output "packetfabric_cs_google_hosted_connection_data" {
  value = data.packetfabric_cs_google_hosted_connection.current
}

# Create a Oracle Hosted Connection 
resource "packetfabric_cs_oracle_hosted_connection" "cs_conn1_hosted_oracle" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  vc_ocid     = var.pf_cs_oracle_vc_ocid
  region      = var.pf_cs_oracle_region
  port        = packetfabric_port.port_1a.id
  pop         = var.pf_cs_pop6
  zone        = var.pf_cs_zone6
  vlan        = var.pf_cs_vlan6
  labels      = var.pf_labels
  po_number   = var.pf_po_number
}
output "packetfabric_cs_oracle_hosted_connection" {
  value     = packetfabric_cs_oracle_hosted_connection.cs_conn1_hosted_oracle
  sensitive = true
}

data "packetfabric_cs_oracle_hosted_connection" "current" {
  provider         = packetfabric
  cloud_circuit_id = packetfabric_cs_oracle_hosted_connection.cs_conn1_hosted_oracle.id
}
output "packetfabric_cs_oracle_hosted_connection_data" {
  value = data.packetfabric_cs_oracle_hosted_connection.current
}

resource "packetfabric_cs_ibm_hosted_connection" "cs_conn1_hosted_ibm" {
  provider    = packetfabric
  ibm_bgp_asn = var.ibm_bgp_asn
  description = "${var.resource_name}-${random_pet.name.id}"
  pop         = var.pf_cs_pop7
  port        = packetfabric_port.port_1a.id
  vlan        = var.pf_cs_vlan7
  speed       = var.pf_cs_speed1
  labels      = var.pf_labels
  po_number   = var.pf_po_number
  lifecycle {
    ignore_changes = [
      zone,
      ibm_bgp_cer_cidr,
      ibm_bgp_ibm_cidr
    ]
  }
}
output "packetfabric_cs_ibm_hosted_connection" {
  value = packetfabric_cs_ibm_hosted_connection.cs_conn1_hosted_ibm
}

data "packetfabric_cs_ibm_hosted_connection" "current" {
  provider         = packetfabric
  cloud_circuit_id = packetfabric_cs_ibm_hosted_connection.cs_conn1_hosted_ibm.id
}
output "packetfabric_cs_ibm_hosted_connection_data" {
  value = data.packetfabric_cs_ibm_hosted_connection.current
}

#######################################
##### MARKETPLACE
#######################################

# Create a Marketplace Service type port
resource "packetfabric_marketplace_service" "marketplace_port" {
  provider     = packetfabric
  name         = "${var.resource_name}-${random_pet.name.id}-port"
  description  = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Lobortis mattis aliquam faucibus purus in massa tempor nec."
  service_type = "port-service"
  sku          = var.pf_sku
  categories   = var.pf_categories
  published    = var.pf_published
  locations    = var.pf_locations
}
output "packetfabric_marketplace_service_port" {
  value = packetfabric_marketplace_service.marketplace_port
}

# Create a Marketplace Service type quick-connect
resource "packetfabric_marketplace_service" "marketplace_quick_connect" {
  provider                = packetfabric
  name                    = "${var.resource_name}-${random_pet.name.id}-quick-connect"
  description             = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Lobortis mattis aliquam faucibus purus in massa tempor nec."
  service_type            = "quick-connect-service"
  sku                     = var.pf_sku
  categories              = var.pf_categories
  published               = var.pf_published
  cloud_router_circuit_id = var.pf_cloud_router_circuit_id
  connection_circuit_ids  = var.pf_connection_circuit_ids
  route_set {
    description = var.pf_route_set_description
    is_private  = var.pf_route_set_is_private
    prefixes {
      prefix     = var.pf_route_set_prefix1
      match_type = var.pf_route_set_match_type1
    }
    prefixes {
      prefix     = var.pf_route_set_prefix2
      match_type = var.pf_route_set_match_type2
    }
  }
}
output "packetfabric_marketplace_service_quick_connect" {
  value = packetfabric_marketplace_service.marketplace_quick_connect
}

# Create a VC Marketplace Connection 
resource "packetfabric_backbone_virtual_circuit_marketplace" "vc_marketplace_conn1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  routing_id  = var.pf_routing_id
  market      = var.pf_market
  interface {
    port_circuit_id = packetfabric_port.port_1a.id
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  bandwidth {
    longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}
output "packetfabric_backbone_virtual_circuit_marketplace" {
  value = packetfabric_backbone_virtual_circuit_marketplace.vc_marketplace_conn1
}

# Create an IX Marketplace Connection 
resource "packetfabric_ix_virtual_circuit_marketplace" "ix_marketplace_conn1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  routing_id  = var.pf_routing_id_ix
  market      = var.pf_market_ix
  asn         = var.pf_asn_ix
  interface {
    port_circuit_id = packetfabric_port.port_1a.id
    untagged        = false
    vlan            = var.pf_vc_vlan1
  }
  bandwidth {
    #longhaul_type     = var.pf_vc_longhaul_type
    speed             = var.pf_vc_speed
    subscription_term = var.pf_vc_subterm
  }
}
output "packetfabric_ix_virtual_circuit_marketplace" {
  value = packetfabric_ix_virtual_circuit_marketplace.ix_marketplace_conn1
}

# Create a AWS Hosted Marketplace Connection 
resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_conn1_marketplace_aws" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  routing_id  = var.pf_routing_id
  market      = var.pf_market
  speed       = var.pf_cs_speed2
  pop         = var.pf_cs_pop2
  zone        = var.pf_cs_zone2
}
output "packetfabric_cs_aws_hosted_marketplace_connection" {
  value = packetfabric_cs_aws_hosted_marketplace_connection.cs_conn1_marketplace_aws
}

# Create a Azure Hosted Marketplace Connection 
resource "packetfabric_cs_azure_hosted_marketplace_connection" "cs_conn1_marketplace_azure" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  azure_service_key = var.azure_service_key
  routing_id        = var.pf_routing_id
  market            = var.pf_market
  speed             = var.pf_cs_speed1 # will be deprecated
}
output "packetfabric_cs_azure_hosted_marketplace_connection" {
  sensitive = true
  value     = packetfabric_cs_azure_hosted_marketplace_connection.cs_conn1_marketplace_azure
}

# Create a GCP Hosted Marketplace Connection 
resource "packetfabric_cs_google_hosted_marketplace_connection" "cs_conn1_marketplace_google" {
  provider                    = packetfabric
  description                 = "${var.resource_name}-${random_pet.name.id}"
  routing_id                  = var.pf_routing_id
  market                      = var.pf_market
  speed                       = var.pf_cs_speed1
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = var.google_vlan_attachment_name
  pop                         = var.pf_cs_pop1
}
output "packetfabric_cs_google_hosted_marketplace_connection" {
  value     = packetfabric_cs_google_hosted_marketplace_connection.cs_conn1_marketplace_google
  sensitive = true
}

# Create a Oracle Hosted Marketplace Connection 
resource "packetfabric_cs_oracle_hosted_marketplace_connection" "cs_conn1_marketplace_oracle" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}"
  vc_ocid     = var.pf_cs_oracle_vc_ocid
  region      = var.pf_cs_oracle_region
  routing_id  = var.pf_routing_id
  market      = var.pf_market
  pop         = var.pf_cs_pop6
}
output "packetfabric_cs_oracle_hosted_marketplace_connection" {
  value     = packetfabric_cs_oracle_hosted_marketplace_connection.cs_conn1_marketplace_oracle
  sensitive = true
}

# Accept the Request AWS
resource "packetfabric_marketplace_service_port_accept_request" "accept_marketplace_request_aws" {
  provider       = packetfabric
  type           = "cloud"
  cloud_provider = "aws" # "aws, azure, google, oracle
  description    = "${var.resource_name}-${random_pet.name.id}"
  interface {
    port_circuit_id = var.pf_market_port_circuit_id
    vlan            = var.pf_cs_vlan2
  }
  vc_request_uuid = packetfabric_cs_aws_hosted_marketplace_connection.cs_conn1_marketplace_aws.id
}

# Accept the Request Backbone VC
resource "packetfabric_marketplace_service_port_accept_request" "accept_marketplace_request_backbone" {
  provider    = packetfabric
  type        = "backbone"
  description = "${var.resource_name}-${random_pet.name.id}"
  interface {
    port_circuit_id = var.pf_market_port_circuit_id
    vlan            = var.pf_vc_vlan1
  }
  vc_request_uuid = packetfabric_backbone_virtual_circuit_marketplace.vc_marketplace_conn1.id
}

# Reject the Request
resource "packetfabric_marketplace_service_port_reject_request" "reject_marketplace_request" {
  provider        = packetfabric
  vc_request_uuid = packetfabric_cs_aws_hosted_marketplace_connection.cs_conn1_marketplace_aws.id
}

# List all Marketplace Service Requests (Port)

data "packetfabric_marketplace_service_port_requests" "sent" {
  provider = packetfabric
  type     = "sent" # sent or received
}
output "packetfabric_marketplace_service_port_requests_sent" {
  value = data.packetfabric_marketplace_service_port_requests.sent
}

data "packetfabric_marketplace_service_port_requests" "received" {
  provider = packetfabric
  type     = "received" # sent or received
}
output "packetfabric_marketplace_service_port_requests_received" {
  value = data.packetfabric_marketplace_service_port_requests.received
}

# List all Marketplace Service Requests (Quick Connect)
data "packetfabric_quick_connect_requests" "quick_connect_sent" {
  provider = packetfabric
  type     = "sent" # sent or received
}
output "packetfabric_quick_connect_requests_quick_connect_sent" {
  value = data.packetfabric_quick_connect_requests.quick_connect_sent
}

data "packetfabric_quick_connect_requests" "quick_connect_received" {
  provider = packetfabric
  type     = "received" # sent or received
}
output "packetfabric_quick_connect_requests_quick_connect_received" {
  value = data.packetfabric_quick_connect_requests.quick_connect_received
}

# Accept the Request Quick Connect
resource "packetfabric_quick_connect_accept_request" "accept_request_quick_connect" {
  provider          = packetfabric
  import_circuit_id = packetfabric_cloud_router_quick_connect.cr_quick_connect.import_circuit_id
}

# Reject the Request
resource "packetfabric_quick_connect_reject_request" "reject_request_quick_connect" {
  provider          = packetfabric
  import_circuit_id = packetfabric_cloud_router_quick_connect.cr_quick_connect.import_circuit_id
  rejection_reason  = "Return filters are too broad."
}

#######################################
##### Dedicated Cloud Connections
#######################################

# AWS Dedicated Connection
resource "packetfabric_cs_aws_dedicated_connection" "pf_cs_conn1_dedicated_aws" {
  provider          = packetfabric
  aws_region        = var.aws_region3
  description       = "${var.resource_name}-${random_pet.name.id}"
  zone              = var.pf_cs_zone3
  pop               = var.pf_cs_pop3
  subscription_term = var.pf_cs_subterm
  service_class     = var.pf_cs_srvclass
  autoneg           = var.pf_cs_autoneg
  speed             = var.pf_cs_speed3
  should_create_lag = var.should_create_lag
  labels            = var.pf_labels
  po_number         = var.pf_po_number
}

# GCP Dedicated Connection
resource "packetfabric_cs_google_dedicated_connection" "pf_cs_conn1_dedicated_google" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  zone              = var.pf_cs_zone4
  pop               = var.pf_cs_pop4
  subscription_term = var.pf_cs_subterm
  service_class     = var.pf_cs_srvclass
  autoneg           = var.pf_cs_autoneg
  speed             = var.pf_cs_speed4
  labels            = var.pf_labels
  po_number         = var.pf_po_number
}
output "packetfabric_cs_google_dedicated_connection" {
  value = packetfabric_cs_google_dedicated_connection.pf_cs_conn1_dedicated_google
}

# Azure Dedicated Connection
resource "packetfabric_cs_azure_dedicated_connection" "pf_cs_conn1_dedicated_azure" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  zone              = var.pf_cs_zone5
  pop               = var.pf_cs_pop5
  subscription_term = var.pf_cs_subterm
  service_class     = var.pf_cs_srvclass
  encapsulation     = var.encapsulation
  port_category     = var.port_category
  speed             = var.pf_cs_speed5
  labels            = var.pf_labels
  po_number         = var.pf_po_number
}

data "packetfabric_cs_dedicated_connections" "current" {
  provider = packetfabric
}
output "packetfabric_cs_dedicated_connection" {
  value = data.packetfabric_cs_dedicated_connections.current
}

#######################################
##### Cloud Router
#######################################

resource "packetfabric_cloud_router" "cr" {
  provider = packetfabric
  asn      = var.pf_cr_asn
  name     = "${var.resource_name}-${random_pet.name.id}"
  capacity = var.pf_cr_capacity
  regions  = var.pf_cr_regions
  labels   = var.pf_labels
}

resource "packetfabric_cloud_router_connection_aws" "crc_1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
  circuit_id  = packetfabric_cloud_router.cr.id
  pop         = var.pf_crc_pop1
  zone        = var.pf_crc_zone1
  speed       = var.pf_crc_speed
  maybe_nat   = var.pf_crc_maybe_nat
  maybe_dnat  = var.pf_crc_maybe_dnat
  is_public   = var.pf_crc_is_public
  labels      = var.pf_labels
  po_number   = var.pf_po_number
}

resource "packetfabric_cloud_router_bgp_session" "crbs_1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_aws.crc_1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.aws_side_asn1
  remote_address = var.aws_remote_address # AWS side
  l3_address     = var.aws_l3_address     # PF side
  orlonger       = var.pf_crbs_orlonger
  # nat { # example source NAT
  #   direction       = "output" # or input
  #   nat_type        = "overload"
  #   pre_nat_sources = ["10.1.1.0/24", "10.1.2.0/24"]
  #   pool_prefixes   = ["192.168.1.50/32", "192.168.1.51/32"]
  # }
  # nat { # example destination NAT
  #   nat_type = "inline_dnat"
  #   dnat_mappings {
  #     private_prefix = "192.168.1.50/32"
  #     public_prefix  = "192.167.1.50/32"
  #   }
  #   dnat_mappings {
  #     private_prefix     = "192.168.2.50/32"
  #     public_prefix      = "192.166.1.50/32"
  #     conditional_prefix = "192.168.2.0/24" # must be a subnet of private_prefix
  #   }
  # }
  prefixes {
    prefix = "0.0.0.0/0"
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = "0.0.0.0/0"
    type   = "in" # Allowed Prefixes from Cloud
  }
  labels    = var.pf_labels
  po_number = var.pf_po_number
}
output "packetfabric_cloud_router_bgp_session_crbs_1" {
  value = packetfabric_cloud_router_bgp_session.crbs_1
}

data "packetfabric_cloud_router_bgp_session" "bgp_session_crbs_1" {
  provider      = packetfabric
  circuit_id    = packetfabric_cloud_router.cr.id
  connection_id = packetfabric_cloud_router_connection_aws.crc_1.id
  depends_on    = [packetfabric_cloud_router_bgp_session.crbs_1]
}
output "packetfabric_cloud_router_bgp_session_crbs_1_data" {
  value = data.packetfabric_cloud_router_bgp_session.bgp_session_crbs_1
}

resource "packetfabric_cloud_router_connection_google" "crc_2" {
  provider                    = packetfabric
  description                 = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  circuit_id                  = packetfabric_cloud_router.cr.id
  google_pairing_key          = var.pf_crc_google_pairing_key
  google_vlan_attachment_name = var.pf_crc_google_vlan_attachment_name
  pop                         = var.pf_crc_pop2
  speed                       = var.pf_crc_speed
  maybe_nat                   = var.pf_crc_maybe_nat
  maybe_dnat                  = var.pf_crc_maybe_dnat
  labels                      = var.pf_labels
  po_number                   = var.pf_po_number
}

resource "packetfabric_cloud_router_connection_ipsec" "crc_3" {
  provider                     = packetfabric
  description                  = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop3}"
  circuit_id                   = packetfabric_cloud_router.cr.id
  pop                          = var.pf_crc_pop3
  speed                        = var.pf_crc_speed
  gateway_address              = var.pf_crc_gateway_address
  ike_version                  = var.pf_crc_ike_version
  phase1_authentication_method = var.pf_crc_phase1_authentication_method
  phase1_group                 = var.pf_crc_phase1_group
  phase1_encryption_algo       = var.pf_crc_phase1_encryption_algo
  phase1_authentication_algo   = var.pf_crc_phase1_authentication_algo
  phase1_lifetime              = var.pf_crc_phase1_lifetime
  phase2_pfs_group             = var.pf_crc_phase2_pfs_group
  phase2_encryption_algo       = var.pf_crc_phase2_encryption_algo
  phase2_authentication_algo   = var.pf_crc_phase2_authentication_algo
  phase2_lifetime              = var.pf_crc_phase2_lifetime
  shared_key                   = var.pf_crc_shared_key
  labels                       = var.pf_labels
}

resource "packetfabric_cloud_router_bgp_session" "crbs_3" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr.id
  connection_id  = packetfabric_cloud_router_connection_ipsec.crc_3.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.vpn_side_asn3
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.vpn_remote_address # On-Prem side
  l3_address     = var.vpn_l3_address     # PF side
  prefixes {
    prefix = "0.0.0.0/0"
    type   = "out" # Allowed Prefixes to Cloud
  }
  prefixes {
    prefix = "0.0.0.0/0"
    type   = "in" # Allowed Prefixes from Cloud
  }
  labels    = var.pf_labels
  po_number = var.pf_po_number
}
output "packetfabric_cloud_router_bgp_session_crbs_3" {
  value = packetfabric_cloud_router_bgp_session.crbs_3
}

data "packetfabric_cloud_router_bgp_session" "bgp_session_crbs_3" {
  provider      = packetfabric
  circuit_id    = packetfabric_cloud_router.cr.id
  connection_id = packetfabric_cloud_router_connection_ipsec.crc_3.id
  depends_on    = [packetfabric_cloud_router_bgp_session.crbs_1]
}
output "packetfabric_cloud_router_bgp_session_crbs_3_data" {
  value = data.packetfabric_cloud_router_bgp_session.bgp_session_crbs_3
}

resource "packetfabric_cloud_router_connection_azure" "crc_4" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
  circuit_id        = packetfabric_cloud_router.cr.id
  azure_service_key = var.pf_crc_azure_service_key
  speed             = var.pf_crc_speed
  maybe_nat         = var.pf_crc_maybe_nat
  maybe_dnat        = var.pf_crc_maybe_dnat
  is_public         = var.pf_crc_is_public
  labels            = var.pf_labels
  po_number         = var.pf_po_number
}

resource "packetfabric_cloud_router_connection_ibm" "crc_5" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop4}"
  circuit_id  = packetfabric_cloud_router.cr.id
  ibm_bgp_asn = var.pf_crc_ibm_bgp_asn
  pop         = var.pf_crc_pop4
  zone        = var.pf_crc_zone4
  maybe_nat   = var.pf_crc_maybe_nat
  maybe_dnat  = var.pf_crc_maybe_dnat
  speed       = var.pf_crc_speed
  labels      = var.pf_labels
  po_number   = var.pf_po_number
  lifecycle {
    ignore_changes = [
      ibm_bgp_cer_cidr,
      ibm_bgp_ibm_cidr
    ]
  }
}

resource "packetfabric_cloud_router_connection_oracle" "crc_6" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-${var.pf_crc_pop5}"
  circuit_id  = packetfabric_cloud_router.cr.id
  region      = var.pf_crc_oracle_region
  vc_ocid     = var.pf_crc_oracle_vc_ocid
  pop         = var.pf_crc_pop5
  zone        = var.pf_crc_zone5
  maybe_nat   = var.pf_crc_maybe_nat
  maybe_dnat  = var.pf_crc_maybe_dnat
  labels      = var.pf_labels
  po_number   = var.pf_po_number
}

resource "packetfabric_cloud_router_connection_port" "crc_7" {
  provider        = packetfabric
  description     = "${var.resource_name}-${random_pet.name.id}"
  circuit_id      = packetfabric_cloud_router.cr.id
  port_circuit_id = packetfabric_port.port_1a.id
  vlan            = var.pf_crc_vlan
  speed           = var.pf_crc_speed
  is_public       = var.pf_crc_is_public
  maybe_nat       = var.pf_crc_maybe_nat
  maybe_dnat      = var.pf_crc_maybe_dnat
  labels          = var.pf_labels
}

data "packetfabric_cloud_router_connections" "all_crc" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id
}
output "packetfabric_cloud_router_connections" {
  value = data.packetfabric_cloud_router_connections.all_crc
}

# Create a Quick Connect Cloud Router Marketplace Connection
resource "packetfabric_cloud_router_quick_connect" "cr_quick_connect" {
  provider              = packetfabric
  cr_circuit_id         = var.pf_cr_circuit_id
  connection_circuit_id = var.pf_connection_circuit_id
  service_uuid          = var.pf_service_uuid
  return_filters {
    prefix     = var.pf_return_filters_prefix1
    match_type = var.pf_return_filters_match_type1
  }
  return_filters {
    prefix     = var.pf_return_filters_prefix2
    match_type = var.pf_return_filters_match_type2
  }
  labels    = var.pf_labels
  po_number = var.pf_po_number
}
output "packetfabric_cloud_router_quick_connect" {
  value = packetfabric_cloud_router_quick_connect.cr_quick_connect
}

#######################################
##### Billing
#######################################

# Get billing information related to the ports created
data "packetfabric_billing" "port_1a" {
  provider   = packetfabric
  circuit_id = packetfabric_port.port_1a.id
}
output "packetfabric_billing_port_1a" {
  value = data.packetfabric_billing.port_1a
}

#######################################
##### Flex Bandwidth
#######################################

resource "packetfabric_flex_bandwidth" "flex1" {
  provider          = packetfabric
  description       = "${var.resource_name}-${random_pet.name.id}"
  subscription_term = var.pf_flex_subscription_term
  capacity          = var.pf_flex_capacity
}
output "packetfabric_flex_bandwidth" {
  value = packetfabric_flex_bandwidth.flex1
}

#######################################
##### Cloud Provider Credentials
#######################################

resource "packetfabric_cloud_provider_credential_aws" "aws_creds1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-aws"
  # using env var PF_AWS_ACCESS_KEY_ID and PF_AWS_SECRET_ACCESS_KEY
}
output "packetfabric_cloud_provider_credential_aws" {
  value     = packetfabric_cloud_provider_credential_aws.aws_creds1
  sensitive = true
}

resource "packetfabric_cloud_provider_credential_google" "google_creds1" {
  provider    = packetfabric
  description = "${var.resource_name}-${random_pet.name.id}-google"
  # using env var GOOGLE_CREDENTIALS
}
output "packetfabric_cloud_provider_credential_google" {
  value     = packetfabric_cloud_provider_credential_google.google_creds1
  sensitive = true
}