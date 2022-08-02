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

resource "packetfabric_cloud_services_azr_backbone" "new" {
  provider = packetfabric
  description = "my-azure-backbone"
  epl =  false
  interface_a {
    port_circuit_id = "PF-AP-XYZ1-1234"
    untagged = false
    vlan = 4
  }
  interface_z {
    port_circuit_id = "PF-AP-ABC1-4321"
    untagged = false
    vlan = 7    
  }
  bandwidth {
    account_uuid = var.pf_account_uuid
    speed = "1gps"
    subscription_term = "1"
  }
}