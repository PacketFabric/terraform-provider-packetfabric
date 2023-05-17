package testutil

import (
	"fmt"
	"log"
	"os"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
)

// ########################################
// ###### RESOURCES & DATA-SOURCES CONSTs
// ########################################

// ###### Resources
const pfPort = "packetfabric_port"
const pfBackboneVirtualCircuit = "packetfabric_backbone_virtual_circuit"
const pfPoinToPoint = "packetfabric_point_to_point"
const pfCloudRouter = "packetfabric_cloud_router"
const pfCloudRouterConnAws = "packetfabric_cloud_router_connection_aws"
const pfCloudRouterConnGoogle = "packetfabric_cloud_router_connection_google"
const pfCloudRouterConnAzure = "packetfabric_cloud_router_connection_azure"
const pfCloudRouterConnPort = "packetfabric_cloud_router_connection_port"
const pfCloudRouterBgpSession = "packetfabric_cloud_router_bgp_session"
const pfCsAwsHostedConn = "packetfabric_cs_aws_hosted_connection"
const pfCsGoogleHostedConn = "packetfabric_cs_google_hosted_connection"
const pfCsAzureHostedConn = "packetfabric_cs_azure_hosted_connection"
const pfCsAwsDedicatedConn = "packetfabric_cs_aws_dedicated_connection"
const pfCsGoogleDedicatedConn = "packetfabric_cs_google_dedicated_connection"
const pfCsAzureDedicatedConn = "packetfabric_cs_azure_dedicated_connection"
const pfLinkAggregationGroup = "packetfabric_link_aggregation_group"
const pfPortLoa = "packetfabric_port_loa"

// ###### Data-sources
const pfDataLocations = "data.packetfabric_locations"
const pfDataLocationsCloud = "data.packetfabric_locations_cloud"
const pfDataLocationsPortAvailability = "data.packetfabric_locations_port_availability"
const pfDataLocationsZones = "data.packetfabric_locations_pop_zones"
const pfDataLocationsRegions = "data.packetfabric_locations_regions"
const pfDataLocationsMarkets = "data.packetfabric_locations_markets"
const pfDataActivityLog = "data.packetfabric_activitylog"
const pfDataPort = "data.packetfabric_ports"
const pfDataBilling = "data.packetfabric_billing"
const pfDatasourceCsAwsHostedConn = "data.packetfabric_cs_aws_hosted_connection"
const pfDatasourceLinkAggregationGroups = "data.packetfabric_link_aggregation_group"

// ########################################
// ###### HARDCODED VALUES
// ########################################

// common
const subscriptionTerm = 1

// packetfabric_port
// packetfabric_point_to_point
const portSpeed = "1Gbps"

var listPortsLab = []string{"LAB1", "LAB2", "LAB4", "LAB6", "LAB8"}

// packetfabric_port_loa
const PortLoaCustomerName = "loa"

// packetfabric_link_aggregation_group
const LinkAggGroupInterval = "fast"

// packetfbaric_backbone_virtual_circuit
const backboneVCspeed = "50Mbps"
const backboneVCepl = false
const backboneVCvlan1Value = 103
const backboneVCvlan2Value = 104
const backboneVClonghaulType = "dedicated"

// packetfabric_cloud_router
const DefaultCloudRouterCapacity = "5Gbps"
const CloudRouterCapacityChange = "10Gbps"
const CloudRouterRegionUS = "US"
const CloudRouterRegionUK = "UK"
const CloudRouterASN = 4556

// packetfabric_cloud_router_connection_aws
// packetfabric_cloud_router_connection_google
const CloudRouterConnSpeed = "100Mbps"

// packetfabric_cloud_router_connection_google
const CloudRouterConnGoogleRegion = "us-west1"
const CloudRouterConnGoogleNetwork = "default"

// packetfabric_cloud_router_connection_port
const CloudRouterConnPortSpeed = "1Gbps"
const CloudRouterConnPortVlan = 101

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
const HostedCloudSpeed = "100Mbps"
const HostedCloudVlan = 100

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
const AzurePeeringBandwidth = 100 // match const CloudRouterConnSpeed and HostedCloudSpeed

// packetfabric_cs_aws_dedicated_connection
// packetfabric_cs_google_dedicated_connection
// packetfabric_cs_azure_dedicated_connection
const DedicatedCloudSpeed = "10Gbps"
const DedicatedCloudServiceClass = "longhaul"
const DedicatedCloudAutoneg = false
const DedicatedCloudShouldCreateLag = false
const DedicatedCloudEncap = "qinq"      // Azure only
const DedicatedCloudPortCat = "primary" // Azure only

