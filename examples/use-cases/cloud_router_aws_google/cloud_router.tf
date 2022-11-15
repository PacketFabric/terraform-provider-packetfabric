# From the PacketFabric side: Create a cloud router
resource "packetfabric_cloud_router" "cr" {
  provider     = packetfabric
  name         = "${var.tag_name}-${random_pet.name.id}"
  account_uuid = var.pf_account_uuid
  asn          = var.pf_cr_asn
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

output "packetfabric_cloud_router" {
  value = packetfabric_cloud_router.cr
}
