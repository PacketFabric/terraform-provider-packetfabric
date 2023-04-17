package testutil

// ACCOUNT
const PF_HOST_KEY = "PF_HOST"
const PF_ACCOUNT_ID_KEY = "PF_ACCOUNT_ID"

//########################################
//###### PORTS / CROSS CONNECT / VC
//########################################

// PORTS
const PF_PORT_MEDIA_KEY = "PF_PORT_MEDIA"
const PF_PORT_POP1_KEY = "PF_PORT_POP1"
const PF_PORT_AVZONE1_KEY = "PF_PORT_AVZONE1"
const PF_PORT_POP2_KEY = "PF_PORT_POP2"
const PF_PORT_AVZONE2_KEY = "PF_PORT_AVZONE2"
const PF_PORT_SUBTERM_KEY = "PF_PORT_SUBTERM"
const PF_PORT_AUTONEG_KEY = "PF_PORT_AUTONEG"
const PF_PORT_SPEED_KEY = "PF_PORT_SPEED"
const PF_PORT_NNI_KEY = "PF_PORT_NNI"
const PF_PORT_ENABLED_KEY = "PF_PORT_ENABLED"

// CROSS CONNECT
const PF_DOCUMENT_UUID1_KEY = "PF_DOCUMENT_UUID1"
const PF_DOCUMENT_UUID2_KEY = "PF_DOCUMENT_UUID2"

// VIRTUAL CIRCUIT
const PF_VC_VLAN1_KEY = "PF_VC_VLAN1"
const PF_VC_VLAN2_KEY = "PF_VC_VLAN2"
const PF_VC_LONGHAUL_TYPE_KEY = "PF_VC_LONGHAUL_TYPE"
const PF_VC_SPEED_KEY = "PF_VC_SPEED"
const PF_VC_SUBTERM_KEY = "PF_VC_SUBTERM"

// BACKBONE VIRTUAL CIRCUIT SPEED BURST
const PF_VC_CIRCUIT_ID_KEY = "PF_VC_CIRCUIT_ID"
const PF_VC_SPEED_BURST_KEY = "PF_VC_SPEED_BURST"

// POINT TO POINT
const PF_PTP_SPEED_KEY = "PF_PTP_SPEED"
const PF_PTP_SUBTERM_KEY = "PF_PTP_SUBTERM"
const PF_PTP_AUTONEG_KEY = "PF_PTP_AUTONEG"
const PF_PTP_MEDIA_KEY = "PF_PTP_MEDIA"
const PF_PTP_POP1_KEY = "PF_PTP_POP1"
const PF_PTP_ZONE1_KEY = "PF_PTP_ZONE1"
const PF_PTP_POP2_KEY = "PF_PTP_POP2"
const PF_PTP_ZONE2_KEY = "PF_PTP_ZONE2"

// ########################################
// ###### HOSTED CLOUD CONNECTIONS
// ########################################

// AZURE HOSTED CONNECTION
const AZURE_SERVICE_KEY = "AZURE_SERVICE_KEY"
const PF_CS_SRC_SVLAN_KEY = "PF_CS_SRC_SVLAN"
const PF_CS_VLAN_PRIVATE_KEY = "PF_CS_VLAN_PRIVATE"
const PF_CS_VLAN_MICROSOFT_KEY = "PF_CS_VLAN_MICROSOFT"

// GCP HOSTED CONNECTION
const GOOGLE_PAIRING_KEY = "GOOGLE_PAIRING_KEY"
const GOOGLE_VLAN_ATTACHMENT_NAME_KEY = "GOOGLE_VLAN_ATTACHMENT_NAME"
const PF_CS_POP1_KEY = "PF_CS_POP1"
const PF_CS_SPEED1_KEY = "PF_CS_SPEED1"
const PF_CS_VLAN1_KEY = "PF_CS_VLAN1"

// AWS HOSTED CONNECTION
const PF_CS_POP2_KEY = "PF_CS_POP2"
const PF_CS_ZONE2_KEY = "PF_CS_ZONE2"
const PF_CS_SPEED2_KEY = "PF_CS_SPEED2"
const PF_CS_VLAN2_KEY = "PF_CS_VLAN2"

