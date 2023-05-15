package testutil

// Patterns:
// Resource schema for required fields only
// - const R<resource_name> = `...`
// Resouce schema for required + optional fields
// - const O<resource_name> = `...`

// Begin of resources templates for required fields only

// Resource: packetfabric_backbone_virtual_circuit
const RResourceBackboneVirtualCircuitVlan = `resource "packetfabric_backbone_virtual_circuit" "%s" {
  provider    = packetfabric
  description = "%s"
  epl         = %t
  interface_a {
    port_circuit_id = %s.id
    vlan            = %v
  }
  interface_z {
    port_circuit_id = %s.id
    vlan            = %v
  }
  bandwidth {
    longhaul_type     = "%s"
    speed             = "%s"
    subscription_term = %v
  }
}`

// Resource: packetfabric_backbone_virtual_circuit_marketplace
const RResourceBackboneVirtualCircuitMarketplace = `resource "packetfabric_backbone_virtual_circuit_marketplace" "%s" {
  provider    = packetfabric
  routing_id  = "%s"
  market      = "%s"
  interface {
    port_circuit_id   = %s.id
    untagged          = %t
    vlan              = %v
  }
  bandwidth {
    account_uuid      = "%s"
  }
}`

// Resource: packetfabric_backbone_virtual_circuit_speed_burst
const RResourceBackboneVirtualCircuitSpeedBurst = `resource "packetfabric_backbone_virtual_circuit_speed_burst" "%s" {
  provider      = packetfabric
  vc_circuit_id = %s.id
  speed         = "%s"
}`

// Resource: packetfabric_cloud_router
const RResourcePacketfabricCloudRouter = `resource "packetfabric_cloud_router" "%s" {
	provider      = packetfabric
	name          = "%s"
  account_uuid  = "%s"
  asn           = %v
	capacity      = "%s"
  regions       = ["%s", "%s"]
  }`

// Resource: packetfabric_cloud_router_connection_aws
const RResourceCloudRouterConnectionAws = `resource "packetfabric_cloud_router_connection_aws" "%s" {
  provider        = packetfabric
  circuit_id      = %s.id
  aws_account_id  = "%s"
  account_uuid    = "%s"
  description     = "%s"
  pop             = "%s"
  speed           = "%s"
}`

// Resource: packetfabric_cloud_router_bgp_session
const RResourceCloudRouterBgpSession = `resource "packetfabric_cloud_router_bgp_session" "%s" {
	provider       = packetfabric
	circuit_id     = %s.id
	connection_id  = %s.id
  remote_address = "%s"
  l3_address     = "%s"
	remote_asn     = %v
	prefixes {
		prefix = "%s"
		type   = "%s"
	}
	prefixes {
		prefix = "%s"
		type   = "%s"
	}
}`

// Resource: packetfabric_cloud_router_connection_azure
const RResourceCloudRouterConnectionAzure = `resource "packetfabric_cloud_router_connection_azure" "%s" {
  provider          = packetfabric
  description       = "%s"
  account_uuid      = "%s"
  circuit_id        = %s.id
  azure_service_key = "%s"
  speed             = "%s"
}`

// Resource: packetfabric_cloud_router_connection_google
const RResourceCloudRouterConnectionGoogle = `resource "packetfabric_cloud_router_connection_google" "%s" {
  provider                    = packetfabric
  description                 = "%s"
  circuit_id                  = %s.id
  google_pairing_key          = "%s"
  google_vlan_attachment_name = "%s"
  pop                         = "%s"
  speed                       = "%s"
}`

// Resource: packetfabric_cloud_router_connection_ibm
const RResourceCloudRouterConnectionIBM = `resource "packetfabric_cloud_router_connection_ibm" "%s" {
  provider    = packetfabric
  description = "%s"
  circuit_id  = %s.id
  ibm_bgp_asn = %v
  pop         = "%s"
  speed       = "%s"
}`

// Resource: packetfabric_cloud_router_connection_ipsec
const RResourceCloudRouterConnectionIpsec = `resource "packetfabric_cloud_router_connection_ipsec" "%s" {
  provider                     = packetfabric
  description                  = "%s"
  circuit_id                   = %s.id
  pop                          = "%s"
  speed                        = "%s"
  gateway_address              = "%s"
  ike_version                  = %v
  phase1_authentication_method = "%s"
  phase1_group                 = "%s"
  phase1_encryption_algo       = "%s"
  phase1_authentication_algo   = "%s"
  phase1_lifetime              = %v
  phase2_pfs_group             = "%s"
  phase2_encryption_algo       = "%s"
  phase2_authentication_algo   = "%s"
  phase2_lifetime              = %v
  shared_key                   = "%s"
}`

