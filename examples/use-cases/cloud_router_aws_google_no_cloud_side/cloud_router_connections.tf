# # just informative
# data "packetfabric_cloud_router_connections" "all_crc" {
#   provider   = packetfabric
#   circuit_id = packetfabric_cloud_router.cr.id

#   depends_on = [
#     packetfabric_cloud_router_bgp_session.crbs_aws,
#     packetfabric_cloud_router_bgp_session.crbs_google
#   ]
# }
# output "packetfabric_cloud_router_connections" {
#   value = data.packetfabric_cloud_router_connections.all_crc
# }