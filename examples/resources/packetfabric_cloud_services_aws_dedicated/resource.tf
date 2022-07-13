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

data "packetfabric_cloud_services_aws_dedicated_conn" "current" {
  provider = packetfabric
}

resource "packetfabric_cloud_services_aws_dedicated" "new" {
  provider = packetfabric
  aws_region = var.pf_cs_aws_d_region
  account_uuid = var.pf_aws_account_id
  description = var.pf_cs_aws_d_descr
  zone = var.pf_cs_aws_d_avzone
  pop = var.pf_cs_aws_d_pop
  subscription_term = var.pf_cs_aws_d_subterm
  service_class = var.pf_cs_aws_d_srvclass
  autoneg = var.pf_cs_aws_d_autoneg
  speed = var.pf_cs_aws_d_speed
  should_create_lag = var.pf_cs_aws_d_createlag
}

output "packetfabric_cloud_services_aws_dedicated_conn" {
  value = data.packetfabric_cloud_services_aws_dedicated_conn.current
}
