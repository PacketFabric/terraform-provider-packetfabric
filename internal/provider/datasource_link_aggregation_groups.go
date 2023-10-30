package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceLinkAggregationGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceLinkAggregationGroupsRead,
		Schema: map[string]*schema.Schema{
			PfLagCircuitId: schemaStringRequired(PfLagCircuitIdDescription),
			PfInterfaces: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfAutoneg:            schemaBoolComputed(PfAutonegDescription),
						PfPortCircuitId:      schemaStringComputed(PfPortCircuitIdDescription7),
						PfState:              schemaStringComputed(PfStateDescription4),
						PfStatus:             schemaStringComputed(PfStatusDescription2),
						PfSpeed:              schemaStringComputed(PfSpeedDescriptionL),
						PfMedia:              schemaStringComputed(PfMediaDescription3),
						PfZone:               schemaStringComputed(PfZoneDescription5),
						PfRegion:             schemaStringComputed(PfRegionDescription3),
						PfMarket:             schemaStringComputed(PfMarketDescription5),
						PfMarketDescription:  schemaStringComputed(PfMarketDescriptionDescription),
						PfPop:                schemaStringComputed(PfPopDescriptionC),
						PfSite:               schemaStringComputed(PfSiteDescription7),
						PfSiteCode:           schemaStringComputed(PfSiteCodeDescription),
						PfOperationalStatus:  schemaStringComputed(PfOperationalStatusDescription2),
						PfAdminStatus:        schemaStringComputed(PfAdminStatusDescription2),
						PfMtu:                schemaIntComputed(PfMtuDescription),
						PfDescription:        schemaStringComputed(PfInterfacesDescription2),
						PfVcMode:             schemaStringComputed(PfVcModeDescription),
						PfIsLag:              schemaBoolComputed(PfIsLagDescription),
						PfIsLagMember:        schemaBoolComputed(PfIsLagMemberDescription),
						PfIsCloud:            schemaBoolComputed(PfIsCloudDescription2),
						PfIsPtp:              schemaBoolComputed(PfIsPtpDescription2),
						PfIsNni:              schemaBoolComputed(PfIsNniDescription),
						PfLagInterval:        schemaStringComputed(PfLagIntervalDescription),
						PfMemberCount:        schemaIntComputed(PfMemberCountDescription),
						PfParentLagCircuitId: schemaStringComputed(PfParentLagCircuitIdDescription),
						PfAccountUuid:        schemaStringComputed(PfAccountUuidDescription4),
						PfSubscriptionTerm:   schemaIntComputed(PfSubscriptionTermDescription7),
						PfDisabled:           schemaBoolComputed(PfDisabledDescription3),
						PfCustomerName:       schemaStringComputed(PfCustomerNameDescription3),
						PfCustomerUuid:       schemaStringComputed(PfCustomerUuidDescription5),
						PfTimeCreated:        schemaStringComputed(PfTimeCreatedDescription6),
						PfTimeUpdated:        schemaStringComputed(PfTimeUpdatedDescription5),
					},
				},
			},
		},
	}
}

func datasourceLinkAggregationGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	lagID := d.Get(PfLagCircuitId)
	interfs, err := c.GetLAGInterfaces(lagID.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfInterfaces, flattenInterfaces(interfs))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(lagID.(string) + "-data")
	return diags

}
