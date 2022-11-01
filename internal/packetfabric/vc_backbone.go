package packetfabric

import (
	"fmt"
)

const vcBackStatusURI = "/v2.1/services/%s/status"

func (c *PFClient) GetBackboneState(vcCircuitID string) (*ServiceState, error) {
	formatedURI := fmt.Sprintf(vcBackStatusURI, vcCircuitID)
	expectedResp := &ServiceState{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, expectedResp)
	if err != nil {
		return nil, err
	}
	return expectedResp, nil
}

func (c *PFClient) IsBackboneComplete(vcCircuitID string) bool {
	status, err := c.GetBackboneState(vcCircuitID)
	if err == nil && status.Status.LastWorkflow.CurrentState == "COMPLETE" {
		return true
	}
	if err != nil {
		// We need to return TRUE in case of error since the server
		// erases the status history after it reachs COMPLETE.
		return true
	}
	return false
}
