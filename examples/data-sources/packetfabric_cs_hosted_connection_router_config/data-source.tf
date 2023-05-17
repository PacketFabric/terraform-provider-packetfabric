data "packetfabric_cs_hosted_connection_router_config" "router_aws_cisco2900" {
  cloud_circuit_id = "PF-AP-LAX1-1002"
  router_type      = "CiscoSystemsInc-2900SeriesRouters-IOS124"
}
resource "local_file" "router_aws_cisco2900_file" {
  filename = "router_config_aws_cisco2900.txt"
  content  = data.packetfabric_cs_hosted_connection_router_config.router_aws_cisco2900.router_config
}