// Resource: packetfabric_cloud_router_connection_oracle
const RResourceCloudRouterconnectionOracle = `resource "packetfabric_cloud_router_connection_oracle" "%s" {
  provider    = packetfabric
  description = "%s"
  circuit_id  = %s.id
  region      = "%s"
  vc_ocid     = "%s"
  pop         = "%s"
}`

// Resource: packetfabric_cloud_router_connection_port
const RResourceCloudRouterConnectionPort = `resource "packetfabric_cloud_router_connection_port" "%s" {
  provider        = packetfabric
  description     = "%s"
  circuit_id      = %s.id
  port_circuit_id = %s.id
  speed           = "%s"
  vlan            = %v
}`

// Resource: packetfabric_cs_aws_dedicated_connection
const RResourceCSAwsDedicatedConnection = `resource "packetfabric_cs_aws_dedicated_connection" "%s" {
  provider          = packetfabric
  aws_region        = "%s"
  description       = "%s"
  pop               = "%s"
  subscription_term = %v
  service_class     = "%s"
  autoneg           = %t
  speed             = "%s"
}`

// Resource: packetfabric_cs_aws_hosted_connection
const RResourceCSAwsHostedConnection = `resource "packetfabric_cs_aws_hosted_connection" "%s" {
  provider       = packetfabric
  port            = %s.id
  aws_account_id  = "%s"
  account_uuid    = "%s"
  description     = "%s"
  pop             = "%s"
  speed           = "%s"
  vlan            = %v
}`

// Resource: packetfabric_cs_aws_hosted_marketplace_connection
const RResourceCSAwsHostedMarketplaceConnection = `resource "packetfabric_cs_aws_hosted_marketplace_connection" "%s" {
  provider    = packetfabric
  routing_id  = %s.id
  market      = "%s"
  speed       = "%s"
  pop         = "%s"
}`

// Resource: packetfabric_cs_azure_dedicated_connection
const RResourceCSAzureDedicatedConnection = `resource "packetfabric_cs_azure_dedicated_connection" "%s" {
  provider          = packetfabric
  description       = "%s"
  pop               = "%s"
  subscription_term = %v
  service_class     = "%s"
  encapsulation     = "%s"
  port_category     = "%s"
  speed             = "%s"
}`

// Resource: packetfabric_cs_azure_hosted_connection
const RResourceCSAzureHostedConnection = `resource "packetfabric_cs_azure_hosted_connection" "%s" {
  provider          = packetfabric
  description       = "%s"
  azure_service_key = "%s"
  port              = %s.id
  speed             = "%s"
  vlan_private      = %v
  vlan_microsoft    = %v
}`

// Resource: packetfabric_cs_azure_hosted_marketplace_connection
const RResourceCSAzureHostedMarketplaceConnection = `resource "packetfabric_cs_azure_hosted_marketplace_connection" "%s" {
  provider          = packetfabric
  description       = "%s"
  azure_service_key = "%s"
  routing_id        = %s.id
  market            = "%s"
  speed             = "%s"
}`

// Resource: packetfabric_cs_google_dedicated_connection
const RResourceCSGoogleDedicatedConnection = `resource "packetfabric_cs_google_dedicated_connection" "%s" {
  provider          = packetfabric
  description       = "%s"
  zone              = "%s"
  pop               = "%s"
  subscription_term = %v
  service_class     = "%s"
  autoneg           = %t
  speed             = "%s"
}`

// Resource: packetfabric_cs_google_hosted_connection
const RResourceCSGoogleHostedConnection = `resource "packetfabric_cs_google_hosted_connection" "%s" {
  provider                    = packetfabric
  description                 = "%s"
  port                        = %s.id
  speed                       = "%s"
  google_pairing_key          = "%s"
  google_vlan_attachment_name = "%s"
  pop                         = "%s"
  vlan                        = %v
}`

// Resource: packetfabric_cs_google_hosted_marketplace_connection
const RResourceCSGGoogleHostedMarketplaceConnection = `resource "packetfabric_cs_google_hosted_marketplace_connection" "%s" {
  provider                    = packetfabric
  description                 = "%s"
  google_pairing_key          = "%s"
  google_vlan_attachment_name = "%s"
  routing_id                  = %s.id
  market                      = "%s"
  speed                       = "%s"
  pop                         = "%s"

}`

// Resource: packetfabric_cs_ibm_hosted_connection
const RResourceCSIBMHostedConnection = `resource "packetfabric_cs_ibm_hosted_connection" "%s" {
  provider    = packetfabric
  ibm_bgp_asn = %v
  description = "%s"
  pop         = "%s"
  port        = %s.id
  vlan        = %v
  speed       = "%s"
}`

// Resource: packetfabric_cs_oracle_hosted_connection
const RResourceCSOracleHostedConnection = `resource "packetfabric_cs_oracle_hosted_connection" "%s" {
  provider    = packetfabric
  description = "%s"
  vc_ocid     = "%s"
  region      = "%s"
  port        = %s.id
  pop         = "%s"
  zone        = "%s"
  vlan        = %v
}`

