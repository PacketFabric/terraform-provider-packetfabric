package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceHighPerformanceInternet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHighPerformanceInternetCreate,
		ReadContext:   resourceHighPerformanceInternetRead,
		UpdateContext: resourceHighPerformanceInternetUpdate,
		DeleteContext: resourceHighPerformanceInternetDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"circuit_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_uuid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
				ValidateFunc: validation.IsUUID,
				Description: "The UUID for the billing account that should be billed. " +
					"Can also be set with the PF_ACCOUNT_ID environment variable.",
			},
			"port_circuit_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"speed": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vlan": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
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
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_routing_v4": resourceHighPerformanceInternetStatic(),
						"static_routing_v6": resourceHighPerformanceInternetStatic(),
						"bgp_v4":            resourceHighPerformanceInternetBGP(),
						"bgp_v6":            resourceHighPerformanceInternetBGP(),
					},
				},
			},
		},
	}
}

func resourceHighPerformanceInternetStatic() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		ForceNew: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"l3_address": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"remote_address": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"address_family": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"prefixes": {
					Type:     schema.TypeSet,
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"prefix": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validateIPAddressWithPrefix,
							},
						},
					},
				},
			},
		},
	}
}
func resourceHighPerformanceInternetBGP() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		ForceNew: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"asn": {
					Type:     schema.TypeInt,
					Required: true,
				},
				"l3_address": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"remote_address": {
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringIsNotEmpty,
				},
				"md5": {
					Type:         schema.TypeString,
					Optional:     true,
					ValidateFunc: validation.StringIsNotEmpty,
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
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"prefix": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validateIPAddressWithPrefix,
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

func resourceHighPerformanceInternetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	highPerformanceInternet := extractHighPerformanceInternet(d)
	resp, err := c.CreateHighPerformanceInternet(&highPerformanceInternet)
	if err != nil || resp == nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.CircuitId)
	return resourceHighPerformanceInternetRead(ctx, d, m)
}

func resourceHighPerformanceInternetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	highPerformanceInternet, err := c.ReadHighPerformanceInternetWithRoutes(d.Id())
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
	return diags
}

func resourceHighPerformanceInternetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	highPerformanceInternet := extractHighPerformanceInternet(d)
	_, err := c.UpdateHighPerformanceInternet(&highPerformanceInternet)
	if err != nil {
		return diag.FromErr(err)
	}
	//d.SetId(resp.CircuitId)
	return diags
}

func resourceHighPerformanceInternetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	response, err := c.DeleteHighPerformanceInternet(d.Id())
	if err != nil || response == nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func extractHighPerformanceInternet(d *schema.ResourceData) packetfabric.HighPerformanceInternet {
	highPerformanceInternet := packetfabric.HighPerformanceInternet{}
	if circuit_id, ok := d.GetOk("circuit_id"); ok {
		highPerformanceInternet.CircuitId = circuit_id.(string)
	}
	if port_circuit_id, ok := d.GetOk("port_circuit_id"); ok {
		highPerformanceInternet.PortCircuitId = port_circuit_id.(string)
	}
	if speed, ok := d.GetOk("speed"); ok {
		highPerformanceInternet.Speed = speed.(string)
	}
	if vlan, ok := d.GetOk("vlan"); ok {
		highPerformanceInternet.Vlan = vlan.(int)
	}
	if description, ok := d.GetOk("description"); ok {
		highPerformanceInternet.Description = description.(string)
	}
	if market, ok := d.GetOk("market"); ok {
		highPerformanceInternet.Market = market.(string)
	}
	if routing_type, ok := d.GetOk("routing_type"); ok {
		highPerformanceInternet.RoutingType = routing_type.(string)
	}
	if state, ok := d.GetOk("state"); ok {
		highPerformanceInternet.State = state.(string)
	}
	if account_uuid, ok := d.GetOk("account_uuid"); ok {
		highPerformanceInternet.AccountUUID = account_uuid.(string)
	}
	if routingConfiguration, ok := d.GetOk("routing_configuration"); ok {
		highPerformanceInternet.RoutingConfiguration = *extractRoutingConfiguration(routingConfiguration.(*schema.Set))
	}
	return highPerformanceInternet
}

func extractRoutingConfiguration(routingConfiguration *schema.Set) *packetfabric.HighPerformanceInternetRoutingConfiguration {
	routingConfigurationList := routingConfiguration.List()
	if 0 == len(routingConfigurationList) {
		return nil
	}
	routingConfigurationMap := routingConfigurationList[0].(map[string]interface{})
	return &packetfabric.HighPerformanceInternetRoutingConfiguration{
		StaticRoutingV4: extractStaticRoutingConfiguration(routingConfigurationMap["static_routing_v4"].(*schema.Set)),
		StaticRoutingV6: extractStaticRoutingConfiguration(routingConfigurationMap["static_routing_v6"].(*schema.Set)),
		BgpV4:           extractBgpRoutingConfiguration(routingConfigurationMap["bgp_v4"].(*schema.Set)),
		BgpV6:           extractBgpRoutingConfiguration(routingConfigurationMap["bgp_v6"].(*schema.Set)),
	}
}

