package testutil

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
)

// ########################################
// ###### RESOURCES & DATA-SOURCES CONSTs
// ########################################

// ###### Resources
const pfPort = "packetfabric_port"
const pfPortLoa = "packetfabric_port_loa"
const pfDocument = "packetfabric_document"
const pfOutboundCrossConnect = "packetfabric_outbound_cross_connect"
const pfLinkAggregationGroup = "packetfabric_link_aggregation_group"
const pfBackboneVirtualCircuit = "packetfabric_backbone_virtual_circuit"
const pfAddSpeedBurst = "packetfabric_backbone_virtual_circuit_speed_burst"
const pfPoinToPoint = "packetfabric_point_to_point"
const pfCloudRouter = "packetfabric_cloud_router"
const pfCloudRouterConnAws = "packetfabric_cloud_router_connection_aws"
const pfCloudRouterConnGoogle = "packetfabric_cloud_router_connection_google"
const pfCloudRouterConnAzure = "packetfabric_cloud_router_connection_azure"
const pfCloudRouterConnIbm = "packetfabric_cloud_router_connection_ibm"
const pfCloudRouterConnOracle = "packetfabric_cloud_router_connection_oracle"
const pfCloudRouterConnPort = "packetfabric_cloud_router_connection_port"
const pfCloudRouterConnIpsec = "packetfabric_cloud_router_connection_ipsec"
const pfCloudRouterBgpSession = "packetfabric_cloud_router_bgp_session"
const pfCsAwsHostedConn = "packetfabric_cs_aws_hosted_connection"
const pfCsGoogleHostedConn = "packetfabric_cs_google_hosted_connection"
const pfCsAzureHostedConn = "packetfabric_cs_azure_hosted_connection"
const pfCsIbmHostedConn = "packetfabric_cs_ibm_hosted_connection"
const pfCsOracleHostedConn = "packetfabric_cs_oracle_hosted_connection"
const pfCsAwsDedicatedConn = "packetfabric_cs_aws_dedicated_connection"
const pfCsGoogleDedicatedConn = "packetfabric_cs_google_dedicated_connection"
const pfCsAzureDedicatedConn = "packetfabric_cs_azure_dedicated_connection"

// ###### Data-sources
const pfDataLocations = "data.packetfabric_locations"
const pfDataLocationsCloud = "data.packetfabric_locations_cloud"
const pfDataLocationsPortAvailability = "data.packetfabric_locations_port_availability"
const pfDataLocationsZones = "data.packetfabric_locations_pop_zones"
const pfDataLocationsRegions = "data.packetfabric_locations_regions"
const pfDataLocationsMarkets = "data.packetfabric_locations_markets"
const pfDataActivityLog = "data.packetfabric_activitylog"
const pfDataPorts = "data.packetfabric_ports"
const pfDataSourcePortVlans = "data.packetfabric_port_vlans"
const pfDataSourcePortDeviceInfo = "data.packetfabric_port_device_info"
const pfDataSourcePortRouterLogs = "data.packetfabric_port_router_logs"
const pfDataLinkAggregationGroups = "data.packetfabric_link_aggregation_group"
const pfDataBilling = "data.packetfabric_billing"
const pfDataCsAwsHostedConn = "data.packetfabric_cs_aws_hosted_connection"
const pfDataCsDedicatedConns = "data.packetfabric_cs_dedicated_connections"
const pfDataCloudRouterConnIpsec = "data.packetfabric_cloud_router_connection_ipsec"

// ########################################
// ###### HARDCODED VALUES
// ########################################

// common
const subscriptionTerm = 1

var labPopsPort = []string{"LAB1", "LAB4", "LAB5", "LAB6"}
var labPopsHostedCloud = []string{"LAB1", "LAB4", "LAB6"}
var labPopsDedicatedCloud = []string{"DEV1", "LAB4", "LAB5"}

// packetfabric_port
// packetfabric_point_to_point
const portSpeed = "1Gbps"

// packetfabric_port_loa
const PortLoaCustomerName = "loa"

// packetfabric_document
const TestFileName = "testdata/test_acc_document.pdf"

// packetfabric_link_aggregation_group
const LinkAggGroupInterval = "fast"

// packetfbaric_backbone_virtual_circuit
const backboneVCspeed = "50Mbps"
const backboneVCepl = false
const backboneVCvlan1Value = 103
const backboneVCvlan2Value = 104
const backboneVClonghaulType = "dedicated"

// packetfabric_backbone_virtual_circuit_speed_burst
const BackboneVcSpeedBurst = "100Mbps"

// packetfabric_cloud_router
const DefaultCloudRouterCapacity = "5Gbps"
const CloudRouterCapacityChange = "10Gbps"
const CloudRouterRegionUS = "US"
const CloudRouterRegionUK = "UK"
const CloudRouterASN = 4556

// packetfabric_cloud_router_connection_aws
// packetfabric_cloud_router_connection_google
const CloudRouterConnSpeed = "100Mbps" // must match const AzurePeeringBandwidth and IbmSpeed

// packetfabric_cloud_router_connection_google
// packetfabric_cs_google_hosted_connection
const GoogleRegion = "us-west1"
const GoogleNetwork = "default"

// packetfabric_cloud_router_connection_ibm
// packetfabric_cs_ibm_hosted_connection
const IbmBgpAsn = 64536
const IbmRegion = "us-east"
const IbmSpeed = 100 // must match const CloudRouterConnSpeed and HostedCloudSpeed

// packetfabric_cloud_router_connection_oracle
// packetfabric_cs_oracle_hosted_connection
const OracleProviderName = "PacketFabric"
const OracleBandwidth = "1 Gbps"
const OracleBgpAsn = 64537
const OracleAuthKey = "dd02c7c2232759874e1c20558"
const OracleBgpPeeringIp1 = "169.254.246.41/30"
const OracleBgpPeeringIp2 = "169.254.246.42/30"
const OracleBgpPeeringIp3 = "169.254.247.41/30"
const OracleBgpPeeringIp4 = "169.254.247.42/30"

// packetfabric_cloud_router_connection_port
const CloudRouterConnPortSpeed = "1Gbps"
const CloudRouterConnPortVlan = 101

// packetfabric_cloud_router_connection_ipsec
const CloudRouterConnIpsecGatewayAddress = "104.198.66.55"
const CloudRouterConnIpsecSpeed = "1Gbps"
const CloudRouterConnIpsecPhase1AuthenticationMethod = "pre-shared-key"
const CloudRouterConnIpsecPhase1Group = "group5"
const CloudRouterConnIpsecPhase1EncryptionAlgo = "aes-256-cbc"
const CloudRouterConnIpsecPhase1AuthenticationAlgo = "sha1"
const CloudRouterConnIpsecPhase2pfsGroup = "group5"
const CloudRouterConnIpsecPhase2EncryptionAlgo = "aes-128-gcm"
const CloudRouterConnIpsecPhase2AuthenticationAlgo = "hmac-sha1-96"
const CloudRouterConnIpsecSharedKey = "superCoolKey"
const CloudRouterConnIpsecIkeVersion = 1
const CloudRouterConnIpsecPhase1Lifetime = 10800
const CloudRouterConnIpsecPhase2Lifetime = 28800

// packetfabric_cloud_router_bgp_session
const CrbsAddressFmly = "v4"
const CloudRouterBgpSessionASN = 64534
const CloudRouterBgpSessionPrefix1 = "10.0.0.0/8"
const CloudRouterBgpSessionType1 = "in"
const CloudRouterBgpSessionPrefix2 = "192.168.1.0/24"
const CloudRouterBgpSessionType2 = "out"
const CloudRouterBgpSessionPrefix3 = "192.168.2.0/24"
const CloudRouterBgpSessionType3 = "out"
const CloudRouterBgpSessionRemoteAddress = "169.254.247.41/30"
const CloudRouterBgpSessionL3Address = "169.254.247.42/30"

