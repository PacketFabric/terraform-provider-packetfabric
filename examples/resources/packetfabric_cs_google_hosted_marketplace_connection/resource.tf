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

resource "packetfabric_cs_google_hosted_marketplace_connection" "new" {
  provider = packetfabric
  routing_id = "PF-1RI-1234"
  market = "ATL"
  pop = "DAL1"
  description = "my-gcp-hosted-mkt-conn"
  google_pairing_key = var.pf_gcp_pair_key
  account_uuid = var.pf_account_uuid
  speed = "100Mbps"
}