func extractStaticRoutingConfiguration(staticRoutingConfiguration *schema.Set) *packetfabric.HighPerformanceInternetStaticConfiguration {
	staticRoutingConfigurationList := staticRoutingConfiguration.List()
	if 0 == len(staticRoutingConfigurationList) {
		return nil
	}
	staticRoutingConfigurationMap := staticRoutingConfigurationList[0].(map[string]interface{})
	return &packetfabric.HighPerformanceInternetStaticConfiguration{
		L3Address:     staticRoutingConfigurationMap["l3_address"].(string),
		RemoteAddress: staticRoutingConfigurationMap["remote_address"].(string),
		AddressFamily: staticRoutingConfigurationMap["address_family"].(string),
		Prefixes:      extractStaticRoutingPrefixes(staticRoutingConfigurationMap["prefixes"].(*schema.Set)),
	}
}

func extractStaticRoutingPrefixes(prefixes *schema.Set) []packetfabric.HighPerformanceInternetStaticRoute {
	staticRoutingPrefixes := make([]packetfabric.HighPerformanceInternetStaticRoute, 0)
	for _, prefix := range prefixes.List() {
		prefixMap := prefix.(map[string]interface{})
		staticRoutingPrefixes = append(staticRoutingPrefixes, packetfabric.HighPerformanceInternetStaticRoute{
			Prefix: prefixMap["prefix"].(string),
		})
	}
	return staticRoutingPrefixes
}

/////////////////////////////////

func extractBgpRoutingConfiguration(bgpRoutingConfiguration *schema.Set) *packetfabric.HighPerformanceInternetBgpConfiguration {
	bgpRoutingConfigurationMap := bgpRoutingConfiguration.List()[0].(map[string]interface{})
	return &packetfabric.HighPerformanceInternetBgpConfiguration{
		Asn:           bgpRoutingConfigurationMap["asn"].(int),
		L3Address:     bgpRoutingConfigurationMap["l3_address"].(string),
		RemoteAddress: bgpRoutingConfigurationMap["remote_address"].(string),
		Md5:           bgpRoutingConfigurationMap["md5"].(string),
		BgpState:      bgpRoutingConfigurationMap["bgp_state"].(string),
		AddressFamily: bgpRoutingConfigurationMap["address_family"].(string),
		Prefixes:      extractBgpRoutingPrefixes(bgpRoutingConfigurationMap["prefixes"].(*schema.Set)),
	}
}

func extractBgpRoutingPrefixes(prefixes *schema.Set) []packetfabric.HighPerformanceInternetBgpPrefix {
	bgpRoutingPrefixes := make([]packetfabric.HighPerformanceInternetBgpPrefix, 0)
	for _, prefix := range prefixes.List() {
		prefixMap := prefix.(map[string]interface{})
		bgpRoutingPrefixes = append(bgpRoutingPrefixes, packetfabric.HighPerformanceInternetBgpPrefix{
			Prefix:          prefixMap["prefix"].(string),
			LocalPreference: prefixMap["local_preference"].(int),
		})
	}
	return bgpRoutingPrefixes
}

func flattenRoutingConfiguration(routingConfiguration *packetfabric.HighPerformanceInternetRoutingConfiguration) []interface{} {
	result := make([]interface{}, 0)
	data := make(map[string]interface{})
	if nil != routingConfiguration.StaticRoutingV4 {
		data["static_routing_v4"] = flattenStaticRoutingConfiguration(routingConfiguration.StaticRoutingV4)
	}
	if nil != routingConfiguration.StaticRoutingV6 {
		data["static_routing_v6"] = flattenStaticRoutingConfiguration(routingConfiguration.StaticRoutingV6)
	}
	if nil != routingConfiguration.BgpV4 {
		data["bgp_v4"] = flattenBGPRoutingConfiguration(routingConfiguration.BgpV4)
	}
	if nil != routingConfiguration.BgpV6 {
		data["bgp_v6"] = flattenBGPRoutingConfiguration(routingConfiguration.BgpV6)
	}
	result = append(result, data)
	return result
}

func flattenStaticRoutingConfiguration(staticRouting *packetfabric.HighPerformanceInternetStaticConfiguration) []interface{} {
	result := make([]interface{}, 0)
	data := make(map[string]interface{})
	result = append(result, data)
	data["l3_address"] = staticRouting.L3Address
	data["remote_address"] = staticRouting.RemoteAddress
	data["address_family"] = staticRouting.AddressFamily
	prefixList := make([]interface{}, 0)
	data["prefixes"] = prefixList
	for _, staticRoute := range staticRouting.Prefixes {
		prefixData := make(map[string]interface{})
		prefixData["prefix"] = staticRoute.Prefix
		prefixList = append(prefixList, prefixData)
	}
	return result
}

func flattenBGPRoutingConfiguration(bgpRouting *packetfabric.HighPerformanceInternetBgpConfiguration) []interface{} {
	result := make([]interface{}, 0)
	data := make(map[string]interface{})
	result = append(result, data)
	data["asn"] = bgpRouting.Asn
	data["l3_address"] = bgpRouting.L3Address
	data["remote_address"] = bgpRouting.RemoteAddress
	data["md5"] = bgpRouting.Md5
	data["bgp_state"] = bgpRouting.BgpState
	data["address_family"] = bgpRouting.AddressFamily
	prefixList := make([]interface{}, 0)
	data["prefixes"] = prefixList
	for _, bgpRoute := range bgpRouting.Prefixes {
		prefixData := make(map[string]interface{})
		prefixData["prefix"] = bgpRoute.Prefix
		prefixData["local_preference"] = bgpRoute.LocalPreference
		prefixList = append(prefixList, prefixData)
	}
	return result
}
