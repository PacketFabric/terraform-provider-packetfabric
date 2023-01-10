# Create a Marketplace Service type quick-connect
resource "packetfabric_marketplace_service" "marketplace_quick_connect" {
  provider                = packetfabric
  name                    = var.pf_name
  description             = var.pf_description
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

# Create a Marketplace Service type port
resource "packetfabric_marketplace_service" "marketplace_port" {
  provider     = packetfabric
  name         = var.pf_name
  description  = var.pf_description
  service_type = "port-service"
  sku          = var.pf_sku
  categories   = var.pf_categories
  published    = var.pf_published
  locations    = var.pf_locations
}
output "packetfabric_marketplace_service_port" {
  value = packetfabric_marketplace_service.marketplace_port
}