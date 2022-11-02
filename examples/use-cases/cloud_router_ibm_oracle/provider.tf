terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.4.0"
    }
    ibm = {
      source  = "IBM-Cloud/ibm"
      version = ">= 1.46.0"
    }
    oci = {
      source  = "oracle/oci"
      version = ">= 4.88.1"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

provider "ibm" {
  region                = var.ibm_region1
  ibmcloud_api_key      = var.ibmcloud_api_key
  iaas_classic_username = var.iaas_classic_username
  iaas_classic_api_key  = var.iaas_classic_api_key
}

provider "oci" {
  region       = var.oracle_region1
  auth         = "APIKey"
  tenancy_ocid = var.tenancy_ocid
  user_ocid    = var.user_ocid
  private_key  = var.private_key
  #private_key_password = var.private_key_password
  fingerprint = var.fingerprint
}
