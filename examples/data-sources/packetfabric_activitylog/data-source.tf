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


data "packetfabric_activitylog" "current" {
  filter {
    user = "alice"
  }
}

output "my-activity-logs" {
  value = packetfabric_activitylog.current
}