type PortDetails struct {
	PFClient              *packetfabric.PFClient
	DesiredSpeed          string
	DesiredPop            string
	DesiredZone           string
	DesiredMedia          string
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
	skipDesiredMarket     *string
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

// packetfabric_cloud_router_connection_port
type RHclCloudRouterConnectionPortResult struct {
	HclResultBase
	Desc        string
	CloudRouter RHclCloudRouterResult
	PortResult  RHclPortResult
	Speed       string
	Vlan        int
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

// packetfabric_cs_aws_dedicated_connection
type RHclCsAwsDedicatedConnectionResult struct {
	HclResultBase
	AwsRegion        string
	Description      string
	Pop              string
	Zone             string
	ShouldCreateLag  bool
	SubscriptionTerm int
	ServiceClass     string
	Autoneg          bool
	Speed            string
}

// packetfabric_cs_google_dedicated_connection
type RHclCsGoogleDedicatedConnectionResult struct {
	HclResultBase
	Desc             string
	Zone             string
	Pop              string
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
	SubscriptionTerm int
	ServiceClass     string
	Encapsulation    string
	PortCategory     string
	Speed            string
}

// ###### Data-sources

// data packetfabric_locations_cloud
type DHclDatasourceLocationsCloudResult struct {
	HclResultBase
}

// data packetfabric_locations_port_availability
type DHclLocationsPortAvailabilityResult struct {
	HclResultBase
}

// data packetfabric_locations
type DHclDatasourceLocationsResult struct {
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

// data packetfabric_activitylog
type DHclActivityLogResult struct {
	HclResultBase
}

// data packetfabric_locations_markets
type DHclLocationsMarketsResult struct {
	HclResultBase
}

// data packetfabric_port
type DHclPortResult struct {
	HclResultBase
}

// data packetfabric_billing
type DHclDatasourceBillingResult struct {
	HclResultBase
}

// data packetfabric_cs_aws_hosted_connection
type DHclCsAwsHostedConnectionResult struct {
	HclResultBase
}

// packetfabric_link_aggregation_group
type DHclDatasourceLinkAggregationGroupsResult struct {
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
	var pop, media, speed, market string
	var err error

	log.Println("Getting pop and zone with available port for desired speed: ", details.DesiredSpeed)
	var skipDesiredMarket *string
	if details.skipDesiredMarket != nil {
		skipDesiredMarket = details.skipDesiredMarket
	}
	pop, _, media, market, err = GetPopAndZoneWithAvailablePort(details.DesiredSpeed, skipDesiredMarket)
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}
	speed = details.DesiredSpeed
	log.Println("Pop, media, market, and speed set to: ", pop, media, market, speed)

	resourceName, hclName := GenerateUniqueResourceName(pfPort)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	log.Println("Generating HCL")
	hcl := fmt.Sprintf(
		RResourcePort,
		hclName,
		uniqueDesc,
		media,
		pop,
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
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

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
	portDetailsZ.skipDesiredMarket = &marketA // Send the market for the second port so it selects a different one to avoid to build a metro VC
	log.Println("Sending the market to the second port: ", *portDetailsZ.skipDesiredMarket)
	portTestResultZ := portDetailsZ.RHclPort(true)

	resourceName, hclName := GenerateUniqueResourceName(pfBackboneVirtualCircuit)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

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

// packetfabric_point_to_point
func RHclPointToPoint() RHclPointToPointResult {

	var speed = portSpeed
	pop1, zone1, media, market1, err := GetPopAndZoneWithAvailablePort(speed, nil)
	if err != nil {
		log.Println("Error getting pop and zone with available port: ", err)
		log.Panic(err)
	}
	log.Println("Pop1, media, and speed set to: ", pop1, zone1, media, market1, speed)

	pop2, zone2, media, market2, err2 := GetPopAndZoneWithAvailablePort(speed, &market1)
	if err2 != nil {
		log.Println("Error getting pop and zone with available port: ", err2)
		log.Panic(err)
	}
	log.Println("Pop2, media, and speed set to: ", pop2, zone2, media, market2, speed)

	uniqueDesc := GenerateUniqueName()
	resourceName, hclName := GenerateUniqueResourceName(pfPoinToPoint)
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

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
	log.Printf("Resource name: %s, description: %s\n", input.HclName, uniqueDesc)

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

	pop, _, _ := popDetails.FindAvailableCloudPopZone()

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnAws)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionAws,
		hclName,
		hclCloudRouterRes.ResourceName,
		os.Getenv("PF_AWS_ACCOUNT_ID"),
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
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
	}
}

// packetfabric_cloud_router_connection_google
func RHclCloudRouterConnectionGoogle() RHclCloudRouterConnectionGoogleResult {

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
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionGoogle,
		CloudRouterConnGoogleRegion,
		CloudRouterConnGoogleNetwork,
		CloudRouterConnGoogleRegion,
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
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

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

// packetfabric_cloud_router_connection_port
func RHclCloudRouterConnectionPort() RHclCloudRouterConnectionPortResult {

	portDetails := CreateBasePortDetails()

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnPort)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

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

// packetfabric_cloud_router_bgp_session
func RHclBgpSession() RHclBgpSessionResult {

	hclCloudConnRes := RHclCloudRouterConnectionAws()

	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterBgpSession)
	log.Printf("Resource name: %s\n", hclName)

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
	pop, _, _ := popDetails.FindAvailableCloudPopZone()

	portDetails := CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCsAwsHostedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	awsHostedConnectionHcl := fmt.Sprintf(
		RResourceCSAwsHostedConnection,
		hclName,
		portTestResult.ResourceName,
		os.Getenv("PF_AWS_ACCOUNT_ID"),
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		HostedCloudSpeed,
		HostedCloudVlan)

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
		Vlan:         HostedCloudVlan,
	}
}

// packetfabric_cs_google_hosted_connection
func RHclCsGoogleHostedConnection() RHclCsHostedCloudGoogleResult {

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
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	googleHostedConnectionHcl := fmt.Sprintf(
		RResourceCSGoogleHostedConnection,
		CloudRouterConnGoogleRegion,
		CloudRouterConnGoogleNetwork,
		CloudRouterConnGoogleRegion,
		hclName,
		portTestResult.ResourceName,
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		HostedCloudSpeed,
		HostedCloudVlan)

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
		Vlan:        HostedCloudVlan,
	}
}

