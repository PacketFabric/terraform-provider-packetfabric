terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 1.2.0"
    }
    google = {
      source  = "hashicorp/google"
      version = ">= 4.56.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.58.0"
    }
  }
}

provider "packetfabric" {}

# Make sure you enabled Compute Engine API
provider "google" {
  project = var.gcp_project_id
  # use GOOGLE_CREDENTIALS environment variable
  region = var.gcp_region1
  zone   = var.gcp_zone1
}

provider "aws" {
  region = var.aws_region1
  # use AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables
}
