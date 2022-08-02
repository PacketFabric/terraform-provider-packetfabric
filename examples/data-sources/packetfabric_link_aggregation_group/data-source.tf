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

resource "packetfabric_link_aggregation_group" "new" {
  provider = packetfabric
  description = "my-LAG"
  interval = "fast"
  members = ["PF-AE-LAX1-1234", "PF-AE-LAX1-5678"]
  pop = "DAL1"
}

data "packetfabric_link_aggregation_group" "current" {
  filter {
    id = packetfabric_link_aggregation_group.new.id
  }
}

output "my-lag-info" {
  value = packetfabric_link_aggregation_group.current
}