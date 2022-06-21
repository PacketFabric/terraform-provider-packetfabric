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
			"locations": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pop": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "POP name",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Region short name",
						},
						"market": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Market code",
						},
						"market_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Market description, Long market name",
						},
						"vendor": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Vendor name",
						},
						"site": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site name",
						},
						"site_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site code",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Type",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Current status of the site",
						},
						"latitude": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's geo location - latitude",
						},
						"longitude": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's geo location - longitude",
						},
						"timezone": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Local timezone of the site",
						},
						"notes": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Additional notes",
						},
						"pcode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "pcode",
						},
						"lead_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's lead time",
						},
						"single_armed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Optional:    true,
							Description: "Indication that site is single armed",
						},
						"address1": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's address - line one",
						},
						"address2": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's address - line two",
						},
						"city": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's address - city name",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's address - state code",
						},
						"postal": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's address - postal code",
						},
						"country": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Site's address - country code",
						},
						"network_provider": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Network provider for ports at this location",
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Date the location was added",
						},
						"enni_supported": {
							Type:        schema.TypeBool,
							Computed:    true,
							Optional:    true,
							Description: "Support for ENNI",
						},
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
	err = d.Set("locations", flattenLocations(&locations))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenLocations(locations *[]packetfabric.Location) []interface{} {
	if locations != nil {
		flattens := make([]interface{}, len(*locations), len(*locations))

		for i, location := range *locations {
			flatten := make(map[string]interface{})
			flatten["pop"] = location.Pop
			flatten["region"] = location.Region
			flatten["market"] = location.Market
			flatten["market_description"] = location.MarketDescription
			flatten["vendor"] = location.Vendor
			flatten["site"] = location.Site
			flatten["site_code"] = location.SiteCode
			flatten["type"] = location.Type
			flatten["status"] = location.Status
			flatten["latitude"] = location.Latitude
			flatten["longitude"] = location.Longitude
			flatten["timezone"] = location.Timezone
			flatten["notes"] = location.Notes
			flatten["pcode"] = location.Pcode
			flatten["lead_time"] = location.LeadTime
			flatten["single_armed"] = location.SingleArmed
			flatten["address1"] = location.Address1
			flatten["address2"] = location.Address2
			flatten["city"] = location.City
			flatten["state"] = location.State
			flatten["postal"] = location.Postal
			flatten["country"] = location.Country
			flatten["network_provider"] = location.NetworkProvider
			flatten["time_created"] = location.TimeCreated
			flatten["enni_supported"] = location.EnniSupported
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
