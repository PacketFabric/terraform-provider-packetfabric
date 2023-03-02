package provider

import (
	"context"
	"fmt"
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
		DeleteContext: resourceThirdPartyVirtualCircuitDelete,
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
			"routing_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The routing ID of the marketplace provider that will be receiving this request.\n\n\tExample: TR-1RI-OQ85",
			},
			"market": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The market code (e.g. \"ATL\" or \"DAL\") in which you would like the marketplace provider to provision their side of the connection.\n\n\tIf the marketplace provider has services published in the marketplace, you can use the PacketFabric portal to see which POPs they are in. Simply remove the number from the POP to get the market code (e.g. if they offer services in \"DAL5\", enter \"DAL\" for the market).",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of this connection.",
			},
			"rate_limit_in": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The upper bound, in Mbps, to limit incoming data by.",
			},
			"rate_limit_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The upper bound, in Mbps, to limit outgoing data by.",
			},
			"bandwidth": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_uuid": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
							ValidateFunc: validation.IsUUID,
							Description: "The UUID for the billing account that should be billed. " +
								"Can also be set with the PF_ACCOUNT_ID environment variable.",
						},
						"speed": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice(speedOptions(), true),
							Description:  "The desired speed of the new connection. Only applicable if `longhaul_type` is \"dedicated\" or \"hourly\".\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\" \"20Gbps\" \"30Gbps\" \"40Gbps\" \"50Gbps\" \"60Gbps\" \"80Gbps\" \"100Gbps\"]",
						},
						"subscription_term": {
							Type:         schema.TypeInt,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.IntInSlice([]int{1, 12, 24, 36}),
							Description:  "The billing term in months. Only applicable if `longhaul_type` is \"dedicated.\"\n\n\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
						},
						"longhaul_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"dedicated", "usage", "hourly"}, true),
							Description:  "Dedicated (no limits or additional charges), usage-based (per transferred GB) or hourly billing. Not applicable for Metro Dedicated.\n\n\tEnum [\"dedicated\" \"usage\" \"hourly\"]",
						},
					},
				},
			},
			"interface": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_circuit_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The circuit ID for the port. This starts with \"PF-AP-\"",
						},
						"vlan": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "Valid VLAN range is from 4-4094, inclusive.",
						},
						"untagged": {
							Type:        schema.TypeBool,
							Required:    true,
							ForceNew:    true,
							Description: "Whether the interface should be untagged.",
						},
					},
				},
			},
			"service_uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "UUID of the marketplace service being requested.",
			},
			"flex_bandwidth_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ID of the flex bandwidth container from which to subtract this VC's speed.",
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
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Marketplace Request.",
		Detail:   fmt.Sprintf("Warning: the Marketplace connection request for %q has been either accepted or rejected.", d.Id()),
	})
	return diags
}

func resourceThirdPartyVirtualCircuitDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if _, err := c.DeleteHostedMktConnection(d.Id()); err != nil {
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
