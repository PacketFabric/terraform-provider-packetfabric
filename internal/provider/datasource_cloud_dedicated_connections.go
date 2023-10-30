package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceDedicatedCloudConnections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDedicatedConRead,
		Schema: map[string]*schema.Schema{
			PfDedicatedConnections: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfUuid:            schemaStringComputed(PfUuidDescription3),
						PfCustomerUuid:    schemaStringComputed(PfCustomerUuidDescription2),
						PfUserUuid:        schemaStringComputed(PfUserUuidDescription),
						PfServiceProvider: schemaStringComputed(PfServiceProviderDescription2),
						PfPortType:        schemaStringComputed(PfPortTypeDescription2),
						PfDeleted:         schemaBoolComputed(PfDeletedDescription),
						PfTimeUpdated:     schemaStringComputed(PfTimeUpdatedDescription),
						PfTimeCreated:     schemaStringComputed(PfTimeCreatedDescription),
						PfCloudCircuitId:  schemaStringComputed(PfUuidDescription3),
						PfAccountUuid:     schemaStringComputed(PfAccountIdDescription),
						PfCloudProvider: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfPop:  schemaStringComputed(PfPopDescriptionA),
									PfSite: schemaStringComputed(PfSiteDescription3),
								},
							},
						},
						PfPop:               schemaStringComputed(PfPopDescriptionA),
						PfSite:              schemaStringComputed(PfSiteDescription3),
						PfServiceClass:      schemaStringComputed(PfVcServiceClassDescription),
						PfDescription:       schemaStringComputed(PfConnectionDescription6),
						PfState:             schemaStringComputed(PfStateDescription5),
						PfSettingsAwsRegion: schemaStringComputed(PfSettingsAwsRegionDescription),
						PfSettings: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfAwsRegion: schemaStringComputed(PfAwsRegionDescription2),
									PfZoneDest:  schemaStringComputed(PfZoneDestDescription),
									PfAutoneg:   schemaBoolComputed(PfAutonegDescription2),
								},
							},
						},
						PfIsCloudRouterConnection: schemaBoolComputed(PfIsCloudRouterConnectionDescription),
						PfSpeed:                   schemaStringComputed(PfSpeedDescriptionJ),
						PfIsLag:                   schemaBoolComputed(PfIsLagDescription),
					},
				},
			},
		},
	}
}

// https://docs.packetfabric.com/api/v2/redoc/#operation/get_connections_dedicated_list
func dataSourceDedicatedConRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	sessions, err := c.GetCurrentCustomersDedicated()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfDedicatedConnections, flattenDedicatedConns(&sessions))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenDedicatedConns(conns *[]packetfabric.DedicatedConnResp) []interface{} {
	fields := stringsToMap(PfUuid, PfPortType, PfDeleted, PfSpeed, PfState, PfCloudCircuitId, PfAccountUuid, PfServiceClass, PfServiceProvider, PfDescription, PfUserUuid, PfCustomerUuid, PfTimeCreated, PfTimeUpdated, PfIsCloudRouterConnection, PfPop, PfSite, PfIsLag)
	if conns != nil {
		flattens := make([]interface{}, len(*conns))
		for i, conn := range *conns {
			flatten := structToMap(&conn, fields)
			flatten[PfSettings] = flattenCloudServiceSettings(&conn.Settings)
			flatten[PfCloudProvider] = flattenCloudServiceProvider(&conn.CloudProvider)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenCloudServiceProvider(provider *packetfabric.CloudServiceProvider) []interface{} {
	fields := stringsToMap(PfPop, PfSite)
	flattens := make([]interface{}, 0)
	if provider != nil {
		flattens = append(flattens, structToMap(provider, fields))
	}
	return flattens
}

func flattenCloudServiceSettings(settings *packetfabric.CloudServiceSettings) []interface{} {
	fields := stringsToMap(PfAwsRegion, PfZoneDest, PfAutoneg)
	flattens := make([]interface{}, 0)
	if settings != nil {
		flattens = append(flattens, structToMap(settings, fields))
	}
	return flattens
}
