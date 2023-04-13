package testutil

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
)

// ########################################
// ###### RESOURCES CONSTs
// ########################################

const pfPort = "packetfabric_port"
const pfCloudRouter = "packetfabric_cloud_router"
const pfCloudRouterConnAws = "packetfabric_cloud_router_connection_aws"
const pfCloudRouterBgpSession = "packetfabric_cloud_router_bgp_session"
const pfCsAwsHostedConn = "packetfabric_cs_aws_hosted_connection"
const pfCsOracleHostedMktConn = "packetfabric_cs_oracle_hosted_marketplace_connection"

// ########################################
// ###### HARDCODED VALUES
// ########################################

const portSubscriptionTerm = 1
const portSpeed = "1Gbps"

// packetfabric_cloud_router
const CrbsAddressFmly = "ivp4"
const CloudRouterCapacity = "1Gbps"
const CloudRouterRegionUS = "US"
const CloudRouterRegionUK = "UK"
const CloudRouterASN = 4556

// packetfabric_cs_aws_hosted_connection
// packetfabric_cloud_router_connection_aws
// packetfabric_cs_oracle_hosted_marketplace_connection
const CloudRouterConnAwsSpeed = "50Mbps"
const CloudRouterConnOracleSpeed = "1Gbps"
const CloudRouterConnAwsVlan = 100

// packetfabric_cloud_router_bg_session
const CloudRouterBgpSessionASN = 64534
const CloudRouterBgpSessionPrefix1 = "10.0.0.0/8"
const CloudRouterBgpSessionType1 = "in"
const CloudRouterBgpSessionPrefix2 = "192.168.0.0/24"
const CloudRouterBgpSessionType2 = "out"

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
	AwsAccountID string
	Desc         string
	Pop          string
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

type RHclCsOracleHostedMarketplaceConnectionResult struct {
	HclResultBase
	RoutingID   string
	Description string
	VcOCID      string
	Region      string
	Market      string
	Pop         string
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
func (details PortDetails) RHclPort() RHclPortResult {
	resourceReferece, resourceName := _generateResourceName(pfPort)
	var pop, media, speed string
	var err error
	var portEnabled bool
	if !details.IsCloudConnection {
		pop, _, media, err = GetPopAndZoneWithAvailablePort(details.DesiredSpeed)
		if err != nil {
			log.Panic(err)
		}
		speed = details.DesiredSpeed
	} else {
		pop = details.DesiredPop
		media = details.DesiredMedia
		speed = details.DesiredSpeed
	}

	// Port must be disabled if test is pointing to PF's dev env.
	if !strings.Contains(os.Getenv(PF_HOST_KEY), "api-beta.dev") {
		portEnabled = details.PortEnabled
	}
	uniqueDesc := _generateUniqueNameOrDesc(pfPort)
	hcl := fmt.Sprintf(
		RResourcePort,
		resourceName,
		uniqueDesc,
		media,
		pop,
		speed,
		portSubscriptionTerm,
		portEnabled)

	return RHclPortResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfPort,
			ResourceName: resourceName,
		},
		ResourceReference: resourceReferece,
		Description:       uniqueDesc,
		Media:             media,
		Pop:               pop,
		Speed:             speed,
		SubscriptionTerm:  portSubscriptionTerm,
		Enabled:           portEnabled,
	}
}

