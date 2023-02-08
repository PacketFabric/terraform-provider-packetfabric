package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceBackbone() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A brief description of this connection.",
		},
		"bandwidth": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"account_uuid": {
						Type:         schema.TypeString,
						Required:     true,
						DefaultFunc:  schema.EnvDefaultFunc("PF_ACCOUNT_ID", nil),
						ValidateFunc: validation.IsUUID,
						Description: "The UUID for the billing account that should be billed. " +
							"Can also be set with the PF_ACCOUNT_ID environment variable.",
					},
					"speed": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice(speedOptions(), true),
						Description:  "The desired speed of the new connection. Only applicable if `longhaul_type` is \"dedicated\" or \"hourly\".\n\n\tEnum: [\"50Mbps\" \"100Mbps\" \"200Mbps\" \"300Mbps\" \"400Mbps\" \"500Mbps\" \"1Gbps\" \"2Gbps\" \"5Gbps\" \"10Gbps\" \"20Gbps\" \"30Gbps\" \"40Gbps\" \"50Gbps\" \"60Gbps\" \"80Gbps\" \"100Gbps\"]",
					},
					"subscription_term": {
						Type:        schema.TypeInt,
						Required:    true,
						Description: "The billing term, in months, for this connection. Only applicable if `longhaul_type` is \"dedicated.\"\n\n\tEnum: [\"1\", \"12\", \"24\", \"36\"]",
					},
					"longhaul_type": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Dedicated (no limits or additional charges), usage-based (per transferred GB) or hourly billing.\n\n\tEnum [\"dedicated\" \"usage\" \"hourly\"]",
					},
				},
			},
		},
		"interface_a": {
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
						Optional:    true,
						Description: "Valid VLAN range is from 4-4094, inclusive.",
					},
					"svlan": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Valid sVLAN.",
					},
					"untagged": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Whether the interface should be untagged.",
					},
				},
			},
		},
		"interface_z": {
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
						Optional:    true,
						Description: "Valid VLAN range is from 4-4094, inclusive.",
					},
					"svlan": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Valid sVLAN.",
					},
					"untagged": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Whether the interface should be untagged.",
					},
				},
			},
		},
		"rate_limit_in": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The upper bound, in Mbps, by which to limit incoming data.",
		},
		"rate_limit_out": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The upper bound, in Mbps, by which to limit outgoing data.",
		},
		"epl": {
			Type:        schema.TypeBool,
			Required:    true,
			ForceNew:    true,
			Description: "If true, the circuit will be an EPL connection rather than an EVPL. Default is false.\n\n\tEPL is an Ethernet Private Line. Typical access ports can only support one EPL connection (meaning one virtual circuit for that port). ENNI ports can support multiple EPL connections.\n\n\tEVPL is an Ethernet Virtual Private Line. A port can support multiple EVPL connections, as bandwidth allows.\n\n\tFor more information on the difference between the two, see [Virtual Circuit Ethernet Features](https://docs.packetfabric.com/reference/specs/ethernet_features/).",
		},
		"flex_bandwidth_id": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "ID of the flex bandwidth container from which to subtract this VC's speed.",
		},
	}
}

func resourceBackboneCreate(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(packetfabric.Backbone) (*packetfabric.BackboneResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	backboneVC := extractBack(d)
	resp, err := fn(backboneVC)
	if err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			if ok := c.IsBackboneComplete(resp.VcCircuitID); ok {
				ticker.Stop()
				createOk <- true
			}
		}
	}()
	<-createOk
	if resp != nil {
		d.SetId(resp.VcCircuitID)
	}
	return diags
}

func resourceServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}, fn func() ([]packetfabric.DedicatedConnResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	dedicatedConns, err := fn()
	if err != nil {
		return diag.FromErr(err)
	}
	for _, conn := range dedicatedConns {
		if conn.CloudCircuitID == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Services Read",
				Detail:   cloudCidNotFoundDetailsMsg,
			})
		} else {
			_ = d.Set("cloud_circuit_id", conn.CloudCircuitID)
		}
	}
	return diags
}

