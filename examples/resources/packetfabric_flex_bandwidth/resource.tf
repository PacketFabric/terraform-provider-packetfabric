resource "packetfabric_flex_bandwidth" "flex1" {
  provider          = packetfabric
  description       = "hello world"
  subscription_term = 1
  capacity          = "200Gbps"
}

output "packetfabric_flex_bandwidth" {
  value = packetfabric_flex_bandwidth.flex1
}