// Resource: packetfabric_cs_oracle_hosted_marketplace_connection
const RResourceCSOracleHostedMarketplaceConnection = `resource "packetfabric_cs_oracle_hosted_marketplace_connection" "%s" {
  provider    = packetfabric
  description = "%s"
  vc_ocid     = "%s"
  region      = "%s"
  routing_id  = %s.id
  market      = "%s"
  pop         = "%s"
}`

// Resource: packetfabric_ix_virtual_circuit_marketplace
const RResourceIXVirtualCircuitMarketplace = `resource "packetfabric_ix_virtual_circuit_marketplace" "%s" {
  provider    = packetfabric
  description = "%s"
  routing_id  = %s.id
  market      = "%s"
  asn         = %v
  interface {
    port_circuit_id = %s.id
    untagged        = %t
    vlan            = %v
  }
  bandwidth {
    longhaul_type     = "%s"
    speed             = "%s"
    subscription_term = %v
  }
}`

// Resource: packetfabric_link_aggregation_group
const RResourceLinkAggregationGroup = `resource "packetfabric_link_aggregation_group" "%s" {
  provider    = packetfabric
  description = "%s"
  interval    = "%s"
  members     = ["%s", "%s"]
  pop         = "%s"
}`

// Resource: packetfabric_marketplace_service_port_accept_request
const RResourceMarketplaceServicePortAcceptRequest = `resource "packetfabric_marketplace_service_port_accept_request" "%s" {
  provider       = packetfabric
  type           = "%s"
  cloud_provider = "%s"
  interface {
    port_circuit_id = %s.id
    vlan            = %v
  }
  vc_request_uuid = "%s"
}`

// Resource: packetfabric_marketplace_service_port_reject_request
const RResourceMarketplaceServicePortRejectRequest = `resource "packetfabric_marketplace_service_port_reject_request" "%s" {
  provider        = packetfabric
  vc_request_uuid = "%s"
}`

// Resource: packetfabric_outbound_cross_connect
const RResourceOutboundCrossConnect = `resource "packetfabric_outbound_cross_connect" "%s" {
  provider      = packetfabric
  description   = "%s"
  document_uuid = "%s"
  port          = %s.id
  site          = "%s"
}`

// Resource: packetfabric_point_to_point
const RResourcePointToPoint = `resource "packetfabric_point_to_point" "%s" {
  provider          = packetfabric
  description       = "%s"
  speed             = "%s"
  media             = "%s"
  subscription_term = %v
  endpoints {
    pop     = "%s"
    zone    = "%s"
    autoneg = %t
  }
  endpoints {
    pop     = "%s"
    zone    = "%s"
    autoneg = %t
  }
}
resource "time_sleep" "wait_30_seconds_%s" {
  depends_on = [%s]
  destroy_duration = "30s"
}`

// Resource: packetfabric_port
const RResourcePort = `resource "packetfabric_port" "%s" {
  provider          = packetfabric
  description       = "%s"
  media             = "%s"
  pop               = "%s"
  speed             = "%s"
  subscription_term = %v
  enabled           = %t
}`

// Resource: packetfabric_port_loa
const RResourcePortLoa = `resource "packetfabric_port_loa" "%s" {
  provider          = packetfabric
  port_circuit_id   = %s.id
  loa_customer_name = "%s"
  destination_email = "%s"
}`

// End of resources templates for required fields only

// Datasource: packetfabric_locations_cloud
const DDataSourceLocationsCloud = `data "packetfabric_locations_cloud" "%s" {
  provider              = packetfabric
  cloud_provider        = "%s"
  cloud_connection_type = "%s"
}`

// Datasource: packetfabric_locations_port_availability
const DDataSourceLocationsPortAvailability = `data "packetfabric_locations_port_availability" "%s" {
  provider  = packetfabric
  pop       = "%s"
}`

// Datasource: packetfabric_locations
const DDatasourceLocations = `data "packetfabric_locations" "%s" {
  provider  = packetfabric
}`

// Datasource: packetfabric_locations_pop_zones
const DDatasourceLocationsPopZones = `data "packetfabric_locations_pop_zones" "%s" {
  provider = packetfabric
  pop      = "%s"
}`

// Datasource: packetfabric_locations_regions
const DDataSourceLocationsRegions = `data "packetfabric_locations_regions" "%s" {
  provider = packetfabric
}`

// Datasource: packetfabric_activitylog
const DDatasourceActivityLog = `data "packetfabric_activitylog" "%s" {
  provider = packetfabric
}`

// Datasource: packetfabric_locations_markets
const DDataSourceLocationsMarkets = `data "packetfabric_locations_markets" "%s" {
  provider = packetfabric
}`

// Datasource: packetfabric_ports
const DDataSourcePorts = `data "packetfabric_ports" "%s" {
  provider          = packetfabric
}`