// packetfabric_cloud_router
func RHclCloudRouter() RHclCloudRouterResult {
	resourceName, hclName := _generateResourceName(pfCloudRouter)
	hcl := fmt.Sprintf(
		RResourcePacketfabricCloudRouter,
		hclName,
		_generateUniqueNameOrDesc(pfCloudRouter),
		os.Getenv(PF_ACCOUNT_ID_KEY),
		CloudRouterASN,
		CloudRouterCapacity,
		CloudRouterRegionUS,
		CloudRouterRegionUK,
		resourceName)

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
	portDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          CloudRouterConnAwsSpeed,
		DesiredProvider:       "aws",
		DesiredConnectionType: "hosted",
		IsCloudConnection:     true,
	}

	pop, _, _ := portDetails._findAvailableCloudPopZoneAndMedia()

	hclCloudRouterRes := RHclCloudRouter()
	resourceName, hclName := _generateResourceName(pfCloudRouterConnAws)

	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionAws,
		hclName,
		hclCloudRouterRes.ResourceName,
		os.Getenv(PF_CRC_AWS_ACCOUNT_ID_KEY),
		os.Getenv(PF_ACCOUNT_ID_KEY),
		_generateUniqueNameOrDesc(pfCloudRouterConnAws),
		pop,
		CloudRouterConnAwsSpeed)
	hcl := fmt.Sprintf("%s\n%s", hclCloudRouterRes.Hcl, crcHcl)
	return RHclCloudRouterConnectionAwsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCloudRouterConnAws,
			ResourceName: resourceName,
		},
		AwsAccountID: os.Getenv(PF_CRC_AWS_ACCOUNT_ID_KEY),
		Pop:          pop,
	}
}

// packetfabric_cloud_router_bgp_session
func RHclBgpSession() RHclBgpSessionResult {

	hclCloudRouterRes := RHclCloudRouter()
	hclCloudConnRes := RHclCloudRouterConnectionAws()

	resourceName, hclName := _generateResourceName(pfCloudRouterBgpSession)
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
func RHclAwsHostedConnection() RHclCloudRouterConnectionAwsResult {

	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}
	portDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          portSpeed,
		DesiredProvider:       "aws",
		DesiredConnectionType: "hosted",
		IsCloudConnection:     true,
	}

	pop, zone, media := portDetails._findAvailableCloudPopZoneAndMedia()
	if pop == "" {
		log.Fatalf("Resource: %s: %s", pfCsAwsHostedConn, "pop cannot be empty")
	}
	portDetails.DesiredPop = pop
	portDetails.DesiredZone = zone
	portDetails.DesiredMedia = media

	resourceName, hclName := _generateResourceName(pfCsAwsHostedConn)
	hclPortResult := portDetails.RHclPort()
	uniqueDesc := _generateUniqueNameOrDesc(pfCsAwsHostedConn)
	awsHostedConnectionHcl := fmt.Sprintf(
		RResourceCSAwsHostedConnection,
		hclName,
		uniqueDesc,
		os.Getenv(PF_CRC_AWS_ACCOUNT_ID_KEY),
		hclPortResult.ResourceReference,
		CloudRouterConnAwsSpeed,
		pop,
		CloudRouterConnAwsVlan)
	hcl := fmt.Sprintf("%s\n%s", hclPortResult.Hcl, awsHostedConnectionHcl)

	return RHclCloudRouterConnectionAwsResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsAwsHostedConn,
			ResourceName: resourceName,
		},
		AwsAccountID: os.Getenv(PF_CRC_AWS_ACCOUNT_ID_KEY),
		Desc:         uniqueDesc,
		Pop:          pop,
	}
}

