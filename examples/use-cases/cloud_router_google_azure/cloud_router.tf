# From the PacketFabric side: Create a cloud router
resource "packetfabric_cloud_router" "cr" {
  provider = packetfabric
  name     = "${var.resource_name}-${random_pet.name.id}"
  labels   = var.pf_labels
  asn      = var.pf_cr_asn
  capacity = var.pf_cr_capacity
  regions  = var.pf_cr_regions
}
# output "packetfabric_cloud_router" {
#   value = packetfabric_cloud_router.cr
# }
