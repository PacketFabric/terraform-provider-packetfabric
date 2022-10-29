terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.3.2"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

# Create random name to use to name objects
resource "random_pet" "name" {}

########################################
###### PORTS/INTERFACES
########################################

# # Create a PacketFabric Ports
# resource "packetfabric_port" "port_1a" {
#   provider          = packetfabric
#   account_uuid      = var.pf_account_uuid
#   autoneg           = var.pf_port_autoneg
#   description       = "${var.tag_name}-${random_pet.name.id}-a"
#   media             = var.pf_port_media
#   nni               = var.pf_port_nni
#   pop               = var.pf_port_pop1
#   speed             = var.pf_port_speed
#   subscription_term = var.pf_port_subterm
#   zone              = var.pf_port_avzone1
# }
# output "packetfabric_port_1a" {
#   value = packetfabric_port.port_1a
# }

# ## 2nd port in the same location same zone to create a LAG
# resource "packetfabric_port" "port_1b" {
#   provider          = packetfabric
#   account_uuid      = var.pf_account_uuid
#   autoneg           = var.pf_port_autoneg
#   description       = "${var.tag_name}-${random_pet.name.id}-b"
#   media             = var.pf_port_media
#   nni               = var.pf_port_nni
#   pop               = var.pf_port_pop1
#   speed             = var.pf_port_speed
#   subscription_term = var.pf_port_subterm
#   zone              = var.pf_port_avzone1
# }
# output "packetfabric_port_1b" {
#   value = packetfabric_port.port_1b
# }

# resource "packetfabric_link_aggregation_group" "lag_1" {
#   provider    = packetfabric
#   description = "${var.tag_name}-${random_pet.name.id}"
#   interval    = "fast" # or slow
#   members     = [packetfabric_port.port_1a.id, packetfabric_port.port_1b.id]
#   #members = [packetfabric_port.port_1a.id]
#   pop = var.pf_port_pop1
# }

# data "packetfabric_link_aggregation_group" "lag_1" {
#   provider            = packetfabric
#   lag_port_circuit_id = packetfabric_link_aggregation_group.lag_1.id
# }

# output "packetfabric_link_aggregation_group" {
#   value = data.packetfabric_link_aggregation_group.lag_1
# }

# resource "packetfabric_port" "port_2" {
#   provider          = packetfabric
#   account_uuid      = var.pf_account_uuid
#   autoneg           = var.pf_port_autoneg
#   description       = "${var.tag_name}-${random_pet.name.id}"
#   media             = var.pf_port_media
#   nni               = var.pf_port_nni
#   pop               = var.pf_port_pop2
#   speed             = var.pf_port_speed
#   subscription_term = var.pf_port_subterm
#   zone              = var.pf_port_avzone2
# }
# output "packetfabric_port_2" {
#   value = packetfabric_port.port_2
# }

# data "packetfabric_port" "ports_all" {
#   provider = packetfabric
# }

# output "packetfabric_ports_all" {
#   value = data.packetfabric_port.ports_all
# }

# #######################################
# ##### Billing
# #######################################

# # Get billing information related to the ports created
# data "packetfabric_billing" "port_1a" {
#   provider   = packetfabric
#   circuit_id = packetfabric_port.port_1a.id
# }
# output "packetfabric_billing_port_1a" {
#   value = data.packetfabric_billing.port_1a
# }

# #######################################
# ##### Virtual Circuit (VC)
# #######################################

# # Create backbone Virtual Circuit
# resource "packetfabric_backbone_virtual_circuit" "vc1" {
#   provider    = packetfabric
#   description = "${var.tag_name}-${random_pet.name.id}"
#   epl         = false
#   interface_a {
#     port_circuit_id = packetfabric_port.port_1a.id
#     untagged        = false
#     vlan            = var.pf_vc_vlan1
#   }
#   interface_z {
#     port_circuit_id = packetfabric_port.port_2.id
#     untagged        = false
#     vlan            = var.pf_vc_vlan2
#   }
#   bandwidth {
#     account_uuid      = var.pf_account_uuid
#     longhaul_type     = var.pf_vc_longhaul_type
#     speed             = var.pf_vc_speed
#     subscription_term = var.pf_vc_subterm
#   }
# }

