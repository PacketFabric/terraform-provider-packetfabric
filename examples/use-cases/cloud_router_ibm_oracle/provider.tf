terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.6.0"
    }
    ibm = {
      source  = "IBM-Cloud/ibm"
      version = ">= 1.53.0"
    }
    oci = {
      source  = "oracle/oci"
      version = ">= 4.121.0"
    }
  }
}

provider "packetfabric" {}

provider "ibm" {
  region = var.ibm_region1
  # use PF_IBM_ACCOUNT_ID, IC_API_KEY, IAAS_CLASSIC_USERNAME, IAAS_CLASSIC_API_KEY environment variables
}

provider "oci" {
  region = var.oracle_region1
  auth   = "APIKey"
  # you can use TF_VAR_ to set below variable https://developer.hashicorp.com/terraform/language/values/variables
  tenancy_ocid = var.tenancy_ocid
  user_ocid    = var.user_ocid
  private_key  = replace("${var.private_key}", "\\n", "\n") # replace() may not be needed depending on the shell
  # private_key_password = var.private_key_password # Passphrase used for the key, if it is encrypted
  fingerprint = var.fingerprint
}
