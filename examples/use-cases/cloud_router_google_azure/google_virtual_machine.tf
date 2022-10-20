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