package provider

import (
	"context"
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
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Point of presence in which the LAG should be located.",
			},
			"members": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "A list of port circuit IDs to include in the LAG. To be included in a LAG, the ports must be at the same site, in the same zone, and have the same speed and media.",
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "The member circuit ID.",
				},
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
	return diags
}

func resourceLinkAggregationGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	_, err := c.GetLAGInterfaces(d.Id())
	if err != nil {
		return diag.FromErr(err)
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
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Link Aggregation Group delete workflow",
		Detail:   resp.WorkflowName,
	})
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