// packetfabric_cs_azure_hosted_connection
func RHclCsAzureHostedConnection() RHclCsHostedCloudAzureResult {

	portDetails := CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCsAzureHostedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

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
		HostedCloudVlan)

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
		VlanPrivate: HostedCloudVlan,
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
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          DedicatedCloudSpeed,
		DesiredProvider:       "aws",
		DesiredConnectionType: "dedicated",
		IsCloudConnection:     true,
	}
	pop, zone, region := popDetails.FindAvailableCloudPopZone()

	hcl := fmt.Sprintf(
		RResourceCSAwsDedicatedConnection,
		hclName,
		region,
		uniqueDesc,
		pop,
		zone,
		DedicatedCloudShouldCreateLag,
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
		AwsRegion:        region,
		Description:      uniqueDesc,
		Pop:              pop,
		Zone:             zone,
		ShouldCreateLag:  DedicatedCloudShouldCreateLag,
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
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          DedicatedCloudSpeed,
		DesiredProvider:       "google",
		DesiredConnectionType: "dedicated",
		IsCloudConnection:     true,
	}
	pop, zone, _ := popDetails.FindAvailableCloudPopZone()

	hcl := fmt.Sprintf(RResourceCSGoogleDedicatedConnection,
		hclName,
		uniqueDesc,
		zone,
		pop,
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
		Zone:             zone,
		Pop:              pop,
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
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          DedicatedCloudSpeed,
		DesiredProvider:       "azure",
		DesiredConnectionType: "dedicated",
		IsCloudConnection:     true,
	}
	pop, _, _ := popDetails.FindAvailableCloudPopZone()

	hcl := fmt.Sprintf(
		RResourceCSAzureDedicatedConnection,
		hclName,
		uniqueDesc,
		pop,
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
		SubscriptionTerm: subscriptionTerm,
		ServiceClass:     DedicatedCloudServiceClass,
		Encapsulation:    DedicatedCloudEncap,
		PortCategory:     DedicatedCloudPortCat,
		Speed:            DedicatedCloudSpeed,
	}
}

// ###### Data-sources

// data.packetfabric_locations_cloud
func DHclDataSourceLocationsCloud(cloudProvider, cloudConnectionType string) DHclDatasourceLocationsCloudResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsCloud)
	log.Printf("Data-source name: %s\n", hclName)
	hcl := fmt.Sprintf(DDataSourceLocationsCloud, hclName, cloudProvider, cloudConnectionType)

	return DHclDatasourceLocationsCloudResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsCloud,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations_port_availability