// packetfabric_cs_aws_hosted_connection
// packetfabric_cs_azure_hosted_connection
// packetfabric_cs_google_hosted_connection
// packetfabric_cs_ibm_hosted_connection
// packetfabric_cs_oracle_hosted_connection
const HostedCloudSpeed = "100Mbps" // must match const AzurePeeringBandwidth and IbmSpeed
const HostedCloudVlan1 = 200
const HostedCloudVlan2 = 201
const HostedCloudVlan3 = 202
const HostedCloudVlan4 = 203
const HostedCloudVlan5 = 204

// packetfabric_cloud_router_connection_azure
// packetfabric_cs_azure_hosted_connection
const AzureLocationProd = "East US"
const AzureLocationDev = "West Central US"
const AzureVnetCidr = "10.7.0.0/16"
const AzureSubnetCidr = "10.7.1.0/24"
const AzurePeeringLocationProd = "New York"
const AzurePeeringLocationDev = "Denver Test"
const AzureServiceProviderNameProd = "PacketFabric"
const AzureServiceProviderNameDev = "Packet Fabric Test"
const AzureExpressRouteTier = "Standard"
const AzureExpressRouteFamily = "MeteredData"
const AzurePeeringBandwidth = 100 // must match const CloudRouterConnSpeed and HostedCloudSpeed

// packetfabric_cs_aws_dedicated_connection
// packetfabric_cs_google_dedicated_connection
// packetfabric_cs_azure_dedicated_connection
const DedicatedCloudSpeed = "10Gbps"
const DedicatedCloudServiceClass = "longhaul"
const DedicatedCloudAutoneg = false
const DedicatedCloudEncap = "dot1q"     // Azure only
const DedicatedCloudPortCat = "primary" // Azure only

type PortDetails struct {
	PFClient              *packetfabric.PFClient
	DesiredSpeed          string
	DesiredPop            string
	DesiredZone           string
	DesiredMedia          *string
	DesiredProvider       string
	DesiredConnectionType string
	DesiredMarket         string
	DesiredRegion         string
	DesiredCity           string
	DesiredState          string
	IsNatCapable          bool
	HasCloudRouter        bool
	AnyType               bool
	IsCloudConnection     bool
	PortEnabled           bool
	SkipDesiredMarket     *string
}

// ########################################
// ###### HCLs RESULTS FOR ASSERTIONS
// ########################################

// ###### Resources

type HclResultBase struct {
	Hcl                    string
	Resource               string
	ResourceName           string
	AdditionalResourceName string
}

// packetfabric_port
type RHclPortResult struct {
	HclResultBase
	ResourceName     string
	Description      string
	Media            string
	Pop              string
	Zone             string
	Speed            string
	SubscriptionTerm int
	Enabled          bool
	Market           string
}

// packetfabric_port_loa
type RHclPortLoaResult struct {
	HclResultBase
	Port             RHclPortResult
	LoaCustomerName  string
	DestinationEmail string
}

// packetfabric_link_aggregation_group
type RHclLinkAggregationGroupResult struct {
	HclResultBase
	Desc     string
	Interval string
	Members  []string
	Pop      string
}

// packetfabric_link_aggregation_group
type DHclLinkAggregationGroupsResult struct {
	HclResultBase
}

// packetfabric_outbound_cross_connect
type RHclOutboundCrossConnectResult struct {
	HclResultBase
	Desc string
	Port RHclPortResult
	Site string
}

// packetfabric_document
type RHclDocumentResult struct {
	HclResultBase
}

// packetfabric_backbone_virtual_circuit
type RHclBackboneVirtualCircuitResult struct {
	HclResultBase
	Desc               string
	Epl                bool
	InterfaceBackboneA InterfaceBackbone
	InterfaceBackboneZ InterfaceBackbone
	BandwidthBackbone
}

type InterfaceBackbone struct {
	PortCircuitID string
	Untagged      bool
	Vlan          int
}

type BandwidthBackbone struct {
	LonghaulType     string
	Speed            string
	SubscriptionTerm int
}

// packetfabric_backbone_virtual_circuit_speed_burst
type RHclBackboneVirtualCircuitSpeedBurstResult struct {
	HclResultBase
	Speed string
}

// packetfabric_point_to_point
type RHclPointToPointResult struct {
	HclResultBase
	Desc             string
	Speed            string
	Media            string
	SubscriptionTerm int
	Pop1             string
	Zone1            string
	Autoneg1         bool
	Pop2             string
	Zone2            string
	Autoneg2         bool
	UpdatedDesc      int
}

// packetfabric_cloud_router
type RHclCloudRouterResult struct {
	HclResultBase
	Asn      int
	Capacity string
	Regions  []string
}
type RHclCloudRouterInput struct {
	ResourceName string
	HclName      string
	Capacity     string
}

// packetfabric_cloud_router_connection_aws
type RHclCloudRouterConnectionAwsResult struct {
	HclResultBase
	CloudRouter  RHclCloudRouterResult
	AwsAccountID string
	AccountUuid  string
	Desc         string
	Pop          string
	Zone         string
	Speed        string
}

// packetfabric_cloud_router_connection_google
type RHclCloudRouterConnectionGoogleResult struct {
	HclResultBase
	CloudRouter RHclCloudRouterResult
	AccountUuid string
	Desc        string
	Pop         string
	Speed       string
}

// packetfabric_cloud_router_connection_azure
type RHclCloudRouterConnectionAzureResult struct {
	HclResultBase
	CloudRouter RHclCloudRouterResult
	AccountUuid string
	Desc        string
	Speed       string
}

// packetfabric_cloud_router_connection_ibm
type RHclCloudRouterConnectionIbmResult struct {
	HclResultBase
	CloudRouter RHclCloudRouterResult
	AccountUuid string
	Desc        string
	Pop         string
	Zone        string
	Speed       string
	IbmBgpAsn   int
}

// packetfabric_cloud_router_connection_oracle
type RHclCloudRouterConnectionOracleResult struct {
	HclResultBase
	CloudRouter RHclCloudRouterResult
	AccountUuid string
	Desc        string
	Pop         string
	Zone        string
}

// packetfabric_cloud_router_connection_port
type RHclCloudRouterConnectionPortResult struct {
	HclResultBase
	Desc        string
	CloudRouter RHclCloudRouterResult
	PortResult  RHclPortResult
	Speed       string
	Vlan        int
}

// packetfabric_cloud_router_connection_ipsec
type RHclCloudRouterConnectionIpsecResult struct {
	HclResultBase
	Port                       RHclPortResult
	Desc                       string
	Pop                        string
	Speed                      string
	GatewayAddress             string
	IkeVersion                 int
	Phase1AuthenticationMethod string
	Phase1Group                string
	Phase1EncryptionAlgo       string
	Phase1AuthenticationAlgo   string
	Phase1Lifetime             int
	Phase2PfsGroup             string
	Phase2EncryptionAlgo       string
	Phase2AuthenticationAlgo   string
	Phase2Lifetime             int
	SharedKey                  string
}

// packetfabric_cloud_router_bgp_session
type RHclBgpSessionResult struct {
	HclResultBase
	CloudRouter     RHclCloudRouterResult
	CloudRouterConn RHclCloudRouterConnectionAwsResult
	AddressFamily   string
	Asn             int
	RemoteAddress   string
	L3Address       string
	Prefix1         string
	Type1           string
	Prefix2         string
	Type2           string
}

// packetfabric_cs_aws_hosted_connection
type RHclCsHostedCloudAwsResult struct {
	HclResultBase
	PortResult   RHclPortResult
	AwsAccountID string
	AccountUuid  string
	Desc         string
	Pop          string
	Zone         string
	Speed        string
	Vlan         int
}

