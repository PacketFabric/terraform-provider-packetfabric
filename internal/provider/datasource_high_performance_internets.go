package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceHighPerformanceInternets() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceHighPerformanceInternetsRead,
		Schema: map[string]*schema.Schema{
			"high_performance_internets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"circuit_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"account_uuid": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The UUID for the billing account that should be billed. " +
								"Can also be set with the PF_ACCOUNT_ID environment variable.",
						},
						"port_circuit_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"speed": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"market": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"routing_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceHighPerformanceInternetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	highPerformanceInternets, err := c.ReadHighPerformanceInternets()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("high_performance_internets", flattenHighPerformanceInternets(&highPerformanceInternets))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenHighPerformanceInternets(highPerformanceInternets *[]packetfabric.HighPerformanceInternet) []interface{} {
	flattens := make([]interface{}, len(*highPerformanceInternets))
	for i, highPerformanceInternet := range *highPerformanceInternets {
		flatten := make(map[string]interface{})
		flattens[i] = flatten
		flatten["circuit_id"] = highPerformanceInternet.CircuitId
		flatten["port_circuit_id"] = highPerformanceInternet.PortCircuitId
		flatten["speed"] = highPerformanceInternet.Speed
		flatten["vlan"] = highPerformanceInternet.Vlan
		flatten["description"] = highPerformanceInternet.Description
		flatten["market"] = highPerformanceInternet.Market
		flatten["routing_type"] = highPerformanceInternet.RoutingType
		flatten["state"] = highPerformanceInternet.State
	}
	return flattens
}
