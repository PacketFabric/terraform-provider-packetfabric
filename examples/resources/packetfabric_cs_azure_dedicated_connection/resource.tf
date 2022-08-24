terraform {
  required_providers {
    packetfabric = {
      source  = "packetfabric/packetfabric"
      version = "~> 0.0.1"
    }
  }
}
provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

resource "packetfabric_cs_azure_dedicated_connection" "new" {
  provider = packetfabric
  description = "my-azure-dedicated-express-conn"
  account_uuid = var.pf_account_uuid
  encapsulation = "dot1q"
  pop = "DAL1"
  port_category = "primary"
  service_class = "longhaul"
  speed = "10Gbps"
  subscription_term = 1
  zone = "A"
}