// packetfabric_cs_google_hosted_connection
type RHclCsHostedCloudGoogleResult struct {
	HclResultBase
	PortResult  RHclPortResult
	AccountUuid string
	Desc        string
	Speed       string
	Pop         string
	Vlan        int
}

// packetfabric_cs_azure_hosted_connection
type RHclCsHostedCloudAzureResult struct {
	HclResultBase
	PortResult  RHclPortResult
	AccountUuid string
	Desc        string
	Speed       string
	VlanPrivate int
}

// packetfabric_cs_ibm_hosted_connection
type RHclCsHostedCloudIbmResult struct {
	HclResultBase
	PortResult  RHclPortResult
	AccountUuid string
	Desc        string
	Pop         string
	Zone        string
	Vlan        int
	Speed       string
	IbmBgpAsn   int
}

// packetfabric_cs_oracle_hosted_connection
type RHclCsHostedCloudOracleResult struct {
	HclResultBase
	PortResult  RHclPortResult
	AccountUuid string
	Desc        string
	Pop         string
	Zone        string
	Vlan        int
}

// packetfabric_cs_aws_dedicated_connection
type RHclCsAwsDedicatedConnectionResult struct {
	HclResultBase
	Description      string
	Pop              string
	Zone             string
	SubscriptionTerm int
	ServiceClass     string
	Autoneg          bool
	Speed            string
}

// packetfabric_cs_google_dedicated_connection
type RHclCsGoogleDedicatedConnectionResult struct {
	HclResultBase
	Desc             string
	Pop              string
	Zone             string
	SubscriptionTerm int
	ServiceClass     string
	Autoneg          bool
	Speed            string
}

// packetfabric_cs_azure_dedicated_connection
type RHclCsAzureDedicatedConnectionResult struct {
	HclResultBase
	Desc             string
	Pop              string
	Zone             string
	SubscriptionTerm int
	ServiceClass     string
	Encapsulation    string
	PortCategory     string
	Speed            string
}

// ###### Data-sources

// data packetfabric_locations_cloud
type DHclLocationsCloudResult struct {
	HclResultBase
}

// data packetfabric_locations_port_availability
type DHclLocationsPortAvailabilityResult struct {
	HclResultBase
}

// data packetfabric_locations
type DHclLocationsResult struct {
	HclResultBase
}

// data packetfabric_locations_pop_zones
type DHclLocationsZonesResult struct {
	HclResultBase
}

// data packetfabric_locations_regions
type DHclLocationsRegionsResult struct {
	HclResultBase
}

// data packetfabric_locations_markets
type DHclLocationsMarketsResult struct {
	HclResultBase
}

// data packetfabric_activitylog
type DHclActivityLogResult struct {
	HclResultBase
}

// data packetfabric_billing
type DHclBillingResult struct {
	HclResultBase
}

// data packetfabric_port
type DHclPortsResult struct {
	HclResultBase
}

// data packetfabric_port_vlans
type DHclPortVlansResult struct {
	HclResultBase
}

// data packetfabric_port_device_info
type DHclPortDeviceInfoResult struct {
	HclResultBase
}

// data packetfabric_port_router_logs
type DHclPortRouterLogsResult struct {
	HclResultBase
}

// data packetfabric_cs_aws_hosted_connection
type DHclCsAwsHostedConnectionResult struct {
	HclResultBase
}

// data packetfabric_cs_dedicated_connections
type DHclDedicatedConnectionsResult struct {
	HclResultBase
}

// data packetfabric_cloud_router_connection_ipsec
type DHclCloudRouterConnIpsecResult struct {
	HclResultBase
}

// Patterns:
// Resource schema for required fields only
// - func RHcl...
// Resouce schema for required + optional fields
// - func OHcl...

// ########################################
// ###### HCLs FOR REQUIRED FIELDS
// ########################################

// ###### Resources

// packetfabric_port
func (details PortDetails) RHclPort(portEnabled bool) RHclPortResult {
	var pop, zone, media, speed, market string
	var err error

	log.Println("Getting pop and zone with available port for desired speed: ", details.DesiredSpeed)
	var SkipDesiredMarket *string
	if details.SkipDesiredMarket != nil {
		SkipDesiredMarket = details.SkipDesiredMarket
	}
	var DesiredMedia *string
	if details.DesiredMedia != nil {
		DesiredMedia = details.DesiredMedia
	}
	pop, zone, media, market, err = GetPopAndZoneWithAvailablePort(details.DesiredSpeed, SkipDesiredMarket, DesiredMedia, false)
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}
	speed = details.DesiredSpeed
	log.Println("Pop,Zone, media, market, and speed set to: ", pop, zone, media, market, speed)

	resourceName, hclName := GenerateUniqueResourceName(pfPort)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfPort, hclName, uniqueDesc)

	log.Println("Generating HCL")
	hcl := fmt.Sprintf(
		RResourcePort,
		hclName,
		uniqueDesc,
		media,
		pop,
		zone,
		speed,
		subscriptionTerm,
		portEnabled)

	log.Println("Returning HCL result")
	return RHclPortResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfPort,
			ResourceName: hclName,
		},
		ResourceName:     resourceName,
		Description:      uniqueDesc,
		Media:            media,
		Pop:              pop,
		Zone:             zone,
		Speed:            speed,
		SubscriptionTerm: subscriptionTerm,
		Enabled:          portEnabled,
		Market:           market,
	}
}

// packetfabric_port_loa
func RHclPortLoa() RHclPortLoaResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	portDetails := PortDetails{
		PFClient:     c,
		DesiredSpeed: portSpeed,
	}

	hclPortResult := portDetails.RHclPort(false)
	resourceName, hclName := GenerateUniqueResourceName(pfPortLoa)
	log.Printf("Resource: %s, Resource name: %s\n", pfPortLoa, hclName)

	email := os.Getenv("PF_USER_EMAIL")

	hcl := fmt.Sprintf(RResourcePortLoa,
		hclName,
		hclPortResult.ResourceName,
		PortLoaCustomerName,
		email,
	)

	return RHclPortLoaResult{
		HclResultBase: HclResultBase{
			Hcl:          fmt.Sprintf("%s\n%s", hclPortResult.Hcl, hcl),
			Resource:     pfPortLoa,
			ResourceName: resourceName,
		},
		LoaCustomerName:  PortLoaCustomerName,
		DestinationEmail: email,
	}
}

// packetfabric_document
func RHclDocumentMSA() RHclDocumentResult {
	resourceName, hclName := GenerateUniqueResourceName(pfDocument)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfDocument, hclName, uniqueDesc)

	hcl := fmt.Sprintf(RResourceDocumentMSA, hclName, TestFileName, uniqueDesc)

	return RHclDocumentResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDocument,
			ResourceName: resourceName,
		},
	}
}

// packetfabric_outbound_cross_connect
func RHclOutboundCrossConnect() RHclOutboundCrossConnectResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	portDetails := PortDetails{
		PFClient:     c,
		DesiredSpeed: portSpeed,
	}

	hclPortResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfOutboundCrossConnect)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfOutboundCrossConnect, hclName, uniqueDesc)

	locations, err := c.ListLocations()
	if err != nil {
		log.Fatal(err)
	}

	var site string
	for _, location := range locations {
		if location.Pop == hclPortResult.Pop {
			site = location.SiteCode
			break
		}
	}

	documentResult := RHclDocumentMSA()

	outboundCrossHcl := fmt.Sprintf(RResourceOutboundCrossConnect,
		hclName,
		uniqueDesc,
		documentResult.ResourceName,
		hclPortResult.ResourceName,
		site)

	hcl := fmt.Sprintf("%s\n%s\n%s", hclPortResult.Hcl, documentResult.Hcl, outboundCrossHcl)

	return RHclOutboundCrossConnectResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfOutboundCrossConnect,
			ResourceName: resourceName,
		},
		Desc: uniqueDesc,
		Site: site,
	}
}

