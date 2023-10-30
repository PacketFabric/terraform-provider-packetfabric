package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLocations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSOurceLocationsRead,
		Schema: map[string]*schema.Schema{
			PfLocations: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfPop:               schemaStringComputed(PfPopDescription4),
						PfRegion:            schemaStringComputed(PfSettingsAwsRegionDescription),
						PfMarket:            schemaStringComputed(PfMarketDescription2),
						PfMarketDescription: schemaStringComputed(PfMarketDescriptionDescription2),
						PfVendor:            schemaStringComputed(PfVendorDescription),
						PfSite:              schemaStringComputed(PfSiteDescription3),
						PfSiteCode:          schemaStringComputed(PfSiteCodeDescription2),
						PfType:              schemaStringComputed(PfTypeCapitalized),
						PfStatus:            schemaStringComputed(PfStatusDescription),
						PfLatitude:          schemaStringComputed(PfLatitudeDescription),
						PfLongitude:         schemaStringComputed(PfLongitudeDescription),
						PfTimezone:          schemaStringComputed(PfTimezoneDescription),
						PfNotes:             schemaStringComputed(PfNotesDescription),
						PfPcode:             schemaIntComputed(PfPcode),
						PfLeadTime:          schemaStringComputed(PfLeadTimeDescription),
						PfSingleArmed:       schemaBoolComputed(PfSingleArmedDescription),
						PfAddress1:          schemaStringComputed(PfAddress1Description),
						PfAddress2:          schemaStringComputed(PfAddress2Description),
						PfCity:              schemaStringComputed(PfCityDescription),
						PfState:             schemaStringComputed(PfStateDescription7),
						PfPostal:            schemaStringComputed(PfPostalDescription),
						PfCountry:           schemaStringComputed(PfCountryDescription),
						PfNetworkProvider:   schemaStringComputed(PfNetworkProviderDescription),
						PfTimeCreated:       schemaStringComputed(PfTimeCreatedDescription9),
						PfEnniSupported:     schemaBoolComputed(PfEnniSupportedDescription),
					},
				},
			},
		},
	}
}

// https://docs.packetfabric.com/api/v2/redoc/#operation/get_location_list
func dataSOurceLocationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	locations, err := c.ListLocations()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfLocations, flattenLocations(&locations))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenLocations(locations *[]packetfabric.Location) []interface{} {
	fields := stringsToMap(PfPop, PfRegion, PfMarket, PfMarketDescription, PfVendor, PfSite, PfSiteCode, PfType, PfStatus, PfLatitude, PfLongitude, PfTimezone, PfNotes, PfPcode, PfLeadTime, PfSingleArmed, PfAddress1, PfAddress2, PfCity, PfState, PfPostal, PfCountry, PfNetworkProvider, PfTimeCreated, PfEnniSupported)
	if locations != nil {
		flattens := make([]interface{}, len(*locations), len(*locations))

		for i, location := range *locations {
			flattens[i] = structToMap(&location, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}
