package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudConnectionRead,
		Schema: map[string]*schema.Schema{
			PfCircuitId:                 schemaStringRequired(PfCircuitIdDescription),
			PfConnectionId:              schemaStringRequired(PfConnectionIdDescription),
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
						PfVlanIdPf:                 schemaIntComputedPlain(),
						PfVlanIdCust:               schemaIntComputedPlain(),
						PfSvlanIdCust:              schemaIntComputedPlain(),
						PfAwsRegion:                schemaStringComputedPlain(),
						PfAwsHostedType:            schemaStringComputedPlain(),
						PfAwsConnectionId:          schemaStringComputedPlain(),
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
						PfPrimaryPublicIp:      schemaStringComputedPlain(),
						PfSecondaryPublicIp:    schemaStringComputedPlain(),
						PfNatPublicIp:          schemaStringComputedPlain(),
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
			PfSubscriptionTerm:     schemaIntComputed(PfSubscriptionTermDescription2),
		},
	}
}

func dataSourceCloudConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk(PfCircuitId)
	if !ok {
		return diag.Errorf(MessageMissingCircuitIdDetail)
	}
	connCID, ok := d.GetOk(PfConnectionId)
	if !ok {
		return diag.Errorf(MesssageCRCIdRequired)
	}
	awsConn, err := c.ReadCloudRouterConnection(cID.(string), connCID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	// Flatten the connection information into a map.
	connInfoMap := flattenCloudConnection(awsConn)

	// Update the resource data with the flattened map.
	for k, v := range connInfoMap {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(connCID.(string) + "-data")
	return diags
}

func flattenCloudConnection(conn *packetfabric.CloudRouterConnectionReadResponse) map[string]interface{} {
	connInfoMap := mapStruct(conn, PfUuid, PfPortType, PfConnectionType, PfPortCircuitId, PfPendingDelete, PfDeleted, PfSpeed, PfState, PfCloudCircuitId, PfAccountUuid, PfServiceClass, PfServiceProvider, PfServiceType, PfDescription, PfCloudProviderConnectionId, PfUserUuid, PfCustomerUuid, PfTimeCreated, PfTimeUpdated, PfPop, PfSite, PfCloudRouterName, PfCloudRouterAsn, PfCloudRouterCircuitId, PfNatCapable, PfDnatCapable, PfZone, PfVlan, PfSubscriptionTerm)
	connInfoMap[PfBgpStateList] = flattenBgpStateList(&conn.BgpStateList)
	connInfoMap[PfCloudProvider] = flattenCloudProvider(&conn.CloudProvider)
	connInfoMap[PfCloudSettings] = flattenCloudSettings(&conn.CloudSettings)

	return connInfoMap
}
