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
  provider          = packetfabric
  name              = "%s"
  account_uuid      = "%s"
  asn               = %v
  capacity          = "%s"
  regions           = ["%s", "%s"]
  subscription_term = %v
}`

// Resource: packetfabric_cloud_router_connection_aws
const RResourceCloudRouterConnectionAws = `resource "packetfabric_cloud_router_connection_aws" "%s" {
  provider          = packetfabric
  circuit_id        = %s.id
  aws_account_id    = "%s"
  account_uuid      = "%s"
  description       = "%s"
  pop               = "%s"
  zone              = "%s"
  subscription_term = %v
  speed             = "%s"
}`

// Resource: packetfabric_cloud_router_bgp_session
const RResourceCloudRouterBgpSession = `resource "packetfabric_cloud_router_bgp_session" "%s" {
	provider       = packetfabric
	circuit_id     = %s.id
	connection_id  = %s.id
	disabled       = %s
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

// Resource: packetfabric_cloud_router_quick_connect
const RResourceCloudRouterQuickConnect = `resource "packetfabric_cloud_router_quick_connect" "%s" {
  provider              = packetfabric
  cr_circuit_id         = %s.id
  connection_circuit_id = %s.id
  service_uuid          = "%s"
  subscription_term     = %v
  return_filters {
    prefix     = "%s"
    match_type = "%s"
  }
  return_filters {
    prefix     = "%s"
    match_type = "%s"
  }
}`

// Resource: packetfabric_cloud_provider_credential_aws
const RResourceCloudProviderCredentialAws = `resource "packetfabric_cloud_provider_credential_aws" "%s" {
  provider       = packetfabric
  description    = "%s"
  aws_access_key = "%s"
  aws_secret_key = "%s"
}`

// Resource: packetfabric_cloud_provider_credential_google
const RResourceCloudProviderCredentialGoogle = `resource "packetfabric_cloud_provider_credential_google" "%s" {
  provider               = packetfabric
  description            = "%s"
}`

// Resource: packetfabric_cloud_router_connection_azure
const RResourceCloudRouterConnectionAzure = `provider "azurerm" {
  features {
    resource_group {
      prevent_deletion_if_contains_resources = false
    }
  }
}
resource "azurerm_resource_group" "resource_group1" {
  name     = "terraform-test-acc-azure-rg1-%s" 
  location = "%s"
}
resource "azurerm_virtual_network" "virtual_network1" {
  name                = "terraform-test-acc-azure-vnet1"
  location            = "%s"
  resource_group_name = azurerm_resource_group.resource_group1.name
  address_space       = ["%s"]
  tags = {
    environment = "terraform-test-acc-azure1"
  }
}
resource "azurerm_subnet" "subnet1" {
  name                 = "terraform-test-acc-azure-subnet1"
  address_prefixes     = ["%s"]
  resource_group_name  = azurerm_resource_group.resource_group1.name
  virtual_network_name = azurerm_virtual_network.virtual_network1.name
}
resource "azurerm_express_route_circuit" "azure_express_route1" {
  name                  = "terraform-test-acc-azure-express-route1"
  resource_group_name   = azurerm_resource_group.resource_group1.name
  location              = "%s"
  peering_location      = "%s"
  service_provider_name = "%s"
  bandwidth_in_mbps     = %v
  sku {
    tier   = "%s"
    family = "%s"
  }
  tags = {
    environment = "terraform-test-acc-azure1"
  }
}
resource "packetfabric_cloud_router_connection_azure" "%s" {
  provider          = packetfabric
  circuit_id        = %s.id
  account_uuid      = "%s"
  description       = "%s"
  speed             = "%s"
  azure_service_key = azurerm_express_route_circuit.azure_express_route1.service_key
  is_public         = %v
  subscription_term = %v
}`

// Resource: packetfabric_cloud_router_connection_google
const RResourceCloudRouterConnectionGoogle = `variable "gcp_project_id" {
  type = string
}
resource "google_compute_router" "google_router1" {
  provider = google
  project  = var.gcp_project_id # set in Env
  region   = "%s"
  name     = "terraform-test-acc-google-router1"
  network  = "%s"
  bgp {
    asn               = 16550
    advertise_mode    = "DEFAULT"
  }
}
resource "google_compute_interconnect_attachment" "google_interconnect1" {
  provider                 = google
  project                  = var.gcp_project_id # set in Env
  name                     = "terraform-test-acc-google-interconnect1"
  region                   = "%s"
  description              = "terraform Test ACC Interconnect to PacketFabric Network"
  type                     = "PARTNER"
  edge_availability_domain = "%s"
  admin_enabled            = true
  router                   = google_compute_router.google_router1.id
}
resource "packetfabric_cloud_router_connection_google" "%s" {
  provider                    = packetfabric
  circuit_id                  = %s.id
  account_uuid                = "%s"
  description                 = "%s"
  google_pairing_key          = google_compute_interconnect_attachment.google_interconnect1.pairing_key
  google_vlan_attachment_name = google_compute_interconnect_attachment.google_interconnect1.name
  pop                         = "%s"
  speed                       = "%s"
  subscription_term           = %v
}`

// Resource: packetfabric_cloud_router_connection_ibm
const RResourceCloudRouterConnectionIbm = `resource "packetfabric_cloud_router_connection_ibm" "%s" {
  provider          = packetfabric
  circuit_id        = %s.id
  account_uuid      = "%s"
  description       = "%s"
  pop               = "%s"
  zone              = "%s"
  speed             = "%s"
  ibm_bgp_asn       = %v
  subscription_term = %v
}
resource "time_sleep" "wait_ibm_connection1" {
  create_duration = "3m"
}
provider "ibm" {
  region = "%s"
}
data "ibm_dl_gateway" "current1" {
  provider   = ibm
  name       = "%s"
  depends_on = [time_sleep.wait_ibm_connection1]
}
variable "ibm_resource_group" {
  type        = string
}
data "ibm_resource_group" "existing_rg1" {
  provider   = ibm
  name       = var.ibm_resource_group # set in Env
}
resource "ibm_dl_gateway_action" "confirmation1" {
  provider       = ibm
  gateway        = data.ibm_dl_gateway.current1.id
  resource_group = data.ibm_resource_group.existing_rg1.id
  action         = "create_gateway_approve"
  global         = true
  metered        = true
  bgp_asn        = %v
  default_export_route_filter = "permit"
  default_import_route_filter = "permit"
  speed_mbps     = %v
  provisioner "local-exec" {
    when    = destroy
    command = "sleep 30"
  }
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
  phase2_lifetime              = %v
  shared_key                   = "%s"
  subscription_term            = %v
}`

// Resource: packetfabric_cloud_router_connection_oracle
const RResourceCloudRouterConnectionOracle = `variable "parent_compartment_id" {
  type        = string
}
variable "fingerprint" {
  type        = string
  sensitive   = true
}
variable "private_key" {
  type        = string
  sensitive   = true
}
variable "tenancy_ocid" {
  type        = string
  sensitive   = true
}
variable "user_ocid" {
  type        = string
  sensitive   = true
}
variable "pf_cs_oracle_drg_ocid" {
  type        = string
}
variable "pf_cs_oracle_region" {
  type        = string
}
provider "oci" {
  region       = var.pf_cs_oracle_region
  auth         = "APIKey"
  tenancy_ocid = var.tenancy_ocid
  user_ocid    = var.user_ocid
  private_key  = replace("${var.private_key}", "\\n", "\n")
  fingerprint  = var.fingerprint
}
data "oci_core_fast_connect_provider_services" "packetfabric_provider1" {
  provider = oci
  compartment_id = var.parent_compartment_id
  filter {
    name   = "provider_name"
    values = ["%s"]
  }
}
resource "oci_core_virtual_circuit" "fast_connect1" {
  provider = oci
  compartment_id       = var.parent_compartment_id
  display_name         = "terraform-test-acc-oracle-fastconnect1"
  region               = var.pf_cs_oracle_region
  type                 = "PRIVATE"
  gateway_id           = var.pf_cs_oracle_drg_ocid
  bandwidth_shape_name = "%s"
  customer_asn         = %v
  ip_mtu               = "MTU_1500"
  is_bfd_enabled       = false
  cross_connect_mappings {
    bgp_md5auth_key         = "%s"
    customer_bgp_peering_ip = "%s"
    oracle_bgp_peering_ip   = "%s"
  }
  provider_service_id = data.oci_core_fast_connect_provider_services.packetfabric_provider1.fast_connect_provider_services.0.id
}
resource "packetfabric_cloud_router_connection_oracle" "%s" {
  provider          = packetfabric
  circuit_id        = %s.id
  account_uuid      = "%s"
  description       = "%s"
  pop               = "%s"
  zone              = "%s"
  vc_ocid           = oci_core_virtual_circuit.fast_connect1.id
  region            = var.pf_cs_oracle_region
  subscription_term = %v
}`

// Resource: packetfabric_cloud_router_connection_port
const RResourceCloudRouterConnectionPort = `resource "packetfabric_cloud_router_connection_port" "%s" {
  provider          = packetfabric
  description       = "%s"
  circuit_id        = %s.id
  port_circuit_id   = %s.id
  speed             = "%s"
  vlan              = %v
  subscription_term = %v
}`

// Resource: packetfabric_cs_aws_dedicated_connection
const RResourceCSAwsDedicatedConnection = `resource "packetfabric_cs_aws_dedicated_connection" "%s" {
  provider          = packetfabric
  aws_region        = "%s"
  description       = "%s"
  pop               = "%s"
  zone              = "%s"
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
  zone            = "%s"
  speed           = "%s"
  vlan            = %v
}`

// Resource: packetfabric_cs_azure_dedicated_connection
const RResourceCSAzureDedicatedConnection = `resource "packetfabric_cs_azure_dedicated_connection" "%s" {
  provider          = packetfabric
  description       = "%s"
  pop               = "%s"
  zone              = "%s"
  subscription_term = %v
  service_class     = "%s"
  encapsulation     = "%s"
  port_category     = "%s"
  speed             = "%s"
}`

// Resource: packetfabric_cs_azure_hosted_connection
const RResourceCSAzureHostedConnection = `provider "azurerm" {
  features {
    resource_group {
      prevent_deletion_if_contains_resources = false
    }
  }
}
resource "azurerm_resource_group" "resource_group2" {
  name     = "terraform-test-acc-azure-rg2"
  location = "%s"
}
resource "azurerm_virtual_network" "virtual_network2" {
  name                = "terraform-test-acc-azure-vnet2"
  location            = "%s"
  resource_group_name = azurerm_resource_group.resource_group2.name
  address_space       = ["%s"]
  tags = {
    environment = "terraform-test-acc-azure2"
  }
}
resource "azurerm_subnet" "subnet" {
  name                 = "terraform-test-acc-azure-subnet2"
  address_prefixes     = ["%s"]
  resource_group_name  = azurerm_resource_group.resource_group2.name
  virtual_network_name = azurerm_virtual_network.virtual_network2.name
}
resource "azurerm_express_route_circuit" "azure_express_route2" {
  name                  = "terraform-test-acc-azure-express-route2"
  resource_group_name   = azurerm_resource_group.resource_group2.name
  location              = "%s"
  peering_location      = "%s"
  service_provider_name = "%s"
  bandwidth_in_mbps     = %v
  sku {
    tier   = "%s"
    family = "%s"
  }
  tags = {
    environment = "terraform-test-acc-azure2"
  }
}
resource "packetfabric_cs_azure_hosted_connection" "%s" {
  provider          = packetfabric
  port              = %s.id
  account_uuid      = "%s"
  description       = "%s"
  azure_service_key = azurerm_express_route_circuit.azure_express_route2.service_key
  speed             = "%s"
  vlan_private      = %v
}`

// Resource: packetfabric_cs_google_dedicated_connection
const RResourceCSGoogleDedicatedConnection = `resource "packetfabric_cs_google_dedicated_connection" "%s" {
  provider          = packetfabric
  description       = "%s"
  pop               = "%s"
  zone              = "%s"
  subscription_term = %v
  service_class     = "%s"
  autoneg           = %t
  speed             = "%s"
}`

// Resource: packetfabric_cs_google_hosted_connection
const RResourceCSGoogleHostedConnection = `variable "gcp_project_id" {
  type = string
}
resource "google_compute_router" "google_router2" {
  provider = google
  project  = var.gcp_project_id # set in Env
  region   = "%s"
  name     = "terraform-test-acc-google-router2"
  network  = "%s"
  bgp {
    asn               = 16550
    advertise_mode    = "DEFAULT"
  }
}
resource "google_compute_interconnect_attachment" "google_interconnect2" {
  provider                 = google
  project                  = var.gcp_project_id # set in Env
  name                     = "terraform-test-acc-google-interconnect2"
  region                   = "%s"
  description              = "terraform Test ACC Interconnect to PacketFabric Network"
  type                     = "PARTNER"
  edge_availability_domain = "%s"
  admin_enabled            = true
  router                   = google_compute_router.google_router2.id
}
resource "packetfabric_cs_google_hosted_connection" "%s" {
  provider                    = packetfabric
  port                        = %s.id
  google_pairing_key          = google_compute_interconnect_attachment.google_interconnect2.pairing_key
  google_vlan_attachment_name = google_compute_interconnect_attachment.google_interconnect2.name
  account_uuid    = "%s"
  description     = "%s"
  pop             = "%s"
  speed           = "%s"
  vlan            = %v
}`

// Resource: packetfabric_cs_ibm_hosted_connection
const RResourceCSIbmHostedConnection = `resource "packetfabric_cs_ibm_hosted_connection" "%s" {
  provider     = packetfabric
  port         = %s.id
  account_uuid = "%s"
  description  = "%s"
  pop          = "%s"
  zone         = "%s"
  speed        = "%s"
  vlan         = %v
  ibm_bgp_asn  = %v
}
resource "time_sleep" "wait_ibm_connection2" {
  create_duration = "5m"
}
provider "ibm" {
  region = "%s"
}
data "ibm_dl_gateway" "current2" {
  provider   = ibm
  name       = "%s"
  depends_on = [time_sleep.wait_ibm_connection2]
}
variable "ibm_resource_group" {
  type        = string
}
data "ibm_resource_group" "existing_rg2" {
  provider   = ibm
  name       = var.ibm_resource_group # set in Env
}
resource "ibm_dl_gateway_action" "confirmation2" {
  provider       = ibm
  gateway        = data.ibm_dl_gateway.current2.id
  resource_group = data.ibm_resource_group.existing_rg2.id
  action         = "create_gateway_approve"
  global         = true
  metered        = true
  bgp_asn        = %v
  default_export_route_filter = "permit"
  default_import_route_filter = "permit"
  speed_mbps     = %v
  provisioner "local-exec" {
    when    = destroy
    command = "sleep 30"
  }
}`

// Resource: packetfabric_cs_oracle_hosted_connection
const RResourceCSOracleHostedConnection = `variable "parent_compartment_id" {
  type        = string
}
variable "fingerprint" {
  type        = string
  sensitive   = true
}
variable "private_key" {
  type        = string
  sensitive   = true
}
variable "tenancy_ocid" {
  type        = string
  sensitive   = true
}
variable "user_ocid" {
  type        = string
  sensitive   = true
}
variable "pf_cs_oracle_drg_ocid" {
  type        = string
}
provider "oci" {
  region       = "%s"
  auth         = "APIKey"
  tenancy_ocid = var.tenancy_ocid
  user_ocid    = var.user_ocid
  private_key  = replace("${var.private_key}", "\\n", "\n")
  fingerprint  = var.fingerprint
}
data "oci_core_fast_connect_provider_services" "packetfabric_provider2" {
  provider = oci
  compartment_id = var.parent_compartment_id
  filter {
    name   = "provider_name"
    values = ["%s"]
  }
}
resource "oci_core_virtual_circuit" "fast_connect2" {
  provider = oci
  compartment_id       = var.parent_compartment_id
  display_name         = "terraform-test-acc-oracle-fastconnect2"
  region               = "%s"
  type                 = "PRIVATE"
  gateway_id           = var.pf_cs_oracle_drg_ocid
  bandwidth_shape_name = "%s"
  customer_asn         = %v
  ip_mtu               = "MTU_1500"
  is_bfd_enabled       = false
  cross_connect_mappings {
    bgp_md5auth_key         = "%s"
    customer_bgp_peering_ip = "%s"
    oracle_bgp_peering_ip   = "%s"
  }
  provider_service_id = data.oci_core_fast_connect_provider_services.packetfabric_provider2.fast_connect_provider_services.0.id
}
resource "packetfabric_cs_oracle_hosted_connection" "%s" {
  provider     = packetfabric
  port         = %s.id
  account_uuid = "%s"
  description  = "%s"
  pop          = "%s"
  zone         = "%s"
  vlan         = %v
  vc_ocid      = oci_core_virtual_circuit.fast_connect2.id
  region       = "%s"
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
  members     = [%s.id]
  pop         = "%s"
}
`

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
  document_uuid = %s.id
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
}`

// Resource: packetfabric_port
const RResourcePort = `resource "packetfabric_port" "%s" {
  provider          = packetfabric
  description       = "%s"
  media             = "%s"
  pop               = "%s"
  zone              = "%s"
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

const RResourceIpamPrefixCommon = `
	admin_contact_uuid = "%s"
	tech_contact_uuid  = "%s"

	ipj_details {
		currently_used_prefixes {
			prefix        = "128.192.1.0/24"
			ips_in_use    = 33
			description   = "Optional description"
			isp_name      = "Optional ISP Name"
			will_renumber = true
		}
		planned_prefixes {
			prefix        = "8.8.8.0/24"
			description   = "Another optional description"
			location      = "Optional Location"
			usage_30d     = 2
			usage_3m      = 0
			usage_6m      = 2
			usage_1y      = 3
		}
		planned_prefixes {
			prefix        = "4.4.4.0/24"
			usage_30d     = 2
			usage_3m      = 0
			usage_6m      = 2
			usage_1y      = 3
		}
	}
}`

// Resource: packetfabric_ipam_prefix
const RResourceIpamPrefix = `resource "packetfabric_ipam_prefix" "%s" {
	length             = 33
    version            = 4
	bgp_region         = "Antarctica"
` + RResourceIpamPrefixCommon

// Resource: packetfabric_ipam_prefix_confirmation
const RResourceIpamPrefixConfirmation = `resource "packetfabric_ipam_prefix_confirmation" "%s" {
	prefix_uuid   = "%s"
` + RResourceIpamPrefixCommon

// Resource: packetfabric_outbound_cross_connect
const RResourceDocumentMSA = `resource "packetfabric_document" "%s" {
  provider        = packetfabric
  document        = "%s"
  type            = "msa"
  description     = "%s"
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

// Datasource: packetfabric_locations_markets
const DDataSourceLocationsMarkets = `data "packetfabric_locations_markets" "%s" {
  provider = packetfabric
}`

// Datasource: packetfabric_activitylogs
const DDatasourceActivityLogs = `data "packetfabric_activitylogs" "%s" {
  provider = packetfabric
}`

// Datasource: packetfabric_billing
const DDatasourceBilling = `data "packetfabric_billing" "%s" {
  circuit_id        = %s.id
}`

// Datasource: packetfabric_ports
const DDataSourcePorts = `data "packetfabric_ports" "%s" {
  provider   = packetfabric
  depends_on = [%s]
}`

// Datasource: packetfabric_port_vlans
const DDataSourcePortVlans = `data "packetfabric_port_vlans" "%s" {
  provider        = packetfabric
  port_circuit_id = %s.id
}`

// Datasource: packetfabric_port_device_info
const DDataSourcePortDeviceInfo = `data "packetfabric_port_device_info" "%s" {
  provider          = packetfabric
  port_circuit_id   = %s.id
}`

// Datasource: packetfabric_port_router_logs
const DDataSourcePortRouterLogs = `data "packetfabric_port_router_logs" "%s" {
  provider        = packetfabric
  port_circuit_id = %s.id
  time_from       = "%s"
  time_to         = "%s"
}`

// Datasource: packetfabric_outbound_cross_connects
const DDatasourceOutboundCrossConnects = `data "packetfabric_outbound_cross_connects" "%s" {
  provider   = packetfabric
  depends_on = [%s]
}`

// Datasource: packetfabric_link_aggregation_group
const DDatasourceLinkAggregationGroups = `data "packetfabric_link_aggregation_group" "%s" {
  provider       = packetfabric
  lag_circuit_id = %s.id
}`

// Datasource: packetfabric_point_to_points
const DDatasourcePointToPoints = `data "packetfabric_point_to_points" "%s" {
  provider = packetfabric
  depends_on = [%s]
}`

// Datasource: packetfabric_cs_aws_hosted_connection
const DDatasourceCsAwsHostedConn = `data "packetfabric_cs_aws_hosted_connection" "%s" {
  provider          = packetfabric
  cloud_circuit_id  = %s.id
}`

// Datasource: packetfabric_cs_dedicated_connections
const DDatasourceDedicatedConns = `data "packetfabric_cs_dedicated_connections" "%s" {
  provider   = packetfabric
  depends_on = [%s]
}`

// Datasource: packetfabric_cloud_router_connection_ipsec
const DDatasourceCloudRouterConnectionIpsec = `data "packetfabric_cloud_router_connection_ipsec" "%s" {
  provider   = packetfabric
  circuit_id = %s.id
}`

// Datasource: packetfabric_cloud_router_connection
const DDatasourceCloudRouterConnection = `data "packetfabric_cloud_router_connection" "%s" {
  circuit_id     = %s.id
  connection_id  = %s.id
}`

// Datasource: packetfabric_cloud_router_connections
const DDatasourceCloudRouterConnections = `data "packetfabric_cloud_router_connections" "%s" {
  circuit_id = %s.id
  depends_on = [%s]
}`

// Datasource: packetfabric_cloud_router_bgp_session
const DDatasourceBgpSession = `data "packetfabric_cloud_router_bgp_session" "%s" {
  provider       = packetfabric
  circuit_id     = %s.id
  connection_id  = %s.id
  depends_on = [%s]
}`