// ORACLE HOSTED CONNECTION
const PF_CS_POP6_KEY = "PF_CS_POP6"
const PF_CS_ZONE6_KEY = "PF_CS_ZONE6"
const PF_CS_VLAN6_KEY = "PF_CS_VLAN6"
const PF_CS_ORACLE_REGION_KEY = "PF_CS_ORACLE_REGION"
const PF_CS_ORACLE_VC_OCID_KEY = "PF_CS_ORACLE_VC_OCID"

// IBM HOSTED CONNECTION
const PF_CS_POP7_KEY = "PF_CS_POP7"
const PF_CS_VLAN7_KEY = "PF_CS_VLAN7"
const IBM_BGP_ASN_KEY = "IBM_BGP_ASN"

// MARKEPTLACE
const PF_ROUTING_ID_KEY = "PF_ROUTING_ID"
const PF_MARKET_KEY = "PF_MARKET"
const PF_MARKET_PORT_CIRCUIT_ID_KEY = "PF_MARKET_PORT_CIRCUIT_ID"
const PF_ROUTING_ID_IX_KEY = "PF_ROUTING_ID_IX"
const PF_MARKET_IX_KEY = "PF_MARKET_IX"
const PF_ASN_IX_KEY = "PF_ASN_IX"

// ########################################
// ###### DEDICATED CLOUD CONNECTIONS
// ########################################

// AWS DEDICATED CONNECTION
const PF_CS_POP3_KEY = "PF_CS_POP3"
const PF_CS_ZONE3_KEY = "PF_CS_ZONE3"
const PF_CS_SPEED3_KEY = "PF_CS_SPEED3"
const AWS_REGION3_KEY = "AWS_REGION3"

// GOOGLE DEDICATED CONNECTION
const PF_CS_POP4_KEY = "PF_CS_POP4"
const PF_CS_ZONE4_KEY = "PF_CS_ZONE4"
const PF_CS_SPEED4_KEY = "PF_CS_SPEED4"

// AZURE DEDICATED CONNECTION
const PF_CS_POP5_KEY = "PF_CS_POP5"
const PF_CS_ZONE5_KEY = "PF_CS_ZONE5"
const PF_CS_SPEED5_KEY = "PF_CS_SPEED5"
const ENCAPSULATION_KEY = "ENCAPSULATION"
const PORT_CATEGORY_KEY = "PORT_CATEGORY"

// DEDICATED ALL CLOUDS
const PF_CS_SRVCLASS_KEY = "PF_CS_SRVCLASS"
const PF_CS_AUTONEG_KEY = "PF_CS_AUTONEG"
const SHOULD_CREATE_LAG_KEY = "SHOULD_CREATE_LAG"
const PF_CS_SUBTERM_KEY = "PF_CS_SUBTERM"

// ########################################
// ###### CLOUD ROUTER
// ########################################
const PF_CR_ASN_KEY = "PF_CR_ASN"
const PF_CR_CAPACITY_KEY = "PF_CR_CAPACITY"
const PF_CR_REGIONS_KEY = "PF_CR_REGIONS"
const PF_CR_MEMBER_KEY = "PF_CR_MEMBER"

// CLOUD ROUTER CONNECTIONS
const PF_CRC_MAYBE_NAT_KEY = "PF_CRC_MAYBE_NAT"
const PF_CRC_IS_PUBLIC_KEY = "PF_CRC_IS_PUBLIC"

// CLOUD ROUTER CONNECTION AWS
const PF_CRC_SPEED_KEY = "PF_CRC_SPEED"
const PF_CRC_POP1_KEY = "PF_CRC_POP1"
const PF_CRC_ZONE1_KEY = "PF_CRC_ZONE1"
const PF_CRC_AWS_ACCOUNT_ID_KEY = "PF_CRC_AWS_ACCOUNT_ID"