// used for Backbone VC
func resourceServiceSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	settings := extractServiceSettings(d)
	backboneVC := extractBack(d)

	if _, err := c.UpdateServiceSettings(d.Id(), settings); err != nil {
		return diag.FromErr(err)
	}
	updateOk := make(chan bool)
	defer close(updateOk)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			if ok := c.IsBackboneComplete(d.Id()); ok {
				ticker.Stop()
				updateOk <- true
			}
		}
	}()
	<-updateOk
	if d.HasChange("bandwidth") {
		billing := packetfabric.BillingUpgrade{
			Speed:            backboneVC.Bandwidth.Speed,
			SubscriptionTerm: backboneVC.Bandwidth.SubscriptionTerm,
		}
		if _, err := c.ModifyBilling(d.Id(), billing); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceBackboneDelete(ctx context.Context, d *schema.ResourceData, m interface{}, fn func(string) (*packetfabric.BackboneDeleteResp, error)) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCircuitID, ok := d.GetOk("id"); ok {
		resp, err := fn(vcCircuitID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Backbone Delete result",
			Detail:   resp.Message,
		})
		return diags
	}
	return diag.Errorf("please provide a valid VC Circuit ID for deletion")
}

func extractBack(d *schema.ResourceData) packetfabric.Backbone {
	backboneVC := packetfabric.Backbone{
		Description: d.Get("description").(string),
		Epl:         d.Get("epl").(bool),
	}
	for _, interfA := range d.Get("interface_a").(*schema.Set).List() {
		backboneVC.Interfaces = append(backboneVC.Interfaces, extractBackboneInterface(interfA.(map[string]interface{})))
	}
	for _, interfZ := range d.Get("interface_z").(*schema.Set).List() {
		backboneVC.Interfaces = append(backboneVC.Interfaces, extractBackboneInterface(interfZ.(map[string]interface{})))
	}
	for _, bw := range d.Get("bandwidth").(*schema.Set).List() {
		backboneVC.Bandwidth = extractBandwidth(bw.(map[string]interface{}))
	}
	if rateLimitIn, ok := d.GetOk("rate_limit_in"); ok {
		backboneVC.RateLimitIn = rateLimitIn.(int)
	}
	if rateLimitOut, ok := d.GetOk("rate_limit_out"); ok {
		backboneVC.RateLimitOut = rateLimitOut.(int)
	}
	if flexBandID, ok := d.GetOk("flex_bandwidth_id"); ok {
		backboneVC.FlexBandwidthID = flexBandID.(string)
	}
	return backboneVC
}

func extractBandwidth(bw map[string]interface{}) packetfabric.Bandwidth {
	bandwidth := packetfabric.Bandwidth{}
	bandwidth.AccountUUID = bw["account_uuid"].(string)
	longhaulType := bw["longhaul_type"]
	if longhaulType != nil {
		bandwidth.LonghaulType = longhaulType.(string)
	}
	if subsTerm := bw["subscription_term"]; subsTerm != nil {
		bandwidth.SubscriptionTerm = subsTerm.(int)
	}
	if speed := bw["speed"]; speed != nil {
		bandwidth.Speed = speed.(string)
	}
	return bandwidth
}

func extractBackboneInterface(interf map[string]interface{}) packetfabric.BackBoneInterface {
	backboneInter := packetfabric.BackBoneInterface{}
	backboneInter.PortCircuitID = interf["port_circuit_id"].(string)
	if vlan := interf["vlan"]; vlan != nil {
		backboneInter.Vlan = vlan.(int)
	}
	if untagged := interf["untagged"]; untagged != nil {
		backboneInter.Untagged = untagged.(bool)
	}

	return backboneInter
}

func speedOptions() []string {
	return []string{
		"50Mbps", "100Mbps", "200Mbps", "300Mbps",
		"400Mbps", "500Mbps", "1Gbps", "2Gbps",
		"5Gbps", "10Gbps", "20Gbps", "30Gbps",
		"40Gbps", "50Gbps", "60Gbps", "80Gbps", "100Gbps"}
}
