package packetfabric

import "fmt"

const streamingEventsURI = "/v2/events"
const streamingEventsURIGet = "/v2/events/%s"

type StreamingEventsCreateResponse struct {
	SubscriptionUUID string `json:"subscription_uuid"`
}

type StreamingEventData struct {
	User      string `json:"user"`
	LogLevel  string `json:"log_level"`
	Category  string `json:"category"`
	Event     string `json:"event"`
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp"`
}

type StreamingEventsGetResponse struct {
	Event string              `json:"event"`
	Data  *StreamingEventData `json:"data"`
}

type StreamingEventsPayload struct {
	Type   string   `json:"type"`
	Events []string `json:"events"`
	VCS    []string `json:"vcs,omitempty"`
	IFDs   []string `json:"ifds,omitempty"`
}

func (c *PFClient) CreateStreamingEvent(streamingEventsData StreamingEventsPayload) (*StreamingEventsCreateResponse, error) {
	resp := &StreamingEventsCreateResponse{}
	_, err := c.sendRequest(streamingEventsURI, postMethod, streamingEventsData, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *PFClient) GetStreamingEvent(subscriptionUUID string) (*StreamingEventsGetResponse, error) {
	formattedURI := fmt.Sprintf(streamingEventsURIGet, subscriptionUUID)
	resp := &StreamingEventsGetResponse{}
	_, err := c.sendRequest(formattedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
