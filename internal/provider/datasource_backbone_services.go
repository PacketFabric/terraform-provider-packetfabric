package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceBackboneServices() *schema.Resource {
	return &schema.Resource{
		ReadContext: backboneServiceRead,
		Schema: map[string]*schema.Schema{
			PfBackboneServices: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfVcCircuitId:  schemaStringComputed(PfVcCircuitIdDescription),
						PfCustomerUuid: schemaStringComputed(PfCustomerUuidDescription),
						PfState:        schemaStringComputed(PfStateDescription),
						PfServiceType:  schemaStringComputed(PfServiceTypeDescription),
						PfServiceClass: schemaStringComputed(PfServiceClassDescription2),
						PfMode:         schemaStringComputed(PfModeDescription),
						PfConnected:    schemaBoolComputed(PfConnectedDescription),
						PfBandwidth: {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: PfBandwidthDescription,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfAccountUuid:      schemaStringComputed(PfAccountUuidDescription3),
									PfLonghaulType:     schemaStringComputed(PfLonghaulTypeDescription2),
									PfSubscriptionTerm: schemaIntComputed(PfSubscriptionTermDescription5),
									PfSpeed:            schemaStringComputed(PfSpeedDescriptionE),
								},
							},
						},
						PfDescription:     schemaStringComputed(PfServiceDescription3),
						PfRateLimitIn:     schemaIntComputed(PfRateLimitInDescription),
						PfRateLimitOut:    schemaIntComputed(PfRateLimitOutDescription),
						PfFlexBandwidthId: schemaStringComputed(PfFlexBandwidthIdDescription2),
						PfTimeCreated:     schemaStringComputed(PfTimeCreatedDescription),
						PfTimeUpdated:     schemaStringComputed(PfTimeUpdatedDescription),
						PfInterfaces: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfPortCircuitId:      schemaStringComputed(PfPortCircuitIdDescription2),
									PfPop:                schemaStringComputed(PfPopDescription8),
									PfSite:               schemaStringComputed(PfSiteDescription6),
									PfSiteName:           schemaStringComputed(PfSiteNameDescription),
									PfCustomerSiteCode:   schemaStringComputed(PfCustomerSiteCodeDescription),
									PfCustomerSiteName:   schemaStringComputed(PfCustomerSiteNameDescription),
									PfSpeed:              schemaStringComputed(PfSpeedDescriptionG),
									PfMedia:              schemaStringComputed(PfMediaDescription2),
									PfZone:               schemaStringComputed(PfZoneDescription3),
									PfDescription:        schemaStringComputed(PfInterfacesDescription),
									PfVlan:               schemaIntComputed(PfVlanDescription3),
									PfSvlan:              schemaIntComputed(PfSvlanDescription),
									PfUntagged:           schemaBoolComputed(PfUntaggedDescription3),
									PfProvisioningStatus: schemaStringComputed(PfProvisioningStatusDescription),
									PfAdminStatus:        schemaStringComputed(PfAdminStatusDescription),
									PfOperationalStatus:  schemaStringComputed(PfOperationalStatusDescription),
									PfCustomerUuid:       schemaStringComputed(PfCustomerUuidDescription4),
									PfCustomerName:       schemaStringComputed(PfCustomerNameDescription2),
									PfRegion:             schemaStringComputed(PfRegionDescription5),
									PfIsCloud:            schemaBoolComputed(PfIsCloudDescription),
									PfIsPtp:              schemaBoolComputed(PfIsPtpDescription),
									PfTimeCreated:        schemaStringComputed(PfTimeCreatedDescription3),
									PfTimeUpdated:        schemaStringComputed(PfTimeUpdatedDescription4),
								},
							},
						},
					},
				},
			},
		},
	}
}

func backboneServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	services, err := c.GetServices()
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set(PfBackboneServices, flattenBackboneService(&services)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(string(uuid.New().String()))
	return diags
}

func flattenBackboneService(services *[]packetfabric.Services) []interface{} {
	fields := stringsToMap(PfVcCircuitId, PfCustomerUuid, PfState, PfServiceType, PfServiceClass, PfMode, PfConnected, PfBandwidth, PfDescription, PfRateLimitIn, PfRateLimitOut, PfFlexBandwidthId, PfTimeCreated, PfTimeUpdated)
	flattens := make([]interface{}, len(*services))
	for i, service := range *services {
		flatten := structToMap(&service, fields)
		flatten[PfInterfaces] = flattenBackBoneInterfaces(&service.Interfaces)
		flattens[i] = flatten
	}
	return flattens
}

func flattenBackBoneInterfaces(interfs *[]packetfabric.ServiceInterface) []interface{} {
	fields := stringsToMap(PfPortCircuitId, PfPop, PfSite, PfSiteName, PfSpeed, PfMedia, PfZone, PfDescription, PfVlan, PfSvlan, PfUntagged, PfProvisioningStatus, PfAdminStatus, PfOperationalStatus, PfCustomerUuid, PfCustomerName, PfRegion, PfIsCloud, PfIsPtp, PfTimeCreated, PfTimeUpdated)
	if interfs != nil {
		flattens := make([]interface{}, len(*interfs))
		for i, interf := range *interfs {
			flattens[i] = structToMap(interf, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenServiceBandwidth(bandw *packetfabric.Bandwidth) []interface{} {
	flattens := make([]interface{}, 0)
	fields := stringsToMap(PfAccountUuid, PfSubscriptionTerm, PfSpeed, PfLonghaulType)
	if bandw != nil {
		flattens = append(flattens, structToMap(bandw, fields))
	}
	return flattens
}