# # Show all Virtual Circuits
# data "packetfabric_virtual_circuits" "all_vcs" {
#   provider   = packetfabric
# }
# output "packetfabric_virtual_circuits" {
#   value = data.packetfabric_virtual_circuits.all_vcs
# }

# #######################################
# ##### Virtual Circuit Speed Burst
# #######################################

# resource "packetfabric_backbone_virtual_circuit_speed_burst" "boost" {
#   provider      = packetfabric
#   vc_circuit_id = var.pf_vc_circuit_id
#   speed         = var.pf_vc_speed_burst
# }

# output "packetfabric_backbone_virtual_circuit_speed_burst" {
#   value = packetfabric_backbone_virtual_circuit_speed_burst.boost
# }

# #######################################
# ##### Point to Point
# #######################################

# resource "packetfabric_point_to_point" "ptp1" {
#   provider          = packetfabric
#   description       = "${var.tag_name}-${random_pet.name.id}"
#   speed             = var.pf_ptp_speed
#   media             = var.pf_ptp_media
#   account_uuid      = var.pf_account_uuid
#   subscription_term = var.pf_ptp_subterm
#   endpoints {
#     pop     = var.pf_ptp_pop1
#     zone    = var.pf_ptp_zone1
#     autoneg = var.pf_ptp_autoneg
#   }
#   endpoints {
#     pop     = var.pf_ptp_pop2
#     zone    = var.pf_ptp_zone2
#     autoneg = var.pf_ptp_autoneg
#   }
# }

# #######################################
# ##### Cross Connect
# #######################################

# ## Get the site filtering on the pop using packetfabric_locations

# # List PacketFabric locations
# data "packetfabric_locations" "locations_all" {
#   provider = packetfabric
#   # check https://github.com/PacketFabric/terraform-provider-packetfabric/issues/63 to use filter
#   # filter {
#   #   pop = var.pf_port_pop1
#   # }
# }
# # output "packetfabric_locations" {
# #   value = data.packetfabric_locations.locations_all
# # }

# locals {
#   all_locations = data.packetfabric_locations.locations_all.locations[*]
#   helper_map = { for val in local.all_locations :
#   val["pop"] => val }
#   pf_port_site1 = local.helper_map["${var.pf_port_pop1}"]["site_code"]
#   pf_port_site2 = local.helper_map["${var.pf_port_pop2}"]["site_code"]
#   
#   pop_in_market = toset([for each in data.packetfabric_locations.locations_all.locations[*] : each.pop if each.market == "WDC"])
# }
# output "pf_port_site1" {
#   value = local.pf_port_site1
# }
# output "pf_port_site2" {
#   value = local.pf_port_site2
# }
# output "packetfabric_location_pop_in_market" {
#   value = local.pop_in_market
# }

# # Create Cross Connect
# resource "packetfabric_outbound_cross_connect" "crossconnect_1" {
#   provider      = packetfabric
#   description   = "${var.tag_name}-${random_pet.name.id}"
#   document_uuid = var.pf_document_uuid1
#   port          = packetfabric_port.port_1a.id
#   site          = local.pf_port_site1
# }
# output "packetfabric_outbound_cross_connect1" {
#   value = packetfabric_outbound_cross_connect.crossconnect_1
# }

# resource "packetfabric_outbound_cross_connect" "crossconnect_2" {
#   provider      = packetfabric
#   description   = "${var.tag_name}-${random_pet.name.id}"
#   document_uuid = var.pf_document_uuid2
#   port          = packetfabric_port.port_1a.id
#   site          = local.pf_port_site2
# }
# output "packetfabric_outbound_cross_connect2" {
#   value = packetfabric_outbound_cross_connect.crossconnect_2
# }

# data "packetfabric_outbound_cross_connect" "crossconnect_1" {
#   provider = packetfabric
# }

