package testutil

import (
	"fmt"
	"log"
	"os"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
)

// ########################################
// ###### RESOURCES CONSTs
// ########################################

const pfPort = "packetfabric_port"
const pfBackboneVirtualCircuit = "packetfabric_backbone_virtual_circuit"
const pfPoinToPoint = "packetfabric_point_to_point"
const pfCloudRouter = "packetfabric_cloud_router"
const pfCloudRouterConnAws = "packetfabric_cloud_router_connection_aws"
const pfCloudRouterConnPort = "packetfabric_cloud_router_connection_port"
const pfCloudRouterBgpSession = "packetfabric_cloud_router_bgp_session"
const pfCsAwsHostedConn = "packetfabric_cs_aws_hosted_connection"
const pfDataSourceLocationsCloud = "data.packetfabric_locations_cloud"
const pfDataLocationsPortAvailability = "data.packetfabric_locations_port_availability"
const pfDataLocations = "data.packetfabric_locations"
const pfDataZones = "data.packetfabric_locations_pop_zones"
const pfDataLocationsRegions = "data.packetfabric_locations_regions"
const pfDataActivityLog = "data.packetfabric_activitylog"
const pfDataLocationsMarkets = "data.packetfabric_locations_markets"

// ########################################
// ###### HARDCODED VALUES
// ########################################

// common
const subscriptionTerm = 1
const updatedSubscriptionTerm = 12

// packetfabric_port
// packetfabric_point_to_point
const portSpeed = "1Gbps"

var listPortsLab = []string{"LAB1", "LAB2", "LAB4", "LAB6", "LAB8"}

// packetfbaric_backbone_virtual_circuit
const backboneVCspeed = "50Mbps"
const backboneVCepl = false
const backboneVCvlan1Value = 103
const backboneVCvlan2Value = 104
const backboneVClonghaulType = "dedicated"

// packetfabric_cloud_router
const CrbsAddressFmly = "ivp4"
const CloudRouterCapacity = "1Gbps"
const CloudRouterRegionUS = "US"
const CloudRouterRegionUK = "UK"
const CloudRouterASN = 4556

// packetfabric_cloud_router_connection_aws
const CloudRouterConnAwsSpeed = "50Mbps"

// packetfabric_cloud_router_connection_port
const CloudRouterConnPortSpeed = "1Gbps"
const CloudRouterConnPortVlan = 101

// packetfabric_cloud_router_bg_session
const CloudRouterBgpSessionASN = 64534
const CloudRouterBgpSessionPrefix1 = "10.0.0.0/8"
const CloudRouterBgpSessionType1 = "in"
const CloudRouterBgpSessionPrefix2 = "192.168.0.0/24"
const CloudRouterBgpSessionType2 = "out"

// packetfabric_cs_aws_hosted_connection
const HostedCloudSpeed = "100Mbps"
const HostedCloudVlan = 100

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

type HclResultBase struct {
	Hcl          string
	Resource     string
	ResourceName string
}

// packetfabric_port
type RHclPortResult struct {
	HclResultBase
	ResourceReference string
	Description       string
	Media             string
	Pop               string
	Speed             string
	SubscriptionTerm  int
	Enabled           bool
	Market            string
}

// packetfabric_cloud_router
type RHclCloudRouterResult struct {
	HclResultBase
	Asn      int
	Capacity string
	Regions  []string
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

// packetfabric_cloud_router_connection_port
type RHclCloudRouterConnectionPortResult struct {
	HclResultBase
	Desc              string
	CloudRouterResult RHclCloudRouterResult
	PortResult        RHclPortResult
	Speed             string
	Vlan              int
}

// packetfabric_cloud_router_bgp_session
type RHclBgpSessionResult struct {
	HclResultBase
	CloudRouter        RHclCloudRouterResult
	CloudRouterConnAws RHclCloudRouterConnectionAwsResult
	AddressFamily      string
	Asn                int
	Prefix1            string
	Type1              string
	Prefix2            string
	Type2              string
}

// packetfabric_cs_aws_hosted_connection
type RHclHostedCloudAwsResult struct {
	HclResultBase
	PortResult   RHclPortResult
	AwsAccountID string
	AccountUuid  string
	Desc         string
	Pop          string
	Speed        string
	Vlan         int
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

// Patterns:
// Resource schema for required fields only
// - func RHcl...
// Resouce schema for required + optional fields
// - func OHcl...

// ########################################
// ###### HCLs FOR REQUIRED FIELDS
// ########################################
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

	resourceReference, hclName := GenerateUniqueResourceName(pfPort)
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
		ResourceReference: resourceReference,
		Description:       uniqueDesc,
		Media:             media,
		Pop:               pop,
		Speed:             speed,
		SubscriptionTerm:  subscriptionTerm,
		Enabled:           portEnabled,
		Market:            market,
	}
}

// packetfabric_cloud_router
func RHclCloudRouter() RHclCloudRouterResult {
	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouter)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	hcl := fmt.Sprintf(
		RResourcePacketfabricCloudRouter,
		hclName,
		uniqueDesc,
		os.Getenv("PF_ACCOUNT_ID"),
		CloudRouterASN,
		CloudRouterCapacity,
		CloudRouterRegionUS,
		CloudRouterRegionUK)

	return RHclCloudRouterResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouter,
			ResourceName: resourceName,
		},
		Asn:      CloudRouterASN,
		Capacity: CloudRouterCapacity,
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
		DesiredSpeed:          CloudRouterConnAwsSpeed,
		DesiredProvider:       "aws",
		DesiredConnectionType: "hosted",
		IsCloudConnection:     true,
	}

	pop, _ := popDetails.FindAvailableCloudPopZone()

	hclCloudRouterRes := RHclCloudRouter()
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
		CloudRouterConnAwsSpeed)
	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, crcHcl)
	return RHclCloudRouterConnectionAwsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouterConnAws,
			ResourceName: resourceName,
		},
		AwsAccountID: os.Getenv("PF_AWS_ACCOUNT_ID"),
		AccountUuid:  os.Getenv("PF_ACCOUNT_ID"),
		Speed:        CloudRouterConnAwsSpeed,
		Pop:          pop,
	}
}

