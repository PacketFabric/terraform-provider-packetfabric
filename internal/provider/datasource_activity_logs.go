package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceActivityLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceActivityLogsRead,
		Schema: map[string]*schema.Schema{
			PfActivityLogs: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: PfActivityLogsDescription,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfLogUuid:      schemaStringComputed(PfLogUuidDescription),
						PfUser:         schemaStringComputed(PfUserDescription2),
						PfLevel:        schemaIntComputed(PfLevelDescription),
						PfCategory:     schemaStringComputed(PfCategoryDescription),
						PfEvent:        schemaStringComputed(PfEventDescription),
						PfMessage:      schemaStringComputed(PfMessageDescription),
						PfTimeCreated:  schemaStringComputed(PfTimeCreatedDescription2),
						PfLogLevelName: schemaStringComputed(PfLogLevelNameDescription),
					},
				},
			},
		},
	}
}

func datasourceActivityLogsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	activityLogs, err := c.GetActivityLogs()
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set(PfActivityLogs, flattenActivityLogs(&activityLogs)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenActivityLogs(logs *[]packetfabric.ActivityLog) []interface{} {
	fields := stringsToMap(PfLogUuid, PfUser, PfLevel, PfCategory, PfEvent, PfMessage, PfTimeCreated, PfLogLevelName)

	if logs != nil {
		flattens := make([]interface{}, len(*logs), len(*logs))
		for i, log := range *logs {
			flattens[i] = structToMap(log, fields)
		}
		return flattens
	}
	return make([]interface{}, 0)
}
