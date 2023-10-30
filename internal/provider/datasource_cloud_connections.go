package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudConnections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionsRead,
		Schema: map[string]*schema.Schema{
			PfCircuitId: schemaStringRequired(PfCircuitIdDescription),
			PfCloudConnections: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPortType:                  schemaStringComputed(PfPortTypeDescription),
						PfConnectionType:            schemaStringComputed(PfConnectionTypeDescription),
						PfPortCircuitId:             schemaStringComputed(PfPortCircuitIdDescription6),
						PfPendingDelete:             schemaBoolComputed(PfPendingDeleteDescription),
						PfDeleted:                   schemaBoolComputed(PfDeletedDescription2),
						PfSpeed:                     schemaStringComputed(PfSpeedDescriptionI),
						PfState:                     schemaStringComputed(PfStateDescription2),
						PfCloudCircuitId:            schemaStringComputed(PfCloudCircuitIdDescription),
						PfAccountUuid:               schemaStringComputed(PfAccountUuidDescription3),
						PfServiceClass:              schemaStringComputed(PfServiceClassDescription),
						PfServiceProvider:           schemaStringComputed(PfServiceProviderDescription),
						PfServiceType:               schemaStringComputed(PfServiceTypeDescription2),
						PfDescription:               schemaStringComputed(PfConnectionDescription3),
						PfUuid:                      schemaStringComputed(PfUuidDescription2),
						PfCloudProviderConnectionId: schemaStringComputed(PfCloudProviderConnectionIdDescription),
						PfCloudSettings: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfVlanIdPf:      schemaIntComputedPlain(),
									PfVlanIdCust:    schemaIntComputedPlain(),
									PfSvlanIdCust:   schemaIntComputedPlain(),
									PfAwsRegion:     schemaStringComputedPlain(),
									PfAwsHostedType: schemaStringComputedPlain(),
									PfAwsConnectionId: {
										Type:       schema.TypeString,
										Computed:   true,
										Deprecated: PfDeprecatedField,
									},
									PfAwsAccountId:             schemaStringComputedPlain(),
									PfGoogleVlanAttachmentName: schemaStringComputedPlain(),
									PfGooglePairingKey:         schemaStringComputedPlain(),
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
									PfVlanIdPrivate:        schemaIntComputedPlain(),
									PfVlanIdMicrosoft:      schemaIntComputedPlain(),
									PfAzureServiceKey:      schemaStringComputedPlain(),
									PfAzureServiceTag:      schemaIntComputedPlain(),
									PfAzureConnectionType:  schemaStringComputedPlain(),
									PfOracleRegion:         schemaStringComputedPlain(),
									PfVcOcid:               schemaStringComputedPlain(),
									PfPortCrossConnectOcid: schemaStringComputedPlain(),
									PfPortCompartmentOcid:  schemaStringComputedPlain(),
									PfAccountId:            schemaStringComputedPlain(),
									PfGatewayId:            schemaStringComputedPlain(),
									PfPortId:               schemaStringComputedPlain(),
									PfName:                 schemaStringComputedPlain(),
									PfBgpAsn:               schemaIntComputedPlain(),
									PfBgpCerCidr:           schemaStringComputedPlain(),
									PfBgpIbmCidr:           schemaStringComputedPlain(),
									PfPublicIp:             schemaStringComputedPlain(),
									PfNatPublicIp:          schemaStringComputedPlain(),
									PfPrimaryPublicIp:      schemaStringComputedPlain(),
									PfSecondaryPublicIp:    schemaStringComputedPlain(),
								},
							},
						},
						PfUserUuid:     schemaStringComputed(PfUserUuidDescription),
						PfCustomerUuid: schemaStringComputed(PfCustomerUuidDescription2),
						PfTimeCreated:  schemaStringComputed(PfTimeCreatedDescription),
						PfTimeUpdated:  schemaStringComputed(PfTimeUpdatedDescription),
						PfCloudProvider: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfPop:  schemaStringComputed(PfPopDescription),
									PfSite: schemaStringComputed(PfSiteDescription),
								},
							},
						},
						PfPop:  schemaStringComputed(PfPopDescription),
						PfSite: schemaStringComputed(PfSiteDescription),
						PfBgpStateList: {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: PfBgpStateListDescription,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfBgpSettingsUuid: schemaStringComputed(PfBgpSettingsUuidDescription),
									PfBgpState:        schemaStringComputed(PfBgpStateDescription),
								},
							},
						},
						PfCloudRouterName:      schemaStringComputed(PfCloudRouterNameDescription),
						PfCloudRouterAsn:       schemaIntComputed(PfCloudRouterAsnDescription),
						PfCloudRouterCircuitId: schemaStringComputed(PfCloudRouterCircuitIdDescription2),
						PfNatCapable:           schemaBoolComputed(PfNatCapableDescription2),
						PfDnatCapable:          schemaBoolComputed(PfDnatCapableDescription),
						PfZone:                 schemaStringComputed(PfZoneDescription4),
						PfVlan:                 schemaIntComputed(PfVlanDescription4),
						PfDesiredNat:           schemaStringComputed(PfDesiredNatDescription),
						PfSubscriptionTerm:     schemaIntComputed(PfSubscriptionTermDescription),
					},
				},
			},
		},
	}
}

func dataSourceCloudConnectionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk(PfCircuitId)
	if !ok {
		return diag.Errorf(MessageMissingCircuitIdDetail)
	}
	awsConns, err := c.ListAwsRouterConnections(cID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfCloudConnections, flattenCloudConnnections(&awsConns))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(cID.(string) + "-data")
	return diags
}

func flattenCloudConnnections(conns *[]packetfabric.CloudRouterConnectionReadResponse) []interface{} {
	fields := stringsToMap(PfUuid, PfPortType, PfConnectionType, PfPortCircuitId, PfPendingDelete, PfDeleted, PfSpeed, PfState, PfCloudCircuitId, PfAccountUuid, PfServiceClass, PfServiceProvider, PfServiceType, PfDescription, PfCloudProviderConnectionId, PfUserUuid, PfCustomerUuid, PfTimeCreated, PfTimeUpdated, PfPop, PfSite, PfCloudRouterName, PfCloudRouterAsn, PfCloudRouterCircuitId, PfNatCapable, PfDnatCapable, PfZone, PfVlan, PfSubscriptionTerm)
	if conns != nil {
		flattens := make([]interface{}, len(*conns), len(*conns))
		for i, conn := range *conns {
			flatten := structToMap(&conn, fields)
			flatten[PfBgpStateList] = flattenBgpStateList(&conn.BgpStateList)
			flatten[PfCloudProvider] = flattenCloudProvider(&conn.CloudProvider)
			flatten[PfCloudSettings] = flattenCloudSettings(&conn.CloudSettings)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenCloudProvider(provider *packetfabric.CloudProvider) []interface{} {
	fields := stringsToMap(PfPop, PfSite)
	flattens := make([]interface{}, 0)
	if provider != nil {
		flattens = append(flattens, structToMap(provider, fields))
	}
	return flattens
}

func flattenBgpStateList(BgpStateList *[]packetfabric.BgpStateObj) []interface{} {
	fields := stringsToMap(PfBgpSettingsUuid, PfBgpState)
	if BgpStateList != nil {
		flattens := make([]interface{}, len(*BgpStateList), len(*BgpStateList))

		for i, bgpStateObj := range *BgpStateList {
			flattens[i] = structToMap(&bgpStateObj, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenCloudSettings(setts *packetfabric.CloudSettings) []interface{} {
	fields := stringsToMap(PfAccountId, PfAwsAccountId, PfAwsConnectionId, PfAwsHostedType, PfAwsRegion, PfAzureConnectionType, PfAzureServiceKey, PfAzureServiceTag, PfBgpAsn, PfBgpCerCidr, PfBgpIbmCidr, PfGatewayId, PfGooglePairingKey, PfGoogleVlanAttachmentName, PfName, PfNatPublicIp, PfOracleRegion, PfPortCompartmentOcid, PfPortCrossConnectOcid, PfPortId, PfPublicIp, PfSvlanIdCust, PfVcOcid, PfVlanIdCust, PfVlanIdMicrosoft, PfVlanIdPf, PfVlanIdPrivate, PfPrimaryPublicIp, PfSecondaryPublicIp)
	flattens := make([]interface{}, 0)
	if setts != nil {
		flattens = append(flattens, structToMap(setts, fields))
	}
	return flattens
}
