package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceActivityLog() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceActivityLogRead,
		Schema: map[string]*schema.Schema{
			"activity_logs": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "The active logs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The log UUID.",
						},
						"user": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The log User.",
						},
						"level": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The log level.",
						},
						"category": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The log Category.",
						},
						"event": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The log Event.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The log Message.",
						},
						"time_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The log time created.",
						},
					},
				},
			},
		},
	}
}

func datasourceActivityLogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	activityLogs, err := c.GetActivityLogs()
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("activity_logs", flattenActivityLogs(activityLogs)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenActivityLogs(logs *[]packetfabric.ActivityLog) []interface{} {
	if logs != nil {
		flattens := make([]interface{}, len(*logs), len(*logs))
		for i, log := range *logs {
			flatten := make(map[string]interface{})
			flatten["log_uuid"] = log.LogUUID
			flatten["user"] = log.User
			flatten["level"] = log.Level
			flatten["category"] = log.Category
			flatten["event"] = log.Event
			flatten["messge"] = log.Message
			flatten["time_created"] = log.TimeCreated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
