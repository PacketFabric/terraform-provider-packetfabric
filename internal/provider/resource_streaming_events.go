package provider

import (
	"context"
	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const VCSDescription = "Specific logical interfaces you wish to subscribe to, each identified by a combination of the virtual circuit ID and port circuit ID " +
	"associated with the logical interface. If none are supplied, then all logical interfaces to which the customer has access are assumed."

func resourceStreamingEvents() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStreamingEventsCreate,
		ReadContext:   resourceStreamingEventsRead,
		DeleteContext: resourceStreamingEventsDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID for created bundle of event streams",
			},
			"streams": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of events to subscribe to.",
						},
						"events": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Categories of events to subscribe to. If not specified, then all event categories are assumed.",
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},
						"vcs": {
							Type:        schema.TypeList,
							Description: VCSDescription,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
							Optional: true,
						},
						"ifds": {
							Type:        schema.TypeList,
							Description: "Specific ports you wish to subscribe to, identified by port circuit IDs. If none are supplied, then all ports to which the customer has access are assumed.",
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
							Optional: true,
						},
					},
				},
				ForceNew: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func getStringListData(data []interface{}) []string {
	var StringListData []string
	for _, i := range data {
		StringListData = append(StringListData, i.(string))
	}
	return StringListData
}

func resourceStreamingEventsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx

	var diags diag.Diagnostics
	var streamsData []packetfabric.StreamData
	streams := d.Get("streams")

	for _, stream := range streams.(*schema.Set).List() {
		streamData := stream.(map[string]interface{})
		payload := packetfabric.StreamData{
			Type:   streamData["type"].(string),
			Events: getStringListData(streamData["events"].([]interface{})),
		}

		if len(streamData["vcs"].([]interface{})) > 0 {
			payload.VCS = getStringListData(streamData["vcs"].([]interface{}))
		}

		if len(streamData["ifds"].([]interface{})) > 0 {
			payload.IFDs = getStringListData(streamData["ifds"].([]interface{}))
		}

		streamsData = append(streamsData, payload)
	}

	resp, err := c.CreateStreamingEvent(packetfabric.StreamingEventsPayload{Streams: streamsData})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.SubscriptionUUID)
	return diags
}

func resourceStreamingEventsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceStreamingEventsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return diag.Diagnostics{}
}
