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
  }
}

provider "packetfabric" {}

# Define default profile
provider "aws" {
  region = var.aws_region1
  # use AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables
}
# Profile for Region1
provider "aws" {
  region = var.aws_region1
  alias  = "region1"
  # use AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables
}
# Profile for Region2
provider "aws" {
  region = var.aws_region2
  # use AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables
  alias = "region2"
}