func RHclCsOracleHostedMarketplaceConnection() RHclCsOracleHostedMarketplaceConnectionResult {
	c, err := _createPFClient()
	if err != nil {
		log.Panic(err)
	}

	portDetails := PortDetails{
		PFClient:              c,
		DesiredSpeed:          CloudRouterConnOracleSpeed,
		DesiredProvider:       "oracle",
		DesiredConnectionType: "hosted",
		IsCloudConnection:     true,
	}

	pop, _, _ := portDetails._findAvailableCloudPopZoneAndMedia()
	if pop == "" {
		log.Fatalf("Resource: %s: %s", pfCsAwsHostedConn, "pop cannot be empty")
	}

	uniqueDesc := _generateUniqueNameOrDesc(pfCsOracleHostedMktConn)
	resourceName, hclName := _generateResourceName(pfCsOracleHostedMktConn)

	hcl := fmt.Sprintf(
		RResourceCSOracleHostedMarketplaceConnection,
		hclName,
		uniqueDesc,
		os.Getenv(PF_CS_ORACLE_MKT_CONN_OCID_KEY),
		os.Getenv(PF_CS_ORACLE_MKT_CONN_REGION_KEY),
		os.Getenv(PF_CS_ORACLE_MKT_CONN_ROUTING_ID_KEY),
		os.Getenv(PF_CS_ORACLE_MKT_CONN_MARKET_KEY),
		pop,
	)
	return RHclCsOracleHostedMarketplaceConnectionResult{
		HclResultBase: HclResultBase{
			Hcl:          hcl,
			Resource:     pfCsOracleHostedMktConn,
			ResourceName: resourceName,
		},
		Description: uniqueDesc,
		RoutingID:   os.Getenv(PF_CS_ORACLE_MKT_CONN_ROUTING_ID_KEY),
		VcOCID:      os.Getenv(PF_CS_ORACLE_MKT_CONN_OCID_KEY),
		Region:      os.Getenv(PF_CS_ORACLE_MKT_CONN_REGION_KEY),
		Market:      os.Getenv(PF_CS_ORACLE_MKT_CONN_MARKET_KEY),
		Pop:         pop,
	}
}

func (details PortDetails) _findAvailableCloudPopZoneAndMedia() (pop, zone, media string) {
	popsAvailable, _ := details.FetchCloudPops()
	popsToSkip := make([]string, 0)
	for _, popAvailable := range popsAvailable {
		if len(popsToSkip) == len(popsAvailable) {
			log.Fatal(errors.New("there's no port available on any pop"))
		}
		if _contains(popsToSkip, pop) {
			continue
		}
		if zoneAvailable, mediaAvailable, availabilityErr := details.GetAvailableCloudPort(popAvailable); availabilityErr != nil {
			popsToSkip = append(popsToSkip, popAvailable)
			continue
		} else {
			pop = popAvailable
			media = mediaAvailable
			zone = zoneAvailable
			return
		}
	}
	return
}

func _generateResourceName(resource string) (resourceName, hclName string) {
	hclName = GenerateUniqueResourceName()
	resourceName = fmt.Sprintf("%s.%s", resource, hclName)
	return
}

func _generateUniqueNameOrDesc(targetResource string) (unique string) {
	unique = fmt.Sprintf("pf_testacc_%s_%s", targetResource, strings.ReplaceAll(uuid.NewString(), "-", "_"))
	return
}

func (details PortDetails) FetchCloudPops() (popsAvailable []string, err error) {
	if details.DesiredProvider == "" {
		err = errors.New("please provide a valid cloud provider to fetch pop")
	}
	if details.PFClient == nil {
		err = errors.New("please create PFClient to fetch cloud pop")
		return
	}
	if cloudLocations, locErr := details.PFClient.GetCloudLocations(
		details.DesiredProvider,
		details.DesiredConnectionType,
		details.IsNatCapable,
		details.HasCloudRouter,
		details.AnyType,
		details.DesiredPop,
		details.DesiredCity,
		details.DesiredState,
		details.DesiredMarket,
		details.DesiredRegion); locErr != nil {
		err = locErr
		return
	} else {
		for _, loc := range cloudLocations {
			popsAvailable = append(popsAvailable, loc.Pop)
		}
	}
	return
}

func _contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (details PortDetails) GetAvailableCloudPort(desiredPop string) (zone, media string, err error) {
	if desiredPop == "" {
		err = errors.New("please provide a valid pop")
		return
	}
	if details.PFClient == nil {
		err = errors.New("please create a PFClient to fetch available cloud port")
		return
	}

	var ports []packetfabric.PortAvailability
	if ports, err = details.PFClient.GetLocationPortAvailability(desiredPop); err != nil {
		return
	}
	for _, port := range ports {
		if port.Count > 0 && port.Speed == details.DesiredSpeed {
			zone = port.Zone
			media = port.Media
			return
		}
	}
	err = errors.New("there's no port available for the requested speed")
	return
}