# output "packetfabric_outbound_cross_connect" {
#   value = data.packetfabric_outbound_cross_connect.crossconnect_1
# }

# #######################################
# ##### ACTIVITY LOG
# #######################################

# data "packetfabric_activitylog" "current" {
#   provider = packetfabric
# }
# output "my-activity-logs" {
#   value = data.packetfabric_activitylog.current
# }

# #######################################
# ##### HOSTED CLOUD CONNECTIONS
# #######################################

# # Create a AWS Hosted Connection 
# resource "packetfabric_cs_aws_hosted_connection" "cs_conn1_hosted_aws" {
#   provider       = packetfabric
#   description    = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid   = var.pf_account_uuid
#   aws_account_id = var.pf_aws_account_id
#   port           = packetfabric_port.port_1a.id
#   speed          = var.pf_cs_speed2
#   pop            = var.pf_cs_pop2
#   vlan           = var.pf_cs_vlan2
#   zone           = var.pf_cs_zone2
# }

# output "packetfabric_cs_aws_hosted_connection" {
#   value = packetfabric_cs_aws_hosted_connection.cs_conn1_hosted_aws
# }

# data "packetfabric_cs_aws_hosted_connection" "current" {
#   provider         = packetfabric
#   cloud_circuit_id = packetfabric_cs_aws_hosted_connection.cs_conn1_hosted_aws.id
# }

# output "packetfabric_cs_aws_hosted_connection_data" {
#   value = data.packetfabric_cs_aws_hosted_connection.current
# }

# # Create a Azure Hosted Connection 
# resource "packetfabric_cs_azure_hosted_connection" "cs_conn1_hosted_azure" {
#   provider          = packetfabric
#   description       = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid      = var.pf_account_uuid
#   azure_service_key = var.azure_service_key
#   port              = packetfabric_port.port_1a.id
#   speed             = var.pf_cs_speed1 # will be deprecated
#   vlan_private      = var.pf_cs_vlan_private
#   #vlan_microsoft = var.pf_cs_vlan_microsoft
# }

# output "packetfabric_cs_azure_hosted_connection" {
#   sensitive = true
#   value     = packetfabric_cs_azure_hosted_connection.cs_conn1_hosted_azure
# }

# data "packetfabric_cs_azure_hosted_connection" "current" {
#   provider = packetfabric
# }

# output "packetfabric_cs_azure_hosted_connection_data" {
#   value = data.packetfabric_cs_azure_hosted_connection.current
# }

# # Create a GCP Hosted Connection 
# resource "packetfabric_cs_google_hosted_connection" "cs_conn1_hosted_google" {
#   provider                    = packetfabric
#   description                 = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid                = var.pf_account_uuid
#   port                        = packetfabric_port.port_1a.id
#   speed                       = var.pf_cs_speed1
#   google_pairing_key          = var.google_pairing_key
#   google_vlan_attachment_name = "${var.tag_name}-${random_pet.name.id}"
#   pop                         = var.pf_cs_pop1
#   vlan                        = var.pf_cs_vlan1
# }

# # type terraform output packetfabric_cs_google_hosted_connection to display the value
# output "packetfabric_cs_google_hosted_connection" {
#   value     = packetfabric_cs_google_hosted_connection.cs_conn1_hosted_google
#   sensitive = true
# }

# data "packetfabric_cs_google_hosted_connection" "current" {
#   provider = packetfabric
# }

# output "packetfabric_cs_google_hosted_connection" {
#   value = data.packetfabric_cs_google_hosted_connection.current
# }

# # Create a Oracle Hosted Connection 
# resource "packetfabric_cs_oracle_hosted_connection" "cs_conn1_hosted_oracle" {
#   provider     = packetfabric
#   description  = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid = var.pf_account_uuid
#   vc_ocid      = var.pf_cs_oracle_vc_ocid
#   region       = var.pf_cs_oracle_region
#   port         = packetfabric_port.port_1a.id
#   pop          = var.pf_cs_pop6
#   zone         = var.pf_cs_zone6
#   vlan         = var.pf_cs_vlan6
# }

