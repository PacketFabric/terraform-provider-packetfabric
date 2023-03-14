# Sample to create a compartment and Virtual Network

# # create a compartment
# resource "oci_identity_compartment" "compartment_1" {
#   provider       = oci
#   compartment_id = var.parent_compartment_id
#   name           = "${var.tag_name}-${random_pet.name.id}"
#   description    = "Compartment demo 1"
#   enable_delete  = true
# }

# # create a Virtual Network
# resource "oci_core_vcn" "subnet_1" {
#   provider       = oci
#   compartment_id = oci_identity_compartment.compartment_1.id
#   display_name   = "${var.tag_name}-${random_pet.name.id}"
#   cidr_block     = var.oracle_subnet_cidr1
# }
# # output "oci_core_vcn" {
# #   value = oci_core_vcn.subnet_1
# # }