// packetfabric_link_aggregation_group
func RHclLinkAggregationGroup() RHclLinkAggregationGroupResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	portDetails1 := PortDetails{
		PFClient:     c,
		DesiredSpeed: portSpeed,
	}

	hclPortResult1 := portDetails1.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfLinkAggregationGroup)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfLinkAggregationGroup, hclName, uniqueDesc)

	linkAggGroupHcl := fmt.Sprintf(RResourceLinkAggregationGroup,
		hclName,
		uniqueDesc,
		LinkAggGroupInterval,
		hclPortResult1.ResourceName,
		hclPortResult1.Pop,
		resourceName,
	)

	hcl := fmt.Sprintf("%s\n%s", hclPortResult1.Hcl, linkAggGroupHcl)

	return RHclLinkAggregationGroupResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfLinkAggregationGroup,
			ResourceName: resourceName,
		},
		Desc:     uniqueDesc,
		Interval: LinkAggGroupInterval,
		Members: []string{
			hclPortResult1.ResourceName,
		},
		Pop: hclPortResult1.Pop,
	}
}

// packetfabric_backbone_virtual_circuit
func RHclBackboneVirtualCircuitVlan() RHclBackboneVirtualCircuitResult {

	portDetailsA := CreateBasePortDetails()
	portTestResultA := portDetailsA.RHclPort(true)
	// Get the market from the first port
	marketA := portTestResultA.Market
	log.Println("Market from first port: ", marketA)

	portDetailsZ := CreateBasePortDetails()
	portDetailsZ.SkipDesiredMarket = &marketA // Send the market for the second port so it selects a different one to avoid to build a metro VC
	log.Println("Sending the market to the second port: ", *portDetailsZ.SkipDesiredMarket)
	portTestResultZ := portDetailsZ.RHclPort(true)

	resourceName, hclName := GenerateUniqueResourceName(pfBackboneVirtualCircuit)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfBackboneVirtualCircuit, hclName, uniqueDesc)

	backboneVirtualCircuitHcl := fmt.Sprintf(
		RResourceBackboneVirtualCircuitVlan,
		hclName,
		uniqueDesc,
		backboneVCepl,
		portTestResultA.ResourceName,
		backboneVCvlan1Value,
		portTestResultZ.ResourceName,
		backboneVCvlan2Value,
		backboneVClonghaulType,
		backboneVCspeed,
		subscriptionTerm,
	)

	hcl := fmt.Sprintf("%s\n%s\n%s", portTestResultA.Hcl, portTestResultZ.Hcl, backboneVirtualCircuitHcl)

	return RHclBackboneVirtualCircuitResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfBackboneVirtualCircuit,
			ResourceName: resourceName,
		},
		Desc: uniqueDesc,
		Epl:  backboneVCepl,
		InterfaceBackboneA: InterfaceBackbone{
			Vlan:          backboneVCvlan1Value,
			PortCircuitID: portTestResultA.ResourceName,
		},
		InterfaceBackboneZ: InterfaceBackbone{
			Vlan:          backboneVCvlan2Value,
			PortCircuitID: portTestResultZ.ResourceName,
		},
		BandwidthBackbone: BandwidthBackbone{
			LonghaulType:     backboneVClonghaulType,
			Speed:            backboneVCspeed,
			SubscriptionTerm: subscriptionTerm,
		},
	}
}

// packetfabric_backbone_virtual_circuit_speed_burst
func RHclBackboneVirtualCircuitSpeedBurst() RHclBackboneVirtualCircuitSpeedBurstResult {
	resourceName, hclName := GenerateUniqueResourceName(pfAddSpeedBurst)
	log.Printf("Resource: %s, Resource name: %s\n", pfAddSpeedBurst, hclName)

	backboneVirtualCircuitResult := RHclBackboneVirtualCircuitVlan()
	speedBurstHcl := fmt.Sprintf(RResourceBackboneVirtualCircuitSpeedBurst, hclName, backboneVirtualCircuitResult.ResourceName, BackboneVcSpeedBurst)

	return RHclBackboneVirtualCircuitSpeedBurstResult{
		HclResultBase: HclResultBase{
			Hcl:          fmt.Sprintf("%s\n%s", speedBurstHcl, backboneVirtualCircuitResult.Hcl),
			Resource:     pfAddSpeedBurst,
			ResourceName: resourceName,
		},
		Speed: BackboneVcSpeedBurst,
	}
}

// packetfabric_point_to_point
func RHclPointToPoint() RHclPointToPointResult {

	var speed = portSpeed
	pop1, zone1, media, market1, err := GetPopAndZoneWithAvailablePort(speed, nil, nil, false)
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}
	log.Println("Pop1, media, and speed set to: ", pop1, zone1, media, market1, speed)

	pop2, zone2, _, market2, err2 := GetPopAndZoneWithAvailablePort(speed, &market1, &media, false)
	if err2 != nil {
		log.Println("Error getting pop and zone with available port: ", err2)
		log.Panic(err2)
	}
	log.Println("Pop2, media, and speed set to: ", pop2, zone2, media, market2, speed)

	uniqueDesc := GenerateUniqueName()
	resourceName, hclName := GenerateUniqueResourceName(pfPoinToPoint)
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfPoinToPoint, hclName, uniqueDesc)

	hcl := fmt.Sprintf(RResourcePointToPoint,
		hclName,
		uniqueDesc,
		speed,
		media,
		subscriptionTerm,
		pop1,
		zone1,
		false,
		pop2,
		zone2,
		false,
		hclName,
		resourceName)

	return RHclPointToPointResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfPoinToPoint,
			ResourceName: resourceName,
		},
		Desc:             uniqueDesc,
		Speed:            speed,
		Media:            media,
		SubscriptionTerm: subscriptionTerm,
		Pop1:             pop1,
		Zone1:            zone1,
		Autoneg1:         false,
		Pop2:             pop2,
		Zone2:            zone2,
		Autoneg2:         false,
	}
}

// packetfabric_cloud_router
func DefaultRHclCloudRouterInput() RHclCloudRouterInput {
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouter)
	return RHclCloudRouterInput{
		ResourceName: resourceName,
		HclName:      hclName,
		Capacity:     DefaultCloudRouterCapacity,
	}
}
func RHclCloudRouter(input RHclCloudRouterInput) RHclCloudRouterResult {
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCloudRouter, input.HclName, uniqueDesc)

	hcl := fmt.Sprintf(
		RResourcePacketfabricCloudRouter,
		input.HclName,
		uniqueDesc,
		os.Getenv("PF_ACCOUNT_ID"),
		CloudRouterASN,
		input.Capacity,
		CloudRouterRegionUS,
		CloudRouterRegionUK)

	return RHclCloudRouterResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouter,
			ResourceName: input.ResourceName,
		},
		Asn:      CloudRouterASN,
		Capacity: input.Capacity,
		Regions:  []string{CloudRouterRegionUS, CloudRouterRegionUK},
	}
}

// packetfabric_cloud_router_connection_aws
func RHclCloudRouterConnectionAws() RHclCloudRouterConnectionAwsResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          CloudRouterConnSpeed,
		DesiredProvider:       "aws",
		DesiredConnectionType: "hosted",
		HasCloudRouter:        true,
		IsCloudConnection:     true,
	}

	pop, zone, _ := popDetails.FindAvailableCloudPopZone()

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnAws)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCloudRouterConnAws, hclName, uniqueDesc)

	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionAws,
		hclName,
		hclCloudRouterRes.ResourceName,
		os.Getenv("PF_AWS_ACCOUNT_ID"),
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		zone,
		CloudRouterConnSpeed)

	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, crcHcl)

	return RHclCloudRouterConnectionAwsResult{
		HclResultBase: HclResultBase{
			Hcl:                    hcl,
			Resource:               pfCloudRouterConnAws,
			ResourceName:           resourceName,
			AdditionalResourceName: hclCloudRouterRes.ResourceName,
		},
		AwsAccountID: os.Getenv("PF_AWS_ACCOUNT_ID"),
		AccountUuid:  os.Getenv("PF_ACCOUNT_ID"),
		Speed:        CloudRouterConnSpeed,
		Pop:          pop,
		Zone:         zone,
	}
}

