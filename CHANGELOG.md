## 1.6.0  (May 24, 2023)

BREAKING CHANGES:

* data-source: [RENAMED] packetfabric_point_to_points (was packetfabric_point_to_point) (#448)
* data-source: [RENAMED] packetfabric_cloud_routers (was packetfabric_cloud_router) (#517)
* data-source: [RENAMED] packetfabric_outbound_cross_connects (was packetfabric_outbound_cross_connect) (#469)
* data-source: [RENAMED] packetfabric_activitylogs (was packetfabric_activitylog) (#469)
* resources: [UPDATED] changed zone from optional to required (#517)
    * packetfabric_port
    * packetfabric_cloud_router_connection_aws
    * packetfabric_cloud_router_connection_ibm
    * packetfabric_cloud_router_connection_oracle
    * packetfabric_cs_aws_hosted_connection
    * packetfabric_cs_ibm_hosted_connection
    * packetfabric_cs_oracle_hosted_connection
    * packetfabric_cs_aws_dedicated_connection
    * packetfabric_cs_google_dedicated_connection
    * packetfabric_cs_azure_dedicated_connection

FEATURES:

* resource: packetfabric_document (#485)
* data-source: packetfabric_documents (#485)

IMPROVEMENTS/ENHANCEMENTS:

* Add Early Termination Liability (ETL) info to all applicable resources (#492)
* Improve service status checks for Port and Backbone VC (#367)
* Provide IBM Gateway ID in IBM Hosted Cloud and Cloud Router Connection resources (#490)
* Provide Azure Private/Microsoft VLAN ID and connection type in Azure Cloud Router Connection resource (#489)
* Handle cloud router deletion when billing is not created yet (#244)
* Add port_circuit_id as computed field in packetfabric_point_to_point (#448)
* Use ptp_circuit_id instead of UUID in packetfabric_point_to_point (#448)
* Refactor status check for packetfabric_point_to_point (#448)
* Add deprecate field to speed in Azure Hosted Cloud resource (#499)
* packetfabric_link_aggregation_group enable/disable LAG interface (#330)
* Add warning and more info to the Google Cloud Router Connection Resource (#506)
* Updates Optional/Computed fields in all data-sources (#440)
* Add AWS_ACCOUNT_ID env var as an option to set AWS Account ID in AWS Cloud Router Connection & Hosted Cloud (#506)
* Add is_lag to packetfabric_cs_dedicated_connections data-source (#524)

BUG FIXES:

* Correct read operation in packetfabric_cloud_router_bgp_session for Azure Connection (#489)
* Fix errors when applying labels in packetfabric_point_to_point (#492)
* Fix panic: Invalid address to set: []string{"cloud_circuit_id"} in all hosted and dedicated cloud resources (#418)
* Fix panic: Invalid address to set: []string{"pop"} in Azure Hosted Cloud resource (#389)
* Fix panic: Invalid address to set: []string{"speed"} in packetfabric_cs_oracle_hosted_connection (#426)
* Fix error: updating cloud_settings for AWS and Google Hosted Cloud resources (#456)
* Fix Terraform Import of BGP session when used with Azure Cloud Router Connection (#369)
* Removed should_create_lag in read() function for Google and Azure Dedicated Cloud resources (#422)
* Ignore Labels order in all resources supporting labels (#509)
* Update IPsec Cloud Router Connection status check (#420)
* Handling config diff due to endpoints order in packetfabric_point_to_point (#510)
* Handling vlan creating config diff with terraform plan/import in packetfabric_cloud_router_connection_port (#388)
* Handling zone, po_number creating config diff with terraform plan/import in all resources (#418, #426)
* Fix packetfabric_point_to_points data-source not returning anything (#470)
* Fix packetfabric_outbound_cross_connects data-source not returning anything (#469)

ACCEPTANCE TESTING:

* ACC Test for packetfabric_port resource (#356)
* ACC Test for packetfabric_port_loa resource (#447)
* ACC Test for packetfabric_outbound_cross_connect resource (#446)
* ACC Test for packetfabric_link_aggregation_group resource (#451)
* ACC Test for packetfabric_backbone_virtual_circuit resource (#454)
* ACC Test for packetfabric_backbone_virtual_circuit_speed_burst resource (#452)
* ACC Test for packetfabric_point_to_point resource (#448)
* ACC Test for packetfabric_cloud_router_connection_aws resource (#424)
* ACC Test for packetfabric_cloud_router_connection_google resource (#450)
* ACC Test for packetfabric_cloud_router_connection_azure resource (#389)
* ACC Test for packetfabric_cloud_router_connection_ibm resource (#419)
* ACC Test for packetfabric_cloud_router_connection_oracle resource (#426)
* ACC Test for packetfabric_cloud_router_connection_port resource (#388)
* ACC Test for packetfabric_cloud_router_connection_ipsec resource (#420)
* ACC Test for packetfabric_cloud_router_bgp_session resource (#369)
* ACC Test for packetfabric_cloud_provider_credential_google resource (#526)
* ACC Test for packetfabric_cloud_provider_credential_aws resource (#526)
* ACC Test for packetfabric_cs_aws_hosted_connection resource (#418)
* ACC Test for packetfabric_cs_google_hosted_connection resource (#456)
* ACC Test for packetfabric_cs_azure_hosted_connection resource (#389)
* ACC Test for packetfabric_cs_ibm_hosted_connection resource (#419)
* ACC Test for packetfabric_cs_oracle_hosted_connection resource (#426)
* ACC Test for packetfabric_cs_aws_dedicated_connection resource (#432)
* ACC Test for packetfabric_cs_google_dedicated_connection resource (#422)
* ACC Test for packetfabric_cs_azure_dedicated_connection resource (#425)
* ACC Test for packetfabric_activitylogs data-source (#465)
* ACC Test for packetfabric_locations_markets data-source (#464)
* ACC Test for packetfabric_locations_regions data-source (#463)
* ACC Test for packetfabric_locations_pop_zones data-source (#462)
* ACC Test for packetfabric_locations data-source (#461)
* ACC Test for packetfabric_locations_port_availability data-source (#460)
* ACC Test for packetfabric_locations_cloud data-source (#453)
* ACC Test for packetfabric_ports data-source (#441)
* ACC Test for packetfabric_port_device_info data-source (#457)
* ACC Test for packetfabric_port_vlans data-source (#458)
* ACC Test for packetfabric_port_router_logs data-source (#459)
* ACC Test for packetfabric_outbound_cross_connects data-source (#469)
* ACC Test for packetfabric_link_aggregation_group data-source (#486)
* ACC Test for packetfabric_point_to_point data-source (#467)
* ACC Test for packetfabric_billing data-source (#439)
* ACC Test for packetfabric_cs_aws_hosted_connection data-source (#439)
* ACC Test for packetfabric_cs_dedicated_connections data-source (#440)
* ACC Test for packetfabric_cloud_router_connection data-source (#470)
* ACC Test for packetfabric_cloud_router_connections data-source (#470)
* ACC Test for packetfabric_cloud_router_connection_ipsec data-source (#466)
* ACC Test for packetfabric_cloud_router_bgp_session (#471)

## 1.5.0  (May 3, 2023)

BREAKING CHANGES:

* Prefix Order and Community fields are deprecated in packetfabric_cloud_router_bgp_session resource and data-source (#436)

FEATURES:

* resource: [UPDATED] packetfabric_cloud_router_connection_aws - adding cloud side provisioning (#436)
* resource: [UPDATED] packetfabric_cloud_router_connection_google - adding cloud side provisioning (#436)
* data-source: packetfabric_cloud_router_connection (#486)

IMPROVEMENTS/ENHANCEMENTS:

* Provide AWS Direct Connect Connection ID in AWS Hosted Cloud and Cloud Router Connection resources (#484)
* Make google_vlan_attachment_name optional when Cloud Settings are used for google hosted cloud (#415)
* Add check/error when user try to update port_circuit_id in Backbone VC resource (#479)
* Add check/error when user try to update cloud settings which cannot be updated-in-place in AWS and Google Hosted Cloud resources (#436)
* Use circuit ID + "-data" for the following data-sources's IDs instead of a random uuid (#436) 
    * packetfabric_cloud_router_bgp_session
    * packetfabric_cloud_router_connection_ipsec
    * packetfabric_cloud_router_connections
    * packetfabric_cs_aws_hosted_connection
    * packetfabric_cs_azure_hosted_connection
    * packetfabric_cs_google_hosted_connection
    * packetfabric_cs_ibm_hosted_connection
    * packetfabric_cs_oracle_hosted_connection
    * packetfabric_cs_hosted_connection_router_config
    * packetfabric_link_aggregation_group
    * packetfabric_port_router_logs
    * packetfabric_port_vlans
* Automate creation of the Release Note from changelog (#478)

## 1.4.0  (April 19, 2023)

BREAKING CHANGES:

* data-source: [RENAMED] packetfabric_ports (was packetfabric_port) (#429)

FEATURES:

* resource: packetfabric_streaming_events (#212)

IMPROVEMENTS/ENHANCEMENTS:

* Update BGP session deletion warning (#431)

BUG FIXES:

* [ERROR] setting state: labels: '': source data must be an array or slice, got struct (#427)
* Check if zone or autoneg are set in packetfabric_port resource (#433)
* Check if po_number is set in the resource using po_number (#435)
* Check if zone is set in packetfabric_cloud_router_connection_aws/oracle/ibm resources (#438)
* Check if ibm_bgp_cer_cidr and ibm_bgp_ibm_cidr are set in packetfabric_cloud_router_connection_ibm resource (#438)
* Check if phase2_authentication_algo is set in packetfabric_cloud_router_connection_ipsec resource (#438)
* Check if vlan is set in packetfabric_cloud_router_connection_port resource (#438)

ACCEPTANCE TESTING:

* ACC Test for packetfabric_cloud_router resource (#368)

## 1.3.0  (April 5, 2023)

FEATURES:

* resource: packetfabric_cloud_provider_credential_aws (#376)
* resource: packetfabric_cloud_provider_credential_google (#376)
* resource: [UPDATED] packetfabric_cs_aws_hosted_connection (#408)
* resource: [UPDATED] packetfabric_cs_google_hosted_connection (#414)
* data-source: [UPDATED] packetfabric_cs_aws_hosted_connection (#408)
* data-source: [UPDATED] packetfabric_cs_google_hosted_connection (#414)
* data-source: packetfabric_cs_hosted_connection_router_config (#409)

IMPROVEMENTS/ENHANCEMENTS:

* Add validation on BGP prefix in/ou in BGP session resource (#393)
* Update Read function for packetfabric_cloud_router_bgp_session (#365)
* Add defaults and additional validation to packetfabric_cloud_router_bgp_session (#365)
* Add cloud_router_circuit_id to QuickConnect import request (#400)
* Add pending_approval to QuickConnect return filters response (#400)
* Add validation for subscription_term and longhaul_type (#401)
* Add validation for Marketplace Service creation on Categories (#410)
* Add new labels page under Guides in Terraform Registry Documentation (#396)

BUG FIXES:

* Handling svlan unset creating config diff with terraform plan after VC creation (#387)
* Correct pop set in the cloud router connection read functions (#399)
* Adding back importing page under Guides (#394)
* Fix po_number reading in all PacketFabric hosted connection (#408)

## 1.2.1  (March 21, 2023)

IMPROVEMENTS/ENHANCEMENTS:

* Update bgp_state in cloud router connection and BGP session data-sources (#384)
* Update default value for autoneg in packetfabric_port resource in the doc (#383)
* Increase status check interval in Cloud Router, Cloud Services, Backbone VC, PTP, Port (#390)

## 1.2.0  (March 13, 2023)

FEATURES:

* resource: packetfabric_user (#371)
* Add Object Labels (#375) and PO Numbers (#374) to the following resources: 
    * packetfabric_backbone_virtual_circuit
    * packetfabric_cloud_router
    * packetfabric_cloud_router_connection_aws
    * packetfabric_cloud_router_connection_azure
    * packetfabric_cloud_router_connection_google
    * packetfabric_cloud_router_connection_ibm
    * packetfabric_cloud_router_connection_ipsec
    * packetfabric_cloud_router_connection_oracle
    * packetfabric_cloud_router_connection_port
    * packetfabric_cloud_router_quick_connect (only Object Label)
    * packetfabric_cs_aws_dedicated_connection
    * packetfabric_cs_aws_hosted_connection
    * packetfabric_cs_azure_dedicated_connection
    * packetfabric_cs_azure_hosted_connection
    * packetfabric_cs_google_dedicated_connection
    * packetfabric_cs_google_hosted_connection
    * packetfabric_cs_ibm_hosted_connection
    * packetfabric_cs_oracle_hosted_connection
    * packetfabric_link_aggregation_group
    * packetfabric_point_to_point
    * packetfabric_port

IMPROVEMENTS/ENHANCEMENTS:

* Update IBM hosted cloud and cloud router status check and use case examples (#370)
* Update Azure Dedicated Cloud status check (#373)
* Added new use case example Cloud Router with AWS, Oracle and PacketFabric Port (#378)
* Update Oracle Cloud Router Connect API endpoint from v2 to v2.1 (#381)

BUG FIXES:

* Update maybe_nat and maybe_dnat in all packetfabric_cloud_router_connection resources (#377)
* Fix update issue with autoneg in packetfabric_port interfaces 10Gbps (#379)

## 1.1.0  (March 2, 2023)

IMPROVEMENTS/ENHANCEMENTS:

* Add l3_address to bgp response in packetfabric_cloud_router_bgp_session (#346)
* Add Terraform Import support for Cloud Router Connection aws, google, azure, ibm, oracle, port and ipsec resources (#340)
* Add Terraform (partial) Import support for Cloud Router BGP session resource (#340)
* Add Terraform Import support for Port resource (#347)
* Add Terraform Import support for Flex Bandwidth resource (#348)
* Add Terraform Import support for Point to Point resource (#349)
* Add Terraform Import support for Dediacted Cloud aws, google, azure (#352)
* Add Terraform Import support for Hosted Cloud aws, google, azure, ibm, oracle, port and ipsec resources (#352)
* Add Terraform Import support for LAG resource (#361)
* Add rate_limit_in/out, flex_bandwidth_id to packetfabric_backbone_virtual_circuit read function (#353)
* Handle location request for Hosted Cloud AWS, Google and IBM (#352)
* Add missing status check AWS and Oracle hosted cloud (#352)
* Deprecate order in packetfabric_cloud_router_bgp_session resource and data-source (#359)
* Add port status check for update operations (#366)

BUG FIXES:

* Fix update issue with untagged in packetfabric_backbone_virtual_circuit interfaces (#364)
* Fix update issue with autoneg in packetfabric_port interfaces (#366)

## 1.0.7  (February 22, 2023)

BUG FIXES:

* Error: Status: 400, Metro VCs do not support upgrades or renewals (#350)

## 1.0.6  (February 20, 2023)

BUG FIXES:

* resource packetfabric_cs_aws_dedicated_connection: autoneg: Default cannot be set with Required (#345)

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
