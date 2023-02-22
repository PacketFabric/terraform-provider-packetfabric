## 1.0.7  (February 22, 2023)

BUG FIXES:

* Error: Status: 400, Metro VCs do not support upgrades or renewals (#350)

## 1.0.6  (February 20, 2023)

BUG FIXES:

* resource packetfabric_cs_aws_dedicated_connection: autoneg: Default cannot be set with Required (#399)

## 1.0.5  (February 20, 2023)

BREAKING CHANGES:

* Change region in packetfabric_cloud_router to required (#341)
* Change pop in packetfabric_cs_aws_hosted_marketplace_connection to required (#344)

IMPROVEMENTS/ENHANCEMENTS:

* Add default asn value to 4556 in packetfabric_cloud_router (#341)
* Add default should_create_lag value for packetfabric_cs_aws_dedicated_connection (#341)
* Add default multihop_ttl value for packetfabric_cloud_router_bgp_session (#341)
* Add default autoneg value for packetfabric_cs_google_dedicated_connection (#341)
* Add default nni and autoneg value for packetfabric_port (#341)

BUG FIXES:

* Errors with "usage" based backbone virtual circuits (#399)

SECURITY:

* CVE-2022-27664 golang.org/x/net/http2 Denial of Service vulnerability (#343)
* CVE-2022-41723 Uncontrolled Resource Consumption (#343)

## 1.0.4  (February 15, 2023)

IMPROVEMENTS/ENHANCEMENTS:

* Update Read functions in Terraform resources to improve Terraform import (#335)
* Add Terraform Import support for IX Virtual Circuit (#333)
* Cleanup rename functions in internal/packetfabric/cloud_router_connection.go and more (#191)
* Add svlan to packetfabric_backbone_virtual_circuit data-source (#335)
* Update Read functions for all Marketplace resources (IX, Cloud, Backbone) (#335)

BUG FIXES:

*  Azure BGP session update: if l3Address is not set, set it based on the values of primarySubnet and secondarySubnet (#336)

## 1.0.3  (February 14, 2023)

BUG FIXES:

* Remove unecessary disable BGP session when deleting packetfabric_cloud_router_bgp_session (#334)
* Add delay after Azure Cloud Router Connection is created (#334)

## 1.0.2  (February 12, 2023)

BUG FIXES:

* Add missing status check for Azure and Port Cloud Router Connection (#332)

## 1.0.1  (February 10, 2023)

IMPROVEMENTS/ENHANCEMENTS:

* Cloud Router documentation updates

BUG FIXES:

* Set remote_address to omitempty in the BGP session schema to prevent sending and empty value (#331)

## 1.0.0  (February 8, 2023)

BREAKING CHANGES:

* resource: [REMOVED] packetfabric_billing_modify_order (updates handled within each resources)

IMPROVEMENTS/ENHANCEMENTS:

* Handle subscription_term, speed, service_class changes in existing resources (#324)
    * packetfabric_port
    * packetfabric_backbone_virtual_circuit
    * packetfabric_point_to_point
    * packetfabric_cs_<aws/azure/google/oracle/ibm>_hosted_connection
    * packetfabric_cs_<aws/azure/google>_dedicated_connection
    * packetfabric_cloud_router_connection_<aws/google/azure/ibm/ipsec/port>
* Add Virtual Circuit Burst status check (#324)
* Implement user inputs checks & validation in all resources (#324)
* Removed deprecated scope attribute in packetfabric_cloud_router resource and data-source (#324)
* Add missing computed ptp_circuit_id attribute in packetfabric_point_to_point (#324)

BUG FIXES:

* Fix packetfabric_cs_<aws/azure/google>_dedicated_connection update (#324)
* Fix packetfabric_link_aggregation_group update (#324)

## 0.9.0  (February 6, 2023)

FEATURES:

* resource: packetfabric_marketplace_service
* resource: packetfabric_cloud_router_quick_connect
* resource: packetfabric_quick_connect_accept_request
* resource: packetfabric_quick_connect_reject_request
* data-source: packetfabric_quick_connect_requests

## 0.8.0  (February 1, 2023)

FEATURES:

* resource: packetfabric_flex_bandwidth (#318)

IMPROVEMENTS/ENHANCEMENTS:

* Add FlexBandwith attribute to Backbone Virtual Circuit resource and data-source (#319)

## 0.7.0  (January 25, 2023)

BREAKING CHANGES:

* resource: [UPDATED] packetfabric_cloud_router_bgp_session (pre_nat_sources and pool_prefixes moved under nat)

FEATURES:

* Add Destination NAT support to packetfabric_cloud_router_bgp_session and packetfabric_cloud_router_connection resources and data-sources (#306)

IMPROVEMENTS/ENHANCEMENTS:

* Improve BGP session update and delete operations in packetfabric_cloud_router_bgp_session resource (#306)

BUG FIXES:

* Fix packetfabric_activitylog data-source (#315)
* Fix docs for marketplace service port resources & data-source (#316)

## 0.6.0  (January 17, 2023)

BREAKING CHANGES:	

* resource: [RENAMED] packetfabric_marketplace_service_port_accept_request (was packetfabric_marketplace_service_accept_request) (#313)
* resource: [RENAMED] packetfabric_marketplace_service_port_reject_request (was packetfabric_marketplace_service_reject_request) (#313)
* data-source: [RENAMED] packetfabric_marketplace_service_port_requests (was packetfabric_marketplace_service_requests) (#313)
* data-source: [UPDATED] packetfabric_cloud_router_bgp_session (require circuit_id and connection_id) (#302)

BUG FIXES:

* Update required/optional parameters packetfabric_backbone_virtual_circuit (#312)
* Deprecate device_id and site_id and correct optics_diagnostics_lane_values.lane_index in packetfabric_port_device_info (#314)

IMPROVEMENTS/ENHANCEMENTS:

* Add zone, desired_nat, vlan to Cloud Router Connection response in packetfabric_cloud_router_connections data-source (#302)
* Add prefixes to create bgp response in packetfabric_cloud_router_bgp_session data-source (#302)

## 0.5.1  (December 7, 2022)

IMPROVEMENTS/ENHANCEMENTS:

* PSE-10317: add dnat_capable to cr connection response (#298)

BUG FIXES:

* Delete Error: could not find BGP session associated with the provided Cloud Router ID: Status: 404 (#266)
* packetfabric_cloud_router - Intermittent 'Context Canceled' error on plan/apply (#284)

## 0.5.0  (December 6, 2022)

FEATURES:

* resource: packetfabric_cs_ibm_hosted_connection
* resource: packetfabric_billing_modify_order (change the subscription term, billing plan, or speed of a connection)
* resource: packetfabric_port_loa
* resource: [UPDATED] packetfabric_port (added enable/disable option)
* data-source: packetfabric_cs_ibm_hosted_connection
* data-source: packetfabric_locations_cloud
* data-source: packetfabric_locations_regions
* data-source: packetfabric_locations_markets
* data-source: packetfabric_locations_pop_zones
* data-source: packetfabric_locations_port_availability
* data-source: packetfabric_port_vlans
* data-source: packetfabric_port_router_logs
* data-source: packetfabric_port_device_info

IMPROVEMENTS/ENHANCEMENTS:

* Add environment variables for Cloud Providers Account information (#276)
* Add support Terraform import for AWS, Google, Oracle and IBM hosted clouds, AWS dedicated cloud and point to point resources (#259)

## 0.4.2  (November 21, 2022)

IMPROVEMENTS/ENHANCEMENTS:

* Enhancement use of PacketFabric environment variables (#269)

## 0.4.1  (November 16, 2022)

BREAKING CHANGES:

* data-source: [REMOVED] packetfabric_cs_aws_dedicated_connection (use packetfabric_cs_dedicated_connections instead)
* data-source: [RENAMED] packetfabric_cs_azure_dedicated_connection (use packetfabric_cs_dedicated_connections instead)
* data-source: [RENAMED] packetfabric_cs_google_dedicated_connection (use packetfabric_cs_dedicated_connections instead)

IMPROVEMENTS/ENHANCEMENTS:

* Add is_awaiting_onramp to AWS and Google hosted cloud resources and data-sources (#261)
* Add support Terraform Import to hosted/dedicated cloud (#259)

BUG FIXES:

* Remove Zone in packetfabric_cs_azure_hosted_marketplace_connection (#240)
* Use-case examples and Location Data Source links on the docs registry are broken (#241)
* Doc: packetfabric_marketplace_requests data-source should not exist (#243)
* Update: Error: Status: 404 Requested URL /v2/services/cloud-routers/PF-L3-CUST-1752618/connections/PF-L3-CON-1752622/bgp not found (#264)

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
* Update Cloud Router Connection response (#194)

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