# output "packetfabric_cs_oracle_hosted_connection" {
#   value     = packetfabric_cs_oracle_hosted_connection.cs_conn1_hosted_oracle
#   sensitive = true
# }

# data "packetfabric_cs_oracle_hosted_connection" "current" {
#   provider         = packetfabric
#   cloud_circuit_id = packetfabric_cs_oracle_hosted_connection.cs_conn1_hosted_oracle.id
# }

# output "packetfabric_cs_oracle_hosted_connection_data" {
#   value = data.packetfabric_cs_oracle_hosted_connection.current
# }

# #######################################
# ##### MARKETPLACE
# #######################################

# # Create a VC Marketplace Connection 
# resource "packetfabric_backbone_virtual_circuit_marketplace" "vc_marketplace_conn1" {
#   provider    = packetfabric
#   description = "${var.tag_name}-${random_pet.name.id}"
#   routing_id  = var.pf_routing_id
#   market      = var.pf_market
#   interface {
#     port_circuit_id = packetfabric_port.port_1a.id
#     untagged        = false
#     vlan            = var.pf_vc_vlan1
#   }
#   bandwidth {
#     account_uuid      = var.pf_account_uuid
#     longhaul_type     = var.pf_vc_longhaul_type
#     speed             = var.pf_vc_speed
#     subscription_term = var.pf_vc_subterm
#   }
# }
# output "packetfabric_backbone_virtual_circuit_marketplace" {
#   value = packetfabric_backbone_virtual_circuit_marketplace.vc_marketplace_conn1
# }

# # Create an IX Marketplace Connection 
# resource "packetfabric_ix_virtual_circuit_marketplace" "ix_marketplace_conn1" {
#   provider    = packetfabric
#   description = "${var.tag_name}-${random_pet.name.id}"
#   routing_id  = var.pf_routing_id_ix
#   market      = var.pf_market_ix
#   asn         = var.pf_asn_ix
#   interface {
#     port_circuit_id = packetfabric_port.port_1a.id
#     untagged        = false
#     vlan            = var.pf_vc_vlan1
#   }
#   bandwidth {
#     account_uuid      = var.pf_account_uuid
#     longhaul_type     = var.pf_vc_longhaul_type
#     speed             = var.pf_vc_speed
#     subscription_term = var.pf_vc_subterm
#   }
# }
# output "packetfabric_ix_virtual_circuit_marketplace" {
#   value = packetfabric_ix_virtual_circuit_marketplace.ix_marketplace_conn1
# }

# # Create a AWS Hosted Marketplace Connection 
# resource "packetfabric_cs_aws_hosted_marketplace_connection" "cs_conn1_marketplace" {
#   provider       = packetfabric
#   description    = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid   = var.pf_account_uuid
#   aws_account_id = var.pf_aws_account_id
#   routing_id     = var.pf_routing_id
#   market         = var.pf_market
#   speed          = var.pf_cs_speed2
#   pop            = var.pf_cs_pop2
#   zone           = var.pf_cs_zone2
# }
# output "packetfabric_cs_aws_hosted_marketplace_connection" {
#   value = packetfabric_cs_aws_hosted_marketplace_connection.cs_conn1_marketplace
# }

# # Create a Azure Hosted Marketplace Connection 
# resource "packetfabric_cs_azure_hosted_marketplace_connection" "cs_conn1_marketplace_azure" {
#   provider          = packetfabric
#   description       = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid      = var.pf_account_uuid
#   azure_service_key = var.azure_service_key
#   routing_id        = var.pf_routing_id
#   market            = var.pf_market
#   speed             = var.pf_cs_speed1 # will be deprecated
# }
# output "packetfabric_cs_azure_hosted_marketplace_connection" {
#   sensitive = true
#   value     = packetfabric_cs_azure_hosted_marketplace_connection.cs_conn1_marketplace_azure
# }

