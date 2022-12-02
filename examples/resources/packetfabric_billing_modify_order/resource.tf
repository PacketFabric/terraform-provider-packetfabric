resource "packetfabric_billing_modify_order" "upgrade_backbone_vc" {
  provider             = packetfabric
  circuit_id           = packetfabric_backbone_virtual_circuit.vc1
  billing_product_type = "longhaul_dedicated"
  subscription_term    = 1
  speed                = "1Gbps"
}
output "packetfabric_billing_modify_order_upgrade_backbone_vc" {
  value = packetfabric_billing_modify_order.upgrade_backbone_vc
}

# Upgrade Cloud Router Connection
resource "packetfabric_billing_modify_order" "upgrade_cloud_router_connection" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr1.id
  speed      = "100Mbps"
}
output "packetfabric_billing_modify_order_upgrade_crc" {
  value = packetfabric_billing_modify_order.upgrade_cloud_router_connection
}