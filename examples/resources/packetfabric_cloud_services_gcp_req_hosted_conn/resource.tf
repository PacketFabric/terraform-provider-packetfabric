terraform {
  required_providers {
    packetfabric = {
      source  = "elx1/packetfabric/packetfabric"
      version = "~> 0.0.1"
    }
  }
}
provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

resource "packetfabric_cloud_services_gcp_req_hosted_conn" "new" {
  provider = packetfabric
  description = "my-gcp-hosted-conn"
  account_uuid = var.pf_account_uuid
  port = "PF-AP-XYZ1-1234"
  speed = "50Mbps"
  googe_pairing_key = var.pf_gcp_pair_key
  google_vlan_attachment_name = "rrv-nms-dev-1a2b-vl1-us-west1-1"
  pop = "DAL1"
  vlan = 43
}