// packetfabric_cloud_router_connection_google
func RHclCloudRouterConnectionGoogle() RHclCloudRouterConnectionGoogleResult {

	var edgeAvailabilityDomain string

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          CloudRouterConnSpeed,
		DesiredProvider:       "google",
		DesiredConnectionType: "hosted",
		HasCloudRouter:        true,
		IsCloudConnection:     true,
	}

	pop, _, _ := popDetails.FindAvailableCloudPopZone()

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnGoogle)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCloudRouterConnGoogle, hclName, uniqueDesc)

	if pop == "LAB1" {
		edgeAvailabilityDomain = "AVAILABILITY_DOMAIN_2"
	} else {
		edgeAvailabilityDomain = "AVAILABILITY_DOMAIN_1"
	}
	log.Printf("Edge Availability Domain %s\n", edgeAvailabilityDomain)

	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionGoogle,
		GoogleRegion,
		GoogleNetwork,
		GoogleRegion,
		edgeAvailabilityDomain,
		hclName,
		hclCloudRouterRes.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		CloudRouterConnSpeed)

	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, crcHcl)

	return RHclCloudRouterConnectionGoogleResult{
		HclResultBase: HclResultBase{
			Hcl:                    hcl,
			Resource:               pfCloudRouterConnGoogle,
			ResourceName:           resourceName,
			AdditionalResourceName: hclCloudRouterRes.ResourceName,
		},
		Desc:        uniqueDesc,
		AccountUuid: os.Getenv("PF_ACCOUNT_ID"),
		Speed:       CloudRouterConnSpeed,
		Pop:         pop,
	}
}

// packetfabric_cloud_router_connection_azure
func RHclCloudRouterConnectionAzure() RHclCloudRouterConnectionAzureResult {

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnAzure)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCloudRouterConnAzure, hclName, uniqueDesc)

	host := os.Getenv("PF_HOST")
	AzureLocation, AzurePeeringLocation, AzureServiceProviderName := setAzureLocations(host)

	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionAzure,
		AzureLocation,
		AzureLocation,
		AzureVnetCidr,
		AzureSubnetCidr,
		AzureLocation,
		AzurePeeringLocation,
		AzureServiceProviderName,
		AzurePeeringBandwidth,
		AzureExpressRouteTier,
		AzureExpressRouteFamily,
		hclName,
		hclCloudRouterRes.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		CloudRouterConnSpeed)

	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, crcHcl)

	return RHclCloudRouterConnectionAzureResult{
		HclResultBase: HclResultBase{
			Hcl:                    hcl,
			Resource:               pfCloudRouterConnAzure,
			ResourceName:           resourceName,
			AdditionalResourceName: hclCloudRouterRes.ResourceName,
		},
		Desc:        uniqueDesc,
		AccountUuid: os.Getenv("PF_ACCOUNT_ID"),
		Speed:       CloudRouterConnSpeed,
	}
}

// packetfabric_cloud_router_connection_ibm
func RHclCloudRouterConnectionIbm() RHclCloudRouterConnectionIbmResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          CloudRouterConnSpeed,
		DesiredProvider:       "ibm",
		DesiredConnectionType: "hosted",
		HasCloudRouter:        true,
		IsCloudConnection:     true,
	}

	pop, zone, _ := popDetails.FindAvailableCloudPopZone()

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnIbm)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCloudRouterConnIbm, hclName, uniqueDesc)

	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionIbm,
		hclName,
		hclCloudRouterRes.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		zone,
		CloudRouterConnSpeed,
		IbmBgpAsn,
		IbmRegion,
		uniqueDesc,
		IbmBgpAsn,
		IbmSpeed)

	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, crcHcl)

	return RHclCloudRouterConnectionIbmResult{
		HclResultBase: HclResultBase{
			Hcl:                    hcl,
			Resource:               pfCloudRouterConnIbm,
			ResourceName:           resourceName,
			AdditionalResourceName: hclCloudRouterRes.ResourceName,
		},
		Desc:        uniqueDesc,
		AccountUuid: os.Getenv("PF_ACCOUNT_ID"),
		Pop:         pop,
		Zone:        zone,
		Speed:       CloudRouterConnSpeed,
		IbmBgpAsn:   IbmBgpAsn,
	}
}

// packetfabric_cloud_router_connection_oracle
func RHclCloudRouterConnectionOracle() RHclCloudRouterConnectionOracleResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          CloudRouterConnSpeed,
		DesiredProvider:       "oracle",
		DesiredConnectionType: "hosted",
		HasCloudRouter:        true,
		IsCloudConnection:     true,
	}

	pop, zone, region := popDetails.FindAvailableCloudPopZone()

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnOracle)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCloudRouterConnOracle, hclName, uniqueDesc)

	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionOracle,
		region,
		OracleProviderName,
		region,
		OracleBandwidth,
		OracleBgpAsn,
		OracleAuthKey,
		OracleBgpPeeringIp1,
		OracleBgpPeeringIp2,
		hclName,
		hclCloudRouterRes.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		zone,
		region)

	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, crcHcl)

	return RHclCloudRouterConnectionOracleResult{
		HclResultBase: HclResultBase{
			Hcl:                    hcl,
			Resource:               pfCloudRouterConnOracle,
			ResourceName:           resourceName,
			AdditionalResourceName: hclCloudRouterRes.ResourceName,
		},
		Desc:        uniqueDesc,
		AccountUuid: os.Getenv("PF_ACCOUNT_ID"),
		Pop:         pop,
		Zone:        zone,
	}
}

// packetfabric_cloud_router_connection_port
func RHclCloudRouterConnectionPort() RHclCloudRouterConnectionPortResult {

	portDetails := CreateBasePortDetails()

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnPort)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCloudRouterConnPort, hclName, uniqueDesc)

	crConnPortHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionPort,
		hclName,
		uniqueDesc,
		hclCloudRouterRes.ResourceName,
		portTestResult.ResourceName,
		CloudRouterConnPortSpeed,
		CloudRouterConnPortVlan,
	)

	hcl := fmt.Sprintf("%s\n%s\n%s", portTestResult.Hcl, hclCloudRouterRes.Hcl, crConnPortHcl)

	return RHclCloudRouterConnectionPortResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouterConnPort,
			ResourceName: resourceName,
		},
		CloudRouter: hclCloudRouterRes,
		PortResult:  portTestResult,
		Desc:        uniqueDesc,
		Speed:       CloudRouterConnPortSpeed,
		Vlan:        CloudRouterConnPortVlan,
	}
}

