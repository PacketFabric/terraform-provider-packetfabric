terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.3.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.58.0"
    }
    oci = {
      source  = "oracle/oci"
      version = ">= 4.111.0"
    }
  }
}

provider "packetfabric" {}

provider "aws" {
  region = var.aws_region1
  # use AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables
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
