package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourcePortRouterLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourcePortRouterLogsRead,
		Description: PfPortRouterLogsDescription,
		Schema: map[string]*schema.Schema{
			PfPortCircuitId: schemaStringRequiredNotEmpty(PfPortCircuitIdDescription),
			PfTimeFrom:      schemaStringRequiredNotEmpty(PfTimeFromDescription),
			PfTimeTo:        schemaStringRequiredNotEmpty(PfTimeToDescription),
			PfPortRouterLogs: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						PfDeviceName:   schemaStringComputed(PfDeviceNameDescription),
						PfIfaceName:    schemaStringComputed(PfIfaceNameDescription),
						PfMessage:      schemaStringComputed(PfMessageDescription2),
						PfSeverity:     schemaIntComputed(PfSeverityDescription),
						PfSeverityName: schemaStringComputed(PfSeverityNameDescription),
						PfTimestamp:    schemaStringComputed(PfTimestampDescription),
					},
				},
			},
		},
	}
}

func datasourcePortRouterLogsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	portCID, ok := d.GetOk(PfPortCircuitId)
	if !ok {
		return diag.Errorf(MessageMissingPortCiruitId)
	}
	var err error
	timeFrom := d.Get(PfTimeFrom)
	timeTo := d.Get(PfTimeTo)
	logs, err := c.GetPortRouterLogs(portCID.(string), timeFrom.(string), timeTo.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set(PfPortRouterLogs, flattenPortRouterLogs(&logs))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(portCID.(string) + "-data")
	return diags
}

func flattenPortRouterLogs(logs *[]packetfabric.PortRouterLogs) []interface{} {
	fields := stringsToMap(PfDeviceName, PfIfaceName, PfMessage, PfSeverity, PfSeverityName, PfTimestamp)
	flattens := make([]interface{}, len(*logs))
	for i, log := range *logs {
		flattens[i] = structToMap(&log, fields)
	}
	return flattens
}
