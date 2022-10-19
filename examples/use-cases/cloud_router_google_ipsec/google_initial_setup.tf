resource "google_compute_network" "vpc_1" {
  provider                = google
  name                    = "${var.tag_name}-${random_pet.name.id}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet_1" {
  provider      = google
  name          = "${var.tag_name}-${random_pet.name.id}"
  ip_cidr_range = var.google_subnet_cidr1
  region        = var.gcp_region1
  network       = google_compute_network.vpc_1.id
}

output "google_compute_network" {
  value = google_compute_network.vpc_1
}

resource "google_compute_firewall" "ssh-rule" {
  provider = google
  name     = "allow-icmp-ssh-http-locust-iperf"
  network  = google_compute_network.vpc_1.name
  allow {
    protocol = "icmp"
  }
  allow {
    protocol = "tcp"
    ports    = ["22", "80", "8089", "5001"]
  }
  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_instance" "vm_1" {
  provider     = google
  name         = "${var.tag_name}-${random_pet.name.id}-vm1"
  machine_type = "e2-micro"
  zone         = var.gcp_zone1
  tags         = ["${var.tag_name}-${random_pet.name.id}"]
  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2204-lts"
    }
  }
  network_interface {
    subnetwork = google_compute_subnetwork.subnet_1.name
    access_config {}
  }
  metadata_startup_script = file("./user-data-ubuntu.sh")
  metadata = {
    sshKeys = "ubuntu:${var.public_key}"
  }
}

data "google_compute_instance" "vm_1" {
  provider = google
  name     = "${var.tag_name}-${random_pet.name.id}-vm1"
  zone     = var.gcp_zone1
  depends_on = [
    google_compute_instance.vm_1
  ]
}

output "google_private_ip_vm_1" {
  description = "Private ip address for VM for Region 1"
  value       = data.google_compute_instance.vm_1.network_interface.0.network_ip
}

output "google_public_ip_vm_1" {
  description = "Public ip address for VM for Region 1 (ssh user: ubuntu)"
  value       = data.google_compute_instance.vm_1.network_interface.0.access_config.0.nat_ip
}

# # Verify Terraform gcloud module works in your environment
# module "gcloud_version" {
#   # https://registry.terraform.io/modules/terraform-google-modules/gcloud/google/latest
#   source  = "terraform-google-modules/gcloud/google"
#   version = "~> 2.0"
#   # when running locally with gcloud already installed
#   service_account_key_file = var.GOOGLE_CREDENTIALS
#   skip_download            = true
#   # when running in a CI/CD pipeline without glcoud installed
#   # use_tf_google_credentials_env_var = true
#   # skip_download                     = false

#   # https://cloud.google.com/sdk/gcloud/reference/compute/routers/update-bgp-peer
#   create_cmd_entrypoint = "gcloud"
#   create_cmd_body       = "version"

#   # no destroy needed
#   destroy_cmd_entrypoint = "echo"
#   destroy_cmd_body       = "skip"
# }