func RHclCloudRouterConnectionPort() RHclCloudRouterConnectionPortResult {

	portDetails := CreateBasePortDetails()
	cloudRouterResult := RHclCloudRouter()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterConnPort)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	crConnPortHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionPort,
		hclName,
		uniqueDesc,
		cloudRouterResult.ResourceName,
		portTestResult.ResourceReference,
		CloudRouterConnPortSpeed,
		CloudRouterConnPortVlan,
	)

	hcl := fmt.Sprintf("%s\n%s\n%s", portTestResult.Hcl, cloudRouterResult.Hcl, crConnPortHcl)

	return RHclCloudRouterConnectionPortResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouterConnPort,
			ResourceName: resourceName,
		},
		CloudRouterResult: cloudRouterResult,
		PortResult:        portTestResult,
		Desc:              uniqueDesc,
		Speed:             CloudRouterConnPortSpeed,
		Vlan:              CloudRouterConnPortVlan,
	}
}

// packetfabric_cloud_router_bgp_session
func RHclBgpSession() RHclBgpSessionResult {

	hclCloudRouterRes := RHclCloudRouter()
	hclCloudConnRes := RHclCloudRouterConnectionAws()

	resourceName, hclName := GenerateUniqueResourceName(pfCloudRouterBgpSession)
	log.Printf("Resource name: %s\n", hclName)

	bgpSessionHcl := fmt.Sprintf(
		RResourceCloudRouterBgpSession,
		hclName,
		hclCloudRouterRes.ResourceName,
		hclCloudConnRes.HclResultBase.ResourceName,
		CrbsAddressFmly,
		CloudRouterBgpSessionASN,
		CloudRouterBgpSessionPrefix1,
		CloudRouterBgpSessionType1,
		CloudRouterBgpSessionPrefix2,
		CloudRouterBgpSessionType2)
	hcl := fmt.Sprintf("%s\n%s\n%s", hclCloudRouterRes.Hcl, hclCloudConnRes.Hcl, bgpSessionHcl)
	return RHclBgpSessionResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouterBgpSession,
			ResourceName: resourceName,
		},
		CloudRouter:        hclCloudRouterRes,
		CloudRouterConnAws: hclCloudConnRes,
		AddressFamily:      CrbsAddressFmly,
		Asn:                CloudRouterBgpSessionASN,
		Prefix1:            CloudRouterBgpSessionPrefix1,
		Type1:              CloudRouterBgpSessionType1,
		Prefix2:            CloudRouterBgpSessionPrefix2,
		Type2:              CloudRouterBgpSessionType2,
	}
}

// packetfabric_cs_aws_hosted_connection
func RHclAwsHostedConnection() RHclHostedCloudAwsResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	popDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          HostedCloudSpeed,
		DesiredProvider:       "aws",
		DesiredConnectionType: "hosted",
	}
	pop, _ := popDetails.FindAvailableCloudPopZone()

	portDetails := CreateBasePortDetails()
	portTestResult := portDetails.RHclPort(false)

	resourceName, hclName := GenerateUniqueResourceName(pfCsAwsHostedConn)
	uniqueDesc := GenerateUniqueName()
	log.Printf("Resource name: %s, description: %s\n", hclName, uniqueDesc)

	awsHostedConnectionHcl := fmt.Sprintf(
		RResourceCSAwsHostedConnection,
		hclName,
		portTestResult.ResourceReference,
		os.Getenv("PF_AWS_ACCOUNT_ID"),
		os.Getenv("PF_ACCOUNT_ID"),
		uniqueDesc,
		pop,
		HostedCloudSpeed,
		HostedCloudVlan)

	hcl := fmt.Sprintf("%s\n%s", portTestResult.Hcl, awsHostedConnectionHcl)

	return RHclHostedCloudAwsResult{
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
		portTestResultA.ResourceReference,
		backboneVCvlan1Value,
		portTestResultZ.ResourceReference,
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
			PortCircuitID: portTestResultA.ResourceReference,
		},
		InterfaceBackboneZ: InterfaceBackbone{
			Vlan:          backboneVCvlan2Value,
			PortCircuitID: portTestResultZ.ResourceReference,
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

func DHclDataSourceLocationsCloud(cloudProvider, cloudConnectionType string) DHclDatasourceLocationsCloudResult {

	resourceName, hclName := GenerateUniqueResourceName(pfDataSourceLocationsCloud)
	log.Printf("Data-source name: %s\n", hclName)
	hcl := fmt.Sprintf(DDataSourceLocationsCloud, hclName, cloudProvider, cloudConnectionType)

	return DHclDatasourceLocationsCloudResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataSourceLocationsCloud,
			ResourceName: resourceName,
		},
	}
}

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

func DHclDataSourceZones() DHclLocationsZonesResult {

	pop, _, _, _, _ := GetPopAndZoneWithAvailablePort(portSpeed, nil)

	resourceName, hclName := GenerateUniqueResourceName(pfDataZones)
	log.Printf("Data-source name: %s\n", hclName)

	hcl := fmt.Sprintf(DDatasourceLocationsPopZones, hclName, pop)

	return DHclLocationsZonesResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfDataZones,
			ResourceName: resourceName,
		},
	}
}

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
