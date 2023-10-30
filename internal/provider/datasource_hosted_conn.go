package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudServicesConnInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudServicesRead,
		Schema: map[string]*schema.Schema{
			PfCloudCircuitId:            schemaStringRequired(PfCloudCircuitIdDescription2),
			PfCustomerUuid:              schemaStringComputed(PfCustomerUuidDescription3),
			PfUserUuid:                  schemaStringComputed(PfUserUuidDescription2),
			PfState:                     schemaStringComputed(PfStateDescription6),
			PfServiceProvider:           schemaStringComputed(PfServiceProviderDescription3),
			PfServiceClass:              schemaStringComputed(PfServiceClassDescription4),
			PfPortType:                  schemaStringComputed(PfPortTypeDescription3),
			PfSpeed:                     schemaStringComputed(PfSpeedDescription),
			PfDescription:               schemaStringComputed(PfConnectionDescription5),
			PfCloudProviderPop:          schemaStringComputed(PfCloudProviderPopDescription),
			PfCloudProviderRegion:       schemaStringComputed(PfCloudProviderRegionDescription),
			PfCloudProviderConnectionId: schemaStringComputed(PfCloudProviderConnectionIdDescription),
			PfSettings: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfVlanIdPf:      schemaIntComputedPlain(),
						PfVlanIdCust:    schemaIntComputedPlain(),
						PfSvlanIdCust:   schemaIntComputedPlain(),
						PfAwsRegion:     schemaStringComputedPlain(),
						PfAwsHostedType: schemaStringComputedPlain(),
						PfAwsAccountId:  schemaStringComputedPlain(),
						PfAwsConnectionId: {
							Type:       schema.TypeString,
							Computed:   true,
							Deprecated: PfDeprecatedField,
						},
						PfGoogleVlanAttachmentName: schemaStringComputedPlain(),
						PfGooglePairingKey:         schemaStringComputedPlain(),
						PfVlanIdPrivate:            schemaIntComputedPlain(),
						PfVlanIdMicrosoft:          schemaIntComputedPlain(),
						PfAzureServiceKey:          schemaStringComputedPlain(),
						PfAzureServiceTag:          schemaIntComputedPlain(),
						PfAzureConnectionType:      schemaStringComputedPlain(),
						PfOracleRegion:             schemaStringComputedPlain(),
						PfVcOcid:                   schemaStringComputedPlain(),
						PfPortCrossConnectOcid:     schemaStringComputedPlain(),
						PfPortCompartmentOcid:      schemaStringComputedPlain(),
						PfAccountId:                schemaStringComputedPlain(),
						PfGatewayId:                schemaStringComputedPlain(),
						PfPortId:                   schemaStringComputedPlain(),
						PfName:                     schemaStringComputedPlain(),
						PfBgpAsn:                   schemaIntComputedPlain(),
						PfBgpCerCidr:               schemaStringComputedPlain(),
						PfBgpIbmCidr:               schemaStringComputedPlain(),
					},
				},
			},
			PfCloudSettings: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfVlanIdPf:      schemaIntComputedPlain(),
						PfVlanIdCust:    schemaIntComputedPlain(),
						PfSvlanIdCust:   schemaIntComputedPlain(),
						PfAwsRegion:     schemaStringComputedPlain(),
						PfAwsHostedType: schemaStringComputedPlain(),
						PfAwsAccountId:  schemaStringComputedPlain(),
						PfAwsConnectionId: {
							Type:       schema.TypeString,
							Computed:   true,
							Deprecated: PfDeprecatedField,
						},
						PfCredentialsUuid:              schemaStringComputedPlain(),
						PfMtu:                          schemaIntComputedPlain(),
						PfAwsDxLocation:                schemaStringComputedPlain(),
						PfAwsDxBandwidth:               schemaStringComputedPlain(),
						PfAwsDxJumboFrameCapable:       schemaBoolComputedPlain(),
						PfAwsDxAwsDevice:               schemaStringComputedPlain(),
						PfAwsDxAwsLogicalDeviceId:      schemaStringComputedPlain(),
						PfAwsDxHasLogicalRedundancy:    schemaBoolComputedPlain(),
						PfAwsDxMacSecCapable:           schemaBoolComputedPlain(),
						PfAwsDxEncryptionMode:          schemaStringComputedPlain(),
						PfAwsVifType:                   schemaStringComputedPlain(),
						PfAwsVifId:                     schemaStringComputedPlain(),
						PfAwsVifBgpPeerId:              schemaStringComputedPlain(),
						PfAwsVifDirectConnectGwId:      schemaStringComputedPlain(),
						PfGoogleRegion:                 schemaStringComputedPlain(),
						PfGoogleProjectId:              schemaStringComputedPlain(),
						PfGoogleVlanAttachmentName:     schemaStringComputedPlain(),
						PfGoogleEdgeAvailabilityDomain: schemaIntComputedPlain(),
						PfGoogleDataplaneVersion:       schemaIntComputedPlain(),
						PfGoogleInterfaceName:          schemaStringComputedPlain(),
						PfGooglePairingKey:             schemaStringComputedPlain(),
						PfGoogleCloudRouterName:        schemaStringComputedPlain(),
						PfCloudState: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfAwsDxConnectionState:           schemaStringComputedPlain(),
									PfAwsDxPortEncryptionStatus:      schemaStringComputedPlain(),
									PfAwsVifState:                    schemaStringComputedPlain(),
									PfGoogleInterconnectState:        schemaStringComputedPlain(),
									PfGoogleInterconnectAdminEnabled: schemaBoolComputedPlain(),
									PfBgpState:                       schemaStringComputedPlain(),
								},
							},
						},
						PfBgpSettings: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfAdvertisedPrefixes:       schemaStringListComputedPlain(),
									PfCustomerAsn:              schemaIntComputedPlain(),
									PfRemoteAsn:                schemaIntComputedPlain(),
									PfL3Address:                schemaStringComputedPlain(),
									PfRemoteAddress:            schemaStringComputedPlain(),
									PfAddressFamily:            schemaStringComputedPlain(),
									PfMd5:                      schemaStringComputedPlain(),
									PfCustomerRouterIp:         schemaStringComputedPlain(),
									PfRemoteRouterIp:           schemaStringComputedPlain(),
									PfGoogleAdvertiseMode:      schemaStringComputedPlain(),
									PfGoogleKeepaliveInterval:  schemaIntComputedPlain(),
									PfGoogleAdvertisedIpRanges: schemaStringListComputedPlain(),
								},
							},
						},
					},
				},
			},
			PfTimeCreated:      schemaStringComputed(PfTimeCreatedDescription),
			PfTimeUpdated:      schemaStringComputed(PfTimeUpdatedDescription),
			PfPop:              schemaStringComputed(PfPopDescription2),
			PfSite:             schemaStringComputed(PfSiteDescription2),
			PfIsAwaitingOnramp: schemaBoolComputed(PfIsAwaitingOnrampDescription),
		},
	}
}

