package provider

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceStreamingEvents() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceStreamingEventsRead,
		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:        schema.TypeString,
				Description: "UUID of the subscription bundle",
				Required:    true,
			},
			"events": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
				Optional: true,
			},
			"stream_time": {
				Type:        schema.TypeInt,
				Description: "How long stream logs should be captured for in minutes.",
				Default:     1,
				Optional:    true,
			},
		},
	}
}

const streamingEventsURIGet = "/v2/events/%s"

type Decoder struct {
	src  io.Reader
	buff *bufio.Reader
}

func NewDecoder(src io.Reader) *Decoder {
	return &Decoder{
		src:  src,
		buff: bufio.NewReader(src),
	}
}

func (d *Decoder) Decode() (packetfabric.StreamingEventsGetResponse, error) {
	event := packetfabric.StreamingEventsGetResponse{}
	for {
		line, err := d.buff.ReadBytes('\n')
		if err != nil {
			return event, fmt.Errorf("the following error occurred while reading bytes: %s", err)
		}
		if len(line) == 1 { // empty line signifies end of event
			return event, nil
		}

		fields := bytes.SplitN(line, []byte(":"), 2)
		field := string(bytes.TrimSpace(fields[0]))

		switch field {
		case "event":
			event.Event = string(bytes.TrimSpace(fields[1]))
		case "data":
			data := packetfabric.StreamingEventData{}
			_ = json.Unmarshal(bytes.TrimSpace(fields[1]), &data)
			event.Data = &data
		}
	}
}

func datasourceStreamingEventsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	formattedURI := fmt.Sprintf("%s%s", c.HostURL, fmt.Sprintf(streamingEventsURIGet, d.Get("subscription_id").(string)))
	req, _ := http.NewRequestWithContext(c.Ctx, "GET", formattedURI, nil)
	token := &c.Token

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("pf-request-source", "terraform")

	resp, err1 := c.HTTPClient.Do(req)
	if err1 != nil {
		return diag.FromErr(err1)
	}

	if resp.StatusCode == http.StatusBadRequest ||
		resp.StatusCode == http.StatusUnauthorized ||
		resp.StatusCode == http.StatusNotFound {
		body, _ := ioutil.ReadAll(resp.Body)
		return diag.FromErr(fmt.Errorf("Status: %d, body: %s", resp.StatusCode, body))
	}

	defer resp.Body.Close()
	decoder := NewDecoder(resp.Body)
	for {
		event, err3 := decoder.Decode()
		if err3 != nil {
			break
		}

		if event.Event != "" {
			tflog.Debug(ctx, fmt.Sprintf("Setting event: %s", event.Event))
		}
	}
	return diags
}

func setFields(d *schema.ResourceData, event packetfabric.StreamingEventsGetResponse) {
	events := d.Get("events")
	var eventsData []string
	for _, eventData := range events.([]interface{}) {
		eventsData = append(eventsData, eventData.(string))
	}
	eventsData = append(eventsData, event.Event)
	_ = d.Set("events", eventsData)
}
