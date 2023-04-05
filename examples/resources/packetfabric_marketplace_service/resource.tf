# Create a Marketplace Service type quick-connect
resource "packetfabric_marketplace_service" "marketplace_quick_connect" {
  provider     = packetfabric
  name         = "hello world"
  description  = "hello world quick connect"
  service_type = "quick-connect-service"
  sku          = "001234567"
  categories = [
    "cloud-computing",
    "security",
    "web-hosting"
  ]
  published               = true
  cloud_router_circuit_id = "PF-L3-CUST-2839140"
  connection_circuit_ids = [
    "PF-L3-CON-2853999",
    "PF-L3-CON-2888830"
  ]
  route_set {
    description = "hello world route set"
    is_private  = false
    prefixes {
      prefix     = "185.56.153.165/32"
      match_type = "orlonger"
    }
    prefixes {
      prefix     = "185.56.153.166/32"
      match_type = "orlonger"
    }
  }
}

# Create a Marketplace Service type port
resource "packetfabric_marketplace_service" "marketplace_port" {
  provider     = packetfabric
  name         = "hello world port"
  description  = "hello world"
  service_type = "port-service"
  sku          = "001234567"
  categories = [
    "cloud-computing",
    "security",
    "web-hosting"
  ]
  published = true
  locations = [
    "PDX1",
    "SFO2"
  ]
}