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
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The Link Aggregation Group description.",
			},
			"interval": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(intervalOptions(), false),
				Description:  "Interval at which LACP packets are sent: fast/slow",
			},
			"pop": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Point of presence.",
			},
			"members": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of member port identifiers. All members must have the same speed and media.",
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "The member",
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
	resp, err := c.DeleteLinkAggregationGroup(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Link Aggregation Group delete workflow",
		Detail:   resp.WorkflowName,
	})
	return diags
}

func extractLAG(d *schema.ResourceData) packetfabric.LinkAggregationGroup {
	lag := packetfabric.LinkAggregationGroup{}
	lag.Description = d.Get("description").(string)
	lag.Interval = d.Get("interval").(string)
	lag.Pop = d.Get("pop").(string)
	lag.Members = d.Get("members").([]string)
	return lag
}

func intervalOptions() []string {
	return []string{"fast", "slow"}
}
