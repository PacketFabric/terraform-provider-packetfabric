terraform {
  required_providers {
    packetfabric = {
      source  = "packetfabric/packetfabric"
      version = "~> 0.0.0"
    }
  }
}

provider "packetfabric" {
  host = var.pf_api_server
  token = var.pf_api_key
}

data "packetfabric_cloud_services_gcp_dedicated_conn" "current" {
  provider = packetfabric
  filter {
  	cloud_circuit_id = "PF-AP-LAX1-1234"
  }
}

output "connection_info" {
  value = data.packetfabric_cloud_services_gcp_dedicated_conn.current
}