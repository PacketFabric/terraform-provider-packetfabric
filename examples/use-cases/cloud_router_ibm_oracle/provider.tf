terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.0.4"
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

provider "packetfabric" {}

provider "ibm" {
  region = var.ibm_region1
  # use PF_IBM_ACCOUNT_ID environment variable
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
