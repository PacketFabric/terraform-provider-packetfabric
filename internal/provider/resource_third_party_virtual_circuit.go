package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceThirdPartyVirtualCircuit() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceThirdPartyVirtualCircuitCreate,
		ReadContext:   resourceThirdPartyVirtualCircuitRead,
		UpdateContext: resourceThirdPartyVirtualCircuitUpdate,
		DeleteContext: resourceThirdPartyVirtualCircuitDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"routing_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The routing ID of the customer to whom this VC will be connected.",
			},
			"market": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The market that the VC will be requested in.",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The description of the third-party VC.",
			},
			"rate_limit_in": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The upper bound, in Mbps, to limit incoming data by.",
			},
			"rate_limit_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The upper bound, in Mbps, to limit outgoing data by.",
			},
			"bandwidth": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_uuid": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The UUID for the billing account that should be billed.",
						},
						"speed": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(ixVcSpeedOptions(), true),
							Description:  "The desired speed of the new connection. Only applicable if `longhaul_type` is \"dedicated\" or \"hourly\".\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\" \"20Gbps\" \"30Gbps\" \"40Gbps\" \"50Gbps\" \"60Gbps\" \"80Gbps\" \"100Gbps\"]",
						},
						"subscription_term": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
							Description:  "The billing term, in months, for this connection. Only applicable if `longhaul_type` is \"dedicated.\"\n\n\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
						},
						"longhaul_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"dedicated", "usage", "hourly"}, true),
							Description:  "Dedicated (no limits or additional charges), usage-based (per transferred GB) or hourly billing.\n\n\tEnum [\"dedicated\" \"usage\" \"hourly\"]",
						},
					},
				},
			},
			"interface": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_circuit_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The circuit ID for the port. This starts with \"PF-AP-\"",
						},
						"vlan": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Valid VLAN range is from 4-4094, inclusive.",
						},
						"untagged": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether the interface should be untagged.",
						},
					},
				},
			},
			"service_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "UUID of the marketplace service being requested.",
			},
			"flex_bandwidth_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The circuit ID of the flex bandwidth container.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceThirdPartyVirtualCircuitCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	thidPartyVC := extractThirdPartyVC(d)
	resp, err := c.CreateThirdPartyVC(thidPartyVC)
	time.Sleep(30 * time.Second)
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		d.SetId(resp.VcRequestUUID)
	}
	return diags
}

func resourceThirdPartyVirtualCircuitRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if _, err := c.GetVCRequest(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceThirdPartyVirtualCircuitUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceUpdateMarketplace(ctx, d, m)
}

func resourceThirdPartyVirtualCircuitDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if _, err := c.DeleteVCRequest(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func extractThirdPartyVC(d *schema.ResourceData) packetfabric.ThirdPartyVC {
	thidPartyVC := packetfabric.ThirdPartyVC{}
	if routingID, ok := d.GetOk("routing_id"); ok {
		thidPartyVC.RoutingID = routingID.(string)
	}
	if market, ok := d.GetOk("market"); ok {
		thidPartyVC.Market = market.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		thidPartyVC.Description = description.(string)
	}
	if rateLimitIn, ok := d.GetOk("rate_limit_in"); ok {
		thidPartyVC.RateLimitIn = rateLimitIn.(int)
	}
	if rateLimitOut, ok := d.GetOk("rate_limit_out"); ok {
		thidPartyVC.RateLimitOut = rateLimitOut.(int)
	}
	for _, bw := range d.Get("bandwidth").(*schema.Set).List() {
		thidPartyVC.Bandwidth = extractBandwidth(bw.(map[string]interface{}))
	}
	for _, interf := range d.Get("interface").(*schema.Set).List() {
		thidPartyVC.Interface = extractThirdPartyInterf(interf.(map[string]interface{}))
	}
	if serviceUUID, ok := d.GetOk("service_uuid"); ok {
		thidPartyVC.ServiceUUID = serviceUUID.(string)
	}
	if flexBandwith, ok := d.GetOk("flex_bandwidth_id"); ok {
		thidPartyVC.FlexBandwidthID = flexBandwith.(string)
	}
	return thidPartyVC
}

func extractThirdPartyInterf(interf map[string]interface{}) packetfabric.Interface {
	interfResp := packetfabric.Interface{}
	if portCID := interf["port_circuit_id"]; portCID != nil {
		interfResp.PortCircuitID = portCID.(string)
	}
	if vlan := interf["vlan"]; vlan != nil {
		interfResp.Vlan = vlan.(int)
	}
	if svlan := interf["svlan"]; svlan != nil {
		interfResp.Svlan = svlan.(int)
	}
	if untagged := interf["untagged"]; untagged != nil {
		interfResp.Untagged = untagged.(bool)
	}
	return interfResp
}