// packetfabric_cloud_router_connection_ipsec
func RHclCloudRouterConnectionIpsec() RHclCloudRouterConnectionIpsecResult {

	pop, _, _, _, err := GetPopAndZoneWithAvailablePort(CloudRouterConnIpsecSpeed, nil, nil, true)
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}

	hclCloudRouterResult := RHclCloudRouter(DefaultRHclCloudRouterInput())

	uniqueDesc := GenerateUniqueName()
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnIpsec)
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCloudRouterConnIpsec, hclName, uniqueDesc)

	cloudRouterIpsecHcl := fmt.Sprintf(RResourceCloudRouterConnectionIpsec,
		hclName,
		uniqueDesc,
		hclCloudRouterResult.ResourceName,
		pop,
		CloudRouterConnIpsecSpeed,
		CloudRouterConnIpsecGatewayAddress,
		CloudRouterConnIpsecIkeVersion,
		CloudRouterConnIpsecPhase1AuthenticationMethod,
		CloudRouterConnIpsecPhase1Group,
		CloudRouterConnIpsecPhase1EncryptionAlgo,
		CloudRouterConnIpsecPhase1AuthenticationAlgo,
		CloudRouterConnIpsecPhase1Lifetime,
		CloudRouterConnIpsecPhase2pfsGroup,
		CloudRouterConnIpsecPhase2EncryptionAlgo,
		CloudRouterConnIpsecPhase2AuthenticationAlgo,
		CloudRouterConnIpsecPhase2Lifetime,
		CloudRouterConnIpsecSharedKey,
	)

	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterResult.Hcl, cloudRouterIpsecHcl)

	return RHclCloudRouterConnectionIpsecResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouterConnIpsec,
			ResourceName: resourceName,
		},
		Desc:                       uniqueDesc,
		Pop:                        pop,
		Speed:                      CloudRouterConnIpsecSpeed,
		GatewayAddress:             CloudRouterConnIpsecGatewayAddress,
		IkeVersion:                 CloudRouterConnIpsecIkeVersion,
		Phase1AuthenticationMethod: CloudRouterConnIpsecPhase1AuthenticationMethod,
		Phase1Group:                CloudRouterConnIpsecPhase1Group,
		Phase1EncryptionAlgo:       CloudRouterConnIpsecPhase1EncryptionAlgo,
		Phase1AuthenticationAlgo:   CloudRouterConnIpsecPhase1AuthenticationAlgo,
		Phase1Lifetime:             CloudRouterConnIpsecPhase1Lifetime,
		Phase2PfsGroup:             CloudRouterConnIpsecPhase2pfsGroup,
		Phase2EncryptionAlgo:       CloudRouterConnIpsecPhase2EncryptionAlgo,
		Phase2AuthenticationAlgo:   CloudRouterConnIpsecPhase2AuthenticationAlgo,
		Phase2Lifetime:             CloudRouterConnIpsecPhase2Lifetime,
		SharedKey:                  CloudRouterConnIpsecSharedKey,
	}
}

// packetfabric_cloud_router_bgp_session
func RHclBgpSession() RHclBgpSessionResult {

	hclCloudConnRes := RHclCloudRouterConnectionAws()

	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterBgpSession)
	log.Printf("Resource: %s, Resource name: %s\n", pfCloudRouterBgpSession, hclName)

	bgpSessionHcl := fmt.Sprintf(
		RResourceCloudRouterBgpSession,
		hclName,
		hclCloudConnRes.HclResultBase.AdditionalResourceName,
		hclCloudConnRes.HclResultBase.ResourceName,
		CloudRouterBgpSessionRemoteAddress,
		CloudRouterBgpSessionL3Address,
		CloudRouterBgpSessionASN,
		CloudRouterBgpSessionPrefix1,
		CloudRouterBgpSessionType1,
		CloudRouterBgpSessionPrefix2,
		CloudRouterBgpSessionType2)
	hcl := fmt.Sprintf("%s\n%s", hclCloudConnRes.Hcl, bgpSessionHcl)

	return RHclBgpSessionResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouterBgpSession,
			ResourceName: resourceName,
		},
		CloudRouter:     hclCloudConnRes.CloudRouter,
		CloudRouterConn: hclCloudConnRes,
		RemoteAddress:   CloudRouterBgpSessionRemoteAddress,
		L3Address:       CloudRouterBgpSessionL3Address,
		Asn:             CloudRouterBgpSessionASN,
		Prefix1:         CloudRouterBgpSessionPrefix1,
		Type1:           CloudRouterBgpSessionType1,
		Prefix2:         CloudRouterBgpSessionPrefix2,
		Type2:           CloudRouterBgpSessionType2,
	}
}

// packetfabric_cs_aws_hosted_connection
func RHclCsAwsHostedConnection() RHclCsHostedCloudAwsResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          HostedCloudSpeed,
		DesiredProvider:       "aws",
		DesiredConnectionType: "hosted",
		IsCloudConnection:     true,
	}
	pop, zone, _ := popDetails.FindAvailableCloudPopZone()

	portDetails := CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCsAwsHostedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCsAwsHostedConn, hclName, uniqueDesc)

	awsHostedConnectionHcl := fmt.Sprintf(
		RResourceCSAwsHostedConnection,
		hclName,
		portTestResult.ResourceName,
		os.Getenv("PF_AWS_ACCOUNT_ID"),
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		zone,
		HostedCloudSpeed,
		HostedCloudVlan1)

	hcl := fmt.Sprintf("%s\n%s", portTestResult.Hcl, awsHostedConnectionHcl)

	return RHclCsHostedCloudAwsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsAwsHostedConn,
			ResourceName: resourceName,
		},
		PortResult:   portTestResult,
		AwsAccountID: os.Getenv("PF_AWS_ACCOUNT_ID"),
		AccountUuid:  os.Getenv("PF_ACCOUNT_ID"),
		Speed:        HostedCloudSpeed,
		Pop:          pop,
		Zone:         zone,
		Vlan:         HostedCloudVlan1,
	}
}

// packetfabric_cs_google_hosted_connection
func RHclCsGoogleHostedConnection() RHclCsHostedCloudGoogleResult {
	var edgeAvailabilityDomain string

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          HostedCloudSpeed,
		DesiredProvider:       "google",
		DesiredConnectionType: "hosted",
		IsCloudConnection:     true,
	}
	pop, _, _ := popDetails.FindAvailableCloudPopZone()

	portDetails := CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCsGoogleHostedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCsGoogleHostedConn, hclName, uniqueDesc)

	if pop == "LAB1" {
		edgeAvailabilityDomain = "AVAILABILITY_DOMAIN_2"
	} else {
		edgeAvailabilityDomain = "AVAILABILITY_DOMAIN_1"
	}
	log.Printf("Edge Availability Domain %s\n", edgeAvailabilityDomain)

	googleHostedConnectionHcl := fmt.Sprintf(
		RResourceCSGoogleHostedConnection,
		GoogleRegion,
		GoogleNetwork,
		GoogleRegion,
		edgeAvailabilityDomain,
		hclName,
		portTestResult.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		HostedCloudSpeed,
		HostedCloudVlan2)

	hcl := fmt.Sprintf("%s\n%s", portTestResult.Hcl, googleHostedConnectionHcl)

	return RHclCsHostedCloudGoogleResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsGoogleHostedConn,
			ResourceName: resourceName,
		},
		PortResult:  portTestResult,
		AccountUuid: os.Getenv("PF_ACCOUNT_ID"),
		Speed:       HostedCloudSpeed,
		Pop:         pop,
		Vlan:        HostedCloudVlan2,
	}
}

// packetfabric_cs_azure_hosted_connection
func RHclCsAzureHostedConnection() RHclCsHostedCloudAzureResult {

	portDetails := CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCsAzureHostedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCsAzureHostedConn, hclName, uniqueDesc)

	host := os.Getenv("PF_HOST")
	AzureLocation, AzurePeeringLocation, AzureServiceProviderName := setAzureLocations(host)

	azureHostedConnectionHcl := fmt.Sprintf(
		RResourceCSAzureHostedConnection,
		AzureLocation,
		AzureLocation,
		AzureVnetCidr,
		AzureSubnetCidr,
		AzureLocation,
		AzurePeeringLocation,
		AzureServiceProviderName,
		AzurePeeringBandwidth,
		AzureExpressRouteTier,
		AzureExpressRouteFamily,
		hclName,
		portTestResult.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		HostedCloudSpeed,
		HostedCloudVlan3)

	hcl := fmt.Sprintf("%s\n%s", portTestResult.Hcl, azureHostedConnectionHcl)

	return RHclCsHostedCloudAzureResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsAzureHostedConn,
			ResourceName: resourceName,
		},
		PortResult:  portTestResult,
		AccountUuid: os.Getenv("PF_ACCOUNT_ID"),
		Speed:       HostedCloudSpeed,
		VlanPrivate: HostedCloudVlan3,
	}
}

