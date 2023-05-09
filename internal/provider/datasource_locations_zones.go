package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceLocationsZones() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLocationsZonesRead,
		Schema: map[string]*schema.Schema{
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The location pop.",
			},
			"locations_zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
				Description: "The list of zones in the given POP.",
			},
		},
	}
}

func dataSourceLocationsZonesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	pop, ok := d.GetOk("pop")
	if !ok {
		return diag.Errorf("please provide a valid pop")
	}
	var diags diag.Diagnostics
	if zones, err := c.GetLocationsZones(pop.(string)); err != nil {
		return diag.FromErr(err)
	} else {
		_ = d.Set("locations_zones", zones)
	}
	d.SetId(uuid.New().String())
	return diags
}
