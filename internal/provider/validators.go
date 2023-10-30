package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"net"
	"regexp"
)

///

func validateIpVersion() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfAddressFamily4v, PfAddressFamily6v}, true)
}

func validateSubscriptionTerm() schema.SchemaValidateFunc {
	return validation.IntInSlice([]int{1, 12, 24, 36})
}

func validatePoNumber() schema.SchemaValidateFunc {
	return validation.StringLenBetween(1, 32)
}

func validateProvider() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		PfAws, PfAzure, PfPacket, PfGoogle, PfIbm,
		PfOracle, PfSalesforce, PfWebex,
	}, true)
}

func validateLongHaul() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfDedicated, PfUsage, PfHourly}, true)
}

func validateVifType() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfDirectConnect, PfPrivate, PfTransit}, false)
}

func validateTTl() schema.SchemaValidateFunc {
	return validation.IntBetween(1, 4)
}

func validateAsPrepend() schema.SchemaValidateFunc {
	return validation.IntBetween(1, 5)
}

func validateVlan() schema.SchemaValidateFunc {
	return validation.IntBetween(4, 4094)
}

func validatePrefixes(prefixesList []interface{}) error {
	inCount, outCount := 0, 0
	for _, prefix := range prefixesList {
		prefixMap := prefix.(map[string]interface{})
		prefixType := prefixMap[PfType].(string)

		if prefixType == PfIn {
			inCount++
		} else if prefixType == PfOut {
			outCount++
		}
	}
	if inCount == 0 || outCount == 0 {
		return fmt.Errorf(MessageMissingMinimumPrefix)
	}
	return nil
}

func validateIPAddressWithPrefix(val interface{}, key string) (warns []string, errs []error) {
	value := val.(string)
	_, _, err := net.ParseCIDR(value)
	if err != nil {
		errs = append(errs, fmt.Errorf("%q is not a valid IP address with prefix: %s", key, value))
	}
	return
}

func validateBfdInterval() schema.SchemaValidateFunc {
	return validation.IntBetween(3, 30000)
}

func validateBfdMultiplier() schema.SchemaValidateFunc {
	return validation.IntBetween(2, 16)
}

func validateCapacity() schema.SchemaValidateFunc {
	return validation.StringInSlice(speedFlexOptions(), true)
}

func validateCategories() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		PfCloudComputing, PfContentDeliveryNetwork,
		PfEdgeComputing, PfSdWan, PfDataStorage,
		PfDeveloperPlatform, PfInternetServiceProvider,
		PfSecurity, PfVideoConferencing, PfVoiceAndMessaging,
		PfWebHosting, PfInternetOfThings, PfPrivateConnectivity,
		PfBareMetalHosting,
	}, true)
}

func validateCloudProvider() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfAws, PfGoogle, PfOracle, PfAzure}, true)
}

func validateIO() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfOut, PfIn}, true)
}

func validateDirection() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfOutput, PfInput}, true)
}

func validateEncapsulation() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfDot1q, PfQinq}, true)
}

func validateEvents() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		PfAuth, PfDocument, PfLagInterface, PfLogicalInterface,
		PfPhysicalInterface, PfOutboundCrossConnect,
		PfPointToPoint, PfRateLimit, PfUser, PfVirtualCircuit,
		PfErrors, PfEtherstats, PfMetrics, PfOptical,
	}, false)
}

func validateFirstName() schema.SchemaValidateFunc {
	return validation.StringLenBetween(1, 255)
}

func validateGoogleEdgeAvailabilityDomain() schema.SchemaValidateFunc {
	return validation.IntInSlice([]int{1, 2})
}

func validateGoogleKeepaliveInterval() schema.SchemaValidateFunc {
	return validation.IntBetween(20, 60)
}

func validateGoogleRegion() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		PfAsiaEast1, PfAsiaEast2, PfAsiaNortheast1, PfAsiaNortheast2,
		PfAsiaNortheast3, PfAsiaSouth1, PfAsiaSoutheast1,
		PfAsiaSoutheast2, PfAustraliaSoutheast1, PfEuropeNorth1,
		PfEuropeWest1, PfEuropeWest2, PfEuropeWest3, PfEuropeWest4,
		PfEuropeWest6, PfNorthamericaNortheast1, PfSouthamericaEast1,
		PfUsCentral1, PfUsEast1, PfUsEast4, PfUsWest1, PfUsWest2,
		PfUsWest3, PfUsWest4,
	}, false)
}

func validateGroup() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfAdmin, PfRegular, PfReadOnly, PfSupport, PfSales}, false)
}

func validateIbmAccountId() schema.SchemaValidateFunc {
	return validation.StringLenBetween(1, 32)
}

func validateIkeVersion() schema.SchemaValidateFunc {
	return validation.IntInSlice([]int{1, 2})
}

func validateInterval() schema.SchemaValidateFunc {
	return validation.StringInSlice(intervalOptions(), false)
}

func validateLastName() schema.SchemaValidateFunc {
	return validation.StringLenBetween(1, 255)
}

func validateLogin() schema.SchemaValidateFunc {
	return validation.StringLenBetween(1, 255)
}

func validateMatchType() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfMatchTypeDefault, PfOrlonger}, true)
}

func validateMedia() schema.SchemaValidateFunc {
	return validation.StringInSlice(pointToPointMediaOptions(), true)
}

func validateMtu() schema.SchemaValidateFunc {
	return validation.IntInSlice([]int{1500, 9001})
}

