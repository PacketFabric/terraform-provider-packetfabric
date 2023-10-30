package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceCloudLocations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudLocationsRead,
		Schema: map[string]*schema.Schema{
			PfCloudProvider:       schemaStringRequiredValidate(PfCloudProviderDescription, validateProvider()),
			PfCloudConnectionType: schemaStringRequiredValidate(PfCloudConnectionTypeDescription, validateCloudConnectionType()),
			PfNatCapable:          schemaBoolOptionalDefault(PfNatCapableDescription, false),
			PfHasCloudRouter:      schemaBoolOptionalDefault(PfHasCloudRouterDescription, false),
			PfAnyType:             schemaBoolOptionalDefault(PfAnyTypeDescription, false),
			PfPop:                 schemaStringOptionalValidate(PfPopDescription6, validation.StringIsNotEmpty),
			PfCity:                schemaStringOptionalValidate(PfCityDescription2, validation.StringIsNotEmpty),
			PfState:               schemaStringOptionalValidate(PfStateDescription8, validation.StringIsNotEmpty),
			PfMarket:              schemaStringOptionalValidate(PfMarketDescription4, validation.StringIsNotEmpty),
			PfRegion:              schemaStringOptionalValidate(PfRegionDescription2, validation.StringIsNotEmpty),
			PfCloudLocations: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: PfCloudLocationsDescription,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPop:                               schemaStringComputed(PfPopDescription),
						PfRegion:                            schemaStringComputed(PfRegionDescription3),
						PfMarket:                            schemaStringComputed(PfMarketDescription5),
						PfMarketDescription:                 schemaStringComputed(PfMarketDescriptionDescription),
						PfZones:                             schemaStringListComputedNotEmpty(PfEmptyString),
						PfVendor:                            schemaStringComputed(PfVendorDescription2),
						PfSite:                              schemaStringComputed(PfSiteDescription4),
						PfSiteCode:                          schemaStringComputed(PfSiteCodeDescription3),
						PfType:                              schemaStringComputed(PfTypeDescription3),
						PfStatus:                            schemaStringComputed(PfStatusDescription3),
						PfLatitude:                          schemaStringComputed(PfLatitudeDescription2),
						PfLongitude:                         schemaStringComputed(PfLongitudeDescription2),
						PfTimezone:                          schemaStringComputed(PfTimezoneDescription2),
						PfNotes:                             schemaStringComputed(PfNotesDescription2),
						PfPcode:                             schemaFloatComputed(PfPcodeDescription),
						PfLeadTime:                          schemaStringComputed(PfLeadTimeDescription2),
						PfSingleArmed:                       schemaBoolComputed(PfSingleArmedDescription2),
						PfAddress1:                          schemaStringComputed(PfAddress1Description2),
						PfAddress2:                          schemaStringComputed(PfAddress2Description2),
						PfCity:                              schemaStringComputed(PfCityDescription3),
						PfState:                             schemaStringComputed(PfStateDescription9),
						PfPostal:                            schemaStringComputed(PfPostalDescription2),
						PfCountry:                           schemaStringComputed(PfCountryDescription2),
						PfCloudProvider:                     schemaStringComputed(PfCloudProviderDescription2),
						PfCloudConnectionRegion:             schemaStringComputed(PfCloudConnectionRegionDescription),
						PfCloudConnectionHostedType:         schemaStringComputed(PfCloudConnectionHostedTypeDescription),
						PfCloudConnectionRegionDescription2: schemaStringComputed(PfCloudConnectionRegionDescriptionDescription),
						PfNetworkProvider:                   schemaStringComputed(PfNetworkProviderDescription2),
						PfTimeCreated:                       schemaStringComputed(PfTimeCreatedDescriptionA),
						PfEnniSupported:                     schemaBoolComputed(PfEnniSupportedDescription2),
					},
				},
			},
		},
	}
}

func dataSourceCloudLocationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cp, ok := d.GetOk(PfCloudProvider)
	if !ok {
		return diag.Errorf(MessageMissingCP)
	}
	ccType, ok := d.GetOk(PfCloudConnectionType)
	if !ok {
		return diag.Errorf(MessageMissingCPType)
	}
	natCapable, hasCloudRouter, anyType := _extractOptionalLocationBoolValues(d)
	pop, city, state, market, region := _extractOptionalLocationStringValues(d)
	locations, err := c.GetCloudLocations(cp.(string), ccType.(string),
		natCapable, hasCloudRouter, anyType, pop, city, state, market, region)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfCloudLocations, flattenCloudLocations(&locations))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(cp.(string))
	return diags
}

func flattenCloudLocations(locs *[]packetfabric.CloudLocation) []interface{} {
	fields := stringsToMap(PfPop, PfRegion, PfMarket, PfMarketDescription, PfVendor, PfSite, PfSiteCode, PfType, PfStatus, PfLatitude, PfLongitude, PfTimezone, PfNotes, PfPcode, PfLeadTime, PfSingleArmed, PfAddress1, PfAddress2, PfZones, PfCity, PfState, PfPostal, PfCountry, PfCloudProvider, PfNetworkProvider, PfTimeCreated, PfEnniSupported)
	flattens := make([]interface{}, len(*locs))
	for i, loc := range *locs {
		flatten := structToMap(loc, fields)
		// TODO: do we need these "nil guards"? should they be built into structToMap
		//if loc.Timezone != nil {
		//	flatten["timezone"] = loc.Timezone
		//}
		//if loc.Notes != nil {
		//	flatten["notes"] = loc.Notes
		//}
		//if loc.Pcode != nil {
		//	flatten["pcode"] = loc.Pcode
		//}
		//if loc.Address2 != nil {
		//	flatten["address2"] = loc.Address2
		//}
		flatten[PfCloudConnectionRegion] = loc.CloudConnectionDetails.Region
		flatten[PfCloudConnectionHostedType] = loc.CloudConnectionDetails.HostedType
		flatten[PfCloudConnectionRegionDescription2] = loc.CloudConnectionDetails.RegionDescription
		flattens[i] = flatten
	}
	return flattens
}

func _extractOptionalLocationBoolValues(d *schema.ResourceData) (natCapable, hasCloudRouter, anyType bool) {
	natCapable = false
	hasCloudRouter = false
	anyType = false
	if nat, ok := d.GetOk(PfNatCapable); ok {
		natCapable = nat.(bool)
	}
	if cloudRouter, ok := d.GetOk(PfHasCloudRouter); ok {
		hasCloudRouter = cloudRouter.(bool)
	}
	if any, ok := d.GetOk(PfAnyType); ok {
		anyType = any.(bool)
	}
	return
}

func _extractOptionalLocationStringValues(d *schema.ResourceData) (pop, city, state, market, region string) {
	pop = PfEmptyString
	city = PfEmptyString
	state = PfEmptyString
	market = PfEmptyString
	region = PfEmptyString
	if p, ok := d.GetOk(PfPop); ok {
		pop = p.(string)
	}
	if c, ok := d.GetOk(PfCity); ok {
		city = c.(string)
	}
	if s, ok := d.GetOk(PfState); ok {
		state = s.(string)
	}
	if m, ok := d.GetOk(PfMarket); ok {
		market = m.(string)
	}
	if r, ok := d.GetOk(PfRegion); ok {
		region = r.(string)
	}
	return
}
