package provider

import (
	"context"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceBackbone() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		CreateContext: resourceBackboneCreate,
		UpdateContext: resourceBackboneUpdate,
		ReadContext:   resourceBackboneRead,
		DeleteContext: resourceBackboneDelete,
		Schema: map[string]*schema.Schema{
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
							Optional:     true,
							ValidateFunc: validation.StringInSlice(speedOptions(), true),
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
							Description:  "Dedicated (no limits or additional charges), usage-based (per transferred GB) or hourly billing. Not applicable for Metro Dedicated.\n\n\tEnum [\"dedicated\" \"usage\" \"hourly\"]",
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
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntBetween(4, 4094),
							Description:  "Valid VLAN range is from 4-4094, inclusive. ",
						},
						"svlan": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntBetween(4, 4094),
							Description:  "Valid sVLAN range is from 4-4094, inclusive. ",
						},
						"untagged": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the interface should be untagged. ",
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
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntBetween(4, 4094),
							Description:  "Valid VLAN range is from 4-4094, inclusive. ",
						},
						"svlan": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      0,
							ValidateFunc: validation.IntBetween(4, 4094),
							Description:  "Valid sVLAN range is from 4-4094, inclusive. ",
						},
						"untagged": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the interface should be untagged. ",
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
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "If true, the circuit will be an EPL connection rather than an EVPL. Default is false.\n\n\tEPL is an Ethernet Private Line. Typical access ports can only support one EPL connection (meaning one virtual circuit for that port). ENNI ports can support multiple EPL connections.\n\n\tEVPL is an Ethernet Virtual Private Line. A port can support multiple EVPL connections, as bandwidth allows.\n\n\tFor more information on the difference between the two, see [Virtual Circuit Ethernet Features](https://docs.packetfabric.com/reference/specs/ethernet_features/).\n\n\t",
			},
			"flex_bandwidth_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ID of the flex bandwidth container from which to subtract this VC's speed.",
			},
			"po_number": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
				Description:  "Purchase order number or identifier of a service.",
			},
			"labels": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Label value linked to an object.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceBackboneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	backboneVC := extractBack(d)
	resp, err := c.CreateBackbone(backboneVC)
	if err != nil {
		return diag.FromErr(err)
	}
	createOk := make(chan bool)
	defer close(createOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
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

		if labels, ok := d.GetOk("labels"); ok {
			diagnostics, created := createLabels(c, d.Id(), labels)
			if !created {
				return diagnostics
			}
		}
	}
	return diags
}

func resourceBackboneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	resp, err := c.GetBackboneByVcCID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if resp != nil {
		_ = d.Set("description", resp.Description)
		if resp.Mode == "epl" {
			_ = d.Set("epl", true)
		} else {
			_ = d.Set("epl", false)
		}
		// Create a new schema set for the bandwidth attribute
		bandwidthSet := schema.NewSet(
			func(i interface{}) int { return 0 },
			[]interface{}{},
		)
		// Add the bandwidth values to the set
		if resp.Bandwidth.LonghaulType == "dedicated" {
			bandwidth := map[string]interface{}{
				"account_uuid":      resp.Bandwidth.AccountUUID,
				"longhaul_type":     resp.Bandwidth.LonghaulType,
				"subscription_term": resp.Bandwidth.SubscriptionTerm,
				"speed":             resp.Bandwidth.Speed,
			}
			bandwidthSet.Add(bandwidth)
		}
		// metro dedicated doesn't need longhaul_type
		if resp.Bandwidth.LonghaulType == "" {
			bandwidth := map[string]interface{}{
				"account_uuid":      resp.Bandwidth.AccountUUID,
				"subscription_term": resp.Bandwidth.SubscriptionTerm,
				"speed":             resp.Bandwidth.Speed,
			}
			bandwidthSet.Add(bandwidth)
		}
		if resp.Bandwidth.LonghaulType == "usage" {
			bandwidth := map[string]interface{}{
				"account_uuid":  resp.Bandwidth.AccountUUID,
				"longhaul_type": resp.Bandwidth.LonghaulType,
			}
			bandwidthSet.Add(bandwidth)
		}
		if resp.Bandwidth.LonghaulType == "hourly" {
			bandwidth := map[string]interface{}{
				"account_uuid":  resp.Bandwidth.AccountUUID,
				"longhaul_type": resp.Bandwidth.LonghaulType,
				"speed":         resp.Bandwidth.Speed,
			}
			bandwidthSet.Add(bandwidth)
		}
		// Set the bandwidth attribute to the schema set
		_ = d.Set("bandwidth", bandwidthSet)

		if len(resp.Interfaces) == 2 {
			interfaceA := make(map[string]interface{})
			interfaceA["port_circuit_id"] = resp.Interfaces[0].PortCircuitID
			interfaceA["vlan"] = resp.Interfaces[0].Vlan
			interfaceA["svlan"] = resp.Interfaces[0].Svlan
			interfaceA["untagged"] = resp.Interfaces[0].Untagged
			_ = d.Set("interface_a", []interface{}{interfaceA})

			interfaceZ := make(map[string]interface{})
			interfaceZ["port_circuit_id"] = resp.Interfaces[1].PortCircuitID
			interfaceZ["vlan"] = resp.Interfaces[1].Vlan
			interfaceZ["svlan"] = resp.Interfaces[1].Svlan
			interfaceZ["untagged"] = resp.Interfaces[1].Untagged
			_ = d.Set("interface_z", []interface{}{interfaceZ})
		}
		if _, ok := d.GetOk("rate_limit_in"); ok {
			_ = d.Set("rate_limit_in", resp.RateLimitIn)
		}
		if _, ok := d.GetOk("rate_limit_out"); ok {
			_ = d.Set("rate_limit_out", resp.RateLimitOut)
		}
		if _, ok := d.GetOk("flex_bandwidth_id"); ok {
			_ = d.Set("flex_bandwidth_id", resp.AggregateCapacityID)
		} else {
			_ = d.Set("flex_bandwidth_id", nil)
		}
		_ = d.Set("po_number", resp.PONumber)
	}

	if _, ok := d.GetOk("labels"); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set("labels", labels)
	}
	return diags
}

