package packetfabric

const streamingEventsURI = "/v2/events"

type StreamingEventsCreateResponse struct {
	SubscriptionUUID string `json:"subscription_uuid"`
}

type StreamingEventData struct {
	User         string `json:"user"`
	LogLevel     string `json:"log_level"`
	Category     string `json:"category"`
	Event        string `json:"event"`
	Message      string `json:"message"`
	TimeStamp    string `json:"timestamp"`
	Lag          string `json:"lag,omitempty"`
	WorkflowID   string `json:"workflow_id,omitempty"`
	LagCircuitID string `json:"lag_circuit_id,omitempty"`
	CircuitID    string `json:"circuit_id,omitempty"`
}

type StreamingEventsGetResponse struct {
	Event string              `json:"event"`
	Data  *StreamingEventData `json:"data"`
}

type StreamData struct {
	Type   string   `json:"type"`
	Events []string `json:"events"`
	VCS    []string `json:"vcs,omitempty"`
	IFDs   []string `json:"ifds,omitempty"`
}

type StreamingEventsPayload struct {
	Streams []StreamData `json:"streams"`
}

func (c *PFClient) CreateStreamingEvent(streamingEventsData StreamingEventsPayload) (*StreamingEventsCreateResponse, error) {
	resp := &StreamingEventsCreateResponse{}
	_, err := c.sendRequest(streamingEventsURI, postMethod, streamingEventsData, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
