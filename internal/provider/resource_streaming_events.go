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
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of events to subscribe to.",
			},
			"events": {
				Type:        schema.TypeList,
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func getStringListData(d *schema.ResourceData, key string) []string {
	var data []string
	for _, i := range d.Get(key).(*schema.Set).List() {
		data = append(data, i.(string))
	}
	return data
}

func resourceStreamingEventsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	streamingType := d.Get("type")
	events := getStringListData(d, "events")

	payload := packetfabric.StreamingEventsPayload{
		Type:   streamingType.(string),
		Events: events,
	}

	if _, ok := d.GetOk("vcs"); ok {
		payload.VCS = getStringListData(d, "vcs")
	}
	if _, ok := d.GetOk("ifds"); ok {
		payload.IFDs = getStringListData(d, "ifds")
	}

	resp, err := c.CreateStreamingEvent(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.SubscriptionUUID)
	return diags
}

func resourceStreamingEventsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics

	_, err := c.GetStreamingEvent(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