// https://docs.packetfabric.com/api/v2/redoc/#operation/get_connection
func dataSourceCloudServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	circuitID, ok := d.GetOk(PfCloudCircuitId)
	if !ok {
		return diag.Errorf(MessageRequiredCloudCircuitId)
	}

	connInfo, err := c.GetCloudConnInfo(circuitID.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	// Flatten the connection information into a map.
	connInfoMap := flattenCloudConnInfo(connInfo)

	// Update the resource data with the flattened map.
	for k, v := range connInfoMap {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(circuitID.(string) + "-data")

	return diags
}

func flattenCloudConnInfo(connInfo *packetfabric.CloudConnInfo) map[string]interface{} {
	fields := stringsToMap(PfCloudCircuitId, PfCustomerUuid, PfUserUuid, PfState, PfServiceProvider, PfServiceClass, PfPortType, PfSpeed, PfDescription, PfTimeCreated, PfTimeUpdated, PfPop, PfSite, PfIsAwaitingOnramp, PfCloudProviderConnectionId)
	connInfoMap := structToMap(connInfo, fields)
	connInfoMap[PfCloudProviderPop] = connInfo.CloudProvider.Pop
	connInfoMap[PfCloudProviderRegion] = connInfo.CloudProvider.Region

	settingsList := make([]interface{}, 0)
	if connInfo.Settings != nil {
		settingsList = append(settingsList, flattenCloudConnInfoSettings(connInfo.Settings))
	}
	connInfoMap[PfSettings] = settingsList

	cloudSettingsList := make([]interface{}, 0)
	if connInfo.CloudSettings != nil {
		cloudSettingsList = append(cloudSettingsList, flattenCloudConnInfoCloudSettings(connInfo.CloudSettings))
	}
	connInfoMap[PfCloudSettings] = cloudSettingsList

	return connInfoMap
}

func flattenCloudConnInfoSettings(settings *packetfabric.Settings) map[string]interface{} {
	return mapStruct(settings, PfVlanIdPf, PfVlanIdCust, PfSvlanIdCust, PfAwsRegion, PfAwsHostedType, PfAwsAccountId, PfAwsConnectionId, PfGoogleVlanAttachmentName, PfGooglePairingKey, PfVlanIdPrivate, PfVlanIdMicrosoft, PfAzureServiceKey, PfAzureServiceTag, PfOracleRegion, PfVcOcid, PfPortCrossConnectOcid, PfPortCompartmentOcid, PfAccountId, PfGatewayId, PfPortId, PfName, PfBgpAsn, PfBgpCerCidr, PfBgpIbmCidr)
}

func flattenCloudConnInfoCloudSettings(cloudSettings *packetfabric.CloudSettings) map[string]interface{} {
	cloudSettingsMap := mapStruct(cloudSettings, PfVlanIdPf, PfVlanIdCust, PfSvlanIdCust, PfAwsRegion, PfAwsHostedType, PfAwsAccountId, PfAwsConnectionId, PfCredentialsUuid, PfMtu, PfAwsDxLocation, PfAwsDxBandwidth, PfAwsDxJumboFrameCapable, PfAwsDxAwsDevice, PfAwsDxAwsLogicalDeviceId, PfAwsDxHasLogicalRedundancy, PfAwsDxMacSecCapable, PfAwsDxEncryptionMode, PfAwsVifType, PfAwsVifId, PfAwsVifBgpPeerId, PfAwsVifDirectConnectGwId, PfGoogleRegion, PfGoogleProjectId, PfGoogleVlanAttachmentName, PfGoogleEdgeAvailabilityDomain, PfGoogleDataplaneVersion, PfGoogleInterfaceName, PfGooglePairingKey, PfGoogleCloudRouterName)

	cloudStateList := make([]interface{}, 0)
	if cloudSettings.CloudState != nil {
		cloudStateList = append(cloudStateList, flattenCloudConnInfoCloudState(cloudSettings.CloudState))
	}
	cloudSettingsMap[PfCloudState] = cloudStateList

	bgpSettingsList := make([]interface{}, 0)
	if cloudSettings.BgpSettings != nil {
		bgpSettingsList = append(bgpSettingsList, flattenCloudConnInfoBGPSettings(cloudSettings.BgpSettings))
	}
	cloudSettingsMap[PfBgpSettings] = bgpSettingsList

	return cloudSettingsMap
}

func flattenCloudConnInfoBGPSettings(bgpSettings *packetfabric.BgpSettings) map[string]interface{} {
	return mapStruct(bgpSettings, PfCustomerAsn, PfL3Address, PfRemoteAddress, PfAddressFamily, PfMd5, PfAdvertisedPrefixes, PfCustomerRouterIp, PfRemoteRouterIp, PfGoogleAdvertiseMode, PfGoogleKeepaliveInterval, PfGoogleAdvertisedIpRanges)
}

func flattenCloudConnInfoCloudState(cloudState *packetfabric.CloudState) map[string]interface{} {
	return mapStruct(cloudState, PfAwsDxConnectionState, PfAwsDxPortEncryptionStatus, PfAwsVifState, PfBgpState, PfGoogleInterconnectState, PfGoogleInterconnectAdminEnabled)
}
