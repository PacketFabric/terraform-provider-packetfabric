package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourcePortRouterLogs() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourcePortRouterLogsRead,
		Schema: map[string]*schema.Schema{
			"port_circuit_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Port identifier",
			},
			"time_from": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The ISO 8601 formatted datetime with optional timezone information, to filter from. Timezone defaults to UTC. Example: time_from=2020-05-23 00:00:00",
			},
			"time_to": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The ISO 8601 formatted datetime with optional timezone information, to filter from. Timezone defaults to UTC. Example: time_to=2020-05-23 00:00:00",
			},
			"port_router_logs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The device name.",
						},
						"iface_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The interface name.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The log message.",
						},
						"severity": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The log severity level.",
						},
						"severity_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The log severity name.",
						},
						"timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The log timestamp.",
						},
					},
				},
			},
		},
		Description: "The list of Port router logs.",
	}
}

func datasourcePortRouterLogsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	portCID, ok := d.GetOk("port_circuit_id")
	if !ok {
		return diag.Errorf("please provide a valid port circuit ID")
	}
	var err error
	timeFrom := d.Get("time_from")
	timeTo := d.Get("time_to")
	logs, err := c.GetPortRouterLogs(portCID.(string), timeFrom.(string), timeTo.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("port_router_logs", flattenPortRouterLogs(&logs))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(portCID.(string) + "-data")
	return diags
}

func flattenPortRouterLogs(logs *[]packetfabric.PortRouterLogs) []interface{} {
	flattens := make([]interface{}, len(*logs))
	for i, log := range *logs {
		flatten := make(map[string]interface{})
		flatten["device_name"] = log.DeviceName
		flatten["iface_name"] = log.IfaceName
		flatten["message"] = log.Message
		flatten["severity"] = log.Severity
		flatten["severity_name"] = log.SeverityName
		flatten["timestamp"] = log.Timestamp
		flattens[i] = flatten
	}
	return flattens
}
