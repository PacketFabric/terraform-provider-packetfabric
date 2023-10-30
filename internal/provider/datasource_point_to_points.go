package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourcePointToPoints() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourcePointToPointsRead,
		Schema: map[string]*schema.Schema{
			PfPointToPoints: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: PfPointToPointsDescription2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPtpUuid:      schemaStringComputed(PfPtpUuidDescription),
						PfPtpCircuitId: schemaStringComputed(PfPtpCircuitIdDescription),
						PfDescription:  schemaStringComputed(PfPointToPointsDescription),
						PfSpeed:        schemaStringComputed(PfSpeedDescription9),
						PfMedia:        schemaStringComputed(PfMediaDescription),
						PfState:        schemaStringComputed(PfStateDescriptionA),
						PfBilling: {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									PfAccountUuid:      schemaStringComputed(PfAccountUuidDescription),
									PfSubscriptionTerm: schemaIntComputed(PfSubscriptionTermDescription6),
									PfContractedSpeed:  schemaStringComputed(PfContractedSpeedDescription),
								},
							},
						},
						PfTimeCreated:  schemaStringComputed(PfTimeCreatedDescriptionB),
						PfTimeUpdated:  schemaStringComputed(PfTimeUpdatedDescription8),
						PfDeleted:      schemaBoolComputed(PfDeletedDescription3),
						PfServiceClass: schemaStringComputed(PfServiceClassDescription5),
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

func datasourcePointToPointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	ptps, err := c.ListPointToPoints()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set(PfPointToPoints, flattenPointToPoints(ptps)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenPointToPoints(ptps *[]packetfabric.PointToPointResp) []interface{} {
	fields := stringsToMap(PfPtpUuid, PfPtpCircuitId, PfDescription, PfSpeed, PfMedia, PfState, PfTimeCreated, PfTimeUpdated, PfDeleted, PfServiceClass)
	if ptps != nil {
		flattens := make([]interface{}, len(*ptps))
		for i, ptp := range *ptps {
			flatten := structToMap(&ptp, fields)
			flatten[PfBilling] = flattenBilling(&ptp.Billing)
			flatten[PfInterfaces] = flattenPtpInterf(&ptp.Interfaces)
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenBilling(billing *packetfabric.Billing) []interface{} {
	flattens := make([]interface{}, 0)
	if billing != nil {
		flattens = append(flattens, mapStruct(billing, PfAccountUuid, PfSubscriptionTerm, PfContractedSpeed))
	}
	return flattens
}

func flattenPtpInterf(interfs *[]packetfabric.Interfaces) []interface{} {
	fields := stringsToMap(PfPortCircuitId, PfPop, PfSite, PfSiteName, PfCustomerSiteCode, PfCustomerSiteName, PfSpeed, PfMedia, PfZone, PfDescription, PfVlan, PfUntagged, PfProvisioningStatus, PfAdminStatus, PfOperationalStatus, PfCustomerUuid, PfCustomerName, PfRegion, PfIsCloud, PfIsPtp, PfTimeCreated, PfTimeUpdated)
	if interfs != nil {
		flattens := make([]interface{}, len(*interfs), len(*interfs))
		for i, interf := range *interfs {
			flattens[i] = structToMap(&interf, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}
