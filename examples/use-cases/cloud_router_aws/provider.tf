terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.3.2"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.23.0"
    }
  }
}

provider "packetfabric" {
  host  = var.pf_api_server
  token = var.pf_api_key
}

# Define default profile
provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region1
}
# Profile for Region1
provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region1
  alias      = "region1"
}
# Profile for Region2
provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region2
  alias      = "region2"
}

# Create random name to use to name objects
resource "random_pet" "name" {}
