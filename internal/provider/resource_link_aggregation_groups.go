package provider

import (
	"context"
	"errors"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLinkAggregationGroups() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLinkAggregationGroupsCreate,
		ReadContext:   resourceLinkAggregationGroupsRead,
		UpdateContext: resourceLinkAggregationGroupsUpdate,
		DeleteContext: resourceLinkAggregationGroupsDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "A brief description of the LAG.",
			},
			"interval": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(intervalOptions(), false),
				Description:  "The LACP interval determines the frequency in which LACP control packets (LACP PDUs) are sent. If you specify fast, they are sent at 1 second intervals. If you specify slow, they are sent at 30 second intervals.\n\n\tEnum: \"fast\" \"slow\"",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Point of presence in which the LAG should be located.",
			},
			"members": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "A list of port circuit IDs to include in the LAG. To be included in a LAG, the ports must be at the same site, in the same zone, and have the same speed and media.",
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "The member circuit ID.",
				},
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
			"enabled": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				Description: "Change LAG Admin Status. Set it to true when LAG is enabled, false when LAG is disabled. ",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceLinkAggregationGroupsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	lag := extractLAG(d)
	resp, err := c.CreateLinkAggregationGroup(lag)
	// Adding this workaround due to a system delay.
	time.Sleep(45 * time.Second)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resp.PortCircuitID)

	if labels, ok := d.GetOk("labels"); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}

	if len(lag.Members) > 0 {
		for _, member := range lag.Members {
			_, err := c.CreateLagMember(d.Id(), member)
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}

	enabled := d.Get("enabled").(bool)
	if !enabled {
		_, err := c.DisableLinkAggregationGroup(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}

	diagnostics, updated := updatePort(c, d)
	if !updated {
		return diagnostics
	}
	return diags
}

func resourceLinkAggregationGroupsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	_, err := c.UpdateLinkAggregationGroup(d.Id(), d.Get("description").(string), d.Get("interval").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("po_number") {
		diagnostics, updated := updatePort(c, d)
		if !updated {
			return diagnostics
		}
	}

	if d.HasChange("labels") {
		labels := d.Get("labels")
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}

	if d.HasChange("enabled") {
		enabled := d.Get("enabled").(bool)
		if !enabled {
			_, err := c.DisableLinkAggregationGroup(d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
		} else {
			_, err := c.EnableLinkAggregationGroup(d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return diags
}

func updatePort(c *packetfabric.PFClient, d *schema.ResourceData) (diag.Diagnostics, bool) {
	lag, err := c.GetPortByCID(d.Id())
	if err != nil {
		return diag.FromErr(err), false
	}

	poNumber, ok := d.GetOk("po_number")
	if !ok {
		return diag.FromErr(errors.New("please enter a purchase order number")), true
	}

	portUpdateData := packetfabric.PortUpdate{
		Description: lag.Description,
		PONumber:    poNumber.(string),
	}
	_, err2 := c.UpdatePort(d.Id(), portUpdateData)
	if err2 != nil {
		return diag.FromErr(err), false
	}
	return diag.Diagnostics{}, true
}

func resourceLinkAggregationGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	lag, err := c.GetPortByCID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	_ = d.Set("description", lag.Description)
	_ = d.Set("pop", lag.Pop)
	_ = d.Set("interval", lag.LagInterval)
	_ = d.Set("po_number", lag.PONumber)

	interfaces, err := c.GetLAGInterfaces(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	members := make([]string, len(*interfaces), len(*interfaces))
	for index, interf := range *interfaces {
		members[index] = interf.PortCircuitID
	}
	_ = d.Set("members", members)

	if _, ok := d.GetOk("labels"); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set("labels", labels)
	}
	return diags
}

func resourceLinkAggregationGroupsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	time.Sleep(45 * time.Second)
	resp, err := c.DeleteLinkAggregationGroup(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	time.Sleep(45 * time.Second)
	return diags
}

func extractLAG(d *schema.ResourceData) packetfabric.LinkAggregationGroup {
	lag := packetfabric.LinkAggregationGroup{}
	if description, ok := d.GetOk("description"); ok {
		lag.Description = description.(string)
	}
	if interval, ok := d.GetOk("interval"); ok {
		lag.Interval = interval.(string)
	}
	if pop, ok := d.GetOk("pop"); ok {
		lag.Pop = pop.(string)
	}
	lag.Members = extractMembers(d)
	return lag
}

func intervalOptions() []string {
	return []string{"fast", "slow"}
}

func extractMembers(d *schema.ResourceData) []string {
	if members, ok := d.GetOk("members"); ok {
		membersResult := make([]string, 0)
		for _, member := range members.([]interface{}) {
			membersResult = append(membersResult, member.(string))
		}
		return membersResult
	}
	return make([]string, 0)
}
