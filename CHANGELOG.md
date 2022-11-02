## 0.4.0  (November 2, 2022)

BREAKING CHANGES:

* resource: [REMOVED] packetfabric_cs_aws_provision_marketplace (replaced by packetfabric_marketplace_service_accept_request)
* resource: [REMOVED] packetfabric_cs_azure_provision_marketplace (replaced by  packetfabric_marketplace_service_accept_request)
* resource: [REMOVED] packetfabric_cs_google_provision_marketplace (replaced by  packetfabric_marketplace_service_accept_request)
* resource: [REMOVED] packetfabric_cloud_router_bgp_prefixes (use packetfabric_cloud_router_bgp_session instead)
* data-source: [REMOVED] packetfabric_cloud_router_bgp_prefixes (use packetfabric_cloud_router_bgp_session instead)
* data-source: [RENAMED] packetfabric_cs_aws_dedicated_connection (was packetfabric_cs_aws_dedicated_connection_conn)

FEATURES:

* resource: packetfabric_cloud_router_connection_azure
* resource: packetfabric_cloud_router_connection_oracle
* resource: packetfabric_cloud_router_connection_ibm
* resource: packetfabric_cloud_router_connection_port
* resource: packetfabric_ix_virtual_circuit_marketplace
* resource: packetfabric_backbone_virtual_circuit_marketplace
* resource: packetfabric_backbone_virtual_circuit_speed_burst
* resource: packetfabric_point_to_point
* resource: packetfabric_cs_oracle_hosted_connection
* resource: packetfabric_cs_oracle_hosted_marketplace_connection
* resource: packetfabric_marketplace_service_accept_request
* resource: packetfabric_marketplace_service_reject_request
* data-source: packetfabric_cs_oracle_hosted_connection
* data-source: packetfabric_point_to_point
* data-source: packetfabric_virtual_circuits
* data-source: packetfabric_marketplace_service_requests

IMPROVEMENTS/ENHANCEMENTS:

* prefixes object attribute missing from the packetfabric_cloud_router_bgp_session resource (#138)
* Add a 30sec delay when deleting a Cloud Router Connection (#157)
* Wait till packetfabric_backbone_virtual_circuit is active (#172)
* The published_quote_line_uuid attribute is missing in resourceRouterConnectionAws() (#158)

BUG FIXES:

* Error: Post "https://api.packetfabric.com/v2/services/backbone": context canceled (#165)
* packetfabric_port update: Error: autoneg is a required field (#181)
* packetfabric_cs_<aws/google/azure>_hosted_marketplace_connection delete not working (#91)
* packetfabric_cloud_router_connection_ipsec: accept null value for phase2_authentication_algo (#192)

## 0.3.1 (October 6, 2022)

BUG FIXES:

* Resource packetfabric_cloud_router_connection_aws does not correctly recognize state when refreshing plans (#149)

## 0.3.0 (September 30, 2022)

BREAKING CHANGES:	

* resource: [RENAMED] packetfabric_cloud_router_connection_aws (was packetfabric_aws_cloud_router_connection)
* data-source: [RENAMED] packetfabric_cloud_router_connections (was packetfabric_aws_cloud_router_connection)

FEATURES:

* resource: packetfabric_cloud_router_connection_google
* resource: packetfabric_cloud_router_connection_ipsec
* data-source: packetfabric_cloud_router_connection_ipsec

IMPROVEMENTS/ENHANCEMENTS:

* Removed the scope attribute when creating a new Cloud Router (#40)

BUG FIXES:

* Address bugs bgp session/prefix deletion (#20)
* Urgent: packetfabric_cloud_router_bgp_prefixes skip delete when destroy applied (#140)

## 0.2.3 (September 1, 2022)

IMPROVEMENTS/ENHANCEMENTS:

* Regenerating documentation

## 0.2.2 (September 1, 2022)

BREAKING CHANGES:

* data-source: [REMOVED] packetfabric_cs_aws_hosted_marketplace_connection (#97)
* data-source: [REMOVED] packetfabric_cs_google_hosted_marketplace_connection (#97)
* data-source: [REMOVED] packetfabric_cs_azure_marketplace_connection (#97)

IMPROVEMENTS/ENHANCEMENTS:

* Change-Me's Need to be Consistent documentation [#120]
* Add CHANGELOG.md documentation [#45]
* More meaningful examples of resources usage documentation [#14]

BUG FIXES:

* Add 30sec delay at port creation/deletion in packetfabric_port resource [#111]
* Security: RandomAlphaNumeric and CryptoRandomAlphaNumeric are not as random as they should be #1 [#108]
* Resource packetfabric_port: remove disable port before deletion [#102]
* Remove all hosted marketplace data sources (AWS/Azure/Google) [#97]
* packetfabric_cs_aws_dedicated_connection data source does not exist [#72]

## 0.2.1 (August 30, 2022)

BREAKING CHANGES:
* resource: [REMOVED] packetfabric_cloud_services_aws_create_backbone_dedicated_cr
* resource: [REMOVED] packetfabric_cloud_services_aws_hosted_connections
* resource: [REMOVED] packetfabric_cloud_services_aws_hosted_connection
* resource: [REMOVED] packetfabric_cloud_services_aws_hosted_mkt_conn 
* resource: [REMOVED] packetfabric_cloud_services_google_backbone
* resource: [REMOVED] packetfabric_cloud_services_azr_backbone
* resource: [REMOVED] packetfabric_cs_aws_hosted_marketplace_connection (#91)
* resource: [REMOVED] packetfabric_cs_google_hosted_marketplace_connection (#91)
* resource: [REMOVED] packetfabric_cs_azure_hosted_marketplace_connection (#91)
* resource: packetfabric_interface [REPLACED BY] packetfabric_port
* resource: packetfabric_cloud_services_aws_req_hosted_conn [REPLACED BY] packetfabric_cs_aws_hosted_connection
* resource: packetfabric_cloud_services_aws_req_dedicated_con [REPLACED BY] packetfabric_cs_aws_dedicated_connection (#72)
* resource: packetfabric_cloud_services_aws_provision_mkt [REPLACED BY] packetfabric_cs_aws_provision_marketplace (#72)
* data-source: packetfabric_cloud_services_aws_connection_info [REPLACED BY] packetfabric_cs_aws_hosted_connection  (#72)
* data-source: packetfabric_cloud_services_aws_dedicated_conn [REPLACED BY] packetfabric_cs_aws_dedicated_connection (#72)

FEATURES:

* resource: packetfabric_link_aggregation_group 
* resource: packetfabric_backbone_virtual_circuit
* resource: packetfabric_outbound_cross_connect
* resource: packetfabric_cs_google_hosted_connection
* resource: packetfabric_cs_google_dedicated_connection
* resource: packetfabric_cs_google_provision_marketplace
* resource: packetfabric_cs_azure_hosted_connection
* resource: packetfabric_cs_azure_dedicated_connection
* resource: packetfabric_cs_azure_provision_marketplace
* data-source: packetfabric_link_aggregation_group
* data-source: packetfabric_port
* data-source: packetfabric_backbone_virtual_circuit (#60 v0.3.0) 
* data-source: packetfabric_activitylog
* data-source: packetfabric_outbound_cross_connect
* data-source: packetfabric_cs_google_hosted_connection
* data-source: packetfabric_cs_google_dedicated_connection
* data-source: packetfabric_cs_azure_hosted_connection
* data-source: packetfabric_cs_azure_dedicated_connection

IMPROVEMENTS/ENHANCEMENTS:

* Rename packetfabric_interface to packetfabric_port documentation [#75]
* Data source Cloud Service AWS/Google/Azure hosted/dedicated connection change [#68]
* No examples for packetfabric_billing and packetfabric_locations documentation [#47]
* cloud_router_bgp_session: NAT settings (prefixes and nat) documentation improvement documentation [#43]
* Data Source aws_cloud_connections: cloud_settings should not be an array [#41]
* aws_cloud_router_connection fix resources and examples documentation [#19]

BUG FIXES:

* resource & data source packetfabric_link_aggregation_group not working [#93]
* packetfabric_cs_azure_provision_marketplace should be using vlan_private and vlan_microsoft [#92]
* packetfabric_cs_aws_hosted_marketplace_connection resource: "zone" is not expected here. [#90]
* packetfabric_cs_\<aws/azure/google>_\<hosted/dedicated>_connection delete aren't working [#83]
* packetfabric_backbone_virtual_circuit delete isn't working [#76]
* packetfabric_cs_aws_dedicated_connection: Error: Provider produced inconsistent result after apply [#71]
* Cannot destroy packetfabric_backbone_virtual_circuit [#70]
* packetfabric_cloud_router: panic: interface conversion: interface {} is []interface {}, not \*schema.Set [#69]
* Destroy not working: OutboundCrossConnect not found [#66]
* packetfabric_cs_google_hosted_connection: Error: Plugin did not respond [#65]
* packetfabric_cs_azure_hosted_connection: json: cannot unmarshal string into Go struct field [#64]
* interface conversion: interface {} is nil, not int on src_svlan [#62]
* Rename Data Source packetfabric_aws_cloud_router_connection to packetfabric_cloud_router_connection [#61]
* packetfabric_interface data source missing [#59]
* make test fails on all branches [#58]
* output packetfabric_cloud_services_gcp_req_hosted_conn1 sensitive = true [#57]
* packetfabric_cloud_services_gcp_req_hosted_conn: panic: interface conversion: interface {} is nil, not string [#56]
* packetfabric_cs_google_dedicated_connection resource missing [#55]
* packetfabric_cs_aws_hosted_connection: src_svlan should be optional [#54]
* Add missing examples for AWS documentation  [#53]
* bandwidth: longhaul_type must be specified for for a longhaul virtual circuit [#52]
* packetfabric_interface status checks needs to be updated ("These ports are not active") [#51]
* AWS, Azure, GCP resources & data source updates (delete/rename) [#50]
* packetfabric_outbound_cross_connect: Error: Provider produced inconsistent result after apply [#49]
* Cloud router creation with 2 regions (US and UK) creates only with 1 region [#30]
* Can’t destroy an interface PacketFabric terraform provider [#22]
* Not possible to setup BFD setting with cloud_router_bgp_session [#21]
* Unable to destroy cloud_router_bgp_session resource [#16]
* Error during cloud_services_aws_hosted_connection & cloud_services_aws_req_hosted_conn resources creation [#15]
* Invalid JSON region field mapping for Cloud Router [#12]
* Resource and data source names do not follow Terraform naming best practices [#9]

## 0.1.0 (June 23, 2022)
Initial Release

FEATURES:

* resource: packetfabric_cloud_router
* resource: packetfabric_aws_cloud_router_connection
* resource: packetfabric_cloud_router_bgp_prefixes
* resource: packetfabric_cloud_router_bgp_session
* resource: packetfabric_cloud_services_aws_create_backbone_dedicated_cr
* resource: packetfabric_cloud_services_aws_hosted_connection
* resource: packetfabric_cloud_services_aws_req_hosted_conn
* resource: packetfabric_cloud_services_aws_req_dedicated_con
* resource: packetfabric_cloud_services_aws_hosted_mkt_conn
* resource: packetfabric_cloud_services_aws_provision_mkt
* data-source: packetfabric_cloud_router
* data-source: packetfabric_aws_cloud_router_connection
* data-source: packetfabric_cloud_router_bgp_prefixes
* data-source: packetfabric_cloud_router_bgp_session
* data-source: packetfabric_billing
* data-source: packetfabric_locations
* data-source: packetfabric_cloud_services_aws_connection_info
* data-source: packetfabric_cloud_services_aws_dedicated_conn 
* data-source: packetfabric_aws_services_hosted_requested_mkt_conn
