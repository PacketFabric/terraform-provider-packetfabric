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

func (c *PFClient) IsBackboneComplete(vcCircuitID string) (result bool) {
	status, err := c.GetBackboneState(vcCircuitID)
	if status.Status.Current.State != "COMPLETE" {
		result = true
	}
	if err != nil {
		result = false
	}
	return
}
