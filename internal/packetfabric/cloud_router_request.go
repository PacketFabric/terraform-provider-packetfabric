package packetfabric

import "fmt"

const cloudRouterRequestsURI = "/v2/services/cloud-routers/requests?request_type=%s"

type CloudRouterRequest struct {
	ImportCircuitID      string          `json:"import_circuit_id,omitempty"`
	CloudRouterCircuitID string          `json:"cloud_router_circuit_id,omitempty"`
	CustomerName         string          `json:"customer_name,omitempty"`
	ServiceUUID          string          `json:"service_uuid,omitempty"`
	State                string          `json:"state,omitempty"`
	TimeCreated          string          `json:"time_created,omitempty"`
	TimeUpdated          string          `json:"time_updated,omitempty"`
	RequestType          string          `json:"request_type,omitempty"`
	RejectionReason      string          `json:"rejection_reason,omitempty"`
	ImportFilters        []ImportFilters `json:"import_filters,omitempty"`
	ReturnFilters        []ReturnFilters `json:"return_filters,omitempty"`
}
type ImportFilters struct {
	Prefix    string `json:"prefix,omitempty"`
	MatchType string `json:"match_type,omitempty"`
	Localpref int    `json:"local_preference,omitempty"`
}
type ReturnFilters struct {
	Prefix          string `json:"prefix,omitempty"`
	MatchType       string `json:"match_type,omitempty"`
	Asprepend       int    `json:"as_prepend,omitempty"`
	Med             int    `json:"med,omitempty"`
	PendingApproval bool   `json:"pending_approval,omitempty"`
}

func (c *PFClient) GetCloudRouterRequests(reqType string) ([]CloudRouterRequest, error) {
	cloudRouterRequests := make([]CloudRouterRequest, 0)
	if _, err := c.sendRequest(fmt.Sprintf(cloudRouterRequestsURI, reqType), getMethod, nil, &cloudRouterRequests); err != nil {
		return nil, err
	}
	return cloudRouterRequests, nil
}