// packetfabric_cs_ibm_hosted_connection
func RHclCsIbmHostedConnection() RHclCsHostedCloudIbmResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          HostedCloudSpeed,
		DesiredProvider:       "ibm",
		DesiredConnectionType: "hosted",
		IsCloudConnection:     true,
	}
	pop, zone, _ := popDetails.FindAvailableCloudPopZone()

	portDetails := CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCsIbmHostedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCsIbmHostedConn, hclName, uniqueDesc)

	IbmHostedConnectionHcl := fmt.Sprintf(
		RResourceCSIbmHostedConnection,
		hclName,
		portTestResult.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		zone,
		HostedCloudSpeed,
		HostedCloudVlan4,
		IbmBgpAsn,
		IbmRegion,
		uniqueDesc,
		IbmBgpAsn,
		IbmSpeed)

	hcl := fmt.Sprintf("%s\n%s", portTestResult.Hcl, IbmHostedConnectionHcl)

	return RHclCsHostedCloudIbmResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsIbmHostedConn,
			ResourceName: resourceName,
		},
		PortResult:  portTestResult,
		AccountUuid: os.Getenv("PF_ACCOUNT_ID"),
		Pop:         pop,
		Zone:        zone,
		Speed:       HostedCloudSpeed,
		Vlan:        HostedCloudVlan4,
		IbmBgpAsn:   IbmBgpAsn,
	}
}

// packetfabric_cs_oracle_hosted_connection
func RHclCsOracleHostedConnection() RHclCsHostedCloudOracleResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          HostedCloudSpeed,
		DesiredProvider:       "oracle",
		DesiredConnectionType: "hosted",
		IsCloudConnection:     true,
	}
	pop, zone, region := popDetails.FindAvailableCloudPopZone()

	portDetails := CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCsOracleHostedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCsOracleHostedConn, hclName, uniqueDesc)

	oracleHostedConnectionHcl := fmt.Sprintf(
		RResourceCSOracleHostedConnection,
		region,
		OracleProviderName,
		region,
		OracleBandwidth,
		OracleBgpAsn,
		OracleAuthKey,
		OracleBgpPeeringIp3,
		OracleBgpPeeringIp4,
		hclName,
		portTestResult.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		zone,
		HostedCloudVlan5,
		region)

	hcl := fmt.Sprintf("%s\n%s", portTestResult.Hcl, oracleHostedConnectionHcl)

	return RHclCsHostedCloudOracleResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsOracleHostedConn,
			ResourceName: resourceName,
		},
		PortResult:  portTestResult,
		AccountUuid: os.Getenv("PF_ACCOUNT_ID"),
		Pop:         pop,
		Zone:        zone,
		Vlan:        HostedCloudVlan5,
	}
}

// packetfabric_cs_aws_dedicated_connection
func RHclCsAwsDedicatedConnection() RHclCsAwsDedicatedConnectionResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	resourceName, hclName := GenerateUniqueResourceName(pfCsAwsDedicatedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCsAwsDedicatedConn, hclName, uniqueDesc)

	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          DedicatedCloudSpeed,
		DesiredProvider:       "aws",
		DesiredConnectionType: "dedicated",
		IsCloudConnection:     true,
	}
	pop, zone, region, err := popDetails.FindAvailableCloudPopZoneDedicated()
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}

	hcl := fmt.Sprintf(RResourceCSAwsDedicatedConnection,
		hclName,
		region,
		uniqueDesc,
		pop,
		zone,
		subscriptionTerm,
		DedicatedCloudServiceClass,
		DedicatedCloudAutoneg,
		DedicatedCloudSpeed,
	)

	return RHclCsAwsDedicatedConnectionResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsAwsDedicatedConn,
			ResourceName: resourceName,
		},
		Description:      uniqueDesc,
		Pop:              pop,
		Zone:             zone,
		SubscriptionTerm: subscriptionTerm,
		ServiceClass:     DedicatedCloudServiceClass,
		Autoneg:          DedicatedCloudAutoneg,
		Speed:            DedicatedCloudSpeed,
	}
}

// packetfabric_cs_google_dedicated_connection
func RHclCsGoogleDedicatedConnection() RHclCsGoogleDedicatedConnectionResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	resourceName, hclName := GenerateUniqueResourceName(pfCsGoogleDedicatedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCsGoogleDedicatedConn, hclName, uniqueDesc)

	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          DedicatedCloudSpeed,
		DesiredProvider:       "google",
		DesiredConnectionType: "dedicated",
		IsCloudConnection:     true,
	}
	pop, zone, _, err := popDetails.FindAvailableCloudPopZoneDedicated()
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}

	hcl := fmt.Sprintf(RResourceCSGoogleDedicatedConnection,
		hclName,
		uniqueDesc,
		pop,
		zone,
		subscriptionTerm,
		DedicatedCloudServiceClass,
		DedicatedCloudAutoneg,
		DedicatedCloudSpeed,
	)

	return RHclCsGoogleDedicatedConnectionResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsGoogleDedicatedConn,
			ResourceName: resourceName,
		},
		Desc:             uniqueDesc,
		Pop:              pop,
		Zone:             zone,
		SubscriptionTerm: subscriptionTerm,
		ServiceClass:     DedicatedCloudServiceClass,
		Autoneg:          DedicatedCloudAutoneg,
		Speed:            DedicatedCloudSpeed,
	}
}

// packetfabric_cs_azure_dedicated_connection
func RHclCsAzureDedicatedConnection() RHclCsAzureDedicatedConnectionResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	resourceName, hclName := GenerateUniqueResourceName(pfCsAzureDedicatedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource: %s, Resource name: %s, description: %s\n", pfCsAzureDedicatedConn, hclName, uniqueDesc)

	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          DedicatedCloudSpeed,
		DesiredProvider:       "azure",
		DesiredConnectionType: "dedicated",
		IsCloudConnection:     true,
	}
	pop, zone, _, err := popDetails.FindAvailableCloudPopZoneDedicated()
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}

	hcl := fmt.Sprintf(
		RResourceCSAzureDedicatedConnection,
		hclName,
		uniqueDesc,
		pop,
		zone,
		subscriptionTerm,
		DedicatedCloudServiceClass,
		DedicatedCloudEncap,
		DedicatedCloudPortCat,
		DedicatedCloudSpeed)

	return RHclCsAzureDedicatedConnectionResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsAzureDedicatedConn,
			ResourceName: resourceName,
		},
		Desc:             uniqueDesc,
		Pop:              pop,
		Zone:             zone,
		SubscriptionTerm: subscriptionTerm,
		ServiceClass:     DedicatedCloudServiceClass,
		Encapsulation:    DedicatedCloudEncap,
		PortCategory:     DedicatedCloudPortCat,
		Speed:            DedicatedCloudSpeed,
	}
}

// ###### Data-sources

// data.packetfabric_locations
func DHclLocations() DHclLocationsResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocations)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataLocations, hclName)

	hcl := fmt.Sprintf(DDatasourceLocations, hclName)

	return DHclLocationsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocations,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations_cloud
func DHclLocationsCloud(cloudProvider, cloudConnectionType string) DHclLocationsCloudResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsCloud)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataLocationsCloud, hclName)
	hcl := fmt.Sprintf(DDataSourceLocationsCloud, hclName, cloudProvider, cloudConnectionType)

	return DHclLocationsCloudResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsCloud,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations_pop_zones