// used for Backbone VC
func resourceBackboneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	settings := extractServiceSettings(d)
	backboneVC := extractBack(d)

	if d.HasChange("bandwidth") {
		billing := packetfabric.BillingUpgrade{
			Speed:            backboneVC.Bandwidth.Speed,
			SubscriptionTerm: backboneVC.Bandwidth.SubscriptionTerm,
		}
		if _, err := c.ModifyBilling(d.Id(), billing); err != nil {
			return diag.FromErr(err)
		}
		updateOk := make(chan bool)
		defer close(updateOk)
		ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
		go func() {
			for range ticker.C {
				if ok := c.IsBackboneComplete(d.Id()); ok {
					ticker.Stop()
					updateOk <- true
				}
			}
		}()
		<-updateOk
	}

	if _, err := c.UpdateServiceSettings(d.Id(), settings); err != nil {
		return diag.FromErr(err)
	}
	updateOk := make(chan bool)
	defer close(updateOk)
	ticker := time.NewTicker(time.Duration(30+c.GetRandomSeconds()) * time.Second)
	go func() {
		for range ticker.C {
			if ok := c.IsBackboneComplete(d.Id()); ok {
				ticker.Stop()
				updateOk <- true
			}
		}
	}()
	<-updateOk

	if d.HasChange("labels") {
		labels := d.Get("labels")
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}
	return diags
}

func resourceBackboneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	if vcCircuitID, ok := d.GetOk("id"); ok {
		_, err := c.DeleteBackbone(vcCircuitID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId("")
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
	if poNumber, ok := d.GetOk("po_number"); ok {
		backboneVC.PONumber = poNumber.(string)
	}
	return backboneVC
}

func extractServiceSettings(d *schema.ResourceData) packetfabric.ServiceSettingsUpdate {
	settUpdate := packetfabric.ServiceSettingsUpdate{}

	if rateLimitIn, ok := d.GetOk("rate_limit_in"); ok {
		settUpdate.RateLimitIn = rateLimitIn.(int)
	}
	if rateLimitOut, ok := d.GetOk("rate_limit_out"); ok {
		settUpdate.RateLimitOut = rateLimitOut.(int)
	}
	if description, ok := d.GetOk("description"); ok {
		settUpdate.Description = description.(string)
	}
	if _, ok := d.GetOk("interface"); ok {
		for _, interf := range d.Get("interface").(*schema.Set).List() {
			settUpdate.Interfaces = append(settUpdate.Interfaces, extractBackboneInterface(interf.(map[string]interface{})))
		}
	}
	if _, ok := d.GetOk("interface_a"); ok {
		for _, interf := range d.Get("interface_a").(*schema.Set).List() {
			settUpdate.Interfaces = append(settUpdate.Interfaces, extractBackboneInterface(interf.(map[string]interface{})))
		}
	}
	if _, ok := d.GetOk("interface_z"); ok {
		for _, interf := range d.Get("interface_z").(*schema.Set).List() {
			// Only include interface_z if it was modified
			if d.HasChange("interface_z") {
				settUpdate.Interfaces = append(settUpdate.Interfaces, extractBackboneInterface(interf.(map[string]interface{})))
			}
		}
	}
	if poNumber, ok := d.GetOk("po_number"); ok {
		settUpdate.PONumber = poNumber.(string)
	}

	return settUpdate
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

func extractBackboneInterface(interf map[string]interface{}) packetfabric.Interfaces {
	backboneInter := packetfabric.Interfaces{}
	if portCID := interf["port_circuit_id"]; portCID != nil {
		backboneInter.PortCircuitID = portCID.(string)
	}
	if vlan := interf["vlan"]; vlan != nil {
		backboneInter.Vlan = vlan.(int)
	}
	if untagged := interf["untagged"]; untagged != nil {
		backboneInter.Untagged = untagged.(bool)
	}
	if svlan := interf["svlan"]; svlan != nil {
		backboneInter.Svlan = svlan.(int)
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