# # Create a GCP Hosted Marketplace Connection 
# resource "packetfabric_cs_google_hosted_marketplace_connection" "cs_conn1_marketplace_google" {
#   provider                    = packetfabric
#   description                 = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid                = var.pf_account_uuid
#   routing_id                  = var.pf_routing_id
#   market                      = var.pf_market
#   speed                       = var.pf_cs_speed1
#   google_pairing_key          = var.google_pairing_key
#   google_vlan_attachment_name = var.google_vlan_attachment_name
#   pop                         = var.pf_cs_pop1
# }
# output "packetfabric_cs_google_hosted_marketplace_connection" {
#   value     = packetfabric_cs_google_hosted_marketplace_connection.cs_conn1_marketplace_google
#   sensitive = true
# }

# # Create a Oracle Hosted Marketplace Connection 
# resource "packetfabric_cs_oracle_hosted_marketplace_connection" "cs_conn1_marketplace_oracle" {
#   provider     = packetfabric
#   description  = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid = var.pf_account_uuid
#   vc_ocid      = var.pf_cs_oracle_vc_ocid
#   region       = var.pf_cs_oracle_region
#   routing_id   = var.pf_routing_id
#   market       = var.pf_market
#   pop          = var.pf_cs_pop6
# }
# output "packetfabric_cs_oracle_hosted_marketplace_connection" {
#   value = packetfabric_cs_oracle_hosted_marketplace_connection.cs_conn1_marketplace_oracle
#   sensitive = true
# }

# # Accept the Request
# resource "packetfabric_marketplace_service_accept_request" "accept_marketplace_request" {
#   provider        = packetfabric
#   type            = "cloud" # backbone, ix or cloud
#   cloud_provider  = "aws" # "aws, azure, google, oracle
#   description     = "${var.tag_name}-${random_pet.name.id}"
#   port_circuit_id = var.pf_market_port_circuit_id
#   vc_request_uuid = packetfabric_cs_aws_hosted_marketplace_connection.cs_conn1_marketplace.id
# }

# # Reject the Request
# resource "packetfabric_marketplace_service_reject_request" "reject_marketplace_request" {
#   provider        = packetfabric
#   delete_reason   = "testing"
#   vc_request_uuid = packetfabric_cs_azure_hosted_marketplace_connection.cs_conn1_marketplace_azure.id
# }

# # List all Marketplace Service Requests (not Cloud Router)
# data "packetfabric_marketplace_service_requests" "sent" {
#   provider = packetfabric
#   type     = "sent" # sent or received
# }
# output "packetfabric_marketplace_service_requests_sent" {
#   value = data.packetfabric_marketplace_service_requests.sent
# }

# data "packetfabric_marketplace_service_requests" "received" {
#   provider = packetfabric
#   type     = "received" # sent or received
# }
# output "packetfabric_marketplace_service_requests_received" {
#   value = data.packetfabric_marketplace_service_requests.received
# }

# #######################################
# ##### DEDICATED CLOUD CONNECTIONS
# #######################################

# # AWS Dedicated Connection
# resource "packetfabric_cs_aws_dedicated_connection" "pf_cs_conn1_dedicated_aws" {
#   provider          = packetfabric
#   aws_region        = var.aws_region3
#   account_uuid      = var.pf_account_uuid
#   description       = "${var.tag_name}-${random_pet.name.id}"
#   zone              = var.pf_cs_zone3
#   pop               = var.pf_cs_pop3
#   subscription_term = var.pf_cs_subterm
#   service_class     = var.pf_cs_srvclass
#   autoneg           = var.pf_cs_autoneg
#   speed             = var.pf_cs_speed3
#   should_create_lag = var.should_create_lag
# }

# data "packetfabric_cs_aws_dedicated_connection" "current" {
#   provider = packetfabric
# }

# output "packetfabric_cs_aws_dedicated_connection" {
#   value = data.packetfabric_cs_aws_dedicated_connection.current
# }

