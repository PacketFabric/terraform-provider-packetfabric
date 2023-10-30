package provider

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLinkAggregationGroups() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLinkAggregationGroupsCreate,
		ReadContext:   resourceLinkAggregationGroupsRead,
		UpdateContext: resourceLinkAggregationGroupsUpdate,
		DeleteContext: resourceLinkAggregationGroupsDelete,
		Timeouts:      schema10MinuteTimeouts(),
		Schema: map[string]*schema.Schema{
			PfId:          schemaStringComputedPlain(),
			PfDescription: schemaStringRequiredNotEmpty(PfLagDescription),
			PfInterval:    schemaStringRequiredValidate(PfIntervalDescription, validateInterval()),
			PfPop:         schemaStringRequiredNewNotEmpty(PfPopDescriptionE),
			PfMembers:     schemaStringListRequiredNewNotEmptyDescribed(PfMembersDescription2, PfMembersDescription),
			PfPoNumber:    schemaStringOptionalValidate(PfPoNumberDescription, validatePoNumber()),
			PfLabels:      schemaStringSetOptional(PfLabelsDescription),
			PfEnabled:     schemaBoolOptionalDefault(PfEnabledDescription, true),
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

	if labels, ok := d.GetOk(PfLabels); ok {
		diagnostics, created := createLabels(c, d.Id(), labels)
		if !created {
			return diagnostics
		}
	}

	enabled := d.Get(PfEnabled).(bool)
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
	_, err := c.UpdateLinkAggregationGroup(d.Id(), d.Get(PfDescription).(string), d.Get(PfInterval).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange(PfPoNumber) {
		diagnostics, updated := updatePort(c, d)
		if !updated {
			return diagnostics
		}
	}

	if d.HasChange(PfLabels) {
		labels := d.Get(PfLabels)
		diagnostics, updated := updateLabels(c, d.Id(), labels)
		if !updated {
			return diagnostics
		}
	}

	if d.HasChange(PfEnabled) {
		enabled := d.Get(PfEnabled).(bool)
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

	poNumber, ok := d.GetOk(PfPoNumber)
	if !ok {
		return diag.FromErr(errors.New(MessageMissingPoNumber)), true
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

	_ = setResourceDataKeys(d, lag, PfDescription, PfPop, PfPoNumber)
	_ = d.Set(PfInterval, lag.LagInterval)

	if lag.Disabled {
		_ = d.Set(PfEnabled, false)
	} else {
		_ = d.Set(PfEnabled, true)
	}

	interfaces, err := c.GetLAGInterfaces(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	members := make([]string, len(*interfaces), len(*interfaces))
	for index, interf := range *interfaces {
		members[index] = interf.PortCircuitID
	}
	_ = d.Set(PfMembers, members)

	if _, ok := d.GetOk(PfLabels); ok {
		labels, err2 := getLabels(c, d.Id())
		if err2 != nil {
			return diag.FromErr(err2)
		}
		_ = d.Set(PfLabels, labels)
	}
	return diags
}

func resourceLinkAggregationGroupsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	host := os.Getenv(PfPfeHost)
	testingInLab := strings.Contains(host, PfDevLab)

	if testingInLab {
		enabled := d.Get(PfEnabled)
		if enabled.(bool) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  MessageDevLag,
			})
			_, err := c.DisableLinkAggregationGroup(d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
		}
		// allow time for LAG to be disabled
		time.Sleep(time.Duration(90) * time.Second)
	}
	time.Sleep(45 * time.Second)
	_, err := c.DeleteLinkAggregationGroup(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	time.Sleep(45 * time.Second)
	return diags
}

func extractLAG(d *schema.ResourceData) packetfabric.LinkAggregationGroup {
	lag := packetfabric.LinkAggregationGroup{}
	if description, ok := d.GetOk(PfDescription); ok {
		lag.Description = description.(string)
	}
	if interval, ok := d.GetOk(PfInterval); ok {
		lag.Interval = interval.(string)
	}
	if pop, ok := d.GetOk(PfPop); ok {
		lag.Pop = pop.(string)
	}
	lag.Members = extractMembers(d)
	return lag
}

func extractMembers(d *schema.ResourceData) []string {
	if members, ok := d.GetOk(PfMembers); ok {
		membersResult := make([]string, 0)
		for _, member := range members.([]interface{}) {
			membersResult = append(membersResult, member.(string))
		}
		return membersResult
	}
	return make([]string, 0)
}
