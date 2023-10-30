package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePortAvailability() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortAvailabilityRead,
		Schema: map[string]*schema.Schema{
			PfPop: schemaStringRequiredNotEmpty(PfPopDescriptionB),
			PfPortsAvailable: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: PfPortAvailabilityDescription,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfZone:    schemaStringComputed(PfZoneDescription2),
						PfSpeed:   schemaStringComputed(PfSpeedDescriptionJ),
						PfMedia:   schemaStringComputed(PfMediaDescription4),
						PfCount:   schemaIntComputed(PfCountDescription),
						PfPartial: schemaBoolComputed(PfPartialDescription),
						PfEnni:    schemaBoolComputed(PfEnniDescription),
					},
				},
			},
		},
	}
}

func dataSourcePortAvailabilityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	pop, ok := d.GetOk(PfPop)
	if !ok {
		return diag.Errorf(MessageMissingPop)
	}
	ports, err := c.GetLocationPortAvailability(pop.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfPortsAvailable, flattenPortsAvailable(&ports))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenPortsAvailable(ports *[]packetfabric.PortAvailability) []interface{} {
	flattens := make([]interface{}, len(*ports))
	fields := stringsToMap(PfZone, PfSpeed, PfMedia, PfCount, PfPartial, PfEnni)
	for i, port := range *ports {
		flattens[i] = structToMap(&port, fields)
	}
	return flattens
}
