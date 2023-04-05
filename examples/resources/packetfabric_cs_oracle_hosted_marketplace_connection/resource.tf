
resource "packetfabric_cs_oracle_hosted_marketplace_connection" "cs_conn1_marketplace_oracle" {
  provider    = packetfabric
  description = "hello world"
  vc_ocid     = var.oracle_vc_ocid
  region      = "us-ashburn-1"
  routing_id  = "PD-WUY-9VB0"
  market      = "HOU"
  pop         = "BOS1"
}