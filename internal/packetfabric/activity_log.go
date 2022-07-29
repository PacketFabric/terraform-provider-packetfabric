package packetfabric

const activityLogsURI = "/v2/activity-logs"

type ActivityLog struct {
	LogUUID     string `json:"log_uuid,omitempty"`
	User        string `json:"user,omitempty"`
	Level       int    `json:"level,omitempty"`
	Category    string `json:"category,omitempty"`
	Event       string `json:"event,omitempty"`
	Message     string `json:"message,omitempty"`
	TimeCreated string `json:"time_created,omitempty"`
}

func (c *PFClient) GetActivityLogs() (*[]ActivityLog, error) {
	expectedResp := make([]ActivityLog, 0)
	_, err := c.sendRequest(activityLogsURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return &expectedResp, nil
}