func DHclZones() DHclLocationsZonesResult {

	pop, _, _, _, err := GetPopAndZoneWithAvailablePort(portSpeed, nil, nil, false)
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsZones)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataLocationsZones, hclName)

	hcl := fmt.Sprintf(DDatasourceLocationsPopZones, hclName, pop)

	return DHclLocationsZonesResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsZones,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations_regions
func DHclLocationsRegions() DHclLocationsRegionsResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsRegions)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataLocationsRegions, hclName)

	hcl := fmt.Sprintf(DDataSourceLocationsRegions, hclName)

	return DHclLocationsRegionsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsRegions,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations_markets
func DHclLocationsMarkets() DHclLocationsMarketsResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsMarkets)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataLocationsMarkets, hclName)

	hcl := fmt.Sprintf(DDataSourceLocationsMarkets, hclName)

	return DHclLocationsMarketsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsMarkets,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations_port_availability
func DHclLocationsPortAvailability() DHclLocationsPortAvailabilityResult {

	pop, _, _, _, err := GetPopAndZoneWithAvailablePort(portSpeed, nil, nil, false)
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsPortAvailability)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataLocationsPortAvailability, hclName)

	hcl := fmt.Sprintf(DDataSourceLocationsPortAvailability, hclName, pop)

	return DHclLocationsPortAvailabilityResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsPortAvailability,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_activitylog
func DHclActivityLog() DHclActivityLogResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataActivityLog)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataActivityLog, hclName)

	hcl := fmt.Sprintf(DDatasourceActivityLog, hclName)

	return DHclActivityLogResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataActivityLog,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_billing
func DHclBilling() DHclBillingResult {

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())

	resourceName, hclName := GenerateUniqueResourceName(pfDataBilling)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataBilling, hclName)

	billingHcl := fmt.Sprintf(DDatasourceBilling,
		hclName,
		hclCloudRouterRes.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, billingHcl)

	return DHclBillingResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataBilling,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_ports
func DHclPorts() DHclPortsResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	portDetails := PortDetails{
		PFClient:     c,
		DesiredSpeed: portSpeed,
	}

	resourceName, hclName := GenerateUniqueResourceName(pfDataPorts)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataPorts, hclName)

	dataPortHcl := fmt.Sprintf(DDataSourcePorts, hclName)

	hcl := fmt.Sprintf("%s\n%s", portDetails.RHclPort(false).Hcl, dataPortHcl)

	return DHclPortsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataPorts,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_port_vlans
func DHclPortVlans() DHclPortVlansResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	portDetails := PortDetails{
		PFClient:     c,
		DesiredSpeed: portSpeed,
	}

	resourceName, hclName := GenerateUniqueResourceName(pfDataSourcePortVlans)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataSourcePortVlans, hclName)

	portResult := portDetails.RHclPort(false)
	portVlansHcl := fmt.Sprintf(
		DDataSourcePortVlans,
		hclName,
		portResult.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", portResult.Hcl, portVlansHcl)

	return DHclPortVlansResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataSourcePortVlans,
			ResourceName: resourceName,
		},
	}
}

// data packetfabric_port_device_info
func DHclPortDeviceInfo() DHclPortDeviceInfoResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	portDetails := PortDetails{
		PFClient:     c,
		DesiredSpeed: portSpeed,
	}

	portResult := portDetails.RHclPort(false)
	resourceName, hclName := GenerateUniqueResourceName(pfDataSourcePortDeviceInfo)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataSourcePortDeviceInfo, hclName)

	portDeviceInfoHcl := fmt.Sprintf(
		DDataSourcePortDeviceInfo,
		hclName,
		portResult.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", portResult.Hcl, portDeviceInfoHcl)

	return DHclPortDeviceInfoResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataSourcePortDeviceInfo,
			ResourceName: resourceName,
		},
	}
}

// data packetfabric_port_router_logs
func DHclPortRouterLogs() DHclPortRouterLogsResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	portDetails := PortDetails{
		PFClient:     c,
		DesiredSpeed: portSpeed,
	}

	now := time.Now()
	timeTo := now.Format("2006-01-02 15:04:05")
	timeFrom := now.Add(-time.Hour).Format("2006-01-02 15:04:05")

	resourceName, hclName := GenerateUniqueResourceName(pfDataSourcePortRouterLogs)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataSourcePortDeviceInfo, hclName)

	portResult := portDetails.RHclPort(true)
	dataSourcePortRouterLogsHcl := fmt.Sprintf(
		DDataSourcePortRouterLogs,
		hclName,
		portResult.ResourceName,
		timeFrom,
		timeTo,
	)

	hcl := fmt.Sprintf("%s\n%s", portResult.Hcl, dataSourcePortRouterLogsHcl)

	return DHclPortRouterLogsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataSourcePortRouterLogs,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_link_aggregation_group
func DHclLinkAggregationGroups() DHclLinkAggregationGroupsResult {
	linkAggregationGroupResult := RHclLinkAggregationGroup()
	resourceName, hclName := GenerateUniqueResourceName(pfDataLinkAggregationGroups)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataLinkAggregationGroups, hclName)

	linkAggregationGroupsHcl := fmt.Sprintf(
		DDatasourceLinkAggregationGroups,
		hclName,
		linkAggregationGroupResult.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", linkAggregationGroupResult.Hcl, linkAggregationGroupsHcl)

	return DHclLinkAggregationGroupsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLinkAggregationGroups,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_cs_aws_hosted_connection
func DHclHostedAwsConn() DHclCsAwsHostedConnectionResult {

	csAwsHostedConnectionResult := RHclCsAwsHostedConnection()

	resourceName, hclName := GenerateUniqueResourceName(pfDataCsAwsHostedConn)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataCsAwsHostedConn, hclName)

	hostedAwsConnHcl := fmt.Sprintf(DDatasourceCsAwsHostedConn,
		hclName,
		csAwsHostedConnectionResult.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", csAwsHostedConnectionResult.Hcl, hostedAwsConnHcl)

	return DHclCsAwsHostedConnectionResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataCsAwsHostedConn,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_cs_dedicated_connections
func DHclDedicatedConnections() DHclDedicatedConnectionsResult {

	csAwsDedicatedConnectionResult := RHclCsAwsDedicatedConnection()

	resourceName, hclName := GenerateUniqueResourceName(pfDataCsDedicatedConns)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataCsDedicatedConns, hclName)

	dedicatedConnHCL := fmt.Sprintf(DDatasourceDedicatedConns, hclName)

	hcl := fmt.Sprintf("%s\n%s", csAwsDedicatedConnectionResult.Hcl, dedicatedConnHCL)

	return DHclDedicatedConnectionsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataCsDedicatedConns,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_cloud_router_connection_ipsec
func DHclCloudRouterConnIpsec() DHclCloudRouterConnIpsecResult {

	cloudRouterConnectionIpsecResult := RHclCloudRouterConnectionIpsec()

	resourceName, hclName := GenerateUniqueResourceName(pfDataCloudRouterConnIpsec)
	log.Printf("Data-source: %s, Data-source name: %s\n", pfDataCsDedicatedConns, hclName)

	dataCloudRouterIpsecHcl := fmt.Sprintf(
		DDatasourceCloudRouterConnectionIpsec,
		hclName,
		cloudRouterConnectionIpsecResult.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", cloudRouterConnectionIpsecResult.Hcl, dataCloudRouterIpsecHcl)

	return DHclCloudRouterConnIpsecResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataCloudRouterConnIpsec,
			ResourceName: resourceName,
		},
	}
}