// CLOUD ROUTER CONNECTION GOOGLE
const PF_CRC_POP2_KEY = "PF_CRC_POP2"
const PF_CRC_GOOGLE_PAIRING_KEY = "PF_CRC_GOOGLE_PAIRING_KEY"
const PF_CRC_GOOGLE_VLAN_ATTACHMENT_NAME_KEY = "PF_CRC_GOOGLE_VLAN_ATTACHMENT_NAME"

// CLOUD ROUTER CONNECTION AZURE
const PF_CRC_POP3_KEY = "PF_CRC_POP3"
const PF_CRC_AZURE_SERVICE_KEY = "PF_CRC_AZURE_SERVICE_KEY"

// CLOUD ROUTER CONNECTION IPSEC
const PF_CRC_IKE_VERSION_KEY = "PF_CRC_IKE_VERSION"
const PF_CRC_PHASE1_AUTHENTICATION_METHOD_KEY = "PF_CRC_PHASE1_AUTHENTICATION_METHOD"
const PF_CRC_PHASE1_GROUP_KEY = "PF_CRC_PHASE1_GROUP"
const PF_CRC_PHASE1_ENCRYPTION_ALGO_KEY = "PF_CRC_PHASE1_ENCRYPTION_ALGO"
const PF_CRC_PHASE1_AUTHENTICATION_ALGO_KEY = "PF_CRC_PHASE1_AUTHENTICATION_ALGO"
const PF_CRC_PHASE1_LIFETIME_KEY = "PF_CRC_PHASE1_LIFETIME"
const PF_CRC_PHASE2_PFS_GROUP_KEY = "PF_CRC_PHASE2_PFS_GROUP"
const PF_CRC_PHASE2_ENCRYPTION_ALGO_KEY = "PF_CRC_PHASE2_ENCRYPTION_ALGO"
const PF_CRC_PHASE2_AUTHENTICATION_ALGO_KEY = "PF_CRC_PHASE2_AUTHENTICATION_ALGO"
const PF_CRC_PHASE2_LIFETIME_KEY = "PF_CRC_PHASE2_LIFETIME"
const PF_CRC_GATEWAY_ADDRESS_KEY = "PF_CRC_GATEWAY_ADDRESS"
const PF_CRC_SHARED_KEY_KEY = "PF_CRC_SHARED_KEY"

// CLOUD ROUTER BGP SESSION IPSEC
const PF_CRBS_AF_KEY = "PF_CRBS_AF"
const PF_CRBS_MHTTL_KEY = "PF_CRBS_MHTTL"
const PF_CRBS_ORLONGER_KEY = "PF_CRBS_ORLONGER"

const VPN_SIDE_ASN3_KEY = "VPN_SIDE_ASN3"
const VPN_REMOTE_ADDRESS_KEY = "VPN_REMOTE_ADDRESS"
const VPN_L3_ADDRESS_KEY = "VPN_L3_ADDRESS"

// CLOUD ROUTER CONNECTION PORT
const PF_CRC_PORT_CIRCUIT_ID_KEY = "PF_CRC_PORT_CIRCUIT_ID"
const PF_CRC_VLAN_KEY = "PF_CRC_VLAN"

// CLOUD ROUTER CONNECTION IBM
const PF_CRC_POP4_KEY = "PF_CRC_POP4"
const PF_CRC_ZONE4_KEY = "PF_CRC_ZONE4"
const PF_CRC_IBM_BGP_ASN_KEY = "PF_CRC_IBM_BGP_ASN"

// CLOUD ROUTER CONNECTION ORACLE
const PF_CRC_POP5_KEY = "PF_CRC_POP5"
const PF_CRC_ZONE5_KEY = "PF_CRC_ZONE5"
const PF_CRC_ORACLE_REGION_KEY = "PF_CRC_ORACLE_REGION"
const PF_CRC_ORACLE_VC_OCID_KEY = "PF_CRC_ORACLE_VC_OCID"

const PF_DTS_TIME_FROM_KEY = "PF_DTS_TIME_FROM"
const PF_DTS_TIME_TO_KEY = "PF_DTS_TIME_TO"
