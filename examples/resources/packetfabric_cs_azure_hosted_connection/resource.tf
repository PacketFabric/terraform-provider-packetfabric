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

resource "packetfabric_cs_azure_hosted_connection" "new" {
  provider = packetfabric
  description = "my-azure-hosted-express-connect"
  account_uuid = var.pf_account_uuid
  azure_service_key = var.pf_azr_srvc_key
  port = "PF-AP-XYZ1-1234"
  speed = "50Mbps"
  src_svlan = "100"
  vlan_microsoft = "6"
  vlan_private = "6"
}
