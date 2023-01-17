package packetfabric

import "fmt"

const activityLogsURI = "/v2/activity-logs"

type ActivityLog struct {
	LogUUID     string `json:"log_uuid,omitempty"`
	User        string `json:"user,omitempty"`
	Level       int    `json:"level,omitempty"`
	Category    string `json:"category,omitempty"`
	Event       string `json:"event,omitempty"`
	Message     string `json:"message,omitempty"`
	TimeCreated string `json:"time_created,omitempty"`
	LevelName   string `json:"log_level_name,omitempty"`
}

func (c *PFClient) GetActivityLogs() ([]ActivityLog, error) {
	expectedResp := make([]ActivityLog, 0)
	_, err := c.sendRequest(activityLogsURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return &expectedResp, nil
}

func (c *PFClient) GetCloudRouterRequests(reqType string) ([]CloudRouterRequest, error) {
	cloudRouterRequests := make([]CloudRouterRequest, 0)
	if _, err := c.sendRequest(fmt.Sprintf(cloudRouterRequestsURI, reqType), getMethod, nil, &cloudRouterRequests); err != nil {
		return nil, err
	}
	return cloudRouterRequests, nil
}