# # GCP Dedicated Connection
# resource "packetfabric_cs_google_dedicated_connection" "pf_cs_conn1_dedicated_google" {
#   provider          = packetfabric
#   account_uuid      = var.pf_account_uuid
#   description       = "${var.tag_name}-${random_pet.name.id}"
#   zone              = var.pf_cs_zone4
#   pop               = var.pf_cs_pop4
#   subscription_term = var.pf_cs_subterm
#   service_class     = var.pf_cs_srvclass
#   autoneg           = var.pf_cs_autoneg
#   speed             = var.pf_cs_speed4
# }

# data "packetfabric_cs_google_dedicated_connection" "current" {
#   provider = packetfabric
# }

# output "packetfabric_cs_google_dedicated_connection" {
#   value = data.packetfabric_cs_google_dedicated_connection.current
# }

# # Azure Dedicated Connection
# resource "packetfabric_cs_azure_dedicated_connection" "pf_cs_conn1_dedicated_azure" {
#   provider          = packetfabric
#   account_uuid      = var.pf_account_uuid
#   description       = "${var.tag_name}-${random_pet.name.id}"
#   zone              = var.pf_cs_zone5
#   pop               = var.pf_cs_pop5
#   subscription_term = var.pf_cs_subterm
#   service_class     = var.pf_cs_srvclass
#   encapsulation     = var.encapsulation
#   port_category     = var.port_category
#   speed             = var.pf_cs_speed5
# }

# data "packetfabric_cs_azure_dedicated_connection" "current" {
#   provider = packetfabric
# }

# output "packetfabric_cs_azure_dedicated_connection" {
#   value = data.packetfabric_cs_azure_dedicated_connection.current
# }

# #######################################
# ##### CLOUD ROUTER
# #######################################

# resource "packetfabric_cloud_router" "cr" {
#   provider     = packetfabric
#   asn          = var.pf_cr_asn
#   name         = "${var.tag_name}-${random_pet.name.id}"
#   account_uuid = var.pf_account_uuid
#   capacity     = var.pf_cr_capacity
#   regions      = var.pf_cr_regions
# }

# resource "packetfabric_cloud_router_connection_aws" "crc_1" {
#   provider       = packetfabric
#   description    = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop1}"
#   circuit_id     = packetfabric_cloud_router.cr.id
#   account_uuid   = var.pf_account_uuid
#   aws_account_id = var.pf_aws_account_id
#   pop            = var.pf_crc_pop1
#   zone           = var.pf_crc_zone1
#   speed          = var.pf_crc_speed
#   maybe_nat      = var.pf_crc_maybe_nat
#   is_public      = var.pf_crc_is_public
# }

# resource "packetfabric_cloud_router_connection_google" "crc_2" {
#   provider                    = packetfabric
#   description                 = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
#   circuit_id                  = packetfabric_cloud_router.cr.id
#   account_uuid                = var.pf_account_uuid
#   google_pairing_key          = var.pf_crc_google_pairing_key
#   google_vlan_attachment_name = var.pf_crc_google_vlan_attachment_name
#   pop                         = var.pf_crc_pop2
#   speed                       = var.pf_crc_speed
#   maybe_nat                   = var.pf_crc_maybe_nat
# }

# resource "packetfabric_cloud_router_connection_ipsec" "crc_3" {
#   provider                     = packetfabric
#   description                  = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop3}"
#   circuit_id                   = packetfabric_cloud_router.cr.id
#   account_uuid                 = var.pf_account_uuid
#   pop                          = var.pf_crc_pop3
#   speed                        = var.pf_crc_speed
#   gateway_address              = var.pf_crc_gateway_address
#   ike_version                  = var.pf_crc_ike_version
#   phase1_authentication_method = var.pf_crc_phase1_authentication_method
#   phase1_group                 = var.pf_crc_phase1_group
#   phase1_encryption_algo       = var.pf_crc_phase1_encryption_algo
#   phase1_authentication_algo   = var.pf_crc_phase1_authentication_algo
#   phase1_lifetime              = var.pf_crc_phase1_lifetime
#   phase2_pfs_group             = var.pf_crc_phase2_pfs_group
#   phase2_encryption_algo       = var.pf_crc_phase2_encryption_algo
#   phase2_authentication_algo   = var.pf_crc_phase2_authentication_algo
#   phase2_lifetime              = var.pf_crc_phase2_lifetime
#   shared_key                   = var.pf_crc_shared_key
# }


