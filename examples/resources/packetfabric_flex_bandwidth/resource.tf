resource "packetfabric_flex_bandwidth" "flex1" {
  provider          = packetfabric
  description       = var.pf_description
  subscription_term = var.pf_flex_subscription_term
  capacity          = var.pf_flex_capacity
}

output "packetfabric_flex_bandwidth" {
  value = packetfabric_flex_bandwidth.flex1
}