func validateMtu2() schema.SchemaValidateFunc {
	return validation.IntInSlice([]int{1500, 1440})
}

func validateNatType() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfNatTypeOverload, PfNatTypeInline}, true)
}

func validatePassword() schema.SchemaValidateFunc {
	return validation.StringLenBetween(8, 64)
}

func validatePhase1AuthenticationAlgo() schema.SchemaValidateFunc {
	return validation.StringInSlice(ipSecPhase1AuthenticationAlgoOptions(), false)
}

func validatePhase1EncryptionAlgo() schema.SchemaValidateFunc {
	return validation.StringInSlice(ipSecPhase1EncryptionAlgoOptions(), false)
}

func validatePhase1Group() schema.SchemaValidateFunc {
	return validation.StringInSlice(ipSecPhasesGroupOptions(), false)
}

func validatePhase1Lifetime() schema.SchemaValidateFunc {
	return validation.IntBetween(180, 86400)
}

func validatePhase2AuthenticationAlgo() schema.SchemaValidateFunc {
	return validation.StringInSlice(ipSecPhase2AuthenticationAlgoOptions(), false)
}

func validatePhase2EncryptionAlgo() schema.SchemaValidateFunc {
	return validation.StringInSlice(ipSecPhase2EncryptionAlgoOptions(), false)
}

func validatePhase2Lifetime() schema.SchemaValidateFunc {
	return validation.IntBetween(180, 86400)
}

func validatePhase2PfsGroup() schema.SchemaValidateFunc {
	return validation.StringInSlice(ipSecPhasesGroupOptions(), false)
}

func validatePhone() schema.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile(`^[0-9 ()+.-]+(\s?(x|ex|ext|ete|extn)?(\.|\.\s|\s)?[\d]{1,9})?$`),
		MessagePhonePatternMismatch,
	)
}

func validatePortCategory() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfPrimary, PfSecondary}, true)
}

func validateCloudConnectionType() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfHosted, PfDedicated}, true)
}

func validateRouterType() schema.SchemaValidateFunc {
	return validation.StringInSlice(getValidRouterTypes(), false)
}

func validateServiceClass() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfLonghaul, PfMetro}, false)
}

func validateServiceClassTrue() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfLonghaul, PfMetro}, true)
}

func validateServiceType() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfServiceTypeDefault, PfQuickConnectService}, true)
}

func validateSpeed() schema.SchemaValidateFunc {
	return validation.StringInSlice(speedOptions(), true)
}

func validateSpeed10or100() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{Pf10Gbps, Pf100Gbps}, true)
}

func validateSpeed1or10or40or100() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{Pf1Gbps, Pf10Gbps, Pf40Gbps, Pf100Gbps}, true)
}

func validateIpSecSpeed() schema.SchemaValidateFunc {
	return validation.StringInSlice(ipSecSpeedOptions(), true)
}

func validateDedicatedSpeed() schema.SchemaValidateFunc {
	return validation.StringInSlice(speedGoogleDeicatedOptions(), false)
}

func validateType1() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfBackbone, PfIx, PfCloud}, true)
}

func validateType2() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfCustomer, PfPort, PfVc}, false)
}

func validateType3() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfLoa, PfMsa}, false)
}

func validateType4() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{PfSent, PfReceived}, true)
}

///

func speedOptions() []string {
	return []string{
		Pf50Mbps, Pf100Mbps, Pf200Mbps, Pf300Mbps, Pf400Mbps,
		Pf500Mbps, Pf1Gbps, Pf2Gbps, Pf5Gbps, Pf10Gbps, Pf20Gbps,
		Pf30Gbps, Pf40Gbps, Pf50Gbps, Pf60Gbps, Pf80Gbps, Pf100Gbps,
	}
}

func ipSecSpeedOptions() []string {
	return []string{
		Pf50Mbps, Pf100Mbps, Pf200Mbps, Pf300Mbps,
		Pf400Mbps, Pf500Mbps, Pf1Gbps, Pf2Gbps,
		Pf5Gbps, Pf10Gbps}
}

func ipSecPhasesGroupOptions() []string {
	return []string{
		PfGroup1, PfGroup14, PfGroup15, PfGroup16,
		PfGroup19, PfGroup2, PfGroup20, PfGroup24,
		PfGroup5}
}

func ipSecPhase1EncryptionAlgoOptions() []string {
	return []string{Pf3desCbc, PfAes128Cbc, PfAes192Cbc, PfAes256Cbc, PfDesCbc}
}

func ipSecPhase1AuthenticationAlgoOptions() []string {
	return []string{PfMd5, PfSha256, PfSha384, PfSha1}
}

func ipSecPhase2EncryptionAlgoOptions() []string {
	return []string{Pf3desCbc, PfAes128Cbc, PfAes128Gcm, PfAes192Cbc, PfAes256Cbc, PfAes192Gcm, PfAes256Cbc, PfAes256Gcm, PfDesCbc}
}

func ipSecPhase2AuthenticationAlgoOptions() []string {
	return []string{PfHmacMd596, PfHmacSha256128, PfHmacSha196}
}

func pointToPointMediaOptions() []string {
	return []string{PfLx, PfEx, PfZx, PfLr, PfEr, PfErDwdm, PfZr, PfZeDwdm, PfLr4, PfEr4, PfCwdm4, PfLr4, PfEr4Lite}
}

func intervalOptions() []string {
	return []string{PfFast, PfSlow}
}