func DHclDataSourceLocationsPortAvailability() DHclLocationsPortAvailabilityResult {

	pop, _, _, _, _ := GetPopAndZoneWithAvailablePort(portSpeed, nil)

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsPortAvailability)
	log.Printf("Data-source name: %s\n", hclName)

	hcl := fmt.Sprintf(DDataSourceLocationsPortAvailability, hclName, pop)

	return DHclLocationsPortAvailabilityResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsPortAvailability,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations
func DHclDataSourceLocations() DHclDatasourceLocationsResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocations)
	log.Printf("Data-source name: %s\n", hclName)

	hcl := fmt.Sprintf(DDatasourceLocations, hclName)

	return DHclDatasourceLocationsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocations,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations_pop_zones
func DHclDataSourceZones() DHclLocationsZonesResult {

	pop, _, _, _, _ := GetPopAndZoneWithAvailablePort(portSpeed, nil)

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsZones)
	log.Printf("Data-source name: %s\n", hclName)

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
func DHclDataSourceLocationsRegions() DHclLocationsRegionsResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsRegions)
	log.Printf("Data-source name: %s\n", hclName)

	hcl := fmt.Sprintf(DDataSourceLocationsRegions, hclName)

	return DHclLocationsRegionsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsRegions,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_activitylog
func DHclDataSourceActivityLog() DHclActivityLogResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataActivityLog)
	log.Printf("Data-source name: %s\n", hclName)

	hcl := fmt.Sprintf(DDatasourceActivityLog, hclName)

	return DHclActivityLogResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataActivityLog,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_locations_markets
func DHclDataSourceLocationsMarkets() DHclLocationsMarketsResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataLocationsMarkets)
	log.Printf("Data-source name: %s\n", hclName)

	hcl := fmt.Sprintf(DDataSourceLocationsMarkets, hclName)

	return DHclLocationsMarketsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataLocationsMarkets,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_ports
func DHclDataSourcePorts() DHclPortResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	portDetails := PortDetails{
		PFClient:     c,
		DesiredSpeed: portSpeed,
	}

	resourceName, hclName := GenerateUniqueResourceName(pfDataPort)
	log.Printf("Data-source name: %s\n", hclName)

	dataPortHcl := fmt.Sprintf(DDataSourcePorts, hclName)

	hcl := fmt.Sprintf("%s\n%s", portDetails.RHclPort(false).Hcl, dataPortHcl)

	return DHclPortResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataPort,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_billing
func DHclDatasourceBilling() DHclDatasourceBillingResult {

	hclCloudRouterRes := RHclCloudRouter(DefaultRHclCloudRouterInput())

	resourceName, hclName := GenerateUniqueResourceName(pfDataBilling)
	log.Printf("Data-source name: %s\n", hclName)

	billingHcl := fmt.Sprintf(DDatasourceBilling,
		hclName,
		hclCloudRouterRes.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, billingHcl)

	return DHclDatasourceBillingResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataBilling,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_cs_aws_hosted_connection
func DHclDatasourceHostedAwsConn() DHclCsAwsHostedConnectionResult {

	csAwsHostedConnectionResult := RHclCsAwsHostedConnection()

	resourceName, hclName := GenerateUniqueResourceName(pfDatasourceCsAwsHostedConn)
	log.Printf("Data-source name: %s\n", hclName)

	hostedAwsConnHcl := fmt.Sprintf(DDatasourceCsAwsHostedConn,
		hclName,
		csAwsHostedConnectionResult.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", csAwsHostedConnectionResult.Hcl, hostedAwsConnHcl)

	return DHclCsAwsHostedConnectionResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDatasourceCsAwsHostedConn,
			ResourceName: resourceName,
		},
	}
}

// data.packetfabric_link_aggregation_group
func DHclDatasourceLinkAggregationGroups() DHclDatasourceLinkAggregationGroupsResult {
	linkAggregationGroupResult := RHclLinkAggregationGroup()
	resourceName, hclName := GenerateUniqueResourceName(pfDatasourceLinkAggregationGroups)

	linkAggregationGroupsHcl := fmt.Sprintf(
		DDatasourceLinkAggregationGroups,
		hclName,
		linkAggregationGroupResult.ResourceName)

	hcl := fmt.Sprintf("%s\n%s", linkAggregationGroupResult.Hcl, linkAggregationGroupsHcl)

	return DHclDatasourceLinkAggregationGroupsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDatasourceLinkAggregationGroups,
			ResourceName: resourceName,
		},
	}
}