# resource "packetfabric_cloud_router_bgp_session" "crbs_3" {
#   provider       = packetfabric
#   circuit_id     = packetfabric_cloud_router.cr.id
#   connection_id  = packetfabric_cloud_router_connection_ipsec.crc_3.id
#   address_family = var.pf_crbs_af
#   multihop_ttl   = var.pf_crbs_mhttl
#   remote_asn     = var.vpn_side_asn3
#   orlonger       = var.pf_crbs_orlonger
#   remote_address = var.vpn_remote_address # On-Prem side
#   l3_address     = var.vpn_l3_address     # PF side
#   prefixes {
#     prefix = "0.0.0.0/0"
#     type   = "out" # Allowed Prefixes to Cloud
#     order  = 0
#   }
#   prefixes {
#     prefix = "0.0.0.0/0"
#     type   = "in" # Allowed Prefixes from Cloud
#     order  = 0
#   }
# }
# output "packetfabric_cloud_router_bgp_session_crbs_3" {
#   value = packetfabric_cloud_router_bgp_session.crbs_3
# }


# resource "packetfabric_cloud_router_connection_azure" "crc_4" {
#   provider          = packetfabric
#   description       = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop2}"
#   circuit_id        = packetfabric_cloud_router.cr.id
#   account_uuid      = var.pf_account_uuid
#   azure_service_key = var.pf_crc_azure_service_key
#   speed             = var.pf_crc_speed
#   maybe_nat         = var.pf_crc_maybe_nat
#   is_public         = var.pf_crc_is_public
# }

# resource "packetfabric_cloud_router_connection_ibm" "crc_5" {
#   provider         = packetfabric
#   description      = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop4}"
#   circuit_id       = packetfabric_cloud_router.cr.id
#   account_uuid     = var.pf_account_uuid
#   ibm_account_id   = var.pf_crc_ibm_account_id
#   ibm_bgp_asn      = var.pf_crc_ibm_bgp_asn
#   ibm_bgp_cer_cidr = var.pf_crc_ibm_bgp_cer_cidr
#   ibm_bgp_ibm_cidr = var.pf_crc_ibm_bgp_ibm_cidr
#   pop              = var.pf_crc_pop4
#   zone             = var.pf_crc_zone4
#   maybe_nat        = var.pf_crc_maybe_nat
#   speed            = var.pf_crc_speed
# }

# resource "packetfabric_cloud_router_connection_oracle" "crc_6" {
#   provider     = packetfabric
#   description  = "${var.tag_name}-${random_pet.name.id}-${var.pf_crc_pop5}"
#   circuit_id   = packetfabric_cloud_router.cr.id
#   account_uuid = var.pf_account_uuid
#   region       = var.pf_crc_oracle_region
#   vc_ocid      = var.pf_crc_oracle_vc_ocid
#   pop          = var.pf_crc_pop5
#   zone         = var.pf_crc_zone5
#   maybe_nat    = var.pf_crc_maybe_nat
# }

# resource "packetfabric_cloud_router_connection_port" "crc_7" {
#   provider        = packetfabric
#   description     = "${var.tag_name}-${random_pet.name.id}"
#   circuit_id      = packetfabric_cloud_router.cr.id
#   account_uuid    = var.pf_account_uuid
#   port_circuit_id = packetfabric_port.port_1a.id
#   vlan            = var.pf_crc_vlan
#   speed           = var.pf_crc_speed
#   is_public       = var.pf_crc_is_public
#   maybe_nat       = var.pf_crc_maybe_nat
# }

# data "packetfabric_cloud_router_connections" "all_crc" {
#   provider   = packetfabric
#   circuit_id = packetfabric_cloud_router.cr.id
# }
# output "packetfabric_cloud_router_connections" {
#   value = data.packetfabric_cloud_router_connections.all_crc
# }