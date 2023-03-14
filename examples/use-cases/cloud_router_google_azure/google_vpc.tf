resource "google_compute_network" "vpc_1" {
  provider                = google
  name                    = "${var.tag_name}-${random_pet.name.id}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet_1" {
  provider      = google
  name          = "${var.tag_name}-${random_pet.name.id}"
  ip_cidr_range = var.gcp_subnet_cidr1
  region        = var.gcp_region1
  network       = google_compute_network.vpc_1.id
}
# output "google_compute_network" {
#   value = google_compute_network.vpc_1
# }