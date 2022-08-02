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

resource "packetfabric_cloud_services_azr_hosted_mkt_conn" "new" {
  provider = packetfabric
  routing_id = "PF-1RI-1234"
  market = "ATL"
  description = "my-azure-hosted-mkt-conn"
  azure_service_key = var.pf_azr_srvc_key
  account_uuid = var.pf_account_uuid
  zone = "A"
  speed = "100Mbps"
  service_uuid = "7138cc00-c611-4dec-a05e-change-me"
}