# From the Google side: Create a Google Cloud Router with ASN 16550.
resource "google_compute_router" "google_router_1" {
  provider = google
  name     = "${var.resource_name}-${random_pet.name.id}"
  network  = google_compute_network.vpc_1.id
  bgp {
    # You must select or create a Cloud Router with its Google ASN set to 16550. This is a Google requirement for all Partner Interconnects.
    asn            = var.gcp_side_asn1
    advertise_mode = "CUSTOM"
    advertised_ip_ranges {
      range = var.gcp_subnet_cidr1
    }
  }
}