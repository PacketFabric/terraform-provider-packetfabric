package testutil

import (
	"fmt"
	"os"
)

// ########################################
// ###### RESOURCES CONSTs
// ########################################

const pfCloudRouter = "packetfabric_cloud_router"
const pfCloudRouterConnAws = "packetfabric_cloud_router_connection_aws"
const pfCloudRouterBgpSession = "packetfabric_cloud_router_bgp_session"

// ########################################
// ###### HCLs FOR REQUIRED FIELDS
// ########################################

// packetfabric_cloud_router
func HclCloudRouter() (hcl string, resourceName string) {
	resourceName, hclName := _generateResourceName(pfCloudRouter)
	hcl = fmt.Sprintf(
		RResourcePacketfabricCloudRouter,
		hclName,
		os.Getenv(PF_CR_DESCR_KEY),
		os.Getenv(PF_ACCOUNT_ID_KEY),
		os.Getenv(PF_CR_CAPACITY_KEY),
		os.Getenv(PF_CR_MEMBER_KEY))
	return
}

// packetfabric_cloud_router_connection_aws
func HclCloudRouterConnectionAws() (hcl string, crResourceName, resourceName string) {

	hclCloudRouter, crResourceName := HclCloudRouter()
	resourceName, hclName := _generateResourceName(pfCloudRouterConnAws)
	pop, _, _ := GetPopAndZoneWithAvailablePort(os.Getenv(PF_CRC_SPEED_KEY), os.Getenv(PF_PORT_MEDIA_KEY))
	crcHcl := fmt.Sprintf(
		RResourceCloudRouterConnectionAws,
		hclName,
		crResourceName,
		os.Getenv(PF_CRC_AWS_ACCOUNT_ID_KEY),
		os.Getenv(PF_ACCOUNT_ID_KEY),
		os.Getenv(PF_CR_DESCR_KEY),
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

func _generateResourceName(resource string) (resourceName, hclName string) {
	hclName = GenerateUniqueResourceName()
	resourceName = fmt.Sprintf("%s.%s", resource, hclName)
	return
}
