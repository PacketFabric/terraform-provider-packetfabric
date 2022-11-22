terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.4.2"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.23.0"
    }
  }
}

provider "packetfabric" {}

# Define default profile
provider "aws" {
  region = var.aws_region1
}
# Profile for Region1
provider "aws" {
  region = var.aws_region1
  alias  = "region1"
}
# Profile for Region2
provider "aws" {
  region = var.aws_region2
  alias  = "region2"
}
