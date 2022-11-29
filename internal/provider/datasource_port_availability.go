package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourcePortAvailability() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortAvailabilityRead,
		Schema: map[string]*schema.Schema{
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The port POP.",
			},
			"ports_available": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The pop zone.",
						},
						"speed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The desired speed of the new connection.\n\t\tEnum: []\"1gps\", \"10gbps\"]",
						},
						"media": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The port media type.",
						},
						"count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The port count.",
						},
						"partial": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True if port is partial.",
						},
						"enni": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True if port is enni.",
						},
					},
				},
				Description: "The list of ports available in the given POP.",
			},
		},
	}
}

func dataSourcePortAvailabilityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	pop, ok := d.GetOk("pop")
	if !ok {
		return diag.Errorf("please provide a valid pop")
	}
	ports, err := c.GetLocationPortAvailability(pop.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ports_available", flattenPortsAvailable(&ports))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenPortsAvailable(ports *[]packetfabric.PortAvailability) []interface{} {
	flattens := make([]interface{}, len(*ports))
	for i, port := range *ports {
		flatten := make(map[string]interface{})
		flatten["zone"] = port.Zone
		flatten["speed"] = port.Speed
		flatten["media"] = port.Media
		flatten["count"] = port.Count
		flatten["partial"] = port.Partial
		flatten["enni"] = port.Enni
		flattens[i] = flatten
	}
	return flattens
}
