package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceHighPerformanceInternet() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceHighPerformanceInternetRead,
		Schema: map[string]*schema.Schema{
			"circuit_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"routing_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_routing_v4": datasourceHighPerformanceInternetStatic(),
						"static_routing_v6": datasourceHighPerformanceInternetStatic(),
						"bgp_v4":            datasourceHighPerformanceInternetBGP(),
						"bgp_v6":            datasourceHighPerformanceInternetBGP(),
					},
				},
			},
		},
	}
}

func datasourceHighPerformanceInternetStatic() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		ForceNew: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"l3_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"remote_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"address_family": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"prefixes": {
					Type:     schema.TypeSet,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"prefix": {
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func datasourceHighPerformanceInternetBGP() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		ForceNew: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"asn": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"l3_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"remote_address": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"md5": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"bgp_state": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"address_family": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"prefixes": {
					Type:     schema.TypeSet,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"prefix": {
								Type:     schema.TypeString,
								Computed: true,
							},
							"local_preference": {
								Type:     schema.TypeInt,
								Optional: true,
							},
						},
					},
				},
			},
		},
	}
}

func datasourceHighPerformanceInternetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	cID, ok := d.GetOk("circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid circuit_id")
	}
	highPerformanceInternet, err := c.ReadHighPerformanceInternetWithRoutes(cID.(string))
	if err != nil || highPerformanceInternet == nil {
		return diag.FromErr(err)
	}
	_ = d.Set("circuit_id", highPerformanceInternet.CircuitId)
	_ = d.Set("port_circuit_id", highPerformanceInternet.PortCircuitId)
	_ = d.Set("speed", highPerformanceInternet.Speed)
	_ = d.Set("vlan", highPerformanceInternet.Vlan)
	_ = d.Set("description", highPerformanceInternet.Description)
	_ = d.Set("market", highPerformanceInternet.Market)
	_ = d.Set("routing_type", highPerformanceInternet.RoutingType)
	_ = d.Set("state", highPerformanceInternet.State)

	routingConfiguration := flattenRoutingConfiguration(&highPerformanceInternet.RoutingConfiguration)
	if err := d.Set("routing_configuration", routingConfiguration); err != nil {
		return diag.Errorf("error setting 'routing_configuration': %s", err)
	}
	d.SetId(uuid.New().String())
	return diags
}
