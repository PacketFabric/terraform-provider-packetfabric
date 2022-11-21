terraform {
  required_providers {
    packetfabric = {
      source  = "PacketFabric/packetfabric"
      version = ">= 0.4.2"
    }
    google = {
      source  = "hashicorp/google"
      version = ">= 4.38.0"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.14.0"
    }
  }
}

provider "packetfabric" {}

# Make sure you enabled Compute Engine API
provider "google" {
  project = var.gcp_project_id
  # https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#credentials-1
  credentials = file(var.gcp_credentials_path) # or use environment variable called GOOGLE_CREDENTIALS
  region      = var.gcp_region1
  zone        = var.gcp_zone1
}

provider "azurerm" {
  features {
    resource_group {
      prevent_deletion_if_contains_resources = false
    }
  }
}
