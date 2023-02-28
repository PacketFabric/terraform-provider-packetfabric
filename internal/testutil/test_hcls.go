package testutil

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

// ########################################
// ###### HCLs FOR REQUIRED FIELDS
// ########################################

// packetfabric_port
func HclPort() (hcl, resourceName, resourceReferece string) {
	resourceReferece, resourceName = _generateResourceName(pfPort)
	pop, _, media, err := GetPopAndZoneWithAvailablePort(os.Getenv(PF_CRC_SPEED_KEY))
	if err != nil {
		log.Panic(err)
	}
	var portEnabled bool
	if os.Getenv(PF_PORT_ENABLED_KEY) != "" {
		portEnabled, err = strconv.ParseBool(os.Getenv(PF_PORT_ENABLED_KEY))
		if err != nil {
			log.Fatal("port enabled must be either true or false")
		}
	}
	hcl = fmt.Sprintf(
		RResourcePort,
		resourceName,
		_generateUniqueNameOrDesc(pfPort),
		media,
		pop,
		os.Getenv(PF_PORT_SPEED_KEY),
		os.Getenv(PF_PORT_SUBTERM_KEY),
		portEnabled)
	return
}

// packetfabric_cloud_router
func HclCloudRouter() (hcl string, resourceName string) {
	resourceName, hclName := _generateResourceName(pfCloudRouter)
	hcl = fmt.Sprintf(
		RResourcePacketfabricCloudRouter,
		hclName,
		_generateUniqueNameOrDesc(pfCloudRouter),
		os.Getenv(PF_ACCOUNT_ID_KEY),
		os.Getenv(PF_CR_CAPACITY_KEY),
		os.Getenv(PF_CR_MEMBER_KEY))
	return
}

// packetfabric_cloud_router_connection_aws
func HclCloudRouterConnectionAws() (hcl string, crResourceName, resourceName string) {

	hclCloudRouter, crResourceName := HclCloudRouter()
	resourceName, hclName := _generateResourceName(pfCloudRouterConnAws)
	pop, _, _, err := GetPopAndZoneWithAvailablePort(os.Getenv(PF_CRC_SPEED_KEY))
	if err != nil {
		log.Fatal(err)
	}
	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionAws,
		hclName,
		crResourceName,
		os.Getenv(PF_CRC_AWS_ACCOUNT_ID_KEY),
		os.Getenv(PF_ACCOUNT_ID_KEY),
		_generateUniqueNameOrDesc(pfCloudRouterConnAws),
		pop,
		os.Getenv(PF_CRC_SPEED_KEY))

	hcl = fmt.Sprintf("%s\n%s", hclCloudRouter, crcHcl)
	return
}

// packetfabric_cloud_router_bgp_session
func HclBgpSession() (hcl string, resourceName string) {

	// hclCR, crResourceName := HclCloudRouter()
	hclRCConn, crResourceName, crConnResourceName := HclCloudRouterConnectionAws()

	resourceName, hclName := _generateResourceName(pfCloudRouterBgpSession)
	bgpSessionHcl := fmt.Sprintf(
		RResourceCloudRouterBgpSession,
		hclName,
		crResourceName,
		crConnResourceName,
		os.Getenv(PF_CRBS_ADDRESS_FMLY_KEY),
		os.Getenv(PF_CRBS_MHTTL_KEY),
		os.Getenv(PF_CRBS_REMOTE_ASN_KEY),
		os.Getenv(PF_CRBS_PRFX1_KEY),
		os.Getenv(PF_CRBS_TYPE1_KEY),
		os.Getenv(PF_CRBS_PRFX2_KEY),
		os.Getenv(PF_CRBS_TYPE2_KEY))
	hcl = fmt.Sprintf("%s\n%s", hclRCConn, bgpSessionHcl)
	return
}

// packetfabric_cs_aws_hosted_connection
func HclAwsHostedConnection() (hcl, resourceName string) {
	pop, _, _, err := GetPopAndZoneWithAvailablePort(os.Getenv(PF_CRC_SPEED_KEY))
	if err != nil {
		log.Fatal(err)
	}
	resourceName, hclName := _generateResourceName(pfCsAwsHostedConn)
	hclPort, _, portResourceReference := HclPort()
	awsHostedConnectionHcl := fmt.Sprintf(
		RResourceCSAwsHostedConnection,
		hclName,
		_generateUniqueNameOrDesc(pfCsAwsHostedConn),
		os.Getenv(PF_CRC_AWS_ACCOUNT_ID_KEY),
		portResourceReference,
		os.Getenv(PF_CS_SPEED2_KEY),
		pop,
		os.Getenv(PF_CS_VLAN2_KEY))
	hcl = fmt.Sprintf("%s\n%s", hclPort, awsHostedConnectionHcl)
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
