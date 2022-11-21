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
			"cloud_provider": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"aws", "azure", "packet", "google", "ibm", "oracle", "salesforce", "webex"}, true),
				Description:  "Filter locations by cloud provider.",
			},
			"cloud_connection_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"hosted", "dedicated"}, true),
				Description:  "Filter locations by cloud connection type.",
			},
			"cloud_locations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pop": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Point of Presence for the cloud provider location\n\t\tExample: LAX1",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Region.",
						},
						"market": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Market.",
						},
						"market_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Port Market description.",
						},
						"zones": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},
						"vendor": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud vendor.",
						},
						"site": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service site.",
						},
						"site_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service site code.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service type.",
						},
						"status": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tye cloud service staus.",
						},
						"latitude": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service location latitude.",
						},
						"longitude": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service location longitude.",
						},
						"timezone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service timezone.",
						},
						"notes": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service notes.",
						},
						"pcode": {
							Type:        schema.TypeFloat,
							Optional:    true,
							Description: "The cloud service PCODE.",
						},
						"lead_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service lead time.",
						},
						"single_armed": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True if cloud service is single armed.",
						},
						"address1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service address 1.",
						},
						"address2": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service address 2.",
						},
						"city": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service city.",
						},
						"state": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service state.",
						},
						"postal": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service postal code.",
						},
						"country": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service country.",
						},
						"cloud_provider": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service provider.",
						},
						"cloud_connection_region": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service connection region.",
						},
						"cloud_connection_hosted_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service connection hosted type.",
						},
						"cloud_connection_region_description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service connection region description.",
						},
						"network_provider": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service network provider.",
						},
						"time_created": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cloud service time created.",
						},
						"enni_supported": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True if enni supported.",
						},
					},
					Description: "The list of list of physical locations optionally filtered by provided parameters.",
				},
			},
		},
	}
}

func dataSourceCloudLocationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cp, ok := d.GetOk("cloud_provider")
	if !ok {
		return diag.Errorf("please provide a valid cloud provider")
	}
	ccType, ok := d.GetOk("cloud_connection_type")
	if !ok {
		return diag.Errorf("please provide a valid cloud connection type")
	}
	locations, err := c.GetCloudLocations(cp.(string), ccType.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("cloud_locations", flattenCloudLocations(&locations))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(cp.(string))
	return diags
}

func flattenCloudLocations(locs *[]packetfabric.CloudLocation) []interface{} {
	flattens := make([]interface{}, len(*locs))
	for i, loc := range *locs {
		flatten := make(map[string]interface{})
		flatten["pop"] = loc.Pop
		flatten["region"] = loc.Region
		flatten["market"] = loc.Market
		flatten["market_description"] = loc.MarketDescription
		flatten["vendor"] = loc.Vendor
		flatten["site"] = loc.Site
		flatten["site_code"] = loc.SiteCode
		flatten["type"] = loc.Type
		flatten["status"] = loc.Status
		flatten["latitude"] = loc.Latitude
		flatten["longitude"] = loc.Longitude
		if loc.Timezone != nil {
			flatten["timezone"] = loc.Timezone
		}
		if loc.Notes != nil {
			flatten["notes"] = loc.Notes
		}
		if loc.Pcode != nil {
			flatten["pcode"] = loc.Pcode
		}
		flatten["lead_time"] = loc.LeadTime
		flatten["single_armed"] = loc.SingleArmed
		flatten["address1"] = loc.Address1
		if loc.Address2 != nil {
			flatten["address2"] = loc.Address2
		}
		flatten["zones"] = loc.Zones
		flatten["city"] = loc.City
		flatten["state"] = loc.State
		flatten["postal"] = loc.Postal
		flatten["country"] = loc.Country
		flatten["cloud_provider"] = loc.CloudProvider
		flatten["cloud_connection_region"] = loc.CloudConnectionDetails.Region
		flatten["cloud_connection_hosted_type"] = loc.CloudConnectionDetails.HostedType
		flatten["cloud_connection_region_description"] = loc.CloudConnectionDetails.RegionDescription
		flatten["network_provider"] = loc.NetworkProvider
		flatten["time_created"] = loc.TimeCreated
		flatten["enni_supported"] = loc.EnniSupported
		flattens[i] = flatten
	}
	return flattens
}
